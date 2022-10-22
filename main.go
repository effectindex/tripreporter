package main

import (
	"errors"
	"flag"
	"github.com/effectindex/tripreporter/api"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	dev = flag.Bool("dev", false, "Run in development mode, alongside `make dev-ui`.")
)

func main() {
	flag.Parse()

	// Load and validate .env

	if err := godotenv.Load(); err != nil {
		log.Fatalf("err loading .env file (copy the .env.example): %v\n", err)
	}

	// "ADDRESS" can be empty, it's the only optional one
	if err := validateEnvKeys("DEV_PORT", "SRV_PORT", "SITE_NAME"); err != nil {
		log.Fatalf("missing .env variables (copy the .env.example): %v\n", err)
	}

	// Setup proxy to webpack hot-reload server (for dev-ui) and regular http server (serves everything)
	api.Setup(*dev)
	s := &http.Server{
		Addr:        os.Getenv("ADDRESS") + ":" + os.Getenv("SRV_PORT"),
		Handler:     api.Handler(),
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
