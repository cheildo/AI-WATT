## ── API build stage ──────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/aiwatt-api ./cmd/api

## ── Migrate build stage ──────────────────────────────────────────────────────
FROM golang:1.22-alpine AS migrate-builder

ENV CGO_ENABLED=0
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

## ── Migrator runtime stage ───────────────────────────────────────────────────
FROM alpine:3.20 AS migrator

RUN apk add --no-cache ca-certificates

COPY --from=migrate-builder /go/bin/migrate /usr/local/bin/migrate
COPY scripts/migrations /migrations

ENTRYPOINT ["migrate", "-path=/migrations"]

## ── API runtime stage ────────────────────────────────────────────────────────
FROM gcr.io/distroless/static-debian12 AS api

COPY --from=builder /bin/aiwatt-api /aiwatt-api
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/aiwatt-api"]
