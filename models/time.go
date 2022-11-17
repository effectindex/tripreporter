package models

import (
	"github.com/jackc/pgtype"
	"go.uber.org/zap"
	"time"
)

var (
	InternetZone = time.FixedZone("InternetZone", 0)
	TimeZero     time.Time
)

type Time struct {
	pgtype.Timestamptz
}

func SetupTime(ctx Context) {
	ctx.Validate()

	if zero, err := time.ParseInLocation(time.RFC3339, "1970-01-01T00:00:00Z", InternetZone); err != nil {
		ctx.Logger.Panicw("Failed to SetupTime", zap.Error(err))
	} else {
		TimeZero = zero
	}
}

// Set returns if the time is set
func (t *Time) Set() bool {
	return t.Time.Unix() != TimeZero.Unix() && !t.Time.IsZero()
}

// New creates a blank "unset" time
func (t *Time) New() {
	t.Time = TimeZero
}

// Now creates the current time
func (t *Time) Now() {
	t.Time = time.Now()
}

// Parse creates a time from a string, and ensures it is not equal to TimeZero
func (t *Time) Parse(s string) {
	panic(ErrorNotImplemented)
}

// String formats models.Time in a format usable by the Postgres Timestamptz type
func (t *Time) String() string {
	return t.Time.Format("2006-01-02 15:04:05Z07:00")
}
