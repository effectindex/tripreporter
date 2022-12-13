package models

import (
	"context"
	"encoding/json"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"github.com/georgysavva/scany/v2/pgxscan"
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
	Password []byte `json:"password_hash" db:"password_hash"`   // Required. Generated from Salt using Argon2ID and is 32 bits long.
	Verified bool   `json:"email_verified" db:"email_verified"` // Optional. Whether email has been verified or not.
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
			MinUniqueNonSymbol: 2,
			Message:            "a-z 0-9 _ -",
			Allowed: allowedChars{
				Symbol: map[string]bool{
					"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true,
					"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true,
				},
				NonSymbol: map[string]bool{
					"_": true, "-": true,
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
					"A": true, "B": true, "C": true, "D": true, "E": true, "F": true, "G": true, "H": true, "I": true, "J": true, "K": true, "L": true, "M": true, "N": true, "O": true, "P": true, "Q": true, "R": true, "S": true, "T": true, "U": true, "V": true, "W": true, "X": true, "Y": true, "Z": true,
					"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true,
					"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true,
				},
				NonSymbol: map[string]bool{
					"_": true, "-": true, "!": true, "@": true, "#": true, "$": true, "%": true, "^": true, "&": true, "*": true, "(": true, ")": true, "+": true, "=": true, " ": true,
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
		query = `where id = $1;`
		queryArg = a.ID.String()
	} else if a.Email != "" {
		query = `where email = $1;`
		queryArg = a.Email
	} else if a.Username != "" {
		query = `where username = $1;`
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
		a.ID = a1[0].ID
		a.Email = a1[0].Email
		a.Username = a1[0].Username
		a.Salt = a1[0].Salt
		a.Password = a1[0].Password
		a.Verified = a1[0].Verified
	}

	return a, nil
}

func (a *Account) Post() (*Account, error) { // TODO: Email verification? / post signup hook?
	a.InitType(a)
	db := a.DB()
	defer db.Commit(context.Background())

	// Init account UUID
	if err := a.InitUUID(a.Logger); err != nil {
		return a, err
	}

	// Validate email, username and password
	if a, err := a.ValidateEmail(); err != nil {
		return a, err
	}
	if a, err := a.ValidateUsername(); err != nil {
		return a, err
	}
	if a, err := a.ValidatePassword(); err != nil {
		return a, err
	}

	// Check if email or username are already taken
	var a1 []*Account
	_ = pgxscan.Select(context.Background(), db, &a1,
		`select * from accounts where email=$1`, a.Email,
	)

	var a2 []*Account
	_ = pgxscan.Select(context.Background(), db, &a2,
		`select * from accounts where username=$1`, a.Username,
	)

	if len(a1) > 0 {
		return a, types.ErrorAccountEmailExists
	}
	if len(a2) > 0 {
		return a, types.ErrorAccountUsernameExists
	}

	// Now we can generate the salt to use for the password
	salt, err := util.GenerateSalt(12, 16, Wordlist.Random(1))

	if err != nil {
		a.Logger.Warnw("Failed to generate salt", "ID", a.ID, zap.Error(err))
		return a, err
	}

	// Set the account's salt and new hashed password properly
	a.Salt = salt
	a.Password = util.GenerateSaltedPasswordHash(a.Password, a.Salt)

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
		a.Password,
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

	fields := make([]interface{}, 0)
	query := "update accounts set"
	qNum := 0

	addQuery := func(s string, i interface{}) {
		qNum++
		query += " " + s + "=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, i)
	}

	if a.Email != "" { // TODO: Validate emails
		addQuery("email", a.Email)
	}

	if a.Username != "" { // TODO: Validate usernames
		addQuery("username", a.Username)
	}

	if len(a.Salt) > 0 {
		addQuery("password_salt", a.Salt)

	}

	if len(a.Password) > 0 {
		addQuery("password_hash", a.Password)
	}

	addQuery("email_verified", a.Verified)

	query = strings.TrimSuffix(query, ",")
	qNum++
	query += " where id=$;" + strconv.Itoa(qNum)
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

	a1 := a.ClearSensitive()
	if _, err := a1.Get(); err != nil {
		return a, err
	} else if !util.SliceEqual(a.Password, a1.Password) {
		return a, types.ErrorAccountPasswordMatch
	}

	if _, err := db.Exec(context.Background(), `delete from accounts where id=$1 and password_hash=$2;`, a.ID, a.Password); err != nil {
		a.Logger.Warnw("Failed to delete account from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	return nil, nil
}

func (a *Account) User() (*User, error) {
	a.Type = "" // We want Get() here to set the new type

	u, err := (&User{Context: a.Context, Unique: a.Unique}).Get()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *Account) ClearSensitive() *Account {
	a.InitType(a)
	return &Account{Context: a.Context, Unique: a.Unique, Email: a.Email, Username: a.Username, Verified: a.Verified}
}

func (a *Account) VerifyPassword(password string) (*Account, error) {
	a.InitType(a)

	if len(a.Salt) == 0 { // should not really be possible in a real scenario
		return a, types.ErrorStringEmpty.PrefixedError("Salt")
	}

	if len(a.Password) == 0 {
		return a, types.ErrorStringEmpty.PrefixedError("Password")
	}

	hash := util.GenerateSaltedPasswordHash([]byte(password), a.Salt)
	if !util.SliceEqual(hash, a.Password) {
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

	if err, ok := err.(types.ErrorString); ok {
		return a, err.PrefixedError("Username")
	}
	return a, err
}

func (a *Account) ValidatePassword() (*Account, error) {
	a.InitType(a)

	err := AccountCfg.Password.Validate(string(a.Password))
	if err, ok := err.(types.ErrorString); ok {
		return a, err.PrefixedError("Password")
	}
	return a, err
}
