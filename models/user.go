package models

import (
	"context"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type User struct { // todo: old name was Profile // todo: visible on public profile:
	Context
	Unique
	Type        string          `json:"type"`
	Created     Time            `json:"created" db:"created"`
	DisplayName string          `json:"display_name" db:"display_name"`   // Optional
	Birth       Time            `json:"date_of_birth" db:"date_of_birth"` // Optional, date in UTC+0, use Age if unset
	Age         int64           `json:"age" db:"age"`                     // Optional, updated by Birth and unfavored if Age set
	Height      decimal.Decimal `json:"height" db:"height"`               // Optional // TODO: Encryption?
	Weight      decimal.Decimal `json:"weight" db:"weight"`               // Optional // TODO: Encryption?
	Medication  UserMedication  `json:"medication"`                       // User's saved medication // TODO: Add to schema
	Preferences UserPreferences `json:"preferences"`                      // User's preferences // TODO: Add to schema
}

func (a *User) Get() error {
	return ErrorNotImplemented
}

func (a *User) Post() error {
	db := a.DB()
	defer db.Commit(context.Background())

	if a.NilUUID() {
		return ErrorUserNotSpecified
	}

	if !a.Created.Set() {
		a.Created.New()
	}

	if _, err := db.Exec(context.Background(),
		`insert into users(account_id, created, display_name, date_of_birth, age, height, weight)
		values($1, $2, $3, $4, $5, $6, $7);`,
		a.ID, a.Created.String(), a.DisplayName, a.Birth.String(), a.Age, a.Height, a.Weight, // TODO: Medication / preferences in DB?
	); err != nil {
		a.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return err
	}

	return nil
}

func (a *User) Patch() error {
	return ErrorNotImplemented
}

func (u *User) Delete() (*User, error) {
	db := u.DB()
	defer db.Commit(context.Background())

	if _, err := (&Account{Context: u.Context, Unique: u.Unique}).Get(); err == nil {
		return u, ErrorUserAccountStillExists
	} else if err != ErrorAccountNotFound {
		return u, err
	}

	// This should not be possible with a proper DB setup, this is only here for cleanup reasons
	// Normally, a user row will be deleted when an account row is deleted.
	if _, err := db.Exec(context.Background(), `delete from users where account_id=$1;`, u.ID); err != nil {
		u.Logger.Warnw("Failed to delete user from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return u, err
	}

	return nil, nil
}

func (a *User) CopyIdentifiers() error {
	return ErrorNotImplemented
}
