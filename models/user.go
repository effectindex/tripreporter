package models

import (
	"github.com/shopspring/decimal"
)

type User struct { // todo: old name was Profile // todo: visible on public profile:
	Context
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
	return ErrorNotImplemented
}

func (a *User) Patch() error {
	return ErrorNotImplemented
}

func (a *User) Delete() error {
	return ErrorNotImplemented
}

func (a *User) CopyIdentifiers() error {
	return ErrorNotImplemented
}
