// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"context"
	"sort"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ReportSource struct {
	types.Context
	Report   uuid.UUID        `json:"report_id" db:"report_id"`         // References the original report ID
	Index    int64            `json:"index" db:"event_index"`           // Order of sources
	Author   bool             `json:"is_author" db:"source_is_author"`  // If the Report.Account is the author of this report
	Imported bool             `json:"imported" db:"source_is_imported"` // If the source was automatically Imported
	Name     string           `json:"name" db:"source_name"`            // Name given when selecting ReportSourceOther
	URL      string           `json:"url" db:"source_url"`              // URL of original source
	Type     ReportSourceType `json:"type" db:"source_type"`            // Type of report given by Report.Account
}

type ReportSourceType int64

const (
	ReportSourceUnknown = iota
	ReportSourceOther
	ReportSourceSelf
	ReportSourceErowid
	ReportSourceBluelight
	ReportSourcePsychonautWiki
	ReportSourceTripSit
	ReportSourceReddit
)

type ReportSources []*ReportSource

func (r ReportSources) Sort() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Index < r[j].Index
	})
}

func (r *ReportSource) Get() ([]*ReportSource, error) {
	db := r.DB()
	defer db.Commit(context.Background())

	if r.Report == uuid.Nil {
		return nil, types.ErrorReportNotSpecified
	}

	var r1 = ReportSources{}
	if err := pgxscan.Select(context.Background(), db, &r1,
		`select * from report_sources where report_id=$1`, r.Report,
	); err != nil {
		r.Logger.Warnw("Failed to get report_sources from DB", zap.Error(err))
		return r1, err
	} else if len(r1) == 0 { // Return if we don't have anything else to parse
		return r1, nil
	}

	r1.Sort()

	return r1, nil
}

// Post will post the ReportSource without finishing the tx! You MUST `db.Rollback` / `db.Commit` this yourself.
func (r *ReportSource) Post(db pgx.Tx) (*ReportSource, error) {
	if _, err := db.Exec(context.Background(),
		`insert into report_sources(
						report_id,
					    source_index,
						source_is_author,
						source_is_imported,
						source_name,
						source_url,
						source_type
					) values($1, $2, $3, $4, $5, $6, $7);`,
		r.Report, r.Index, r.Author, r.Imported, r.Name, r.URL, r.Type,
	); err != nil {
		r.Logger.Warnw("Failed to write report subject to DB", zap.Error(err))
		return r, err
	}

	return r, nil
}

func (r *ReportSource) FromData(r1 *ReportSource) {
	r.Report = r1.Report
	r.Author = r1.Author
	r.Name = r1.Name
	r.Imported = r1.Imported
	r.URL = r1.URL
	r.Type = r1.Type
}
