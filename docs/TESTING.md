# AI WATT — Testing Guide

End-to-end guide for testing the full protocol stack: smart contracts, backend API, Veriflow agent, and frontend dApp.

---

## Table of Contents

1. [Prerequisites](#1-prerequisites)
2. [Repo overview](#2-repo-overview)
3. [Local environment](#3-local-environment)
   - 3.1 [Infrastructure (MySQL + Redis)](#31-infrastructure-mysql--redis)
   - 3.2 [Smart contracts on Hardhat](#32-smart-contracts-on-hardhat)
   - 3.3 [Backend API](#33-backend-api)
   - 3.4 [Frontend dApp](#34-frontend-dapp)
4. [Running the test suites](#4-running-the-test-suites)
   - 4.1 [Contracts (Hardhat)](#41-contracts-hardhat)
   - 4.2 [Backend (Go)](#42-backend-go)
   - 4.3 [Frontend (TypeScript)](#43-frontend-typescript)
   - 4.4 [All at once](#44-all-at-once)
5. [Manual end-to-end flow (local)](#5-manual-end-to-end-flow-local)
6. [Staging environment (XDC Apothem)](#6-staging-environment-xdc-apothem)
   - 6.1 [Wallet and gas setup](#61-wallet-and-gas-setup)
   - 6.2 [Deploy contracts to Apothem](#62-deploy-contracts-to-apothem)
   - 6.3 [Configure the backend](#63-configure-the-backend)
   - 6.4 [Configure the frontend](#64-configure-the-frontend)
   - 6.5 [Manual test flow (staging)](#65-manual-test-flow-staging)
7. [CI/CD pipeline](#7-cicd-pipeline)
8. [Troubleshooting](#8-troubleshooting)

---

## 1. Prerequisites

Install these tools before starting:

| Tool | Version | Install |
|---|---|---|
| Node.js | 20+ | `nvm install 20` |
| Go | 1.22+ | https://go.dev/dl |
| Docker + Docker Compose | any recent | https://docs.docker.com/get-docker |
| go-migrate CLI | latest | `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest` |
| MetaMask | any | https://metamask.io |

> Run `make install` from the repo root to install all Node and Go dependencies in one step.

---

## 2. Repo overview

```
aiwatt/
├── contracts/       # Solidity (Hardhat) — XDC Network
├── backend/         # Go API server — Gin, GORM, MySQL, Redis
├── veriflow-agent/  # Go binary — deployed on borrower hardware
├── frontend/        # React + Vite + wagmi dApp
└── scripts/         # deploy-backend.sh, health-check.sh
```

Each layer has its own test suite and can be tested independently.

---

## 3. Local environment

### 3.1 Infrastructure (MySQL + Redis)

The backend requires MySQL 8 and Redis 7. Start them with Docker:

```bash
docker compose up -d mysql redis
```

Wait until both containers are healthy:

```bash
docker compose ps
# mysql: healthy, redis: healthy
```

Or use the Makefile shortcut which waits automatically:

```bash
make dev-infra
```

### 3.2 Smart contracts on Hardhat

Open a dedicated terminal and start the local Hardhat node:

```bash
cd contracts
npx hardhat node
# Hardhat Network listening on http://127.0.0.1:8545
# Accounts and private keys are printed — save one for the deployer
```

In a second terminal, deploy the mock stablecoins and then all protocol contracts:

```bash
# From repo root
make deploy-mocks            # deploys MockUSDC + MockUSDT → prints addresses
make deploy-contracts-local  # deploys all phases → prints proxy addresses
```

Copy the printed contract addresses — you will need them for the backend and frontend env files.

### 3.3 Backend API

```bash
# 1. Create your env file
cp backend/.env.example backend/.env

# 2. Fill in the minimum required values:
#    DATABASE_URL=mysql://aiwatt:aiwatt@localhost:3306/aiwatt?parseTime=true
#    REDIS_URL=redis://localhost:6379/0
#    JWT_SECRET=any-string-at-least-32-chars-long
#    XDC_RPC_URL=http://localhost:8545
#    VERIFLOW_SIGNER_PRIVATE_KEY=<hardhat account private key>
#    WATT_USD_PROXY_ADDRESS=<from deploy-contracts-local output>
#    MINT_ENGINE_PROXY_ADDRESS=<from deploy-contracts-local output>
#    SWATT_USD_PROXY_ADDRESS=<from deploy-contracts-local output>

# 3. Run database migrations
make migrate-up ENV=local

# 4. Start the server
make dev-backend
# API running at http://localhost:8080
# Swagger UI at http://localhost:8080/swagger/index.html
```

### 3.4 Frontend dApp

```bash
# 1. Create your env file
cp frontend/.env.local.example frontend/.env.local

# 2. Fill in contract addresses (from deploy-contracts-local output):
#    VITE_USDC_ADDRESS=0x...
#    VITE_USDT_ADDRESS=0x...
#    VITE_WATT_ADDRESS=0x...
#    VITE_MINT_ENGINE_ADDRESS=0x...
#    VITE_SWATT_ADDRESS=0x...
#    VITE_CHAIN_ID=31337
#    VITE_RPC_URL=http://localhost:8545
#    VITE_API_URL=http://localhost:8080

# 3. Install deps and start
cd frontend
npm install
npm run dev
# Frontend at http://localhost:5173
```

**Connect MetaMask to Hardhat:**
- Network name: `Hardhat Local`
- RPC URL: `http://localhost:8545`
- Chain ID: `31337`
- Currency symbol: `ETH`

Import one of the Hardhat-printed private keys as a wallet — it starts with 10,000 ETH for gas.

---

## 4. Running the test suites

### 4.1 Contracts (Hardhat)

```bash
cd contracts

# Run all tests (258 tests across 4 phases)
npx hardhat test

# Run with gas cost report
npx hardhat test --reporter gas

# Run tests with coverage
npx hardhat coverage

# Run a single file
npx hardhat test test/MintEngine.test.ts
npx hardhat test test/assets/AssetRegistry.test.ts
```

**Test files:**

| File | What it covers |
|---|---|
| `test/WattUSD.test.ts` | WATT mint, burn, pause, roles |
| `test/sWattUSD.test.ts` | sWATT staking, rewards, lock periods |
| `test/MintEngine.test.ts` | USDC/USDT → WATT minting, fees, limits |
| `test/assets/AssetRegistry.test.ts` | GPU/robot registration, status transitions |
| `test/assets/OCNFT.test.ts` | On-chain NFT minting, URI, access control |
| `test/assets/HealthAttestation.test.ts` | Veriflow attestation writes and reads |
| `test/assets/Phase3Integration.test.ts` | End-to-end Phase 3 (register → attest → mint NFT) |
| `test/credit/LendingPool.test.ts` | Loan lifecycle, collateral, liquidation |
| `test/credit/WEVQueue.test.ts` | Work-energy-value queue priority, scoring |

All tests run against the in-process Hardhat network — no wallet or external connection needed.

### 4.2 Backend (Go)

The Go test suite needs MySQL and Redis running. Start infra first (`make dev-infra`), then:

```bash
cd backend

# Run all tests with race detector
go test -race -count=1 ./...

# Run tests for a specific package
go test -race ./internal/service/...
go test -race ./internal/repository/...
go test -race ./internal/api/handler/...

# Verbose output
go test -race -v ./internal/service/...

# Run a single test by name
go test -race -run TestLoanService_Create ./internal/service/...
```

The test suite requires these env vars (set automatically by the CI, set manually for local):

```bash
export DATABASE_URL=mysql://aiwatt:aiwatt@localhost:3306/aiwatt_test?parseTime=true
export REDIS_URL=redis://localhost:6379/0
export JWT_SECRET=ci-test-secret-32-chars-minimum-ok
export XDC_RPC_URL=http://localhost:8545
export APP_ENV=test
```

> The test database (`aiwatt_test`) is separate from the dev database. Create it and run migrations before the first run:
>
> ```bash
> mysql -u root -prootpassword -e "CREATE DATABASE IF NOT EXISTS aiwatt_test;"
> migrate -path scripts/migrations \
>   -database "mysql://aiwatt:aiwatt@localhost:3306/aiwatt_test?parseTime=true" up
> ```

### 4.3 Frontend (TypeScript)

The frontend has no runtime test suite yet — validation is through static analysis:

```bash
cd frontend

# TypeScript compiler — catches type errors across the full codebase
npx tsc --noEmit

# ESLint — catches style and logic issues
npm run lint

# Both at once
make check
```

### 4.4 All at once

From the repo root:

```bash
make test
# Runs: test-contracts → test-backend → test-frontend in sequence
```

---

## 5. Manual end-to-end flow (local)

With everything running (`make dev`), verify each feature in the browser:

**Faucet — get test tokens**
1. Open `http://localhost:5173/faucet`
2. Connect MetaMask (Hardhat network)
3. Click **Get 10,000 USDC** — approve the transaction in MetaMask
4. Balance should update after ~2 blocks (~4s on Hardhat)
5. Repeat for USDT

**Buy — mint WATT**
1. Navigate to `/buy`
2. Enter an amount of USDC (e.g. `1000`)
3. Review the confirm modal: paying / receiving / fee / rate
4. Confirm → MetaMask popup → approve
5. Check the toast notification and verify WATT balance increases

**Stake — deposit WATT**
1. Navigate to `/stake`
2. Select the **Stake** tab, enter an amount
3. Confirm and approve in MetaMask
4. Switch to **Unstake** tab — your staked balance should appear
5. Unstake — verify the lock-period message and sWATT balance update

**Borrow — open a loan**
1. Navigate to `/borrow`
2. Select a collateral engine (GPU, Robotics, or AI Energy)
3. Enter a loan amount below the displayed maximum
4. Confirm and approve — check the loan summary panel updates

**Dashboard — verify loan appears**
1. Navigate to `/dashboard`
2. Skeleton loaders display for ~900ms, then your active loan row appears
3. Verify loan status, amount, and collateral ratio match what you submitted

**Veriflow — check attestation cards**
1. Navigate to `/veriflow`
2. Cards load with health scores, GPU telemetry, and uptime percentages
3. On-chain hash and block number should appear (from HealthAttestation contract)

**Docs — verify contract addresses**
1. Navigate to `/docs`
2. All contracts deployed locally should have addresses (not "Pending deploy")
3. Copy button should copy the address to clipboard with a success toast

---

## 6. Staging environment (XDC Apothem)

The staging environment runs on **XDC Apothem testnet** (chainId 51). It is the closest environment to production — use it before any mainnet deployment.

### 6.1 Wallet and gas setup

1. Install MetaMask and create a deployer wallet (a fresh wallet is safer)
2. Add Apothem to MetaMask:
   - Network name: `XDC Apothem Testnet`
   - RPC URL: `https://erpc.apothem.network`
   - Chain ID: `51`
   - Currency symbol: `TXDC`
   - Block explorer: `https://explorer.apothem.network`
3. Get free testnet XDC (TXDC) from the Apothem faucet:
   - https://faucet.apothem.network
   - You need ~50 TXDC to deploy all contract phases
4. Export the deployer wallet's private key from MetaMask — you will set it as `DEPLOYER_PRIVATE_KEY`

### 6.2 Deploy contracts to Apothem

```bash
# 1. Set up contracts env
cp contracts/.env.example contracts/.env

# Fill in:
#   DEPLOYER_PRIVATE_KEY=0x<your-private-key>
#   TREASURY_ADDRESS=0x<your-treasury-wallet>
#   BACKEND_SIGNER_ADDRESS=0x<hot-wallet-used-by-backend>

# 2. Deploy all phases to Apothem
make deploy-contracts-staging

# The script prints proxy addresses after each phase.
# Copy them all — you'll need them for backend and frontend.

# 3. Optionally verify contracts on the block explorer
make verify-contracts-staging
```

After deploy, `contracts/.env` is auto-populated with the proxy addresses by the deploy scripts.

### 6.3 Configure the backend

On the staging VPS (or locally pointing at Apothem), update the backend env:

```bash
# /opt/aiwatt/staging.env (on the VPS) or backend/.env (local)

APP_ENV=staging
XDC_RPC_URL=https://erpc.apothem.network
XDC_CHAIN_ID=51

# Proxy addresses from deploy output
WATT_USD_PROXY_ADDRESS=0x974B2bd650c88D290469a471DeB1Ee6aC55AD2d9
MINT_ENGINE_PROXY_ADDRESS=0x37a1d3b072af3055993215DB82e3832dA2d10AEd
SWATT_USD_PROXY_ADDRESS=0x08c14aCc5547f03fC523453e005A243C50A7aa94

# Signer wallet private key (BACKEND_SIGNER_ADDRESS above, must have VERIFLOW_SIGNER role)
VERIFLOW_SIGNER_PRIVATE_KEY=0x...
```

Run migrations and restart:

```bash
make migrate-up ENV=staging
```

### 6.4 Configure the frontend

```bash
cp frontend/.env.staging.example frontend/.env.staging
# Edit .env.staging — contract addresses are pre-filled for known Apothem deployments
# Update VITE_WALLETCONNECT_PROJECT_ID with your project ID from cloud.walletconnect.com

# Build staging bundle
make build-frontend ENV=staging
```

The staging frontend build uses `--mode staging` which reads `frontend/.env.staging`.

### 6.5 Manual test flow (staging)

**1. Testnet Faucet**
- Open `https://staging.aiwatt.io/faucet` (or `localhost:5173/faucet` with staging env vars)
- Connect MetaMask on Apothem network
- Claim 10,000 mock USDC and 10,000 mock USDT
- Verify balances update on-chain (takes ~5–10s on Apothem)

**2. Full mint flow**
- Navigate to Buy — enter `500` USDC
- Approve USDC spending in MetaMask (ERC-20 approval tx)
- Confirm the mint transaction
- Wait for 2 block confirmations (~10–20s)
- Verify WATT balance increases in wallet and in the UI

**3. Stake and unstake**
- Stake 100 WATT — verify sWATT is received
- Check the lock period is applied (read from contract)
- Attempt early unstake — should show a lock-period error
- Advance time if testing locally, or wait the full period on staging

**4. Borrow**
- Select a collateral engine
- Enter a loan amount
- Confirm — verify loan appears in Dashboard
- Check on-chain: `https://explorer.apothem.network/address/<LENDING_POOL_PROXY_ADDRESS>`

**5. Veriflow telemetry**
- Run the Veriflow agent (or POST a mock payload):
  ```bash
  curl -X POST https://api-staging.aiwatt.io/api/v1/veriflow/telemetry \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer <jwt>" \
    -d '{"asset_id": "<asset-uuid>", "gpu_util": 87.3, "uptime": 99.1}'
  ```
- Open `/veriflow` — the asset's health card should appear with a score
- Verify the attestation was written to the HealthAttestation contract (check the block explorer)

**6. API health check**
```bash
# Full retry loop
ENV=staging ./scripts/health-check.sh

# Or manually
curl https://api-staging.aiwatt.io/api/v1/vault/stats
# Expect: {"success":true,"data":{...},"error":null}
```

---

## 7. CI/CD pipeline

Every `git push` to `main` triggers the CI workflow automatically. No manual steps needed unless deploying to production.

### Workflow overview

```
push to main ──► ci.yml ──────────────────────────────── parallel ──►
                  ├── frontend: tsc + lint + build (staging)
                  ├── backend:  vet + build + test (MySQL + Redis services)
                  ├── contracts: compile + hardhat test (258 tests)
                  └── docker:   docker build validation

                 if ci passes ──► deploy-staging.yml ──►
                  ├── build-backend: docker build + push to ghcr.io:staging-latest
                  ├── build-frontend: staging mode build + upload artifact
                  ├── deploy-backend: migrate → SSH → docker run
                  ├── deploy-frontend: rsync dist/ to VPS
                  └── health-check: 10× curl retry loop

push tag v*.*.* ──► deploy-production.yml ──►
                  ├── (same CI gate)
                  ├── build-backend: image tagged with semver + latest
                  ├── build-frontend: PROD_* secrets injected
                  ├── [MANUAL APPROVAL in GitHub UI] ──►
                  ├── deploy-backend: migrate → SSH → docker run
                  ├── deploy-frontend: rsync to production VPS
                  ├── health-check: 15× curl retry loop
                  └── GitHub Release: auto-generated release notes
```

### Trigger a staging deploy

```bash
git push origin main
# Watch progress at: github.com/neurowatt/aiwatt/actions
```

### Trigger a production deploy

```bash
git tag v1.0.0
git push origin v1.0.0
# Then go to GitHub → Actions → approve the "production" environment gate
```

### GitHub Secrets required

Configure these under `Settings → Secrets and variables → Actions`:

| Secret | Used by |
|---|---|
| `VITE_WALLETCONNECT_PROJECT_ID` | CI + both deploy workflows |
| `STAGING_SSH_HOST` | staging deploy |
| `STAGING_SSH_USER` | staging deploy |
| `STAGING_SSH_KEY` | staging deploy |
| `STAGING_DB_URL` | staging migrations |
| `STAGING_REDIS_URL` | staging backend env |
| `STAGING_JWT_SECRET` | staging backend env |
| `STAGING_SIGNER_PRIVATE_KEY` | staging backend env |
| `PROD_SSH_HOST` | production deploy |
| `PROD_SSH_USER` | production deploy |
| `PROD_SSH_KEY` | production deploy |
| `PROD_DB_URL` | production migrations |
| `PROD_USDC_ADDRESS` | production frontend build |
| `PROD_USDT_ADDRESS` | production frontend build |
| `PROD_WATT_ADDRESS` | production frontend build |
| `PROD_MINT_ENGINE_ADDRESS` | production frontend build |
| `PROD_SWATT_ADDRESS` | production frontend build |

---

## 8. Troubleshooting

**`migrate: no migration` or `dirty` database**

The migration state is dirty if a previous run partially applied. Force-reset to the last clean version:

```bash
# Check current version
migrate -path backend/scripts/migrations -database "$DATABASE_URL" version

# Force to a specific version (e.g. 9 if 10 failed)
migrate -path backend/scripts/migrations -database "$DATABASE_URL" force 9

# Re-run from there
migrate -path backend/scripts/migrations -database "$DATABASE_URL" up
```

**`contract call failed` / wrong chain**

The backend and frontend must point to the same network. Check:
- `XDC_RPC_URL` in `backend/.env` matches the deployed contract network
- `VITE_CHAIN_ID` in `frontend/.env.*` matches the MetaMask network
- Contract addresses in both env files are from the same deploy

**MetaMask shows wrong chain / nonce errors**

- Go to MetaMask → Settings → Advanced → Reset Account (clears local nonce cache)
- If using Hardhat, restart the node (`npx hardhat node`) and redeploy — Hardhat's block history resets on restart

**`hardhat test` fails with `provider not connected`**

You have a running `hardhat node` process that is conflicting. Run tests in the default in-process network instead:

```bash
# Tests use the in-process Hardhat network by default — no node needed
npx hardhat test
# NOT: npx hardhat test --network localhost
```

**`go test` fails with `connection refused` (MySQL/Redis)**

Start the infra containers first:

```bash
make dev-infra
# Wait for "MySQL ready" before running tests
```

**Frontend build fails with `VITE_* is not defined`**

Vite only injects env vars that exist in the `.env.*` file at build time. Make sure the correct env file exists and is populated:

```bash
# For local dev
ls frontend/.env.local      # must exist and have VITE_CHAIN_ID etc.

# For staging build
ls frontend/.env.staging

# Rebuild
make build-frontend ENV=staging
```

**Faucet button does nothing / wallet not connected**

- The Faucet page only works on testnet (chainId 51 for staging, 31337 for local)
- MetaMask must be on the correct network — the page shows the active chain at the top
- Switch network in MetaMask, then reload the page

**Attestation not appearing in Veriflow**

The backend signer wallet (`VERIFLOW_SIGNER_PRIVATE_KEY`) must hold the `VERIFLOW_SIGNER` role on the HealthAttestation contract. Grant it after deploy:

```typescript
// In a Hardhat script or console
const ha = await ethers.getContractAt('HealthAttestation', HEALTH_ATTESTATION_PROXY_ADDRESS)
const role = await ha.VERIFLOW_SIGNER_ROLE()
await ha.grantRole(role, BACKEND_SIGNER_ADDRESS)
```
