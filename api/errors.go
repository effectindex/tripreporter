package api

import (
	"net/http"
	"time"

	"github.com/effectindex/tripreporter/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ctx *ErrorHandler
)

type ErrorHandler struct {
	types.Context
}

func SetupContext(c types.Context) {
	ctx = &ErrorHandler{Context: c}
}

func (h *ErrorHandler) Handle(w http.ResponseWriter, r *http.Request, e types.ErrorApi) {
	h.Validate()

	logger := CreateLogger(w)
	defer logger.Sync()

	msg, status := e.ErrorHttp()

	// Here we log messages and errors, depending on the severity of the status
	if status >= 500 {
		// Log errors that our fault as warnings, and tell the client we had an error.
		h.Logger.Warnw("API Internal Error", "status", status, "path", r.URL.Path, "message", msg)
		logger.Errorw(msg, "status", status)
	} else {
		// If the message isn't an error on our end, only log in debug
		h.Logger.Debugw("API Response", "status", status, "path", r.URL.Path, "message", msg)

		// If the message is a client error, warn them, otherwise it's an info
		if status >= 400 {
			logger.Warnw(msg, "status", status)
		} else {
			logger.Infow(msg, "status", status)
		}
	}
}

func CreateLogger(w http.ResponseWriter) *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "time"
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		// Display time in UTC
		t = t.In(time.UTC)

		type appendTimeEncoder interface {
			AppendTimeLayout(time.Time, string)
		}

		if enc, ok := enc.(appendTimeEncoder); ok {
			enc.AppendTimeLayout(t, time.RFC3339Nano)
			return
		}

		enc.AppendString(t.Format(time.RFC3339Nano))
	}

	encoder := zapcore.NewJSONEncoder(config)
	writer := zapcore.AddSync(w)
	core := zapcore.NewCore(encoder, writer, zap.DebugLevel)
	logger := zap.New(core)
	return logger.Sugar()
}
