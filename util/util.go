package util

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
