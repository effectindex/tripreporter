// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package db

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	schemaDefs = []string{
		`create table if not exists accounts (
			id uuid primary key,
			email varchar(255) not null unique,
			username varchar(255) not null unique,
			password_salt bytea not null unique,
			password_hash bytea not null,
			email_verified bool default false,
			CHECK (accounts.email <> ''),
			CHECK (accounts.username <> ''),
			CHECK (accounts.password_salt <> ''),
			CHECK (accounts.password_hash <> '')
		);`,
		`create table if not exists users (
			account_id uuid primary key references accounts(id) on delete cascade,
			created timestamptz not null default to_timestamp('0'),
			display_name varchar(255) not null,
			date_of_birth timestamptz not null,
			height decimal default 0,
			weight decimal default 0
		);`,
		`create table if not exists sessions (
    		account_id uuid references accounts(id) on delete cascade,
    		session_index int not null default 0,
    		session_key uuid not null unique,
    		primary key(account_id, session_index)
    	);`,
		`create table if not exists drugs (
			id uuid primary key,
			account_id uuid references accounts(id) on delete cascade,
			drug_name varchar(4096) not null,
			drug_dosage int not null default 0,
			drug_dosage_unit varchar(4096) not null,
			drug_roa int not null default 0,
			drug_frequency int default 0,
			drug_prescribed int default 0
		);`,
		`create table if not exists reports (
    		id uuid primary key,
    		account_id uuid references accounts(id),
    		creation_time timestamptz not null default to_timestamp('0'),    		
    		modified_time timestamptz not null default to_timestamp('0'),
    		report_date timestamptz default to_timestamp('0'),
    		title varchar(4096) not null, -- 156*16, or 4KiB
    		setting varchar(4096) not null -- 156*16, or 4KiB
    	);`,
		`create table if not exists report_events (
    		report_id uuid references reports(id) on delete cascade,
    		event_index int not null,
    		event_timestamp timestamptz default to_timestamp('0'),
    		event_type int not null,
    		event_section int not null default 0,
    		event_content varchar(10485760), -- 10MiB
    		event_drug uuid references drugs(id) on delete cascade ,
    		primary key(report_id, event_index)
    	);`,
	}
	indexDefs = []string{
		//`create unique index if not exists index_email on users(email);`, // TODO: INDEXES
		//`create index if not exists index_user_id on submissions(user_id);`,
		//`create index if not exists index_submission_time on submissions(created);`,
		//`create index if not exists index_goals_index on goals(index);`,
	}
	triggerDefs = []string{
		//`create or replace function upd_last_edited() returns trigger language plpgsql as // TODO
		//$$ begin
		//	new.last_edited = current_timestamp at time zone 'utc';
		//	return new;
		//end; $$;`,
		//`create trigger upd_last_edited_tgr
		//before update on submissions
		//for each row execute procedure upd_last_edited();`,
	}
)

// SetupDB creates a new PostgreSQL connection
func SetupDB(docker bool, logger *zap.SugaredLogger) *pgxpool.Pool {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if docker {
		dbHost = os.Getenv("DOCKER_POSTGRES_HOST")
	}

	// First try connecting to the database we're going to use
	dbPool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		// Now, try connecting to the default and making a db there
		dbPool, err = pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/postgres", dbUser, dbPass, dbHost, dbPort))
		if err != nil {
			logger.Fatalw("Cannot connect to DB", zap.Error(err))
		}
		_, err = dbPool.Exec(context.Background(), fmt.Sprintf("create database \"%s\"", dbName))
		if err != nil {
			logger.Fatalw("Cannot create database", zap.Error(err))
		}
		dbPool.Close()
		dbPool, err = pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
		if err != nil {
			logger.Fatalw("Cannot use database", zap.Error(err))
		}
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		logger.Fatalw("Cannot connect to DB", zap.Error(err))
	}

	for _, def := range schemaDefs {
		logger.Infof("Executing %s", def)
		_, err = dbPool.Exec(context.Background(), def)
		if err != nil {
			logger.Fatalw("Cannot execute schema", zap.Error(err))
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
			logger.Infow("Failed to create DB trigger", zap.Error(err))
		}
	}
	logger.Infof("Connected to SQL DB at %s:%s", dbHost, dbPort)

	return dbPool
}

// SetupRedis creates a new redis.Client instance
func SetupRedis(docker bool, logger *zap.SugaredLogger) *redis.Client {
	dbHost := os.Getenv("REDIS_HOST")
	dbPort := os.Getenv("REDIS_PORT")
	dbPass := os.Getenv("REDIS_PASS")

	if docker {
		dbHost = os.Getenv("DOCKER_REDIS_HOST")
	}

	dbAddr := dbHost + ":" + dbPort
	rdc := redis.NewClient(&redis.Options{
		Addr:     dbAddr,
		Password: dbPass,
		DB:       0,
	})

	// If connection doesn't work, panic
	if _, err := rdc.Ping(context.Background()).Result(); err != nil {
		logger.Panic("Failed to ping Redis server")
	}

	// We have a working connection
	logger.Infof("Connected to Redis at  %s", dbAddr)
	return rdc
}
