package models

import (
	"context"
	"strconv"
	"strings"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

type User struct { // todo: old name was Profile // todo: visible on public profile:
	types.Context
	Unique
	Created     Timestamp       `json:"created" db:"created"`             // Required, set by default.
	DisplayName string          `json:"display_name" db:"display_name"`   // Optional
	Birth       Timestamp       `json:"date_of_birth" db:"date_of_birth"` // Optional, use Age if unset
	Age         Age             `json:"age"`                              // Optional, updated by Birth and unfavored if Age set
	Height      Decimal         `json:"height" db:"height"`               // Optional // TODO: Encryption?
	Weight      Decimal         `json:"weight" db:"weight"`               // Optional // TODO: Encryption?
	Medication  UserMedication  `json:"medication"`                       // User's saved medication // TODO: Add to schema
	Preferences UserPreferences `json:"preferences"`                      // User's preferences // TODO: Add to schema
}

func (u *User) Get() (*User, error) {
	u.InitType(u)
	db := u.DB()
	defer db.Commit(context.Background())

	if u.NilUUID() {
		return u, types.ErrorUserNotSpecified
	}

	var u1 []*User
	if err := pgxscan.Select(context.Background(), db, &u1,
		`select created, display_name, date_of_birth, height, weight from users where account_id = $1;`, u.ID,
	); err != nil {
		u.Logger.Warnw("Failed to get user from DB", zap.Error(err))
		return u, err
	} else if len(u1) == 0 {
		return u, types.ErrorUserNotFound
	} else if len(u1) > 1 { // This shouldn't happen
		u.Logger.Errorw("Multiple users found for parameters", "users", u1)
		return u, types.ErrorUserNotSpecified
	} else {
		u.Created = u1[0].Created
		u.DisplayName = u1[0].DisplayName
		u.Birth = u1[0].Birth
		u.Height = u1[0].Height
		u.Weight = u1[0].Weight

		if u.Birth.Valid() {
			u.Age.Update(u.Birth)
		}
	}

	return u, nil
}

func (u *User) Post() (*User, error) {
	u.InitType(u)
	db := u.DB()
	defer db.Commit(context.Background())

	if u.NilUUID() {
		return u, types.ErrorUserNotSpecified
	}

	if !u.Created.Valid() {
		u.Created.Now()
	}

	if u.Birth.Valid() {
		u.Age.Update(u.Birth)
	}

	if _, err := db.Exec(context.Background(),
		`insert into users(account_id, created, display_name, date_of_birth, height, weight)
		values($1, $2, $3, $4, $5, $6);`,
		u.ID, u.Created.String(), u.DisplayName, u.Birth.String(), u.Height.String(), u.Weight.String(), // TODO: Medication / preferences in DB?
	); err != nil {
		u.Logger.Warnw("Failed to write account to DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return u, err
	}

	return u, nil
}

func (u *User) Patch() (*User, error) {
	u.InitType(u)
	db := u.DB()

	if u.NilUUID() {
		return u, types.ErrorUserNotSpecified
	}

	fields := make([]interface{}, 0)
	query := "update users set"
	qNum := 0

	addQuery := func(s string, i interface{}) {
		qNum++
		query += " " + s + "=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, i)
	}

	if u.Created.Valid() {
		addQuery("created", u.Created)
	}

	if u.DisplayName != "" { // TODO: Validate display names
		addQuery("display_name", u.DisplayName)
	}

	if u.Birth.Valid() { // TODO: Validate DOB
		u.Age.Update(u.Birth)
		addQuery("date_of_birth", u.Birth)
		addQuery("age", u.Age)
	}

	if u.Age.Valid() && !u.Birth.Valid() {
		return u, types.ErrorUserBirthNotSpecified
	}

	if u.Height.Valid() {
		addQuery("height", u.Height)
	}

	if u.Weight.Valid() {
		addQuery("height", u.Height)
	}

	// TODO: Impl medication
	// TODO: Impl preferences

	query = strings.TrimSuffix(query, ",")
	qNum++
	query += " where id=$;" + strconv.Itoa(qNum)
	fields = append(fields, u.ID)

	_, err := db.Exec(context.Background(), query, fields...)

	if err != nil {
		u.Logger.Warnw("Failed to update user in DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return u, err
	}

	// Want to commit before Get()
	db.Commit(context.Background())
	return u.Get()
}

func (u *User) Delete() (*User, error) {
	u.InitType(u)
	db := u.DB()
	defer db.Commit(context.Background())

	if _, err := (&Account{Context: u.Context, Unique: u.Unique}).Get(); err == nil {
		return u, types.ErrorUserAccountStillExists
	} else if err != types.ErrorAccountNotFound {
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
