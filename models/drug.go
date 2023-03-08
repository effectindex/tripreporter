package models

import "time"

type Drug struct {
	Unique
	Account    Unique                `db:"account_id"`                                // References account that created this drug.
	Name       string                `json:"name" db:"drug_name"`                     // Required
	Dosage     int64                 `json:"dosage,omitempty" db:"drug_dosage"`       // Optional
	DosageUnit DosageUnit            `json:"dosage_unit" db:"drug_dosage_unit"`       // Required, use DosageUnknown if unset
	RoA        RouteOfAdministration `json:"roa" db:"drug_roa"`                       // Required, use RoaOther if unset
	Frequency  time.Duration         `json:"frequency,omitempty" db:"drug_frequency"` // Optional, uses 0 if unset
	Prescribed DrugPrescribed        `json:"prescribed" db:"drug_prescribed"`         // Optional, default false
}
