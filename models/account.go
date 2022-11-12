package models

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"go.uber.org/zap"
)

func (a *Account) Get() error { // TODO: Implement a.verified / other params
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
		return ErrorAccountNotSpecified
	}

	var a1 []*Account
	if err := pgxscan.Select(context.Background(), db, &a1,
		`select id, email, username, password_hash from accounts `+query, queryArg,
	); err != nil {
		a.Logger.Warnw("Failed to get account from DB", zap.Error(err))
		return err
	} else if len(a1) == 0 {
		return ErrorAccountNotFound
	} else if len(a1) > 1 { // This shouldn't happen
		a.Logger.Errorw("Multiple accounts found for parameters", "account", a)
		return ErrorAccountNotSpecified
	} else {
		a.ID = a1[0].ID
		a.Email = a1[0].Email
		a.Username = a1[0].Username
		a.Password = a1[0].Password
	}

	// todo
	return nil
}

func (a *Account) Post() error {
	db := a.DB()
	defer db.Commit(context.Background())

	if err := a.InitUUID(a.Logger); err != nil {
		return err
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
		a.Email,
		a.Username,
		a.Password, // TODO: Salt / verified in DB?
	); err != nil {
		a.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return err
	}

	return nil
}

func (a *Account) Patch() error {
	// todo
	return nil
}

func (a *Account) Delete() error {
	db := a.DB()
	defer db.Commit(context.Background())

	a1 := a.CopyIdentifiers()
	if err := a1.Get(); err != nil {
		return err
	} else if a.Password != a1.Password {
		return ErrorAccountPasswordMatch
	}

	if _, err := db.Exec(context.Background(), `delete from accounts where id=$1 and password_hash=$2`, a.ID, a.Password); err != nil {
		a.Logger.Warnw("Failed to delete account from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return err
	}

	// TODO: User.Delete() should be called here, once implemented
	//if _, err := db.Exec(context.Background(), `delete from users where id=$1`, a.ID.String()); err != nil {
	//	logger.Warnw("Failed to delete account from DB", zap.Error(err))
	//	db.Rollback(context.Background())
	//	return errDelete
	//}

	return nil
}

func (a *Account) CopyIdentifiers() *Account {
	return &Account{Context: a.Context, Unique: Unique{ID: a.ID}, Email: a.Email, Username: a.Username}
}
