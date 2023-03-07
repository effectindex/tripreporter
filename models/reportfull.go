package models

import (
	"context"
	"go.uber.org/zap"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type ReportFull struct {
	types.Context
	Unique
	Created      Timestamp `json:"creation_time" db:"creation_time"`
	LastModified Timestamp `json:"modified_time" db:"modified_time"`
	Date         Timestamp `json:"report_date" db:"report_date"`
	//Source       URL       `json:"source_url" db:"source_url"` // TODO
	//Effects      []Effect  // TODO
	//Submitter    Submitter // TODO
	Title   string       `json:"title" db:"title"`               // Required.
	Setting string       `json:"setting,omitempty" db:"setting"` // Optional.
	Events  ReportEvents `json:"events,omitempty"`               // Saved in report_events table and appended manually.
}

func (r *ReportFull) Get() (*ReportFull, error) {
	r.InitType(r)
	db := r.DB()
	defer db.Commit(context.Background())

	if r.NilUUID() {
		return r, types.ErrorReportNotSpecified
	}

	var r1 []*ReportFull
	if err := pgxscan.Select(context.Background(), db, &r1,
		`select * from reports where id=$1`, r.ID,
	); err != nil {
		r.Logger.Warnw("Failed to get report from DB", zap.Error(err))
		return r, err
	} else if len(r1) == 0 {
		return r, types.ErrorReportNotFound
	} else if len(r1) > 1 { // This shouldn't happen
		r.Logger.Errorw("Multiple reports found for parameters", "reports", r1)
		return r, types.ErrorReportNotSpecified
	}

	var r2 ReportEvents
	if err := pgxscan.Select(context.Background(), db, &r2,
		`select * from report_events where report_id=$1`, r.ID,
	); err != nil {
		r.Logger.Warnw("Failed to get report_events from DB", zap.Error(err))
		return r, err
	}

	for n, i := range r2 {
		if i.Type == ReportEventDrug && !i.Drug.NilUUID() {
			var d1 []*Drug
			if err := pgxscan.Select(context.Background(), db, &d1,
				`select * from drugs where id=$1`, r.ID,
			); err != nil {
				r.Logger.Warnw("Failed to get drug from DB", zap.Error(err))
				return r, err // only return if we error here, as this one matters
			} else if len(d1) == 0 {
				r.Logger.Warnw("No drugs found for parameters", "report event", i)
			} else if len(d1) > 1 { // This shouldn't happen
				r.Logger.Warnw("Multiple drugs found for parameters", "drugs", d1)
			}

			r2[n].Drug = *d1[0]
		}
	}

	r2.Sort()
	r.FromData(r1[0])
	r.Events = r2

	return r, nil
}

func (r *ReportFull) FromData(r1 *ReportFull) {
	r.InitType(r)
	r.ID = r1.ID
	r.Created = r1.Created
	r.LastModified = r1.LastModified
	r.Date = r1.Date
	r.Title = r1.Title
	r.Setting = r1.Setting
	r.Events = r1.Events
}
