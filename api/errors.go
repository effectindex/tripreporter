package api

import (
	"net/http"

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

	if status >= 500 {
		h.Logger.Warnw("API Internal Error", "status", status, "path", r.URL.Path, "message", msg)
	} else {
		h.Logger.Debugw("API Response", "status", status, "path", r.URL.Path, "message", msg)
	}

	logger.Infow(msg, "status", status)
}

func CreateLogger(w http.ResponseWriter) *zap.SugaredLogger {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	writer := zapcore.AddSync(w)
	core := zapcore.NewCore(encoder, writer, zap.DebugLevel)
	logger := zap.New(core)
	return logger.Sugar()
}
