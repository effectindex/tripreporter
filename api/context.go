package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ctx *Context
)

type Context struct {
	types.Context
}

// SetupContext creates a new context for this package, derived from the given context
func SetupContext(c types.Context) {
	c.Validate()
	ctx = &Context{Context: c}
}

// Handle is a wrapper around HandleStatus for pre-written messages
func (c *Context) Handle(w http.ResponseWriter, r *http.Request, m Message) {
	msg, status := m.Message()
	c.HandleStatus(w, r, msg, status)
}

// HandlePrefixed is a wrapper around HandleStatus for pre-written messages with a prefix attached
func (c *Context) HandlePrefixed(w http.ResponseWriter, r *http.Request, prefix string, m Message) {
	msg, status := m.Message()
	c.HandleStatus(w, r, prefix+msg, status)
}

// HandleStatus will write a JSON response to the request, and to our regular ctx.Logger, from the message and status given
func (c *Context) HandleStatus(w http.ResponseWriter, r *http.Request, msg string, status int) {
	logger := CreateLogger(w)
	defer logger.Sync()

	setJsonStatus(w, status)

	// Here we log messages and errors, depending on the severity of the status
	if status >= 500 {
		// Log errors that our fault as warnings, and tell the client we had an error.
		c.Logger.Warnw("API Internal Error", "status", status, "path", r.URL.Path, "message", msg)
		logger.Errorw(msg, "status", status)
	} else {
		// If the message isn't an error on our end, only log in debug
		c.Logger.Debugw("API Response", "status", status, "path", r.URL.Path, "message", msg)

		// If the message is a client error, warn them, otherwise it's an info
		if status >= 400 {
			logger.Warnw(msg, "status", status)
		} else {
			logger.Infow(msg, "status", status)
		}
	}
}

// HandleJson will write a JSON response to the request, with the contents of i, or an error if the marshal failed
func (c *Context) HandleJson(w http.ResponseWriter, r *http.Request, i interface{}, status int) {
	logger := CreateLogger(w)
	defer logger.Sync()

	setJsonStatus(w, status)

	if j, err := json.Marshal(i); err != nil {
		status = http.StatusInternalServerError

		c.Logger.Warnw("API Internal Error", "status", status, "path", r.URL.Path, "i", i, zap.Error(err))
		logger.Errorw(err.Error(), "status", status)
	} else {
		c.Logger.Debugw("API Response", "status", status, "path", r.URL.Path, "json", string(j))
		_, _ = fmt.Fprintf(w, "%s\n", j)
	}
}

// HandleMessage is a wrapper to create a simple http handler that responds with the given message
func (c *Context) HandleMessage(m Message) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.Handle(w, r, m)
	}
}

// HandleFunc is a wrapper to wrap around an arbitrary function
func (c *Context) HandleFunc(fn func(http.ResponseWriter, *http.Request), handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
		handler.ServeHTTP(w, r)
	}
}

// CreateLogger will create a new Zap logger from an http.ResponseWriter, to log to an http request directly.
// You must defer logger.Sync() yourself.
func CreateLogger(w http.ResponseWriter) *zap.SugaredLogger {
	logger := util.CreateZapWriterLogger(
		w, zap.NewProductionEncoderConfig(),
		func(c zapcore.EncoderConfig) zapcore.Encoder {
			return zapcore.NewJSONEncoder(c)
		},
	)

	return logger.Sugar()
}

func setJsonStatus(w http.ResponseWriter, status int) {
	// Set content type and status code properly
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status) // Set after writing header, as this closes the stream
}
