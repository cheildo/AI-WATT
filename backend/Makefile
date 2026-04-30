# ════════════════════════════════════════════════════════════════════════════
# AI WATT — Backend Makefile
# Usage: make <target> [ENV=local|staging|production]
# ════════════════════════════════════════════════════════════════════════════

ENV   ?= local
IMAGE ?= ghcr.io/neurowatt/aiwatt-backend
TAG   ?= latest

MIGRATIONS_DIR := scripts/migrations
ENV_FILE       := .env.$(ENV)

# ── Env file loading ──────────────────────────────────────────────────────────
# Source the env file (if it exists) before running any recipe command.
# set -a exports every variable defined in the file to the child process.
# Usage in recipes: $(LOAD_ENV) <your-command>
#
# Falls back gracefully if the file is missing — useful in CI where vars
# are injected directly into the shell environment.
LOAD_ENV = $(if $(wildcard $(ENV_FILE)),set -a && . ./$(ENV_FILE) && set +a &&,)

# Derive DATABASE_URL for migration targets.
# Prefer the value from the env file; allow shell override for CI.
define db_url
$(shell $(if $(wildcard $(ENV_FILE)), . ./$(ENV_FILE) && echo $$DATABASE_URL, echo $(DATABASE_URL)))
endef

# ── Development ───────────────────────────────────────────────────────────────

.PHONY: dev
dev: ## Start backend (hot-reload with air if installed, else go run) [ENV=local]
	@echo "Loading $(ENV_FILE)..."
	@$(LOAD_ENV) \
	if command -v air > /dev/null 2>&1; then \
		air; \
	else \
		go run ./cmd/api; \
	fi

# ── Testing ───────────────────────────────────────────────────────────────────

.PHONY: test
test: ## Run all Go tests with race detector [ENV=local]
	@$(LOAD_ENV) APP_ENV=test go test -race -count=1 ./...

.PHONY: test-verbose
test-verbose: ## Run tests with verbose output [ENV=local]
	@$(LOAD_ENV) APP_ENV=test go test -race -v -count=1 ./...

# ── Build ─────────────────────────────────────────────────────────────────────

.PHONY: build
build: ## Compile binary for current OS
	CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o tmp/aiwatt-api ./cmd/api

.PHONY: docker-build
docker-build: ## Build Docker image
	docker build \
		--tag $(IMAGE):$(TAG) \
		--tag $(IMAGE):$(shell git rev-parse --short HEAD) \
		.

.PHONY: docker-push
docker-push: ## Push Docker image to registry
	docker push $(IMAGE):$(TAG)
	docker push $(IMAGE):$(shell git rev-parse --short HEAD)

# ── Linting ───────────────────────────────────────────────────────────────────

.PHONY: lint
lint: ## Run go vet + staticcheck
	go vet ./...
	@command -v staticcheck > /dev/null 2>&1 && staticcheck ./... || true

# ── Migrations ────────────────────────────────────────────────────────────────
# DATABASE_URL is read from .env.$(ENV) automatically.
# Override on the command line: make migrate-up DATABASE_URL=mysql://...

.PHONY: migrate-up
migrate-up: ## Run all pending migrations [ENV=local|staging|production]
	@$(LOAD_ENV) migrate -path $(MIGRATIONS_DIR) -database "$$DATABASE_URL" up

.PHONY: migrate-down
migrate-down: ## Roll back last migration [ENV=local|staging|production]
	@$(LOAD_ENV) migrate -path $(MIGRATIONS_DIR) -database "$$DATABASE_URL" down 1

.PHONY: migrate-status
migrate-status: ## Show current migration version [ENV=local|staging|production]
	@$(LOAD_ENV) migrate -path $(MIGRATIONS_DIR) -database "$$DATABASE_URL" version

.PHONY: migrate-force
migrate-force: ## Force migration to version N — usage: make migrate-force ARGS=9 [ENV=local]
	@$(LOAD_ENV) migrate -path $(MIGRATIONS_DIR) -database "$$DATABASE_URL" force $(ARGS)

# ── Cleanup ───────────────────────────────────────────────────────────────────

.PHONY: clean
clean:
	rm -rf tmp/

# ── Help ──────────────────────────────────────────────────────────────────────

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@echo ""
	@echo "AI WATT — Backend"
	@echo "Usage: make <target> [ENV=local|staging|production]"
	@echo ""
	@echo "Env file loaded: .env.\$$ENV  (copy .env.\$$ENV.example if missing)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) \
		| sed 's/:.*##/  →/' \
		| sed 's/^/  /'
	@echo ""
