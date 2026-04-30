# ════════════════════════════════════════════════════════════════════════════
# AI WATT — Frontend Makefile
# ════════════════════════════════════════════════════════════════════════════

ENV ?= local

.PHONY: dev
dev: ## Start Vite dev server [ENV=local|staging|production]
ifeq ($(ENV),staging)
	npm run dev -- --mode staging
else ifeq ($(ENV),production)
	npm run dev -- --mode production
else
	npm run dev
endif

.PHONY: build
build: ## Build for ENV=local|staging|production
ifeq ($(ENV),staging)
	npm run build -- --mode staging
else ifeq ($(ENV),production)
	npm run build -- --mode production
else
	npm run build
endif

.PHONY: preview
preview: ## Preview production build locally
	npm run preview

.PHONY: check
check: lint typecheck ## Run all static checks

.PHONY: typecheck
typecheck: ## Run TypeScript compiler check
	npx tsc --noEmit

.PHONY: lint
lint: ## Run ESLint
	npm run lint

.PHONY: clean
clean:
	rm -rf dist/
