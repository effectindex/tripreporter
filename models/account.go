// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/effectindex/tripreporter/crypto"
	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var (
	AccountCfg AccountConfig
)

type Account struct { // todo: this should be oauth / credentials. allow changing email or logging in with google
	types.Context
	Unique
	Email    string `json:"email" db:"email"`                   // Optional. Make clear that password reset isn't possible if not set.
	Username string `json:"username" db:"username"`             // Required. Generate from wordlist + 3 numbers if left blank.
	Salt     []byte `json:"password_salt" db:"password_salt"`   // Required. Generated from random []byte(16), + wordlist(1) + []byte(32-16-len(word)).
	Hash     []byte `json:"password_hash" db:"password_hash"`   // Required. Generated from Salt using Argon2ID and is 32 bits long.
	Verified bool   `json:"email_verified" db:"email_verified"` // Optional. Whether email has been verified or not.

	Password string `json:"password"`           // Optional. Only used in API requests, is here to so API users aren't confused by `password_hash` when making a new account.
	NewPass  string `json:"new_password"`       // Optional. Only used in PATCH API requests, when changing the password.
	NewUser  *User  `json:"new_user,omitempty"` // Optional. Only used in API requests, when creating an account.
}

type AccountPublic struct {
	types.Context
	Unique
	Email    string `json:"email"`
	Username string `json:"username"`
	Verified bool   `json:"email_verified"`
}

type AccountConfig struct {
	Username StringRestriction `json:"username"`
	Password StringRestriction `json:"password"`
}

func SetupAccountConfig(ctx types.Context) {
	ctx.Validate()

	defaultConfig := AccountConfig{
		Username: StringRestriction{
			MinLength:          3,
			MaxLength:          32,
			MinUniqueTotal:     1,
			MinUniqueSymbol:    0,
			MinUniqueNonSymbol: 0,
			Message:            "a-z 0-9 _ -",
			Allowed: allowedChars{
				Symbol: map[string]bool{
					"_": true, "-": true,
				},
				NonSymbol: map[string]bool{
					"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true,
					"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true,
				},
			},
		},
		Password: StringRestriction{
			MinLength:          8,
			MaxLength:          1024,
			MinUniqueTotal:     5,
			MinUniqueSymbol:    2,
			MinUniqueNonSymbol: 3,
			Message:            "A-Z a-z 0-9 _ - ! @ # $ % ^ & * ( ) + = space",
			Allowed: allowedChars{
				Symbol: map[string]bool{
					"_": true, "-": true, "!": true, "@": true, "#": true, "$": true, "%": true, "^": true, "&": true, "*": true, "(": true, ")": true, "+": true, "=": true, " ": true,
				},
				NonSymbol: map[string]bool{
					"A": true, "B": true, "C": true, "D": true, "E": true, "F": true, "G": true, "H": true, "I": true, "J": true, "K": true, "L": true, "M": true, "N": true, "O": true, "P": true, "Q": true, "R": true, "S": true, "T": true, "U": true, "V": true, "W": true, "X": true, "Y": true, "Z": true,
					"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true,
					"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true,
				},
			},
		},
	}

	if bytes, err := os.ReadFile(os.Getenv("ACCOUNT_CONFIG")); err != nil {
		ctx.Logger.Panicw("Failed to read ACCOUNT_CONFIG file", zap.Error(err))
	} else {
		var cfg AccountConfig
		usingDefault := ""

		if err := json.Unmarshal(bytes, &cfg); err != nil {
			ctx.Logger.Warnw("Failed to unmarshal account config", zap.Error(err))
			AccountCfg = defaultConfig
			usingDefault = "default "
		} else {
			AccountCfg = cfg
		}

		ctx.Logger.Infof("Loaded %saccount config with %v allowed chars", usingDefault, len(AccountCfg.Username.Allowed.NonSymbol)+len(AccountCfg.Username.Allowed.Symbol))
	}
}

func (a *Account) Get() (*Account, error) { // TODO: Implement a.verified / other params
	a.InitType(a)
	db := a.DB()
	defer db.Commit(context.Background())

	var query string
	var queryArg string

	if !a.NilUUID() {
		query = `where id=$1;`
		queryArg = a.ID.String()
	} else if a.Email != "" {
		query = `where email=$1;`
		queryArg = a.Email
	} else if a.Username != "" {
		query = `where username=$1;`
		queryArg = a.Username
	} else {
		return a, types.ErrorAccountNotSpecified
	}

	var a1 []*Account
	if err := pgxscan.Select(context.Background(), db, &a1,
		`select * from accounts `+query, queryArg,
	); err != nil {
		a.Logger.Warnw("Failed to get account from DB", zap.Error(err))
		return a, err
	} else if len(a1) == 0 {
		return a, types.ErrorAccountNotFound
	} else if len(a1) > 1 { // This shouldn't happen
		a.Logger.Errorw("Multiple accounts found for parameters", "accounts", a1)
		return a, types.ErrorAccountNotSpecified
	} else {
		a.FromData(a1[0])
	}

	return a, nil
}

func (a *Account) Post() (*Account, error) { // TODO: Email verification? / post signup hook?
	a.InitType(a)
	db := a.DB()
	defer db.Commit(context.Background())

	// Init account UUID
	if err := a.InitUUIDv4(a.Logger); err != nil {
		return a, err
	}

	// Validate email, username and password
	if a, err := a.ValidateEmail(); err != nil {
		return a, err
	}
	if a, err := a.ValidateUsername(); err != nil {
		return a, err
	}
	if a, err := a.ValidatePassword(a.Password, "Password"); err != nil {
		return a, err
	}

	// Check if email or username are already taken
	if _, err := a.ExistsWithEmail(db); err != nil {
		return a, err
	}
	if _, err := a.ExistsWithUsername(db); err != nil {
		return a, err
	}

	// Now we can generate the salt to use for the password
	salt, err := crypto.GenerateSalt(12, 16, Wordlist.Random(1))

	if err != nil {
		a.Logger.Warnw("Failed to generate salt", "ID", a.ID, zap.Error(err))
		return a, err
	}

	// Set the account's salt and new hashed password properly
	a.Salt = salt
	a.Hash = crypto.GenerateSaltedPasswordHash([]byte(a.Password), a.Salt)

	if _, err := db.Exec(context.Background(),
		`insert into accounts(
			id,
			email,
			username,
			password_salt,
			password_hash,
			email_verified
		)
		values(
			$1,
			$2,
			$3,
			$4,
		    $5,
		    $6
		);`,
		a.ID,
		a.Email,
		a.Username,
		a.Salt,
		a.Hash,
		a.Verified,
	); err != nil {
		a.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	return a, nil
}

func (a *Account) Patch() (*Account, error) {
	a.InitType(a)
	db := a.DB()

	if a.NilUUID() {
		return a, types.ErrorAccountNotSpecified
	}

	// Verify account exists, and that password is being provided to update account info
	a1 := &Account{Context: a.Context}
	a1.FromData(a)
	if _, err := a1.Get(); err != nil {
		return a, err
	}

	if _, err := a1.VerifyPassword(a.Password); err != nil {
		return a, err
	}

	fields := make([]interface{}, 0)
	query := "update accounts set"
	qNum := 0

	addQuery := func(s string, i interface{}) {
		qNum++
		query += " " + s + "=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, i)
	}

	if a.Email != "" {
		if a, err := a.ValidateEmail(); err != nil {
			return a, err
		}
		if a1, err := a.ExistsWithEmail(db); err != nil && a.ID != a1.ID {
			return a, err
		}
		addQuery("email", a.Email)
	}

	if a.Username != "" {
		if a, err := a.ValidateUsername(); err != nil {
			return a, err
		}
		if a1, err := a.ExistsWithUsername(db); err != nil && a.ID != a1.ID {
			return a, err
		}
		addQuery("username", a.Username)
	}

	if len(a.NewPass) > 0 {
		if a, err := a.ValidatePassword(a.NewPass, "New password"); err != nil {
			return a, err
		}
		if len(a.Salt) == 0 {
			salt, err := crypto.GenerateSalt(12, 16, Wordlist.Random(1))
			if err != nil {
				a.Logger.Warnw("Failed to generate salt", "ID", a.ID, zap.Error(err))
				return a, err
			}

			a.Salt = salt
		}
		a.Hash = crypto.GenerateSaltedPasswordHash([]byte(a.NewPass), a.Salt)
	}

	if len(a.Salt) > 0 {
		addQuery("password_salt", a.Salt)
	}

	if len(a.Hash) > 0 {
		addQuery("password_hash", a.Hash)
	}

	addQuery("email_verified", a.Verified)

	query = strings.TrimSuffix(query, ",")
	qNum++
	query += " where id=$" + strconv.Itoa(qNum)
	fields = append(fields, a.ID)

	_, err := db.Exec(context.Background(), query, fields...)

	if err != nil {
		a.Logger.Warnw("Failed to update account in DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	// Want to commit before Get()
	db.Commit(context.Background())
	return a.Get()
}

func (a *Account) Delete() (*Account, error) {
	a.InitType(a)
	db := a.DB()
	defer db.Commit(context.Background())

	if a.NilUUID() {
		return a, types.ErrorAccountNotSpecified
	}

	// Verify account exists, and that password is being provided to update account info
	a1 := &Account{Context: a.Context}
	a1.FromData(a)
	if _, err := a1.Get(); err != nil {
		return a, err
	}

	if _, err := a1.VerifyPassword(a.Password); err != nil {
		return a, err
	}

	if _, err := db.Exec(context.Background(), `delete from accounts where id=$1;`, a.ID); err != nil {
		a.Logger.Warnw("Failed to delete account from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	// TODO: Part of refactoring to pointer-based model
	return a, nil
}

func (a *Account) User() (*User, error) {
	a.Type = "" // We want Get() here to set the new type

	u, err := (&User{Context: a.Context, Unique: a.Unique}).Get()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *Account) FromRefreshToken(token *http.Cookie) (*Account, error) {
	a.InitType(a)

	if token == nil || len(token.Value) == 0 {
		return a, types.ErrorStringEmpty.PrefixedError(types.CookieRefreshToken)
	}

	val, err := a.Cache.Get(context.Background(), token.Value).Result()
	if err == redis.Nil {
		return a, types.ErrorSessionRefreshNotFound
	} else if err != nil {
		a.Logger.Debugw("Failed to get account from refresh token", zap.Error(err))
		return a, err
	}

	id, err := uuid.Parse(val)
	if err != nil { // this means an invalid UUID got put into Redis somehow
		a.Logger.Warnw("Failed to parse valid account UUID from refresh token", zap.Error(err))
		return a, types.ErrorAccountNotFound
	}

	a.ID = id
	return a, nil
}

func (a *Account) FromBody(r *http.Request) (*Account, error) {
	a.InitType(a)

	if r.Body == nil {
		return a, types.ErrorStringEmpty.PrefixedError("Request body")
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return a, err
	}

	if len(body) == 0 {
		return a, types.ErrorStringEmpty.PrefixedError("Request body")
	}

	var a1 *Account
	err = json.Unmarshal(body, &a1)
	if err != nil {
		return a, err
	}

	a.FromData(a1)

	if a.NewUser != nil {
		// Make sure to copy over the context
		a.NewUser.Context = a.Context

		// Don't allow requests to set the user ID, we want to use the account ID.
		a.NewUser.Unique = a.Unique
		a.NewUser.Unique.InitType(a.Logger)

		// Don't allow requests to set created value themselves, we want to set it ourselves.
		a.NewUser.Created.Default()
	}

	return a, nil // TODO: Ensure a.Unique.Type is not mutated by the above
}

func (a *Account) FromData(a1 *Account) {
	a.InitType(a)
	a.ID = a1.ID
	a.Email = a1.Email
	a.Username = a1.Username
	a.Salt = a1.Salt
	a.Hash = a1.Hash
	a.Verified = a1.Verified
	a.Password = a1.Password
	a.NewPass = a1.NewPass
	a.NewUser = a1.NewUser
}

func (a *Account) ClearAll() *AccountPublic {
	a.InitType(a)
	return &AccountPublic{Context: a.Context, Unique: a.Unique}
}

func (a *Account) ClearImmutable() *Account {
	a.InitType(a)
	return &Account{Context: a.Context, Unique: a.Unique, Email: a.Email, Username: a.Username, Password: a.Password, NewPass: a.NewPass, NewUser: a.NewUser}
}

func (a *Account) CopyPublic() *AccountPublic {
	p := &AccountPublic{Context: a.Context, Unique: a.Unique, Email: a.Email, Username: a.Username, Verified: a.Verified}
	p.InitType(p)
	return p
}

func (a *Account) VerifyPassword(password string) (*Account, error) {
	a.InitType(a)

	if len(a.Salt) == 0 { // should not really be possible in a real scenario
		return a, types.ErrorStringEmpty.PrefixedError("Salt")
	}

	if len(password) == 0 {
		return a, types.ErrorStringEmpty.PrefixedError("Password")
	}

	providedHash := crypto.GenerateSaltedPasswordHash([]byte(password), a.Salt)
	if subtle.ConstantTimeCompare(providedHash, a.Hash) != 1 {
		return a, types.ErrorAccountPasswordMatch
	}

	return a, nil
}

func (a *Account) ValidateEmail() (*Account, error) {
	a.InitType(a)

	if len(a.Email) == 0 {
		return a, types.ErrorStringEmpty.PrefixedError("Email")
	}

	addr, err := mail.ParseAddress(a.Email)
	if err != nil {
		return a, err
	}

	domain := strings.Split(addr.Address, "@")
	if len(domain) == 0 {
		return a, types.ErrorAccountEmailDomainEmpty
	}

	if !strings.Contains(domain[len(domain)-1], ".") {
		return a, types.ErrorAccountEmailTLDEmpty
	}

	return a, nil
}

func (a *Account) ValidateUsername() (*Account, error) {
	a.InitType(a)

	err := AccountCfg.Username.Validate(a.Username)

	return a, types.PrefixedError(err, "Username")
}

func (a *Account) ValidatePassword(password, prefix string) (*Account, error) {
	a.InitType(a)

	err := AccountCfg.Password.Validate(password)

	return a, types.PrefixedError(err, prefix)
}

func (a *Account) ExistsWithEmail(db pgx.Tx) (*Account, error) {
	var a1 []*Account
	_ = pgxscan.Select(context.Background(), db, &a1,
		`select * from accounts where email=$1`, a.Email,
	)

	if len(a1) > 0 {
		return a1[0], types.ErrorAccountEmailExists
	}

	return a, nil
}

func (a *Account) ExistsWithUsername(db pgx.Tx) (*Account, error) {
	var a1 []*Account
	_ = pgxscan.Select(context.Background(), db, &a1,
		`select * from accounts where username=$1`, a.Username,
	)

	if len(a1) > 0 {
		return a1[0], types.ErrorAccountUsernameExists
	}

	return a, nil
}
