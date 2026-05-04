# AI WATT — Backend

Golang API server for the AI WATT protocol.

**Stack:** Go 1.22 · Gin · GORM · MySQL 8 · Redis · go-migrate · go-ethereum · zap

---

## Prerequisites

- Go 1.22+
- MySQL 8 and Redis 7 — either via Docker (see repo root) or Homebrew

---

## Running with Docker (recommended)

From the repo root:

```bash
cp .env.example .env
docker compose up -d --build
```

Migrations run automatically in the `aiwatt-migrator` container before the API starts.

---

## Running natively (Homebrew)

```bash
# Start services (if not already running)
brew services start mysql
/opt/homebrew/opt/redis/bin/redis-server /opt/homebrew/etc/redis.conf --daemonize yes

cp .env.example .env.local      # already populated for local use

# Apply migrations
make migrate-up ENV=local

# Start with hot-reload (uses air if installed, else go run)
make dev ENV=local
```

API available at `http://localhost:8080`.

---

## Makefile targets

```bash
make dev             ENV=local      # start server (hot-reload)
make build                          # compile binary → tmp/aiwatt-api
make test                           # go test -race ./...
make migrate-up      ENV=local      # apply all pending migrations
make migrate-down    ENV=local      # roll back last migration
make migrate-status  ENV=local      # show current schema version
make migrate-force   ENV=local ARGS=9  # force version to N
make lint                           # go vet + staticcheck
make swag                           # regenerate Swagger docs
```

The `ENV` flag selects the env file (e.g. `ENV=local` loads `.env.local`).

---

## Database migrations

Migration files live in `scripts/migrations/`. Always use `make migrate-*` — never `GORM AutoMigrate` in production.

```
000001_create_users.up.sql / .down.sql
000002_create_assets.up.sql / .down.sql
...
```

Every migration requires a working `.down.sql`.

---

## Generate Swagger docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
make swag
```

Spec written to `api/openapi.yaml`. Run after any handler annotation change.

---

## Project layout

```
backend/
├── cmd/api/            # main.go — server entry point
├── internal/
│   ├── api/
│   │   ├── handler/    # HTTP handlers (one file per domain)
│   │   ├── dto/        # Request / response structs
│   │   ├── middleware/ # JWT auth, rate limiter, request logger
│   │   └── router.go
│   ├── service/        # Business logic
│   ├── repository/     # GORM queries
│   ├── models/         # GORM structs
│   ├── blockchain/     # XDC client, event indexer, tx manager
│   └── veriflow/       # Telemetry ingestor, scorer, attester
├── pkg/                # logger, jwt, crypto, response envelope
├── configs/            # Config struct + example
├── api/                # openapi.yaml (generated)
└── scripts/migrations/ # go-migrate SQL files
```

---

## Layer rules

```
Request → DTO → Handler → Service → Repository → Model → DB
```

- Handlers bind input to a DTO, call service, map result to a response DTO
- Services contain business logic only — no GORM, no `http.Request`
- Repositories receive and return `models.*` — no DTOs, no business logic
- Models are never returned as JSON — always mapped to a response DTO first

---

## Tests

```bash
make test             # all tests with race detector
make test-verbose     # verbose output
```

---

## Environment variables

See [.env.example](.env.example) for the full list with descriptions.
