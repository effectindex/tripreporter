// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"time"
)

type Age struct {
	Timestamp
}

// Valid returns if the Age is set
func (a *Age) Valid() bool {
	return a != nil && a.Timestamp.Valid()
}

// Get returns the current Age based on current time
func (a *Age) Get() int {
	return a.GetAtTime(time.Now())
}

// GetAtTime will return the Age at a time.Time
func (a *Age) GetAtTime(t time.Time) int {
	if !a.Valid() {
		return -1
	}

	// Flatten day to
	birthday := a.Time
	ty, tm, td := t.Date()
	t = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)

	by, bm, bd := birthday.Date()
	birthday = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)

	// First birthday has not happened yet
	if t.Before(birthday) {
		return 0
	}

	// Get rough age based on given year - birth year
	age := ty - by
	anniversary := birthday.AddDate(age, 0, 0)

	// Correct for if birthday is before or after given date in the year
	if anniversary.After(t) {
		age--
	}

	return age
}

// Update will set the new date of birth for this age
func (a *Age) Update(birth Timestamp) {
	if birth.Valid() {
		*a = Age{birth}
	}
}

// Default will set Age to it's default value
func (a *Age) Default() {
	a.Timestamp.Default()
}

// String formats Age in a format usable by the Postgres Timestamptz type
func (a *Age) String() string {
	return a.Timestamp.String()
}

func (a *Age) Scan(src interface{}) error {
	return a.Timestamp.Scan(src)
}
