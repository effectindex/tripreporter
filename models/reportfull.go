package models

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/effectindex/tripreporter/types"
	"github.com/georgysavva/scany/v2/pgxscan"
	"go.uber.org/zap"
)

type ReportFull struct {
	types.Context
	Unique
	Account      Unique    `json:"account_id" db:"account_id"` // References the account that created this report.
	Created      Timestamp `json:"creation_time" db:"creation_time"`
	LastModified Timestamp `json:"modified_time" db:"modified_time"`
	Date         Timestamp `json:"report_date" db:"report_date"`
	//Source       URL       `json:"source_url" db:"source_url"` // TODO
	//Effects      []Effect  // TODO
	//Submitter    Submitter // TODO
	Title   string       `json:"title" db:"title"`               // Required.
	Setting string       `json:"setting,omitempty" db:"setting"` // Optional.
	Events  ReportEvents `json:"report_events,omitempty"`        // Saved in report_events table and appended manually.
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

func (r *ReportFull) FromBody(r1 *http.Request) (*ReportFull, error) {
	r.InitType(r)

	type ReportFormEvent struct {
		Timestamp  string `json:"timestamp,omitempty"`
		IsDrug     bool   `json:"is_drug,omitempty"`
		Section    int64  `json:"section,string,omitempty"`
		Content    string `json:"content,omitempty"`
		DrugName   string `json:"drug_name,omitempty"`
		DrugDosage string `json:"drug_dosage,omitempty"`
		RoA        int64  `json:"roa,string,omitempty"`
		Prescribed int64  `json:"prescribed,string,omitempty"`
	}

	type ReportForm struct {
		Title          string            `json:"title"`
		Setting        string            `json:"setting,omitempty"`
		ReportDate     string            `json:"report_date"`
		ReportSections []ReportFormEvent `json:"report_sections,omitempty"`
	}

	// We need a report ID in order to parse the report sections.
	// Anything calling this method should init an ID manually, unless making a new report.
	if r.NilUUID() {
		err := r.InitUUID(r.Logger)
		if err != nil {
			return r, err
		}
	}

	rCtx := r1.Context()
	ctxValues, ok := rCtx.Value(ContextValuesKey).(*ContextValues)
	if !ok {
		return r, types.ErrorContextCastFailed
	}

	accountID := ctxValues.SessionClaims.Account.UUID

	if r1.Body == nil {
		return r, types.ErrorStringEmpty.PrefixedError("Request body")
	}

	defer r1.Body.Close()
	body, err := io.ReadAll(r1.Body)
	if err != nil {
		return r, err
	}

	if len(body) == 0 {
		return r, types.ErrorStringEmpty.PrefixedError("Request body")
	}

	var rf *ReportForm
	err = json.Unmarshal(body, &rf)
	if err != nil {
		return r, err
	}

	//
	// Now we should have all the data, we need to turn some types into Go types to make sense.

	// First lets fix the create and last modified timestamps, if they're not valid
	if !r.Created.Valid() {
		r.Created.Now()
	}

	if !r.LastModified.Valid() {
		r.LastModified = r.Created
	}

	// Now we want to trim empty sections from the array
	formSections := rf.ReportSections[:0]
	for _, s := range rf.ReportSections {
		sectionEmpty := false

		if s.IsDrug {
			if s.Section == 0 && len(s.DrugName) == 0 && len(s.DrugDosage) == 0 && s.RoA == 0 && s.Prescribed == 0 {
				sectionEmpty = true
			}
		} else {
			if s.Section == 0 && len(s.Content) == 0 {
				sectionEmpty = true
			}
		}

		// Keep non-empty sections
		if !sectionEmpty {
			formSections = append(formSections, s)
		}
	}

	// Only keep non-empty sections
	rf.ReportSections = formSections

	// Try to find if any of the timestamps have been set
	firstTimestamp := "T00:00:00Z"
	for _, s := range rf.ReportSections {
		if len(s.Timestamp) > 0 {
			firstTimestamp = "T" + s.Timestamp + ":00Z"
			break
		}
	}

	// Now lets parse the rf.ReportDate as an actual Timestamp
	date, err := r.Date.Parse(rf.ReportDate + firstTimestamp)
	if err != nil {
		return r, err
	}

	r.Date = *date

	// Now we can parse the sections properly
	sections := make(ReportEvents, 0)
	accountUnique := Unique{ID: accountID}
	accountUnique.InitType(&Account{})

	for n, s := range rf.ReportSections {
		event := &ReportEvent{
			Unique:  r.Unique,
			Index:   int64(n),
			Type:    ReportEventNote,
			Content: s.Content,
		}

		if s.IsDrug {
			event.Type = ReportEventDrug
			event.Drug = Drug{ // TODO: Frequency is not parsed here
				Account:    accountUnique,
				Name:       s.DrugName,
				RoA:        RouteOfAdministration(s.RoA),
				Prescribed: DrugPrescribed(s.Prescribed),
			}
			event.Drug.ParseDose(s.DrugDosage)
			event.Drug.InitType(event.Drug)

			if err != nil {
				return r, err
			}

			err = event.Drug.InitUUID(r.Logger)
			if err != nil {
				return r, err
			}
		}

		event.Section = ReportEventSection(s.Section)

		sections = append(sections, event)
	}

	// Set title, sections and events now that we have the sections parsed
	r.Title = rf.Title
	r.Setting = rf.Setting
	r.Events = sections

	// We should have a completely parsed ReportFull now
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
