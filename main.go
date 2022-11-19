package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/effectindex/tripreporter/api"
	"github.com/effectindex/tripreporter/db"
	"github.com/effectindex/tripreporter/models"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	dev = flag.Bool("dev", false, "Run in development mode, alongside `make dev-ui`.")
)

func main() {
	flag.Parse()

	// Setup Zap for logging
	var err error
	var logger *zap.Logger

	if *dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sLogger := logger.Sugar()
	ctx := models.Context{Logger: sLogger}

	// Load and validate .env
	if err := godotenv.Load(); err != nil {
		logger.Fatal("err loading .env file (copy the .env.example)", zap.Error(err))
	}

	// "SRV_ADDR" and "REDIS_PASS" can be empty, they're the only optional ones
	if err := validateEnvKeys(
		"SRV_PORT", "DEV_PORT", "SITE_NAME", "WORDLIST", "DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME", "REDIS_HOST",
	); err != nil {
		logger.Fatal("missing .env variables (copy the .env.example)", zap.Error(err))
	}

	// Setup NodeID for uuid generation
	randomID := make([]byte, 6)
	if _, err = rand.Read(randomID); err != nil {
		logger.Fatal("failed to initialize NodeID", zap.Error(err))
	}
	randomID[5] |= 0x01 // Set least significant bit of first true
	uuid.SetNodeID(randomID)
	ctx.Logger.Infof("Initialized random NodeID: %s", hex.EncodeToString(randomID))

	// Setup required connections for postgresql and redis
	sDB := db.SetupDB(ctx.Logger)
	rDB := db.SetupRedis(ctx.Logger)

	defer sDB.Close()
	defer func(rDB *redis.Client) {
		err := rDB.Close()
		if err != nil {
			logger.Fatal("Failed to close Redis", zap.Error(err))
		}
	}(rDB)

	// Set context database now that we have one
	ctx.Database = sDB

	// Setup wordlist
	models.SetupWordlist(ctx)

	// Setup proxy to webpack hot-reload server (for dev-ui) and regular http server (serves everything)
	api.Setup(*dev, ctx.Logger)

	// Setup http server
	s := &http.Server{
		Addr:        os.Getenv("SRV_ADDR") + ":" + os.Getenv("SRV_PORT"),
		Handler:     api.Handler(),
		IdleTimeout: time.Minute,
	}

	if *dev {
		ctx.Logger.Infof("Running on %s in development mode...", s.Addr)
	} else {
		ctx.Logger.Infof("Running on %s in production mode...", s.Addr)
	}

	if err := s.ListenAndServe(); err != nil {
		ctx.Logger.DPanicf("Error in ListenAndServe: %v", err)
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
