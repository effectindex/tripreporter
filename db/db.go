package db

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

var (
	schemaDefs = []string{
		`create table if not exists accounts (
			id uuid primary key,
			email varchar(255) not null unique,
			username varchar(255) not null unique,
			password_hash varchar(255)
		);`,
		`create table if not exists users (
			id uuid references accounts(id),
			created timestamptz default to_timestamp(0),
			display_name varchar(255),
			date_of_birth timestamptz default to_timestamp(0),
			age int default -1,
			height decimal default -1,
			weight decimal default -1
		);`,
	}
	indexDefs = []string{
		//`create unique index if not exists index_email on users(email);`,
		//`create index if not exists index_user_id on submissions(user_id);`,
		//`create index if not exists index_submission_time on submissions(created);`,
		//`create index if not exists index_goals_index on goals(index);`,
	}
	triggerDefs = []string{
		//`create or replace function upd_last_edited() returns trigger language plpgsql as
		//$$ begin
		//	new.last_edited = current_timestamp at time zone 'utc';
		//	return new;
		//end; $$;`,
		//`create trigger upd_last_edited_tgr
		//before update on submissions
		//for each row execute procedure upd_last_edited();`,
	}
)

func SetupDB(logger *zap.SugaredLogger) *pgxpool.Pool {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// First try connecting to the database we're going to use
	dbPool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		// Now, try connecting to the default and making a db there
		dbPool, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/postgres", dbUser, dbPass, dbHost, dbPort))
		if err != nil {
			logger.Fatalw("Cannot connect to DB", zap.Error, err)
		}
		_, err = dbPool.Exec(context.Background(), fmt.Sprintf("create database \"%s\"", dbName))
		if err != nil {
			logger.Fatalw("Cannot create database", zap.Error, err)
		}
		dbPool.Close()
		dbPool, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
		if err != nil {
			logger.Fatalw("Cannot use database", zap.Error, err)
		}
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		logger.Fatalw("Cannot connect to DB", zap.Error, err)
	}

	for _, def := range schemaDefs {
		logger.Infof("Executing %s", def)
		_, err = dbPool.Exec(context.Background(), def)
		if err != nil {
			logger.Fatalw("Cannot execute schema", zap.Error, err)
		}
	}
	for _, def := range indexDefs {
		logger.Infof("Executing %s", def)
		_, err := dbPool.Exec(context.Background(), def)
		if err != nil {
			logger.Infow("Failed to create index on DB. It probably already exists, so you can ignore this.", zap.Error, err)
		}
	}
	for _, def := range triggerDefs {
		logger.Infof("Executing %s", def)
		_, err := dbPool.Exec(context.Background(), def)
		if err != nil {
			logger.Infow("Failed to create DB trigger", zap.Error, err)
		}
	}
	logger.Infof("Connected to SQL DB at %s:%s", dbHost, dbPort)
	return dbPool
}

// SetupRedis create a new redis.Client instance
func SetupRedis(logger *zap.SugaredLogger) *redis.Client {
	rdc := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	// If connection doesn't work, panic
	_, err := rdc.Ping(context.Background()).Result()
	if err != nil {
		logger.Panic("Failed to ping Redis server")
	}

	// We have a working connection
	logger.Infof("Connected to Redis at  %s", os.Getenv("REDIS_HOST"))
	return rdc
}
