# ════════════════════════════════════════════════════════════════════════════
# AI WATT — Root Makefile
# Usage: make <target> [ENV=local|staging|production]
# ════════════════════════════════════════════════════════════════════════════

BOLD  := \033[1m
GREEN := \033[0;32m
GOLD  := \033[0;33m
CYAN  := \033[0;36m
RED   := \033[0;31m
RESET := \033[0m

ENV ?= local

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@echo ""
	@echo "$(BOLD)AI WATT — Build & Deploy$(RESET)"
	@echo "$(CYAN)Usage: make <target> [ENV=local|staging|production]$(RESET)"
	@echo ""
	@echo "$(GOLD)── Development ────────────────────────────────────$(RESET)"
	@grep -E '^(dev|local)[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""
	@echo "$(GOLD)── Testing ─────────────────────────────────────────$(RESET)"
	@grep -E '^test[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""
	@echo "$(GOLD)── Building ─────────────────────────────────────────$(RESET)"
	@grep -E '^build[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""
	@echo "$(GOLD)── Contracts ────────────────────────────────────────$(RESET)"
	@grep -E '^(deploy|verify)[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""
	@echo "$(GOLD)── Database ─────────────────────────────────────────$(RESET)"
	@grep -E '^migrate[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""
	@echo "$(GOLD)── Utilities ────────────────────────────────────────$(RESET)"
	@grep -E '^(lint|clean|install)[^:]*:.*##' $(MAKEFILE_LIST) | sed 's/:.*##/  →/' | sed 's/^/  /'
	@echo ""


# ── Development ──────────────────────────────────────────────────────────────

.PHONY: dev
dev: ## Start stack for ENV (local: Docker infra + migrate + servers; staging/production: servers only)
ifeq ($(ENV),local)
	@echo "$(GREEN)Starting local stack...$(RESET)"
	@$(MAKE) dev-infra
	@$(MAKE) migrate-up ENV=local
endif
	@echo "$(GREEN)Starting backend + frontend (ENV=$(ENV))...$(RESET)"
	@(cd backend && $(MAKE) dev ENV=$(ENV)) &
	@(cd frontend && $(MAKE) dev ENV=$(ENV)) &
	@wait

.PHONY: dev-infra
dev-infra: ## Start Docker infra only (MySQL + Redis)
	@echo "$(GREEN)Starting MySQL + Redis...$(RESET)"
	docker compose up -d mysql redis
	@echo "$(CYAN)Waiting for health checks...$(RESET)"
	@until docker compose exec -T mysql mysqladmin ping -h localhost -u aiwatt -paiwatt --silent 2>/dev/null; do sleep 2; done
	@echo "$(GREEN)MySQL ready$(RESET)"

.PHONY: dev-backend
dev-backend: ## Start backend API only (hot-reload with air if installed)
	@cd backend && $(MAKE) dev

.PHONY: dev-frontend
dev-frontend: ## Start frontend dev server only
	@cd frontend && $(MAKE) dev

.PHONY: local
local: dev ## Alias for dev


# ── Testing ──────────────────────────────────────────────────────────────────

.PHONY: test
test: test-contracts test-backend test-frontend ## Run all tests

.PHONY: test-contracts
test-contracts: ## Run all Hardhat contract tests (258 tests)
	@echo "$(GREEN)Running contract tests...$(RESET)"
	@cd contracts && $(MAKE) test

.PHONY: test-backend
test-backend: ## Run Go backend unit tests
	@echo "$(GREEN)Running backend tests...$(RESET)"
	@cd backend && $(MAKE) test

.PHONY: test-frontend
test-frontend: ## Run frontend type check + lint
	@echo "$(GREEN)Running frontend checks...$(RESET)"
	@cd frontend && $(MAKE) check

.PHONY: test-integration
test-integration: ## Run integration tests against staging (ENV=staging required)
	@echo "$(GREEN)Running integration tests against $(ENV)...$(RESET)"
	@scripts/health-check.sh $(ENV)


# ── Building ─────────────────────────────────────────────────────────────────

.PHONY: build
build: build-backend build-frontend build-agent ## Build all artifacts

.PHONY: build-backend
build-backend: ## Build backend Docker image [ENV=local|staging|production]
	@echo "$(GREEN)Building backend Docker image ($(ENV))...$(RESET)"
	@cd backend && $(MAKE) docker-build ENV=$(ENV)

.PHONY: build-frontend
build-frontend: ## Build frontend static assets [ENV=local|staging|production]
	@echo "$(GREEN)Building frontend ($(ENV))...$(RESET)"
	@cd frontend && $(MAKE) build ENV=$(ENV)

.PHONY: build-agent
build-agent: ## Build Veriflow agent static Linux binary
	@echo "$(GREEN)Building Veriflow agent...$(RESET)"
	@cd veriflow-agent && $(MAKE) build-agent


# ── Contract Deployment ───────────────────────────────────────────────────────

.PHONY: deploy-mocks
deploy-mocks: ## Deploy mock USDC/USDT to local Hardhat node
	@cd contracts && $(MAKE) deploy-mocks

.PHONY: deploy-contracts-local
deploy-contracts-local: ## Deploy all contracts to local Hardhat node
	@cd contracts && $(MAKE) deploy ENV=local

.PHONY: deploy-contracts-staging
deploy-contracts-staging: ## Deploy/upgrade contracts to XDC Apothem testnet
	@echo "$(GOLD)Deploying contracts to Apothem (staging)...$(RESET)"
	@cd contracts && $(MAKE) deploy ENV=staging

.PHONY: deploy-contracts-production
deploy-contracts-production: ## Deploy/upgrade contracts to XDC mainnet [REQUIRES CONFIRMATION]
	@echo "$(RED)$(BOLD)WARNING: Deploying to XDC MAINNET. This is irreversible.$(RESET)"
	@read -p "Type 'deploy-mainnet' to confirm: " confirm && [ "$$confirm" = "deploy-mainnet" ]
	@cd contracts && $(MAKE) deploy ENV=production

.PHONY: verify-contracts-staging
verify-contracts-staging: ## Verify contracts on Apothem block explorer
	@cd contracts && $(MAKE) verify ENV=staging

.PHONY: verify-contracts-production
verify-contracts-production: ## Verify contracts on XDC mainnet block explorer
	@cd contracts && $(MAKE) verify ENV=production


# ── Database Migrations ───────────────────────────────────────────────────────

.PHONY: migrate-up
migrate-up: ## Run all pending migrations [ENV=local|staging|production]
	@echo "$(GREEN)Running migrations ($(ENV))...$(RESET)"
	@cd backend && $(MAKE) migrate-up ENV=$(ENV)

.PHONY: migrate-down
migrate-down: ## Roll back the last migration [ENV=local|staging|production]
	@echo "$(GOLD)Rolling back migration ($(ENV))...$(RESET)"
	@cd backend && $(MAKE) migrate-down ENV=$(ENV)

.PHONY: migrate-status
migrate-status: ## Show current migration version [ENV=local|staging|production]
	@cd backend && $(MAKE) migrate-status ENV=$(ENV)


# ── Utilities ────────────────────────────────────────────────────────────────

.PHONY: lint
lint: ## Lint all sub-repos
	@cd backend  && $(MAKE) lint
	@cd contracts && $(MAKE) lint
	@cd frontend  && $(MAKE) lint

.PHONY: install
install: ## Install all dependencies (Go, npm, migrate CLI)
	@echo "$(GREEN)Installing dependencies...$(RESET)"
	@cd backend   && go mod download
	@cd contracts && npm install
	@cd frontend  && npm install
	@which migrate > /dev/null 2>&1 || \
		go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "$(GREEN)All dependencies installed$(RESET)"

.PHONY: clean
clean: ## Remove all build artifacts
	@rm -rf backend/tmp
	@rm -rf frontend/dist
	@rm -rf veriflow-agent/dist
	@rm -rf contracts/artifacts contracts/cache
	@echo "$(GREEN)Cleaned$(RESET)"

.PHONY: health
health: ## Check health of running services [ENV=local|staging|production]
	@scripts/health-check.sh $(ENV)

.PHONY: logs
logs: ## Tail backend logs (Docker)
	docker compose logs -f backend

.PHONY: ps
ps: ## Show running containers
	docker compose ps
