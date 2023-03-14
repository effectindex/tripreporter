package api

import (
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/effectindex/tripreporter/ui"
	"github.com/effectindex/tripreporter/util"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

var (
	proxy *httputil.ReverseProxy // proxy is used for `make dev-ui`
	dev   = false                // will disable serving /static/ from cache and proxy un-handled requests

	staticFS, _ = fs.Sub(ui.StaticFiles, "dist")
	httpFS      = http.FileServer(http.FS(staticFS))
	staticIcons = []string{"android-chrome-192x192.png", "android-chrome-512x512.png", "apple-touch-icon.png", "favicon-16x16.png", "favicon-32x32.png"}
)

// Setup manages functions that should be ready to use before
func Setup(isDevelopment bool, logger *zap.SugaredLogger) {
	proxy = util.NewProxy("http://localhost:"+os.Getenv("DEV_PORT"), logger)
	dev = isDevelopment
}

// CorsWrapper will wrap h in a CORS handler
func CorsWrapper(h http.Handler, logger *zap.SugaredLogger) http.Handler {
	serveUrl := os.Getenv("VUE_APP_PROD_URL")
	if dev {
		addr := os.Getenv("SRV_ADDR")
		port := os.Getenv("SRV_PORT")
		if len(addr) == 0 {
			addr = "localhost"
		}

		//goland:noinspection HttpUrlsUsage
		serveUrl = fmt.Sprintf("http://%s:%v", addr, port)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{serveUrl},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		Debug:            dev,
	})

	if dev {
		c.Log = &util.StdLogWrapper{Prefix: "[cors] ", Level: zap.DebugLevel, Logger: logger.Desugar()}
	}

	logger.Debugw("Created CorsWrapper", "serveUrl", serveUrl)

	return c.Handler(h)
}

// Handler will handle /api, /static/, and pass the rest off to Router.
func Handler() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/static/").HandlerFunc(Static)

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
	SetupReportEndpoints(v1)
	SetupSessionEndpoints(v1)
	SetupUserEndpoints(v1)

	// Let api.Router do everything else
	router.PathPrefix("/").HandlerFunc(Router)

	return router
}

// Router will route /, /favicon.ico, normal pages, and anything not handled by Handler.
func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if r.URL.Path == "/favicon.ico" || r.URL.Path == "/site.webmanifest" ||
		(strings.HasSuffix(r.URL.Path, ".png") && slices.Contains(staticIcons, r.URL.Path)) {
		rawFile, _ := ui.StaticFiles.ReadFile("dist" + r.URL.Path)
		w.Write(rawFile)
		return
	}

	// Proxy everything else in dev. The static file below isn't auto-injected and doesn't have hot reload capability.
	if dev {
		proxy.ServeHTTP(w, r)
		return
	}

	// Serve the default index.html, where built files are auto-injected with webpack, otherwise.
	// This should only really happen on page paths, while in production.
	ctx.Logger.Debugw("Serving dist/index.html", "path", r.URL.Path)
	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

// Static is used to handle /static/.
func Static(w http.ResponseWriter, r *http.Request) {
	ctx.Logger.Debugw("Serving static", "path", r.URL.Path)

	// Serve /static/ with `make dev-ui` in development mode
	if dev {
		proxy.ServeHTTP(w, r)
		return
	}

	// Otherwise, we serve the embedded static files normally.
	httpFS.ServeHTTP(w, r)
}
