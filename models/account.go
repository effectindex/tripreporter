package models

import (
	"context"
	"strconv"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

type Account struct { // todo: this should be oauth / credentials. allow changing email or logging in with google
	Context
	Unique
	Type     string `json:"type"`
	Email    string `json:"email" db:"email"`                   // Optional. Make clear that password reset isn't possible if not set.
	Username string `json:"username" db:"username"`             // Required. Generate from wordlist + 3 numbers if left blank.
	Password string `json:"password_hash" db:"password_hash"`   // Required
	Verified bool   `json:"email_verified" db:"email_verified"` // Optional. Whether email has been verified or not.
}

func (a *Account) Get() (*Account, error) { // TODO: Implement a.verified / other params
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
		return a, ErrorAccountNotSpecified
	}

	var a1 []*Account
	if err := pgxscan.Select(context.Background(), db, &a1,
		`select id, email, username, password_hash from accounts `+query, queryArg,
	); err != nil {
		a.Logger.Warnw("Failed to get account from DB", zap.Error(err))
		return a, err
	} else if len(a1) == 0 {
		return a, ErrorAccountNotFound
	} else if len(a1) > 1 { // This shouldn't happen
		a.Logger.Errorw("Multiple accounts found for parameters", "accounts", a1)
		return a, ErrorAccountNotSpecified
	} else {
		a.ID = a1[0].ID
		a.Email = a1[0].Email
		a.Username = a1[0].Username
		a.Password = a1[0].Password
	}

	return a, nil
}

func (a *Account) Post() (*Account, error) { // TODO: Email verification? / post signup hook?
	db := a.DB()
	defer db.Commit(context.Background())

	if err := a.InitUUID(a.Logger); err != nil {
		return a, err
	}

	if _, err := db.Exec(context.Background(),
		`insert into accounts(
			id,
			email,
			username,
			password_hash
		)
		values(
			$1,
			$2,
			$3,
			$4
		);`,
		a.ID,
		a.Email,    // TODO: Validate emails
		a.Username, // TODO: Validate usernames
		a.Password, // TODO: Salt / verified in DB?
	); err != nil {
		a.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	return a, nil
}

func (a *Account) Patch() (*Account, error) {
	db := a.DB()

	if a.NilUUID() {
		return a, ErrorAccountNotSpecified
	}

	fields := make([]interface{}, 0)
	query := "update accounts set"
	qNum := 0

	if a.Email != "" { // TODO: Validate emails
		qNum++
		query += " email=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, a.Email)
	}

	if a.Username != "" { // TODO: Validate usernames
		qNum++
		query += " username=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, a.Username)
	}

	if a.Password != "" {
		qNum++
		query += " password_hash=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, a.Password)
	}

	// TODO: Implement a.Verified

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
	db := a.DB()
	defer db.Commit(context.Background())

	a1 := a.CopyIdentifiers()
	if _, err := a1.Get(); err != nil {
		return a, err
	} else if a.Password != a1.Password {
		return a, ErrorAccountPasswordMatch
	}

	if _, err := db.Exec(context.Background(), `delete from accounts where id=$1 and password_hash=$2;`, a.ID, a.Password); err != nil {
		a.Logger.Warnw("Failed to delete account from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return a, err
	}

	return nil, nil
}

func (a *Account) CopyIdentifiers() *Account {
	return &Account{Context: a.Context, Unique: Unique{ID: a.ID}, Email: a.Email, Username: a.Username}
}
