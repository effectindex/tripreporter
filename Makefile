# SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
#
# SPDX-License-Identifier: OSL-3.0

tripreporter: deps-ui build-ui deps-server build-server test-server

##########################################################
# licensing

reuse-lint:
	which reuse || { echo "`reuse` not found! see https://reuse.software/"; exit 1; }
	reuse lint || exit 1

##########################################################
# tests

test-server:
	@while read -r l; do \
  		go test "$$l"; \
	done < <(for f in $$(find . -name '*_test.go'); do dirname "$$f"; done | sort | uniq)

##########################################################
# deps

deps-ui:
	cd ui/; \
	npm i

deps-server:
	go get -u "github.com/cristalhq/jwt/v4"
	go get -u "github.com/georgysavva/scany/v2"
	go get -u "github.com/go-redis/redis"
	go get -u "github.com/google/uuid"
	go get -u "github.com/gorilla/mux"
	go get -u "github.com/jackc/pgtype"
	go get -u "github.com/jackc/pgx/v5"
	go get -u "github.com/joho/godotenv"
	go get -u "github.com/rs/cors"
	go get -u "github.com/shopspring/decimal"
	go get -u "go.uber.org/zap"
	go get -u "golang.org/x/crypto"
	go get -u "golang.org/x/exp/slices"
	go mod tidy

##########################################################
# build

build-ui: deps-ui
	cd ui/; \
	npm run build

build-server:
	go build -o tripreporter .
	chmod +x tripreporter

##########################################################
# Run development frontend + backend.
# For normal development, you will want to run both at
# at the same time. dev-server provides the API and
# dev-ui provides the frontend, both on localhost:3000.

dev-ui:
	cd ui/; \
	npm run serve

dev-server: build-server
	./tripreporter -dev
