// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"github.com/effectindex/tripreporter/types"
	"github.com/google/uuid"
)

// ReportSummary is a clone of Report with fewer fields to be sent, when displaying a User's profile, with Drugs created automatically from Events.
type ReportSummary struct {
	types.Context
	Unique
	Account uuid.UUID       `json:"account_id" db:"account_id"`
	Title   string          `json:"title" db:"title"`
	Date    Timestamp       `json:"report_date" db:"report_date"`
	Drugs   map[string]Drug `json:"drugs"` // [Drug.Name]Drug, Appended manually from Events.
}
