tripreporter: deps-server build-server deps-ui build-ui

##########################################################
# deps

deps-ui:
	cd ui/; \
	npm i

deps-server:
	go get -u "github.com/georgysavva/scany/v2"
	go get -u "github.com/go-redis/redis"
	go get -u "github.com/google/uuid"
	go get -u "github.com/gorilla/mux"
	go get -u "github.com/jackc/pgx/v5"
	go get -u "github.com/joho/godotenv"
	go get -u "github.com/shopspring/decimal"
	go get -u "go.uber.org/zap"
	go get -u "golang.org/x/crypto"
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
