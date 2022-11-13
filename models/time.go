package models

import (
	"go.uber.org/zap"
	"time"
)

var (
	InternetZone = time.FixedZone("InternetZone", 0)
	TimeZero     time.Time
)

type Time struct {
	time.Time
}

func SetupTime(ctx Context) {
	ctx.Validate()

	if zero, err := time.ParseInLocation(time.RFC3339, "0001-01-01T00:00:00Z", InternetZone); err != nil {
		ctx.Logger.Panicw("Failed to SetupTime", zap.Error(err))
	} else {
		TimeZero = zero
	}
}

// Set returns if the time is set
func (t *Time) Set() bool {
	return t.Unix() != TimeZero.Unix()
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
