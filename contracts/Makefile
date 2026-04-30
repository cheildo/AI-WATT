# ════════════════════════════════════════════════════════════════════════════
# AI WATT — Contracts Makefile
# ════════════════════════════════════════════════════════════════════════════

ENV ?= local

ifeq ($(ENV),local)
  NETWORK := localhost
else ifeq ($(ENV),staging)
  NETWORK := apothem
else ifeq ($(ENV),production)
  NETWORK := xdc
endif

.PHONY: test
test: ## Run all Hardhat tests
	npx hardhat test

.PHONY: test-gas
test-gas: ## Run tests with gas reporting
	REPORT_GAS=true npx hardhat test

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	npx hardhat coverage

.PHONY: compile
compile: ## Compile all contracts
	npx hardhat compile

.PHONY: node
node: ## Start local Hardhat node
	npx hardhat node

.PHONY: deploy-mocks
deploy-mocks: ## Deploy mock USDC/USDT (local only)
	npx hardhat run scripts/deploy-mocks.ts --network $(NETWORK)

.PHONY: deploy
deploy: ## Deploy all protocol contracts [ENV=local|staging|production]
	@echo "Deploying to $(NETWORK)..."
	npx hardhat run scripts/deploy-proxy.ts    --network $(NETWORK)
	npx hardhat run scripts/deploy-phase3.ts   --network $(NETWORK)
	npx hardhat run scripts/deploy-phase4.ts   --network $(NETWORK)
	npx hardhat run scripts/deploy-phase5.ts   --network $(NETWORK)

.PHONY: verify
verify: ## Verify contracts on block explorer [ENV=staging|production]
	npx hardhat run scripts/verify.ts --network $(NETWORK)

.PHONY: lint
lint: ## Lint Solidity files
	@npx solhint 'src/**/*.sol' 2>/dev/null || echo "solhint not installed, skipping"

.PHONY: clean
clean:
	npx hardhat clean
