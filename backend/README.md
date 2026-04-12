# AI WATT — Backend

Golang API server for the AI WATT protocol.

**Stack:** Go 1.22 · Gin · GORM · MySQL 8 · Redis · go-migrate · go-ethereum · zap

---

## Prerequisites

- Go 1.22+
- MySQL 8 (or `docker-compose up -d` from repo root)
- Redis 7 (or `docker-compose up -d` from repo root)

---

## Setup

```bash
cp .env.example .env
# Fill in JWT_SECRET, DATABASE_URL, REDIS_URL at minimum
```

### Run database migrations

```bash
# Install migrate CLI (once)
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Apply all migrations
migrate -path scripts/migrations \
        -database "$DATABASE_URL" \
        up
```

### Start the server

```bash
go run ./cmd/api
# API available at http://localhost:8080
# Swagger UI at http://localhost:8080/swagger/index.html
```

---

## Generate Swagger docs

```bash
# Install swag CLI (once)
go install github.com/swaggo/swag/cmd/swag@latest

swag init -g cmd/api/main.go -o api/
```

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

- Handlers never touch GORM directly
- Services never write SQL
- Repositories never contain business logic
- Models are never returned as JSON — always mapped to a response DTO first

---

## Running tests

```bash
go test ./...
```

---

## Environment variables

See [.env.example](.env.example) for the full list.
