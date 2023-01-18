# tripreporter

## Building and running

## Required

Copy `.env.example` to `.env`, and `config/redis.conf.example` to `config/redis.conf`, and (optionally) modify them. 

These files are required for building and running the tripreporter.
This step is necessary for both development and production use.

### For development (no Docker)

Choose one of the following methods to run, `http://localhost:3000` should be accessible afterwards.

If you have issues when first running development, try running a production build first.

```bash
# Build production without Docker. This is required to have caches of static files for `make dev-server`.
make

# Build and run development UI, and server.
# Run both in two different terminals.
make dev-ui
make dev-server
```

### Docker

This is intended for production use.
This will run on `http://localhost:3000` by default.

```bash
docker compose up -d

# If you have issues / want to troubleshoot, use this command to force re-build (add -d to run in background)
docker-compose up --build --force-recreate --no-deps
```
