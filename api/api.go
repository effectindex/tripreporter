package api

import (
	"fmt"
	"github.com/effectindex/tripreporter/ui"
	"github.com/effectindex/tripreporter/util"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
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

// Handler manages all paths, with Router handling anything not defined here.
func Handler() http.Handler {
	mux := http.NewServeMux()

	// let api.Router do everything else
	mux.HandleFunc("/", Router)

	// serve /static/ by cache in production (no hot-reload support)
	if !dev { // if running in development mode, let api.Router reverse proxy it
		staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
		httpFS := http.FileServer(http.FS(staticFS))
		mux.Handle("/static/", httpFS)
	}

	// API functions
	mux.HandleFunc("/api/v1/greeting", greetingAPI)
	return mux
}

// Router will route everything except /static/ and valid /api/ endpoints.
func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	// For API endpoints not already handled in Handler()
	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
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

func greetingAPI(w http.ResponseWriter, r *http.Request) { // TODO: Remove / use individual files for routes
	w.Write([]byte("Hello, there!"))
}
