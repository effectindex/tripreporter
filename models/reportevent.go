// SPDX-FileCopyrightText: 2023 froggie <incoming@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

import (
	"sort"

	"github.com/google/uuid"
)

type ReportEvent struct {
	Report    uuid.UUID          `json:"report_id" db:"report_id"`          // References the original report ID
	Index     int64              `json:"index" db:"event_index"`            // Order of sections
	Timestamp Timestamp          `json:"timestamp" db:"event_timestamp"`    // Timestamp of event
	Type      ReportEventType    `json:"type" db:"event_type"`              // Type of event
	Section   ReportEventSection `json:"section" db:"event_section"`        // Section event is in
	Content   string             `json:"content" db:"event_content"`        // Content of event, if ReportEventNote
	DrugID    uuid.UUID          `json:"drug_id,omitempty" db:"event_drug"` // DrugID of event, corresponds to Drug, if ReportEventDrug
	Drug      Drug               `json:"drug,omitempty"`                    // Drug of event, if ReportEventDrug
}

type ReportEventType int64

const (
	ReportEventUnknown ReportEventType = iota
	ReportEventNote                    = 1
	ReportEventDrug                    = 2
)

type ReportEventSection int64

const (
	ReportEventSectionUnknown ReportEventSection = iota
	ReportEventSectionDescription
	ReportEventSectionOnset
	ReportEventSectionPeak
	ReportEventSectionOffset
)

type ReportEvents []*ReportEvent

func (r ReportEvents) Sort() {
	sort.Slice(r, func(i, j int) bool {
		return r[i].Index < r[j].Index
	})
}
