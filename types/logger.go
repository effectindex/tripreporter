// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package types

import (
	"fmt"
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type StdLogWrapper struct {
	Prefix string
	Level  zapcore.Level
	Logger *zap.Logger
}

func (w *StdLogWrapper) Printf(format string, v ...any) {
	defer w.Logger.Sync()
	message := getMessage(format, v)
	w.Logger.Log(w.Level, w.Prefix+message)
}

// CreateZapWriterLogger creates a new zap.Logger that will write to an io.Writer with a provided encoder.
// You must defer logger.Sync() yourself.
func CreateZapWriterLogger(w io.Writer, c zapcore.EncoderConfig, e func(c zapcore.EncoderConfig) zapcore.Encoder) *zap.Logger {
	config := c
	config.TimeKey = "time"
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// Display time in UTC
		t = t.In(time.UTC)

		type appendTimeEncoder interface {
			AppendTimeLayout(time.Time, string)
		}

		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, time.RFC3339)
			return
		}

		enc.AppendString(t.Format(time.RFC3339))
	}

	encoder := e(config)
	writer := zapcore.AddSync(w)
	core := zapcore.NewCore(encoder, writer, zap.DebugLevel)
	logger := zap.New(core)
	return logger
}

// getMessage format with Sprint, Sprintf, or neither.
func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}
