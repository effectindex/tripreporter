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
	Report   uuid.UUID        `json:"report_id" db:"report_id"`            // References the original report ID
	Index    int64            `json:"index" db:"event_index"`              // Order of sources
	IsAuthor bool             `json:"is_author" db:"source_is_author"`     // If the Report.Account is the author of this report
	Imported bool             `json:"is_imported" db:"source_is_imported"` // If the source was automatically Imported
	Author   string           `json:"author" db:"source_author"`           // Author of original report
	Profile  string           `json:"profile" db:"source_author_profile"`  // Profile of original author on original ReportSourceType
	URL      string           `json:"url" db:"source_url"`                 // URL of original source
	TypeName string           `json:"type_name" db:"source_type_name"`     // TypeName given when selecting ReportSourceOther
	Type     ReportSourceType `json:"type" db:"source_type"`               // Type of report given by Report.Account
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
						source_author,
						source_url,
						source_type_name,
						source_type
					) values($1, $2, $3, $4, $5, $6, $7, $8);`,
		r.Report, r.Index, r.IsAuthor, r.Imported, r.Author, r.URL, r.TypeName, r.Type,
	); err != nil {
		r.Logger.Warnw("Failed to write report subject to DB", zap.Error(err))
		return r, err
	}

	return r, nil
}

func (r *ReportSource) FromData(r1 *ReportSource) {
	r.Report = r1.Report
	r.Index = r1.Index
	r.IsAuthor = r1.IsAuthor
	r.Imported = r1.Imported
	r.Author = r1.Author
	r.URL = r1.URL
	r.TypeName = r1.TypeName
	r.Type = r1.Type
}
