// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ReportSubject struct {
	types.Context
	Report        uuid.UUID   `json:"report_id" db:"report_id"`                          // References the original report ID
	Age           int64       `json:"age" db:"subject_age"`                              // Age of the subject in years, optional // TODO: Implement dynamic age
	Gender        string      `json:"gender" db:"subject_gender"`                        // Gender of the subject, optional
	DisplayUnit   DisplayUnit `json:"display_unit" db:"subject_display_unit"`            // DisplayUnit of the HeightCm and WeightKg, default UnitMetric, optional
	HeightCm      Decimal     `json:"height_cm" db:"subject_height_cm"`                  // HeightCm of the subject, optional
	WeightKg      Decimal     `json:"weight_kg" db:"subject_weight_kg"`                  // WeightKg of the subject, optional
	MedicationIDs []uuid.UUID `json:"medication_ids,omitempty" db:"subject_medications"` // MedicationIDs of the subject, corresponds to Drug, optional
	Medications   []Drug      `json:"medications,omitempty"`                             // Medications of subject, optional
}

func (r *ReportSubject) Get() (*ReportSubject, error) {
	db := r.DB()
	defer db.Commit(context.Background())

	if r.Report == uuid.Nil {
		return r, types.ErrorReportNotSpecified
	}

	var r1 []*ReportSubject
	if err := pgxscan.Select(context.Background(), db, &r1,
		`select * from report_subjects where report_id=$1`, r.Report,
	); err != nil {
		r.Logger.Warnw("Failed to get report subject from DB", zap.Error(err))
		return r, err
	} else if len(r1) == 0 { // Return if we don't have anything else to parse
		return r, nil
	} else if len(r1) > 1 { // This shouldn't happen, but we still want to continue getting a report even if so
		r.Logger.Errorw("Multiple report subjects found for parameters", "report_subjects", r1)
		return r, nil
	}

	for _, m := range r1[0].MedicationIDs {
		if medication, err := (&Drug{Context: r.Context, Unique: Unique{ID: m}}).Get(); err != nil && m != uuid.Nil {
			return r, err
		} else {
			r1[0].Medications = append(r1[0].Medications, *medication)
		}
	}

	r.FromData(r1[0])

	return r, nil
}

// Post will post the ReportSubject without finishing the tx! You MUST `db.Rollback` / `db.Commit` this yourself.
func (r *ReportSubject) Post(db pgx.Tx) (*ReportSubject, error) {
	r.MedicationIDs = make([]uuid.UUID, len(r.Medications))
	for n, medication := range r.Medications {
		if _, err := medication.Post(db); err != nil {
			return r, err
		}

		r.MedicationIDs[n] = medication.ID
	}

	if _, err := db.Exec(context.Background(),
		`insert into report_subjects(
						report_id, 
						subject_age, 
						subject_gender, 
						subject_display_unit, 
						subject_height_cm, 
						subject_weight_kg, 
						subject_medications
					) values($1, $2, $3, $4, $5, $6, $7);`,
		r.Report, r.Age, r.Gender, r.DisplayUnit, r.HeightCm, r.WeightKg, r.MedicationIDs,
	); err != nil {
		r.Logger.Warnw("Failed to write report subject to DB", zap.Error(err))
		return r, err
	}

	return r, nil
}

func (r *ReportSubject) FromData(r1 *ReportSubject) {
	r.Report = r1.Report
	r.Age = r1.Age
	r.Gender = r1.Gender
	r.DisplayUnit = r1.DisplayUnit
	r.HeightCm = r1.HeightCm
	r.WeightKg = r1.WeightKg
	r.MedicationIDs = r1.MedicationIDs
	r.Medications = r1.Medications
}
