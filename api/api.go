package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/effectindex/tripreporter/ui"
	"github.com/effectindex/tripreporter/util"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	proxy *httputil.ReverseProxy // proxy is used for `make dev-ui`
	dev   = false                // will disable serving /static/ from cache and proxy un-handled requests
)

// Setup manages functions that should be ready to use before
func Setup(isDevelopment bool, logger *zap.SugaredLogger) {
	proxy = util.NewProxy("http://localhost:"+os.Getenv("DEV_PORT"), logger)
	dev = isDevelopment
}

// Handler will serve /api, and pass the rest off to Router.
// In production, Handler will also serve /static/.
func Handler() http.Handler {
	router := mux.NewRouter()

	// serve /static/ by cache in production (no hot-reload support)
	if !dev { // if running in development mode, let api.Router reverse proxy it
		staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
		httpFS := http.FileServer(http.FS(staticFS))
		router.PathPrefix("/static/").HandlerFunc(ctx.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx.Logger.Debugw("Serving static", "path", r.URL.Path)
		}, httpFS))
	}

	// Redirect /api with no trailing slash to the documentation url
	router.Handle("/api", http.RedirectHandler(os.Getenv("DOCS_URL"), http.StatusMovedPermanently))

	// API functions
	vX := router.PathPrefix("/api/").Subrouter()
	vX.MethodNotAllowedHandler = ctx.HandleMessage(MsgMethodNotAllowed)
	vX.NotFoundHandler = ctx.HandleMessage(MsgInvalidApiVersion)

	// API v1 methods
	v1 := vX.PathPrefix("/v1").Subrouter()
	v1.MethodNotAllowedHandler = ctx.HandleMessage(MsgMethodNotAllowed)
	v1.NotFoundHandler = ctx.HandleMessage(MsgInvalidEndpoint)

	// API v1 endpoints
	SetupAccountEndpoints(v1)
	SetupSessionEndpoints(v1)
	SetupUserEndpoints(v1)

	// Let api.Router do everything else, including serving /static/ in development
	router.PathPrefix("/").HandlerFunc(Router)

	return router
}

// Router will route /, /favicon.ico and anything not handled by Handler.
// In development, Router will also handle /static/.
func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}

	if dev { // serve everything else in dev to the proxy, if in dev
		proxy.ServeHTTP(w, r)
	} else { // else, serve just dist/index.html and let Vue's JS handle it
		rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
		w.Write(rawFile)
	}
}
