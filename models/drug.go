// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	doseRegex = regexp.MustCompile(`([0-9]+)?(\s*)(.*)`)
)

type Drug struct {
	Unique
	Account    uuid.UUID             `json:"account_id" db:"account_id"`        // References account that created this drug.
	Name       string                `json:"name" db:"drug_name"`               // Required
	Dosage     int64                 `json:"dosage" db:"drug_dosage"`           // Optional
	DosageUnit string                `json:"dosage_unit" db:"drug_dosage_unit"` // Optional, uses "unknown" if unset
	RoA        RouteOfAdministration `json:"roa" db:"drug_roa"`                 // Required, uses RoaUnknown if unset
	Frequency  time.Duration         `json:"frequency" db:"drug_frequency"`     // Optional, uses 0 if unset
	Prescribed DrugPrescribed        `json:"prescribed" db:"drug_prescribed"`   // Optional, uses DrugPrescribedUnknown if unset
}

func (d *Drug) ParseDose(dose string) {
	parts := doseRegex.FindStringSubmatch(dose)
	if len(parts) == 1 {
		return
	}

	dosage := 0
	if len(parts) > 3 {
		d, err := strconv.Atoi(parts[1])
		if err == nil {
			dosage = d
		}
	}

	d.Dosage = int64(dosage)
	d.DosageUnit = parts[3]

	return
}
