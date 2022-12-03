tripreporter: deps-server build-server deps-ui build-ui

# deps

deps-server:
	go get -u "github.com/georgysavva/scany/v2"
	go get -u "github.com/go-redis/redis"
	go get -u "github.com/google/uuid"
	go get -u "github.com/jackc/pgx/v5"
	go get -u "github.com/joho/godotenv"
	go get -u "github.com/shopspring/decimal"
	go get -u "go.uber.org/zap"
	go get -u "golang.org/x/crypto"
	go mod tidy

deps-ui:
	cd ui/; \
	npm i

# build

build-server:
	go build -o tripreporter .
	chmod +x tripreporter

build-ui: deps-ui
	cd ui/; \
	npm run build

# run production

run:
	./tripreporter

all: tripreporter run

# run development (needs both dev and dev-ui separately)

dev-ui: deps-ui
	cd ui/; \
	npm run serve

dev-server: build-server
	./tripreporter -dev
