package models

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"

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

func (u *User) Get() (*User, error) {
	db := u.DB()
	defer db.Commit(context.Background())

	if u.NilUUID() {
		return u, ErrorUserNotSpecified
	}

	var u1 []*User
	if err := pgxscan.Select(context.Background(), db, &u1,
		`select created, display_name, date_of_birth, age, height, weight from users where account_id = $1;`, u.ID,
	); err != nil {
		u.Logger.Warnw("Failed to get user from DB", zap.Error(err))
		return u, err
	} else if len(u1) == 0 {
		return u, ErrorUserNotFound
	} else if len(u1) > 1 { // This shouldn't happen
		u.Logger.Errorw("Multiple users found for parameters", "users", u1)
		return u, ErrorUserNotSpecified
	} else {
		u.Created = u1[0].Created
		u.DisplayName = u1[0].DisplayName
		u.Birth = u1[0].Birth
		u.Age = u1[0].Age
		u.Height = u1[0].Height
		u.Weight = u1[0].Weight
	}

	return u, nil
}

func (u *User) Post() (*User, error) {
	db := u.DB()
	defer db.Commit(context.Background())

	if u.NilUUID() {
		return u, ErrorUserNotSpecified
	}

	if !u.Created.Set() {
		u.Created.New()
	}

	if _, err := db.Exec(context.Background(),
		`insert into users(account_id, created, display_name, date_of_birth, age, height, weight)
		values($1, $2, $3, $4, $5, $6, $7);`,
		u.ID, u.Created.String(), u.DisplayName, u.Birth.String(), u.Age, u.Height, u.Weight, // TODO: Medication / preferences in DB?
	); err != nil {
		u.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return u, err
	}

	return u, nil
}

func (u *User) Patch() (*User, error) {
	return u, ErrorNotImplemented
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
