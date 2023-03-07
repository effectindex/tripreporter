package models

type ReportEvent struct {
	Unique
	ReportEventType    `json:"type" db:"type"`
	Timestamp          `json:"timestamp" db:"timestamp"`
	ReportEventSection `json:"section" db:"section"`
	Content            string `json:"content" db:"content"`
}

type ReportEventType int64

const (
	ReportEventUnknown   = iota
	ReportEventNote      = 1
	ReportEventSubstance = 2
)

type ReportEventSection int64

const (
	ReportEventSectionUnknown = iota
	ReportEventSectionDescription
	ReportEventSectionOnset
	ReportEventSectionPeak
	ReportEventSectionOffset
)
