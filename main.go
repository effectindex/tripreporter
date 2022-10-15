package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/effectindex/tripreporter/ui"
	"github.com/effectindex/tripreporter/util"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

var (
	dev   = flag.Bool("dev", false, "Run in development mode, alongside `make dev-ui`.")
	proxy *httputil.ReverseProxy
)

func main() {
	flag.Parse()

	// Load and validate .env

	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading .env file (copy the .env.example): %v\n", err)
	}

	if err := validateEnvKeys("DEV_PORT", "SRV_PORT"); err != nil {
		log.Fatalf("missing .env variables (copy the .env.example): %v\n", err)
	}

	// Setup proxy to webpack hot-reload server (for dev-ui) and regular http server (serves everything)
	proxy = util.NewProxy("http://localhost:" + os.Getenv("DEV_PORT")) // proxy is used for `make dev-ui`
	s := &http.Server{
		Addr:        "localhost:" + os.Getenv("SRV_PORT"),
		Handler:     Handler(),
		IdleTimeout: time.Minute,
	}

	if *dev {
		log.Printf("Running on %s in development mode...\n", s.Addr)
	} else {
		log.Printf("Running on %s in production mode...\n", s.Addr)
	}

	if err := s.ListenAndServe(); err != nil {
		log.Printf("error in ListenAndServe: %v\n", err)
	}
}

func Handler() http.Handler {
	mux := http.NewServeMux()

	// let Router do everything else
	mux.HandleFunc("/", Router)

	// serve /static/ by cache in production (no hot-reload support)
	if !*dev { // if running in development mode, let Router reverse proxy it
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

	if *dev {
		proxy.ServeHTTP(w, r)
	} else {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
		w.Write(rawFile)
	}
}

func greetingAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, there!"))
}

func validateEnvKeys(keys ...string) error {
	missing := make([]string, 0)
	for _, key := range keys {
		if os.Getenv(key) == "" {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		return errors.New("[" + strings.Join(missing, ", ") + "]")
	}
	return nil
}
