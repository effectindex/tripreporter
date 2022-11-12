package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type RouteOfAdministration int64

const ( // TODO: check if complete
	RoaOther RouteOfAdministration = iota
	RoaOral
	RoaBuccal
	RoaRectal
	RoaInhaled
	RoaSublabial
	RoaIntranasal
	RoaSublingual
	RoaOtherInjection
	RoaBuccalInjection
	RoaIntravenousInjection
	RoaSubcutanousInjection
	RoaIntramuscularInjection
)

type DosageUnit int64

const (
	DosageUnknown DosageUnit = iota
	DosageMicrograms
	DosageMilligrams
	DosageGrams
)

type DisplayUnit int64

const (
	UnitUnknown DisplayUnit = iota
	UnitMetric
	UnitImperial
)

type Account struct { // todo: this should be oauth / credentials. allow changing email or logging in with google // todo: restructure to model
	Context
	Unique
	Type     string `json:"type"`
	Email    string `json:"email" db:"email"`                   // Optional. Make clear that password reset isn't possible if not set.
	Username string `json:"username" db:"username"`             // Required. Generate from wordlist + 3 numbers if left blank.
	Password string `json:"password_hash" db:"password_hash"`   // Required
	Verified bool   `json:"email_verified" db:"email_verified"` // Optional. Whether email has been verified or not.
}

type User struct { // todo: old name was Profile // todo: visible on public profile:
	Context
	Type        string          `json:"type"`
	Created     time.Time       `json:"created,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`  // Optional
	Birth       time.Time       `json:"date_of_birth,omitempty"` // Optional, date in UTC+0, use Age if unset
	Age         int64           `json:"age,omitempty"`           // Optional, updated by Birth and unfavored if Age set
	Height      decimal.Decimal `json:"height,omitempty"`        // Optional // TODO: Encryption?
	Weight      decimal.Decimal `json:"weight,omitempty"`        // Optional // TODO: Encryption?
	Medication  UserMedication  `json:"medication,omitempty"`    // User's saved medication // TODO: Add to schema
	Preferences UserPreferences `json:"preferences,omitempty"`   // User's preferences // TODO: Add to schema
}

type UserMedication struct { // TODO: How do we feel about this being unencrypted?
	Medications []Medication `json:"medications,omitempty"`
}

type Medication struct {
	Name       string                `json:"name,omitempty"`        // Required
	Dosage     int64                 `json:"dosage,omitempty"`      // Optional
	DosageUnit DosageUnit            `json:"dosage_unit,omitempty"` // Required, use DosageUnknown if unset
	Frequency  *time.Duration        `json:"frequency,omitempty"`   // nil == unknown / don't display
	RoA        RouteOfAdministration `json:"roa,omitempty"`         // Required, use RoaOther if unset
	Prescribed bool                  `json:"prescribed,omitempty"`  // Optional, default true
}

type UserPreferences struct { // TODO: How do we feel about this being unencrypted?
	Timezone     time.Location `json:"timezone,omitempty"`      // Default: Europe/London
	HeightFormat DisplayUnit   `json:"height_format,omitempty"` // Default: UnitMetric // Display height in centimeters or feet + inches
	WeightFormat DisplayUnit   `json:"weight_format,omitempty"` // Default: UnitMetric // Display weight in kilograms or pounds
}
