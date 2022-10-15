package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/effectindex/tripreporter/ui"
	"github.com/joho/godotenv"
)

var (
	dev   = flag.Bool("dev", false, "Run in development mode, alongside `make dev-ui`.")
	proxy *httputil.ReverseProxy
)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(target string) *httputil.ReverseProxy {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("error making reverse proxy: %v\n", err) // likely malformed addr
		return nil
	}

	return httputil.NewSingleHostReverseProxy(u)
}

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading .env file (copy the .env.example): %v\n", err)
	}

	if err := validateEnvKeys("DEV_PORT", "SRV_PORT"); err != nil {
		log.Fatalf("missing .env variables (copy the .env.example): %v\n", err)
	}

	proxy = NewProxy("http://localhost:" + os.Getenv("DEV_PORT")) // proxy is used for `make dev-ui`
	s := &http.Server{
		Addr:        "localhost:" + os.Getenv("SRV_PORT"),
		Handler:     router(),
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

func router() http.Handler {
	mux := http.NewServeMux()

	// index
	mux.HandleFunc("/", indexHandler)

	// static files
	if !*dev { // if running in development mode, let indexHandler reverse proxy it
		staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
		httpFS := http.FileServer(http.FS(staticFS))
		mux.Handle("/static/", httpFS)
	}

	// api
	mux.HandleFunc("/api/v1/greeting", greetingAPI)
	return mux
}

// TODO: impl
//func handleErr(w http.ResponseWriter, err error) {
//	w.WriteHeader(http.StatusInternalServerError)
//	fmt.Fprintf(w, "%s: %v\n", http.StatusText(http.StatusInternalServerError), err)
//}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

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
