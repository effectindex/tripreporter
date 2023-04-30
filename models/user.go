// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"
	"strconv"
	"strings"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type User struct {
	types.Context
	Unique
	Account     uuid.UUID       `json:"account_id" db:"account_id"`       // References the account that created this user.
	Username    string          `json:"username" db:"username"`           // References Account.Username.
	DisplayName string          `json:"display_name" db:"display_name"`   // References Account.DisplayName.
	Created     Timestamp       `json:"created" db:"created"`             // Required, set by default.
	Birth       Timestamp       `json:"date_of_birth" db:"date_of_birth"` // Optional, use Age if unset. // TODO: Impl encryption
	Age         Age             `json:"age"`                              // Unused. Updated by Birth and unfavored if Age set. // TODO: Impl encryption
	Height      Decimal         `json:"height" db:"height"`               // Unused // TODO: Impl encryption
	Weight      Decimal         `json:"weight" db:"weight"`               // Unused // TODO: Impl encryption
	Reports     []ReportSummary `json:"reports,omitempty"`                // References all Reports that a user has created.
}

type UserPublic struct {
	types.Context
	Unique
	Created     Timestamp       `json:"created"`
	Username    string          `json:"username"`
	DisplayName string          `json:"display_name"`
	Reports     []ReportSummary `json:"reports,omitempty"`
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
		`SELECT users.*, accounts.username, accounts.display_name FROM users LEFT JOIN accounts ON users.account_id = accounts.id WHERE account_id = $1;`, u.ID,
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
		u.Username = u1[0].Username
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

func (u *User) GetWithReports() (*User, error) {
	u, err := u.Get()
	if err != nil {
		return u, err
	}

	db := u.DB()
	defer db.Commit(context.Background())

	u.Reports = make([]ReportSummary, 0)

	var r1 []*ReportSummary
	if err := pgxscan.Select(context.Background(), db, &r1,
		`SELECT id, account_id, title, report_date FROM reports WHERE account_id=$1`, u.ID,
	); err != nil {
		u.Logger.Warnw("Failed to get reports from DB", zap.Error(err))
		return u, err
	} else if len(r1) == 0 { // We're allowed to have no reports
		return u, nil
	}

	// Collect all reports this user has made, and get all the drugs for each report
	for _, report := range r1 {
		if report == nil {
			continue
		}

		report.Drugs = make(map[string]Drug, 0)

		// Collect events
		var r2 ReportEvents
		if err := pgxscan.Select(context.Background(), db, &r2,
			`SELECT * FROM report_events WHERE report_id=$1`, report.ID,
		); err != nil {
			u.Logger.Warnw("Failed to get report_events from DB", zap.Error(err))
			return u, err
		}

		// Collect event drugs
		for n, i := range r2 {
			if i.Type == ReportEventDrug && i.DrugID != uuid.Nil {
				if drug, err := (&Drug{Context: u.Context, Unique: Unique{ID: i.DrugID}}).Get(); err != nil {
					return u, err
				} else {
					// Add drug to report summary if not already in the list
					if r2[n].Type == ReportEventDrug {
						drugName := drug.Name
						if _, ok := report.Drugs[drugName]; !ok {
							report.Drugs[drugName] = *drug
						}
					}
				}
			}
		}

		u.Reports = append(u.Reports, *report)
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
		`INSERT INTO users(account_id, created, date_of_birth, height, weight) VALUES($1, $2, $3, $4, $5);`,
		u.ID, u.Created.String(), u.Birth.String(), u.Height.String(), u.Weight.String(), // TODO: Medication / preferences in DB?
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
	query := "UPDATE USERS SET"
	qNum := 0

	addQuery := func(s string, i interface{}) {
		qNum++
		query += " " + s + "=$" + strconv.Itoa(qNum) + ","
		fields = append(fields, i)
	}

	if u.Created.Valid() {
		addQuery("created", u.Created)
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

	// TODO: u.Username and u.DisplayName support
	// TODO: Impl medication
	// TODO: Impl preferences

	query = strings.TrimSuffix(query, ",")
	qNum++
	query += " WHERE id=$;" + strconv.Itoa(qNum)
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
	if _, err := db.Exec(context.Background(), `DELETE FROM users WHERE account_id=$1;`, u.ID); err != nil {
		u.Logger.Warnw("Failed to delete user from DB", zap.Error(err))
		_ = db.Rollback(context.Background())
		return u, err
	}

	return nil, nil
}

func (u *User) CopyPublic() *UserPublic {
	p := &UserPublic{Context: u.Context, Unique: u.Unique, Created: u.Created, Username: u.Username, DisplayName: u.DisplayName, Reports: u.Reports}
	u.InitType(u)
	return p
}
