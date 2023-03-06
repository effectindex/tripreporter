package models

import "time"

type Medication struct {
	Name       string                `json:"name"`                // Required
	Dosage     int64                 `json:"dosage,omitempty"`    // Optional
	DosageUnit DosageUnit            `json:"dosage_unit"`         // Required, use DosageUnknown if unset
	Frequency  time.Duration         `json:"frequency,omitempty"` // Optional, uses 0 if unset
	RoA        RouteOfAdministration `json:"roa"`                 // Required, use RoaOther if unset
	Prescribed bool                  `json:"prescribed"`          // Optional, default true
}
