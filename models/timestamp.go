package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/effectindex/tripreporter/types"
	"github.com/jackc/pgtype"
)

type Timestamp struct {
	pgtype.Timestamptz
}

// Valid returns if the time is valid
func (t *Timestamp) Valid() bool {
	if t == nil {
		return false
	}

	return !t.Time.IsZero() && t.Status != pgtype.Undefined
}

// Default will set Timestamp to it's default value
func (t *Timestamp) Default() {
	*t = Timestamp{pgtype.Timestamptz{Status: pgtype.Null}}
}

// Now set the value to the current time
func (t *Timestamp) Now() {
	*t = Timestamp{pgtype.Timestamptz{Time: time.Now(), Status: pgtype.Present}}
}

// Parse creates a Timestamp from a string
func (t *Timestamp) Parse(s string) (*Timestamp, error) {
	if t1, err := time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return nil, err
	} else {
		*t = Timestamp{pgtype.Timestamptz{Time: t1, Status: pgtype.Present}}
		return t, nil
	}
}

// ParseDate creates a Timestamp from a string and a 2006-01-02 formatted date.
func (t *Timestamp) ParseDate(s string) (*Timestamp, error) {
	return t.Parse(s + "T00:00:00Z")
}

// String formats models.Timestamp in a format usable by the Postgres Timestamptz type
func (t *Timestamp) String() string {
	return t.Time.Format(time.RFC3339)
}

// Scan implements the database/sql Scanner interface. This copies pgtype.Timestamptz and modifies it to scan as time.UTC
func (t *Timestamp) Scan(src interface{}) error {
	if src == nil {
		*t = Timestamp{pgtype.Timestamptz{Status: pgtype.Null}}
		return nil
	}

	var err error
	switch src := src.(type) {
	case string:
		err = t.DecodeText(nil, []byte(src))
	case []byte:
		srcCopy := make([]byte, len(src))
		copy(srcCopy, src)
		err = t.DecodeText(nil, srcCopy)
	case time.Time:
		// Scan into UTC and not current timezone
		*t = Timestamp{pgtype.Timestamptz{Time: src.In(time.UTC), Status: pgtype.Present}}
		return nil
	}

	if err != nil {
		return err
	}

	// Scan into UTC and not current timezone
	if t.Valid() {
		*t = Timestamp{pgtype.Timestamptz{Time: t.Time.In(time.UTC), Status: t.Status, InfinityModifier: t.InfinityModifier}}
	}

	return fmt.Errorf("cannot scan %T", src)
}

// MarshalJSON implements the encoding/json Marshaler interface. This copies pgtype.Timestamptz and modifies it to treat pgtype.Undefined as pgtype.Null
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	switch t.Status {
	case pgtype.Null, pgtype.Undefined:
		return []byte("null"), nil
	}

	if t.Status != pgtype.Present {
		return nil, types.ErrorGenericUnknown
	}

	var s string

	switch t.InfinityModifier {
	case pgtype.None:
		s = t.Time.Format(time.RFC3339Nano)
	case pgtype.Infinity:
		s = "infinity"
	case pgtype.NegativeInfinity:
		s = "-infinity"
	}

	return json.Marshal(s)
}
