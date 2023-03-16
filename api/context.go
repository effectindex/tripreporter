package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cristalhq/jwt/v4"
	"github.com/effectindex/tripreporter/models"
	"github.com/effectindex/tripreporter/types"
	"github.com/effectindex/tripreporter/util"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	ctx *Context
)

type Context struct {
	types.Context

	JwtKey     []byte
	JwtBuilder jwt.Builder
}

// SetupContext creates a new context for this package, derived from the given context
func SetupContext(c types.Context) {
	c.Validate()
	ctx = &Context{Context: c}
}

// SetupJwt will set up the JWT file, key and builder
func SetupJwt() {
	if err := godotenv.Load(".jwt.env"); err != nil {
		ctx.Logger.Warnw("err loading .jwt.env file, creating new one", zap.Error(err))
	}

	// Get key from .jwt.env, decode it
	jwtKey := make([]byte, 512/8) // 512-bit key
	key := os.Getenv("JWT_AUTH_KEY")
	decodedKey, err := base64.StdEncoding.DecodeString(key)

	// Check if decoded key is the right length, make a new one and write it if not
	if len(decodedKey) != 512/8 || err != nil {
		ctx.Logger.Warnw("JWT key in .jwt.env missing or not good, creating new one", zap.Error(err))

		jwtKey, err := util.GenerateRandomBytes(512 / 8)
		if err != nil {
			ctx.Logger.Fatalw("Could not make JWT key bytes", zap.Error(err))
		}

		encodedKey := base64.StdEncoding.EncodeToString(jwtKey)
		err = godotenv.Write(map[string]string{"JWT_AUTH_KEY": encodedKey}, ".jwt.env")
		if err != nil {
			ctx.Logger.Fatalw("Failed to write new .jwt.env", zap.Error(err))
		}
	} else {
		jwtKey = decodedKey
	}

	ctx.JwtKey = jwtKey
	ctx.Logger.Infow("Found valid JWT key")

	signer, err := jwt.NewSignerHS(jwt.HS512, ctx.JwtKey)
	if err != nil {
		ctx.Logger.Fatalw("Cannot create JWT signer", zap.Error(err))
	}

	ctx.JwtBuilder = *jwt.NewBuilder(signer)
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

// HandleRedirect is a wrapper to redirect to a destination and set the status code.
// TODO: This is not functional on frontend pages, see #90.
func (c *Context) HandleRedirect(w http.ResponseWriter, r *http.Request, url string, status int) {
	c.Logger.Debugw("API Redirect", "status", status, "path", r.URL.Path, "destination", url)
	http.Redirect(w, r, url, status)
}

// GetCtxValOrHandle will return the requests models.ContextValues and an ok if successful.
// It will return an http.StatusBadRequest if unsuccessful.
func (c *Context) GetCtxValOrHandle(w http.ResponseWriter, r *http.Request) (*models.ContextValues, bool) {
	rCtx := r.Context()
	ctxVal, ok := rCtx.Value(models.ContextValuesKey).(*models.ContextValues)
	if ok {
		return ctxVal, ok
	}

	ctx.HandleStatus(w, r, types.ErrorContextCastFailed.Error(), http.StatusBadRequest)
	return nil, false
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
