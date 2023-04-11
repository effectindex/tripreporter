// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"
	"regexp"
	"strconv"
	"time"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	doseRegex = regexp.MustCompile(`([0-9]+)?(\s*)(.*)`)
)

type Drug struct {
	types.Context
	Unique
	Account    uuid.UUID             `json:"account_id" db:"account_id"`        // References account that created this drug.
	Name       string                `json:"name" db:"drug_name"`               // Required
	Dosage     int64                 `json:"dosage" db:"drug_dosage"`           // Optional
	DosageUnit string                `json:"dosage_unit" db:"drug_dosage_unit"` // Optional, uses "unknown" if unset
	RoA        RouteOfAdministration `json:"roa" db:"drug_roa"`                 // Required, uses RoaUnknown if unset
	Frequency  time.Duration         `json:"frequency" db:"drug_frequency"`     // Optional, uses 0 if unset
	Prescribed DrugPrescribed        `json:"prescribed" db:"drug_prescribed"`   // Optional, uses DrugPrescribedUnknown if unset
}

// Get will get the Drug without committing! You MUST commit this yourself.
func (d *Drug) Get() (*Drug, error) {
	d.InitType(d)
	db := d.DB()

	if d.ID == uuid.Nil {
		return d, types.ErrorDrugNotFound
	}

	var d1 []*Drug
	if err := pgxscan.Select(context.Background(), db, &d1,
		`select * from drugs where id=$1`, d.ID,
	); err != nil {
		d.Logger.Warnw("Failed to get drug from DB", zap.Error(err))
		return d, err
	} else if len(d1) == 0 {
		return d, types.ErrorDrugNotFound
	} else if len(d1) > 1 {
		d.Logger.Warnw("Multiple drugs found for parameters", "drugs", d1)
		return d, types.ErrorDrugNotSpecified
	} else {
		d.FromData(d1[0])
	}

	return d, nil
}

// Post will post the Drug without committing! You MUST commit this yourself.
func (d *Drug) Post() (*Drug, error) {
	d.InitType(d)
	db := d.DB()

	if _, err := db.Exec(context.Background(),
		`insert into drugs(
						id,
						account_id,
						drug_name,
						drug_dosage,
						drug_dosage_unit,
						drug_roa,
						drug_frequency,
						drug_prescribed
					) values($1, $2, $3, $4, $5, $6, $7, $8);`,
		d.ID, d.Account, d.Name, d.Dosage, d.DosageUnit, d.RoA, d.Frequency, d.Prescribed,
	); err != nil {
		d.Logger.Warnw("Failed to write drug to DB", zap.Error(err))
		return d, err
	}

	return d, nil
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
}

func (d *Drug) FromData(d1 *Drug) {
	d.InitType(d)
	d.ID = d1.ID
	d.Account = d1.Account
	d.Name = d1.Name
	d.Dosage = d1.Dosage
	d.DosageUnit = d1.DosageUnit
	d.RoA = d1.RoA
	d.Frequency = d1.Frequency
	d.Prescribed = d1.Prescribed
}
