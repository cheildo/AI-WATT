# PLAN.md — AI WATT Build Tracker

Current phase and task tracking. Update this file as work progresses.

---

## Current Phase: Phase 11 — AttestationWriter

**Status: PENDING**

---

## Completed Phases

### Phase 0 — Scaffold ✅
- Monorepo structure: `backend/`, `contracts/`, `veriflow-agent/`, `frontend/`
- Go module initialized (`github.com/neurowatt/aiwatt-backend`), all deps installed
- Hardhat configured for local + Apothem (chainId:51) + XDC mainnet (chainId:50)
- Docker Compose: MySQL 8 + Redis 7
- Per-sub-repo `.env.example` and `README.md` files
- First go-migrate migration: `000001_create_users`
- Backend full layer scaffold: handler/dto/service/repository/model/blockchain/veriflow
- Veriflow agent standalone binary scaffold

### Phase 1 — WattUSD + MintEngine ✅
- `IWattUSD.sol` + `IMintEngine.sol` interfaces
- `WattUSD.sol` — ERC-20 synthetic dollar, 6 decimals, UUPS, role-gated mint/burn
- `MintEngine.sol` — deposit USDC/USDT -> mint WATT 1:1 minus 0.1% fee, redeem flow
- `MockStablecoin.sol` — testnet USDC/USDT with `faucet()` function
- 51 tests, all passing
- Deployed to XDC Apothem testnet
- Block explorer verification configured (testnet.xdcscan.com, Etherscan API key)

### Phase 10 — Frontend (React dApp) ✅
- **Scaffold**: Vite 6 + React 18 + TypeScript strict + Tailwind CSS 3 with AI WATT design tokens (surface/brand/yield/warn/danger). Inter + JetBrains Mono fonts. Path alias `@/*`.
- **Chain config** (`src/chains.ts`): `xdcMainnet` (chainId 50) + `xdcApothem` (chainId 51) via viem `defineChain`.
- **Providers** (`src/providers.tsx`): `WagmiProvider` (injected + WalletConnect connectors) + `QueryClientProvider` (staleTime 30s).
- **Stores**: `walletStore` (JWT persisted to localStorage), `txStore` (20-item pending tx ring buffer) — both Zustand.
- **Contract ABIs** (`src/contracts/abis.ts`): minimal ABI slices for ERC-20, MintEngine, sWattUSD, LendingPool, WEVQueue.
- **Contract addresses** (`src/contracts/addresses.ts`): read from `VITE_*_ADDRESS` env vars.
- **Contract hooks** (`src/hooks/contracts/`): useWattUSD (balance/allowance/mint/redeem), useSWattUSD (balance/NAV/vaultStats/stake/requestUnstake), useLendingPool (loan/borrowerLoans/originate/repay), useWEVQueue (status/userQueue/requestRedeem/cancelRedeem). All use wagmi `useReadContract`/`useWriteContract` + `useWaitForTransactionReceipt`.
- **API hooks** (`src/hooks/api/`): usePortfolio (loans + balances), useVeriflow (health score 60s refetch, attestation, assets), useActivity (chain events 30s refetch, paginated).
- **Components**: WalletButton, ActionButton (3 variants + spinner), AmountInput (MAX button), DetailsPanel, HealthBadge (colour-coded), HealthCard (score gradient bar), ActivityTable, ProtocolStats (on-chain TVL/NAV), NavBar, SideNav (NavLink active state), ToastProvider (Radix Toast + txStore).
- **Pages (8)**: Buy (mint/redeem WATT), Stake (stake/unstake + WEV queue), Borrow (engine selector E1/E2/E3 + loan details), Portfolio (balances + open loans, wallet-gated), Veriflow (asset health grid + recharts score trend line + attestation history), Activity (paginated event table), Governance (Phase 6 coming soon), Bridge (LayerZero coming soon).
- **Router**: React Router v6, `NavBar + SideNav` layout shell, `/` → `/buy`.
- `tsc --noEmit` and `vite build` both pass clean.
- **Skipped (deferred)**: React Testing Library component tests; Playwright E2E; WalletConnect projectId config; CORS `.env.local` setup doc.

### Phase 9 — Veriflow v1 (TelemetryAgent + IngestionService + ScoringEngine) ✅
- **TelemetryAgent (`veriflow-agent/`)**: Split collector into `nvidia_collector.go` (GPU util/temp/memory/power/ECC uncorrected errors), `system_collector.go` (ipmitool fan + /proc/uptime → UptimePct), `metrics.go` (shared struct). Reporter fixed to sign JSON payload bytes (not formatted string), retry 3x with exponential backoff, structured zap logging. `cmd/main.go` config struct with `REPORT_INTERVAL` env var. `deploy/veriflow-agent.service` systemd template with security hardening. `Makefile` with `build-agent → dist/veriflow-agent-linux-amd64`.
- **Models + Migrations**: `assets.hmac_secret` (migration 000009); `telemetry.ecc_errors` + `uptime_pct` (migration 000010).
- **TelemetryRepo**: Added `GetLastN(ctx, assetID, n)` for rolling scoring window.
- **ScoringEngine (`internal/veriflow/scorer.go`)**: 12-row rolling window (~60 min), 4 scoring rules (25 pts each): GPU util (≥70%/≥40%), temperature (≤75°C/≤85°C), ECC error rate (count/10000, thresholds 0.01%/0.1%), UptimePct (≥99.5%/≥95%). Heartbeat check: last row >15 min → score=0.
- **IngestionService (`internal/service/veriflow_service.go`)**: Looks up `asset.HMACSecret` from DB (no more X-HMAC-Secret header), validates asset ACTIVE, verifies HMAC over unsigned JSON bytes, inserts telemetry, fires async `scoreAndAct`. Post-score: persists health score, sends `NotifyService.Send` alert if score <60, calls `TxManager.UpdateAssetStatusOnChain(FLAGGED)` if score <40 and asset ACTIVE.
- **Skipped (deferred)**: Unit tests for collector (mock nvidia-smi), ScoringEngine rule tests; Phase 11 on-chain attestation write.

### Phase 8 — Blockchain Layer (Golang) ✅
- 8 ABI JSON files extracted from Hardhat artifacts into `internal/blockchain/abis/`
- Go bindings generated with `abigen` for all contracts: WattUSD, sWattUSD, MintEngine, AssetRegistry, OCNFT, HealthAttestation, LendingPool, WEVQueue → `internal/blockchain/contracts/`
- `BlockchainClient`: dials XDC RPC, initialises all 8 contract binding instances, exposes `GetLatestBlock`, `GetTransactionReceipt`, `AllAddresses`, `NAVPerShare`, `IsAssetActive`
- `EventIndexer`: backfills from Redis `indexer:last_block` to chain head on startup, then live `SubscribeFilterLogs`; dispatches by contract address to typed parser; persists all key events to `chain_events` table via `EventRepo.Create`; updates `indexer:last_block` in Redis after each event
- `TxManager`: Redis-cached nonce with mutex (prevents race on concurrent sends), `SuggestGasPrice`, receipt polling with 30s timeout, auto-retry on nonce errors; exposes `MintOCNFT`, `RegisterAssetOnChain`, `UpdateLTVOnChain`, `UpdateAssetStatusOnChain`, `SubmitAttestation`
- Contract addresses loaded from env vars (all `_PROXY_ADDRESS` env vars), `XDC_CHAIN_ID` defaults to 51 (Apothem)
- `AssetService`: injected with `TxManager`; `Register` triggers async `RegisterAssetOnChain` + `MintOCNFT` after DB insert; `UpdateLTV` triggers async `UpdateLTVOnChain`
- `LoanService`: injected with `BlockchainClient` + `assetRepo`; `Create` verifies `AssetRegistry.isActive(assetId32)` on-chain before creating the DB record
- `main.go`: initialises `BlockchainClient`, `TxManager`, `EventIndexer`; starts indexer as background goroutine; all constructors updated
- `go build ./...` passes clean
- **Skipped (deferred):** integration tests against local Hardhat node; TxManager nonce-retry unit tests

### Phase 7 — Backend API (Golang) ✅
- Migrations 000002–000008: assets, loans, repayments, chain_events, telemetry (partitioned), attestations, wev_queue — all with working `.down.sql`
- Models: repayment, attestation, wev_queue (added to existing user/asset/loan/telemetry/chain_event)
- Repositories: user, loan, asset, telemetry, repayment, attestation, wev, event — all backed by GORM
- Services: UserService (EIP-191 wallet login + bcrypt email login + JWT), LoanService, AssetService, MintService (stub), YieldService (stub), WEVService, VeriflowService (HMAC verify + telemetry insert), NotifyService, TreasuryService (stub)
- DTOs: user, loan, asset, mint, yield, veriflow, wev — all with binding/json tags and Swagger annotations
- Handlers: user, loan, asset (+ GetHealth), mint (+ GetVaultStats injected via YieldService), veriflow, wev (new) — Swagger godoc on every endpoint
- Middleware: JWT auth, role-based access, CORS (origin allowlist), Redis rate limiter (60 req/60s per IP, fail-open), request logger
- Router: all routes registered under `/api/v1/` with correct middleware chains; WEV, vault/stats, asset health routes added
- main.go: GORM MySQL init (connection pool), Redis init (ping check), all repos instantiated and wired into services, `ALLOWED_ORIGINS` env var for CORS
- `go build ./...` passes clean
- **Skipped (deferred):** `swag init` Swagger doc generation; service unit tests and handler httptest tests

### Phase 5 — WEVQueue.sol ✅
- `IWEVQueue.sol` + `WEVQueue.sol`
- Standard queue (30-day) and priority queue (3-day, 0.5% WATT fee to jump ahead)
- PROCESSOR_ROLE keeper calls `processBatch` → redeems sWATT via `sWattUSD.redeem`, sends WATT to users
- Priority fee held in contract, added to protocolFees on fulfillment, refunded on cancel
- sWattUSD updated: `maxWithdraw`/`maxRedeem`/`_withdraw` exempt WEVQueue address as owner/caller so batch can bypass WEV threshold
- Inline reentrancy guard (same pattern as LendingPool)
- 48 tests (258 total), including sWattUSD WEV guard bypass integration test
- `_nextRequestId()` uses monotonic nonce — no requestId collision even in same block

### Phase 4 — LendingPool.sol ✅
- `ILendingPool.sol` + `LendingPool.sol`
- Pre-origination gates: asset ACTIVE in AssetRegistry, attestation < 48h, score >= 60, no double-encumbrance
- Pro-rata interest/principal split on repayment: 90% → sWattUSD.receiveYield(), 10% → protocol fees
- Inline reentrancy guard (`uint8 _reentrancyStatus`) — OZ ReentrancyGuard has constructor rejected by hardhat-upgrades
- fullRepay, flagDefaulted (permissionless post-maturity), liquidate (LIQUIDATOR_ROLE → AssetRegistry LIQUIDATED)
- `_settle()` clears asset lock, resets AssetRegistry to ACTIVE, emits LoanSettled
- 43 tests (210 total), including end-to-end lifecycle integration test
- Deploy script grants LENDINGPOOL_ROLE/MINTER_ROLE/YIELD_DISTRIBUTOR_ROLE to LendingPool address

### Phase 3 — AssetRegistry + OCNFT + HealthAttestation ✅
- `IAssetRegistry.sol`, `IOCNFT.sol`, `IHealthAttestation.sol` interfaces
- `AssetRegistry.sol` — PENDING/ACTIVE/FLAGGED/LIQUIDATED lifecycle, LTV in basis points, REGISTRAR_ROLE + LENDINGPOOL_ROLE + ADMIN_ROLE
- `OCNFT.sol` — ERC-721 soulbound NFT, 1:1 tokenId↔assetId, MINTER_ROLE settlement transfer via `address(0)` auth bypass
- `HealthAttestation.sol` — keccak256 health hash + score 0-100, 12h cooldown, newest-first history
- 76 tests (167 total), integration test: register → activate → mint OC-NFT → attest → LendingPool readiness check
- OCNFT soulbound fix: pass `address(0)` as auth to `super._update` to bypass OZ ownership check for MINTER_ROLE
- Deploy script grants REGISTRAR_ROLE/MINTER_ROLE/VERIFLOW_SIGNER to BACKEND_SIGNER_ADDRESS
- LENDINGPOOL_ROLE on AssetRegistry deferred to Phase 4 deploy script

### Phase 2 — sWattUSD ERC-4626 Yield Vault ✅
- `IsWattUSD.sol` interface
- `sWattUSD.sol` — ERC-4626 vault, WATT as underlying asset, NAV per share rises via `receiveYield()`
- WEV threshold guard (`maxWithdraw`/`maxRedeem` capped) — WEVQueue stub for Phase 5
- Inflation attack protection: deploy script seeds 1 WATT initial deposit
- 40 tests passing (91 total across all phases)
- `evmVersion: cancun` set in Hardhat config (required for OZ v5 `mcopy` opcode)
- Deploy to Apothem: run `deploy-proxy.ts` then `verify.ts` (addresses go in `contracts/.env`)

---

## Phase Roadmap

| Phase | Description | Status |
|---|---|---|
| 0 | Monorepo scaffold, tooling, config | ✅ Complete |
| 1 | WattUSD.sol + MintEngine.sol + tests — deployed to Apothem | ✅ Complete |
| 2 | sWattUSD.sol ERC-4626 vault — NAV, deposit, withdraw | ✅ Complete |
| 3 | AssetRegistry.sol + OCNFT.sol + HealthAttestation.sol | ✅ Complete |
| 4 | LendingPool.sol — full loan lifecycle (Engine 2) | ✅ Complete |
| 5 | WEVQueue.sol — sWATT redemption queue | ✅ Complete |
| 6 | Governor.sol + Timelock.sol + WattToken.sol ($WATT) | 🔄 In progress |
| 7 | Backend API — auth, core services, all DTOs, Swagger | ⬜ Pending |
| 8 | Blockchain layer — BlockchainClient, EventIndexer, TxManager | ⬜ Pending |
| 9 | Veriflow v1 — TelemetryAgent + IngestionService + ScoringEngine | ✅ Complete |
| 10 | Frontend — React dApp, wallet connect, all pages | ✅ Complete |
| 11 | AttestationWriter — daily hash writes to XDC | ⬜ Pending |
| 12 | TreasuryService — Engine 3 idle capital sweep | ⬜ Pending |
| 13 | Engine 1 — Pre-delivery PO financing | ⬜ Pending |
| Future | Multi-chain bridge via LayerZero OFT | ⬜ Pending |

---

## Detailed Phase Tasks

---

### Phase 4 — LendingPool.sol

**Goal:** On-chain loan lifecycle for Engine 2 (post-delivery productivity-backed loans). Reads AssetRegistry for LTV and HealthAttestation to verify asset health before disbursement.

**Contract: `ILendingPool.sol`**
- Define all external function signatures before implementation

**Contract: `LendingPool.sol`**
- UUPS upgradeable, `AccessControlUpgradeable`, `ReentrancyGuardUpgradeable`, `PausableUpgradeable`
- Roles: `CURATOR_ROLE` (loan originators), `LIQUIDATOR_ROLE`, `ADMIN_ROLE`
- Struct: `Loan { bytes32 loanId, bytes32 assetId, address borrower, address curator, uint256 principal, uint256 outstanding, uint256 interestRate, LoanStatus status, uint8 engineType, uint256 originatedAt, uint256 maturityAt }`
- Enum: `LoanStatus { PENDING, ACTIVE, REPAYING, SETTLED, DEFAULTED, LIQUIDATED }`
- Key functions:
  - `originateLoan(bytes32 assetId, address borrower, uint256 principal, uint256 interestRate, uint256 termDays, uint8 engineType)` — checks `AssetRegistry.isActive(assetId)`, checks latest `HealthAttestation` score >= 60 and not stale (< 48h), mints WATT to borrower
  - `repay(bytes32 loanId, uint256 amount)` — splits into principal + interest, routes interest to `sWattUSD.receiveYield()`, reduces outstanding
  - `fullRepay(bytes32 loanId)` — settles loan, sets status SETTLED, updates AssetRegistry status to ACTIVE
  - `liquidate(bytes32 loanId)` — LIQUIDATOR_ROLE only, sets status LIQUIDATED, calls AssetRegistry to mark LIQUIDATED
  - `flagDefaulted(bytes32 loanId)` — when maturity passed and loan not repaid
  - `getLoan(bytes32 loanId)`, `getBorrowerLoans(address borrower)`, `getProtocolFees()`, `withdrawFees(address to)` (ADMIN_ROLE)
- Events: `LoanOriginated`, `RepaymentReceived`, `LoanSettled`, `LoanLiquidated`, `LoanDefaulted`
- Custom errors: `AssetNotActive`, `AttestationStale`, `HealthScoreTooLow`, `InsufficientBalance`, `LoanNotActive`, `UnauthorizedCurator`
- Fork reference: MetaStreet Pool.sol accounting patterns — adapt for RWA (no NFT price oracle needed)

**Hardhat tests (`test/credit/LendingPool.test.ts`)**
- Originate loan — happy path
- Originate reverts when asset not in AssetRegistry
- Originate reverts when health score < 60 (mock low attestation)
- Originate reverts when attestation stale > 48h
- Partial repayment — interest splits correctly to sWattUSD.receiveYield()
- Full repayment — status SETTLED, AssetRegistry status resets to ACTIVE
- Liquidation — LIQUIDATOR_ROLE triggers, AssetRegistry updated
- Role guard: non-curator cannot originate
- Integration: AssetRegistry register -> OCNFT mint -> HealthAttestation submit -> originate -> repay -> settle

**Deploy script `scripts/deploy-phase4.ts`**
- Deploy LendingPool proxy, wire to AssetRegistry + HealthAttestation + sWattUSD + WattUSD from `.env`
- Grant `LENDINGPOOL_ROLE` on AssetRegistry to LendingPool address
- Save proxy address to `.env`

---

### Phase 5 — WEVQueue.sol

**Goal:** sWATT redemption queue. Users enter queue, backend keeper processes batch redemptions on schedule. Priority auction for faster exit.

**Contract: `IWEVQueue.sol`**

**Contract: `WEVQueue.sol`**
- UUPS upgradeable, `AccessControlUpgradeable`, `ReentrancyGuardUpgradeable`
- Roles: `PROCESSOR_ROLE` (backend keeper wallet), `ADMIN_ROLE`
- Struct: `RedemptionRequest { bytes32 requestId, address user, uint256 sWattAmount, uint256 priorityFee, uint256 requestedAt, RequestStatus status }`
- Enum: `RequestStatus { QUEUED, PROCESSING, FULFILLED, CANCELLED }`
- Key functions:
  - `requestRedeem(uint256 sWattAmount)` — standard queue, ~30-day estimated wait, pulls sWATT from user
  - `requestPriorityRedeem(uint256 sWattAmount, uint256 priorityFee)` — pays 0.5% fee to jump queue, ~3-day wait
  - `cancelRequest(bytes32 requestId)` — user cancels while QUEUED, returns sWATT
  - `processBatch(bytes32[] calldata requestIds)` — PROCESSOR_ROLE only, burns sWATT, releases WATT to users
  - `getQueueDepth()`, `getUserRequests(address user)`, `getRequest(bytes32 requestId)`, `nextProcessingTimestamp()`
- Wire `sWattUSD.sol` maxWithdraw/maxRedeem to delegate to WEVQueue availability — replace existing stub from Phase 2
- Events: `RedemptionRequested`, `RedemptionFulfilled`, `RedemptionCancelled`, `BatchProcessed`
- Custom errors: `RequestNotFound`, `NotQueued`, `NothingToProcess`, `InsufficientPriorityFee`

**Hardhat tests (`test/credit/WEVQueue.test.ts`)**
- Standard queue entry — sWATT pulled, request created
- Priority queue — fee charged, ordering correct vs standard entries
- Cancel request — sWATT returned
- processBatch — sWATT burned, WATT released to users
- Non-PROCESSOR_ROLE processBatch reverts
- Queue depth accurate across multiple requests and cancellations
- Integration with sWattUSD maxRedeem guard

**Deploy script `scripts/deploy-phase5.ts`**
- Deploy WEVQueue proxy, wire to sWattUSD + WattUSD
- Call `sWattUSD.setWEVQueue(wevQueueAddress)` to activate real queue manager
- Grant `PROCESSOR_ROLE` to keeper wallet, save address to `.env`

---

### Phase 6 — WattToken.sol + Governor.sol + Timelock.sol ⏸ (deferred — do later)

**Goal:** On-chain governance. $WATT holders vote on parameter changes, curator approvals, contract upgrades. All upgrades gated by 48h timelock.

**Contract: `WattToken.sol`**
- UUPS upgradeable, ERC-20Upgradeable, ERC-20VotesUpgradeable, `AccessControlUpgradeable`
- Fixed supply: 1,000,000,000 $WATT minted at initialize() to treasury multisig — no inflation after TGE
- Vesting: team/investor allocations via OZ `VestingWallet.sol` deployed at TGE
- Functions: standard ERC-20 + `delegate(address)`, `getVotes(address)`, `getPastVotes(address, uint256 timepoint)`

**Contract: `AiWattTimelock.sol`**
- Extends `TimelockControllerUpgradeable`
- `minDelay`: 48 hours (172800 seconds)
- `proposers`: Governor contract only
- `executors`: address(0) (anyone can execute after delay)
- Admin renounced after setup

**Contract: `AiWattGovernor.sol`**
- Extends OZ Governor + settings + counting + votes + quorum + timelock control
- Token: WattToken
- `votingDelay`: 1 day, `votingPeriod`: 3 days
- `proposalThreshold`: 100,000 $WATT (0.01% of supply)
- `quorumNumerator`: 4 (4% of circulating supply)
- Proposal types: fee changes, LTV threshold updates, curator whitelisting, contract upgrades

**Hardhat tests (`test/governance/`)**
- `WattToken.test.ts`: mint, transfer, delegation, voting weight, past votes
- `Governor.test.ts`: full lifecycle — propose -> vote -> queue -> execute; quorum not met reverts; timelock delay enforced
- `Timelock.test.ts`: only Governor can propose, early execution reverts

**Deploy script `scripts/deploy-phase6.ts`**
- Deploy WattToken -> AiWattTimelock -> AiWattGovernor in sequence
- Transfer admin of all prior contracts (WattUSD, MintEngine, sWattUSD, AssetRegistry, OCNFT, HealthAttestation, LendingPool, WEVQueue) to Timelock
- Deployer renounces ADMIN_ROLE on all contracts after handover

---

### Phase 7 — Backend API (Golang)

**Goal:** Full REST API. All services, repositories, handlers, DTOs, Swagger docs, middleware. Connected to MySQL and Redis. No blockchain calls yet — those come in Phase 8.

**Migrations (`scripts/migrations/`)**
- `000001_create_users.up.sql` — already exists (Phase 0 scaffold)
- `000002_create_assets.up.sql` — columns: id, asset_id (bytes32 hex), asset_type, borrower_wallet, health_score, ltv, status, loan_id, location, hmac_secret, registered_at, updated_at
- `000003_create_loans.up.sql` — columns: id, loan_id (bytes32 hex), asset_id, borrower_id (FK users), curator_id (FK users), engine_type (1/2/3), principal, outstanding, interest_rate, status, originated_at, maturity_at, settled_at
- `000004_create_repayments.up.sql` — columns: id, loan_id, amount, tx_hash, paid_at
- `000005_create_chain_events.up.sql` — columns: id, event_type, contract_address, tx_hash, block_number, log_index, args_json, created_at
- `000006_create_telemetry.up.sql` — columns: id, asset_id, gpu_utilization, temperature, error_rate, uptime_pct, raw_json, recorded_at — partition by (asset_id, date)
- `000007_create_attestations.up.sql` — columns: id, asset_id, health_score, health_hash, xdc_tx_hash, attested_at
- `000008_create_wev_queue.up.sql` — columns: id, request_id, user_id, swatt_amount, priority_fee, status, requested_at, processed_at
- All migrations must have working .down.sql

**Models (`internal/models/`)**
- user.go, asset.go, loan.go, repayment.go, chain_event.go, telemetry.go, attestation.go, wev_queue.go
- All GORM structs with correct column tags, indexes, and foreign key relationships
- BeforeCreate hook on all models to auto-generate UUID if ID is empty

**Repositories (`internal/repository/`)**
- `user_repo.go`: Create, GetByID, GetByWallet, UpdateKYCStatus, List
- `loan_repo.go`: Create, GetByID, GetByLoanID, GetByBorrower, UpdateStatus, ListActive
- `asset_repo.go`: Create, GetByAssetID, UpdateHealthScore, UpdateStatus, ListByBorrower, ListActive
- `repayment_repo.go`: Create, GetByLoan, SumByLoan
- `telemetry_repo.go`: Insert, GetLatestByAsset, GetHistoryByAsset(assetID, from, to time.Time)
- `attestation_repo.go`: Create, GetLatestByAsset, GetHistory(assetID, limit)
- `wev_repo.go`: Create, GetByRequestID, GetByUser, UpdateStatus, GetPendingBatch
- `event_repo.go`: Insert, GetByType, GetByTxHash

**Services (`internal/service/`)**
- `user_service.go`: Login(wallet, signature, message) — verify EIP-4361 SIWE signature, issue JWT; GetProfile, UpdateKYC, GetRole
- `loan_service.go`: CreateApplication, ApproveLoan, RecordRepayment, GetLoan, ListByBorrower, FlagDefault, Liquidate
- `mint_service.go`: GetMintQuote, RecordMint, RecordRedeem, GetSupplyStats
- `yield_service.go`: GetCurrentNAV, GetAPR(period), GetVaultStats, DistributeYield
- `asset_service.go`: RegisterAsset, GetAsset, ListActive, UpdateHealthScore, PrepareOCNFTMetadata (uploads JSON to IPFS via Pinata, returns CID)
- `wev_service.go`: EnqueueRedemption, CancelRedemption, GetQueueStatus, GetUserQueue
- `treasury_service.go` (stub — full impl Phase 12): GetIdleCapital, GetTBillYield
- `notify_service.go`: SendLoanAlert, SendHealthAlert — email + webhook delivery

**DTOs (`internal/api/dto/`)**
- `user_dto.go`: LoginRequest { wallet, signature, message }, UserResponse { id, wallet, kyc_status, role, created_at }
- `loan_dto.go`: CreateLoanRequest { asset_id, amount, engine_type, term_days }, LoanResponse { loan_id, asset_id, borrower, principal, outstanding, status, engine_type, originated_at, maturity_at }, ListLoansResponse { loans[], total, page }
- `asset_dto.go`: RegisterAssetRequest { asset_type, serial_number, specs_json, location, borrower_wallet }, AssetResponse { asset_id, asset_type, health_score, ltv, status, location }, ListAssetsResponse
- `mint_dto.go`: MintRequest { amount, token }, RedeemRequest { amount }, MintResponse { watt_amount, fee, tx_hash }, SupplyStatsResponse { watt_supply, nav_per_swatt, total_deposited, utilization_rate }
- `yield_dto.go`: VaultStatsResponse { nav, total_assets, deployed_pct, t_bill_reserve, apr_7d, apr_30d }
- `veriflow_dto.go`: TelemetryPayload { asset_id, gpu_utilization, temperature, error_rate, uptime_pct, raw_metrics, hmac_signature, timestamp }, HealthScoreResponse { asset_id, score, status, last_attested }
- `wev_dto.go`: RedemptionRequest { swatt_amount, priority bool }, RedemptionResponse { request_id, estimated_days, status }, QueueStatusResponse { depth_swatt, next_processing, standard_days, priority_days }

**Handlers (`internal/api/handler/`)**
- `user_handler.go`: POST /api/v1/auth/login, GET /api/v1/users/me, PUT /api/v1/users/kyc
- `loan_handler.go`: POST /api/v1/loans, GET /api/v1/loans/:id, GET /api/v1/loans, POST /api/v1/loans/:id/repay, POST /api/v1/loans/:id/approve (CURATOR_ROLE)
- `asset_handler.go`: POST /api/v1/assets, GET /api/v1/assets/:assetId, GET /api/v1/assets, GET /api/v1/assets/:assetId/health
- `mint_handler.go`: GET /api/v1/mint/quote, GET /api/v1/mint/stats, GET /api/v1/vault/stats
- `veriflow_handler.go`: POST /api/v1/veriflow/telemetry (agent posts here), GET /api/v1/veriflow/assets, GET /api/v1/veriflow/assets/:assetId
- `wev_handler.go`: POST /api/v1/wev/redeem, DELETE /api/v1/wev/redeem/:requestId, GET /api/v1/wev/queue, GET /api/v1/wev/queue/me

**Middleware (`internal/api/middleware/`)**
- `auth.go`: JWT validation, extract user, attach to Gin context
- `wallet_auth.go`: SIWE signature verification for login endpoint
- `role.go`: RequireRole(role string) middleware factory
- `rate_limit.go`: Redis-backed — 100 req/min per IP, 500 req/min per wallet
- `logger.go`: zap request logger — method, path, status, latency, wallet
- `cors.go`: allow configured origins

**Swagger**
- Annotate every handler: @Summary, @Param, @Success, @Failure, @Router
- Run `swag init -g cmd/api/main.go -o api/` and commit generated docs
- Verify Swagger UI at http://localhost:8080/swagger/index.html

**Tests**
- Unit tests for all service functions with mocked repositories
- Handler tests for all endpoints with httptest + mock services — verify status codes and response shapes

---

### Phase 8 — Blockchain Layer (Golang)

**Goal:** Connect backend to XDC Network. Go contract bindings, event indexer, transaction manager.

**Contract Bindings (`internal/blockchain/contracts/`)**
- Generate Go bindings with abigen for all deployed contracts:
  - WattUSD.go, MintEngine.go, sWattUSD.go
  - AssetRegistry.go, OCNFT.go, HealthAttestation.go
  - LendingPool.go, WEVQueue.go
  - WattToken.go, AiWattGovernor.go
- Command: `abigen --abi=abis/ContractName.json --pkg=contracts --out=internal/blockchain/contracts/contract_name.go`
- Store ABIs in `internal/blockchain/abis/`
- Store deployed proxy addresses in config (loaded from env)

**BlockchainClient (`internal/blockchain/client.go`)**
- Wraps go-ethereum ethclient.Client
- NewBlockchainClient(rpcURL string)
- Initializes all contract binding instances from config addresses
- GetLatestBlock(), GetTransactionReceipt(txHash), WatchBlocks(ctx, chan)

**EventIndexer (`internal/blockchain/indexer.go`)**
- On startup: backfill historical events from last indexed block (Redis key `indexer:last_block`) to latest
- Live: SubscribeFilterLogs for all contract addresses
- For each event: parse log into typed struct -> call appropriate repository insert
- Events to index:
  - WattUSD: Transfer, Mint, Burn
  - MintEngine: WATTMinted, WATTRedeemed
  - sWattUSD: Deposit, Withdraw, YieldReceived
  - AssetRegistry: AssetRegistered, LTVUpdated, StatusChanged
  - OCNFT: OCNFTMinted, OCNFTBurned
  - HealthAttestation: AttestationSubmitted
  - LendingPool: LoanOriginated, RepaymentReceived, LoanSettled, LoanLiquidated
  - WEVQueue: RedemptionRequested, BatchProcessed
- All events stored in chain_events MySQL table via event_repo
- Updates `indexer:last_block` in Redis after each batch

**TxManager (`internal/blockchain/tx_manager.go`)**
- Backend hot wallet — private key from VERIFLOW_SIGNER_PRIVATE_KEY env var
- SendTransaction(ctx, contractABI, method string, args ...interface{})
- Nonce management: Redis-cached nonce with mutex lock (prevents race on concurrent writes)
- Gas: EstimateGas + 20% buffer
- Retry: on "nonce too low" error, refresh from chain, retry up to 3x
- Receipt confirmation: poll until confirmed or 30s timeout
- Exposed write functions:
  - MintOCNFT(to, assetId, metadataURI)
  - RegisterAssetOnChain(assetId, assetType, initialLTV)
  - UpdateLTVOnChain(assetId, newLTV)
  - UpdateAssetStatusOnChain(assetId, status)
  - SubmitAttestation(assetId, healthHash, score) — used by Phase 11

**Wire into services**
- asset_service.go: after DB record created -> call TxManager.RegisterAssetOnChain() + TxManager.MintOCNFT()
- loan_service.go: on loan approval -> verify AssetRegistry.isActive() via BlockchainClient read first

**Tests**
- Integration tests against Hardhat local node (npx hardhat node)
- Test EventIndexer: emit event on local node -> verify row inserted in DB
- Test TxManager: nonce handling, retry logic with mock RPC errors

---

### Phase 9 — Veriflow v1 (TelemetryAgent + IngestionService + ScoringEngine)

**Goal:** Hardware intelligence layer. TelemetryAgent binary on borrower servers collects GPU metrics and posts signed payloads to backend. Backend ingests, validates, and scores each asset.

**TelemetryAgent (`veriflow-agent/`)**
- Standalone Go binary — CGO_ENABLED=0, single static binary
- Config: env vars or agent.yaml — ASSET_ID, BACKEND_URL, HMAC_SECRET, REPORT_INTERVAL (default 5 min)
- Collector (internal/collector/):
  - `nvidia_collector.go`: shells out to `nvidia-smi --query-gpu=utilization.gpu,temperature.gpu,memory.used,memory.total,ecc.errors.uncorrected.total --format=csv,noheader` — parse CSV output
  - `system_collector.go`: `ipmitool sensor` for power/fan; /proc/loadavg for system load
  - `heartbeat.go`: records last successful collection timestamp
- Reporter (internal/reporter/):
  - Builds TelemetryPayload struct
  - Signs with HMAC-SHA256 using HMAC_SECRET — sets X-HMAC-Signature header
  - POST /api/v1/veriflow/telemetry — retry 3x on network failure with exponential backoff
  - Structured JSON logging via zap
- Systemd service template: veriflow-agent/deploy/veriflow-agent.service
- Build: `make build-agent` -> dist/veriflow-agent-linux-amd64

**IngestionService (`internal/veriflow/ingestor.go`)**
- Implements logic behind POST /api/v1/veriflow/telemetry handler (handler scaffolded in Phase 7):
  1. Verify HMAC: recompute HMAC-SHA256(payload_json, secret), compare to header -> 401 if mismatch
  2. Validate asset exists and is ACTIVE in asset_repo
  3. Insert raw telemetry row into telemetry table
  4. Trigger async: go ScoringEngine.ScoreAsync(assetID)
- HMAC secrets stored in assets.hmac_secret, provisioned at asset onboarding via AssetService

**ScoringEngine (`internal/veriflow/scorer.go`)**
- Score(assetID string) (HealthScore, error):
  1. Read last 12 telemetry rows (last 60 min at 5-min intervals)
  2. Apply rule-based scoring (0-100):
     - GPU utilization >= 70% -> +25 pts; 40-70% -> +15 pts; < 40% -> +0 pts
     - Temperature <= 75C -> +25 pts; 75-85C -> +15 pts; > 85C -> +0 pts
     - ECC error rate <= 0.01% -> +25 pts; 0.01-0.1% -> +10 pts; > 0.1% -> +0 pts
     - Uptime >= 99.5% -> +25 pts; 95-99.5% -> +15 pts; < 95% -> +0 pts
     - Heartbeat missed > 15 min -> score = 0 (overrides all other scores)
  3. Persist score to assets.health_score via asset_repo
  4. If score < 60 -> notify_service.SendHealthAlert()
  5. If score < 40 and loan ACTIVE -> TxManager.UpdateAssetStatusOnChain(assetID, FLAGGED)
  6. LTV recommendation: score >= 80 -> maintain; 60-79 -> reduce LTV 10%; < 60 -> flag for curator review
- GetHealthSummary(assetID) — last score, trend, last attested timestamp
- Runs async after each ingestion (goroutine) and on 1-hour cron schedule as safety net

**Tests**
- TelemetryAgent: unit test collectors with mocked nvidia-smi output; unit test HMAC signing + verification
- Backend: unit test ScoringEngine.Score() with seeded telemetry rows — verify all scoring rules
- Test threshold triggers: score drop < 60 fires notify; < 40 fires on-chain status update

---

### Phase 10 — Frontend (React dApp)

**Goal:** Full production React dApp. Wallet connect, all protocol interactions, Veriflow dashboard, portfolio, governance.

**Setup**
- Init with Vite: `npm create vite@latest . -- --template react-ts`
- Install: wagmi, viem, @tanstack/react-query, zustand, tailwindcss, @radix-ui/react-*, recharts
- Configure XDC Network as custom chain in Wagmi:
```typescript
const xdcMainnet = { id: 50, name: 'XDC Network', nativeCurrency: { name: 'XDC', symbol: 'XDC', decimals: 18 }, rpcUrls: { default: { http: ['https://rpc.xdcrpc.com'] } } }
const xdcApothem = { id: 51, name: 'XDC Apothem Testnet', nativeCurrency: { name: 'TXDC', symbol: 'TXDC', decimals: 18 }, rpcUrls: { default: { http: ['https://erpc.apothem.network'] } } }
```
- Configure React Query for all backend API calls
- Tailwind CSS with AI WATT design tokens in tailwind.config.ts

**Contract hooks (`src/hooks/contracts/`)**
- `useWattUSD.ts`: useMintWatt(amount), useRedeemWatt(amount), useWattBalance(address), useWattAllowance(owner, spender)
- `useSWattUSD.ts`: useStakeWatt(amount), useRequestUnstake(amount), useSWattBalance(address), useNAVPerShare(), useVaultStats()
- `useLendingPool.ts`: useOriginateLoan(params), useRepayLoan(loanId, amount), useLoan(loanId), useBorrowerLoans(address)
- `useWEVQueue.ts`: useRequestRedeem(amount, priority), useCancelRedeem(requestId), useQueueStatus(), useUserQueue(address)
- All hooks use wagmi's useReadContract, useWriteContract, useWaitForTransactionReceipt

**API hooks (`src/hooks/api/`)**
- usePortfolio.ts, useVeriflow.ts, useActivity.ts, useGovernance.ts

**Pages (`src/pages/`)**
- `Buy.tsx`: mint WATT — token selector (USDC/USDT), amount input, details panel, wallet connect prompt
- `Stake.tsx`: stake/unstake WATT <-> sWATT — sub-tabs STAKE/UNSTAKE, NAV display, APR, WEV queue options
- `Borrow.tsx`: loan application — engine selector (E1/E2/E3), amount, term, asset ID, Veriflow requirement notice
- `Portfolio.tsx`: balances, accrued yield, open loans, WEV queue status, Allo points
- `Veriflow.tsx`: asset health grid, telemetry charts (recharts line chart for GPU util over time), attestation history
- `Activity.tsx`: protocol activity table — filterable by event type, asset, engine
- `Governance.tsx`: active proposals, vote buttons, proposal creation, timelock status
- `Bridge.tsx`: coming soon page — LayerZero OFT Phase Future

**Components (`src/components/`)**
- TokenSelector, AmountInput, DetailsPanel, ActionButton (with tx loading/success/error states)
- WalletButton, HealthBadge, HealthCard, ActivityTable
- ProtocolStats, NavBar, SideNav, PortfolioCard, WEVWidget, ToastProvider

**State (`src/stores/`)**
- `walletStore.ts`: connected address, chain ID, connection status
- `txStore.ts`: pending transactions queue with status tracking

**Environment (`frontend/.env.example`)**
```
VITE_API_BASE_URL=http://localhost:8080
VITE_XDC_RPC_URL=https://erpc.apothem.network
VITE_CHAIN_ID=51
VITE_WATT_ADDRESS=
VITE_SWATT_ADDRESS=
VITE_MINT_ENGINE_ADDRESS=
VITE_LENDING_POOL_ADDRESS=
VITE_WEV_QUEUE_ADDRESS=
```

**Tests**
- Component tests with React Testing Library
- Hook tests for contract interactions (mock wagmi hooks)
- E2E smoke test: connect wallet -> mint WATT -> stake -> check portfolio (Playwright)

---

### Phase 11 — AttestationWriter

**Goal:** Scheduled backend service writes daily cryptographic proofs of hardware health to HealthAttestation.sol on XDC. Anchors Veriflow scoring data immutably on-chain.

**AttestationWriter (`internal/veriflow/attester.go`)**
- Cron: `0 0 * * *` (00:00 UTC daily)
- Also triggerable via admin API: POST /api/v1/admin/attestations/run
- For each ACTIVE asset in assets table:
  1. Fetch last 24h telemetry from telemetry_repo
  2. If no telemetry -> skip, alert via notify_service
  3. Build metrics snapshot: averages of all readings over 24h period
  4. Compute healthHash = keccak256(abi.encodePacked(assetId, avgScore, avgGpuUtil, avgTemp, timestamp))
     - MUST use same abi.encodePacked encoding as HealthAttestation.sol to allow on-chain pre-image verification
  5. Call TxManager.SubmitAttestation(assetId, healthHash, score)
  6. On success: persist to attestations table (asset_id, score, health_hash, xdc_tx_hash, attested_at)
  7. On failure: retry once after 5 min -> log error + alert admin

**New API endpoints**
- GET /api/v1/attestations/:assetId — latest attestation + 30-day history
- GET /api/v1/attestations/:assetId/verify?hash=<hash> — verifies hash matches stored pre-image (for external auditors)

**Scheduler (`internal/scheduler.go`)**
- Initialize all cron jobs at startup from cmd/api/main.go:
  - AttestationWriter: daily 00:00 UTC (Phase 11)
  - ScoringEngine sweep: every 1 hour safety net (Phase 9)
  - WEV processor check: alert if queue depth > threshold and batch not run (Phase 5 keepalive)

**Tests**
- Unit test hash computation is deterministic and matches contract encoding
- Test skip logic when no telemetry data available
- Test retry on TxManager failure
- Test verify endpoint validates hash pre-image correctly

---

### Phase 12 — TreasuryService (Engine 3)

**Goal:** Automate Engine 3. Idle capital in MintEngine not deployed to loans is swept into T-bill wrapper to generate base yield. Yield routed back to sWATT vault.

**Contract additions**
- Update `MintEngine.sol`: add `getIdleCapital() returns (uint256)`, add `deployToTBills(uint256 amount)` — TREASURY_ROLE only
- Update `Treasury.sol`: implement `deployToM0(uint256 amount)` — swap idle WATT for M0 $M token, record balance
- Update `Treasury.sol`: implement `harvestYield()` — read accrued $M yield, convert to WATT, call sWattUSD.receiveYield(amount)

**TreasuryService (`internal/service/treasury_service.go`)** — replace Phase 7 stub with full implementation:
- GetIdleCapital(): reads MintEngine.getIdleCapital() via BlockchainClient
- DeployIdleCapital(): if idle > $500k WATT threshold, calls TxManager to invoke MintEngine.deployToTBills(70% of idle) — keeps 30% as liquid buffer
- HarvestTBillYield(): calls TxManager -> Treasury.harvestYield(), updates yield_service with harvested amount
- GetTBillYield(): reads current T-bill APY from M0 Protocol
- GetTreasuryStats(): idle capital, deployed capital, T-bill balance, accrued yield

**Scheduler additions (add to `internal/scheduler.go`)**
- TreasuryService.DeployIdleCapital(): every 6 hours
- TreasuryService.HarvestTBillYield(): daily at 02:00 UTC

**Admin API endpoints**
- GET /api/v1/admin/treasury/stats
- POST /api/v1/admin/treasury/deploy (ADMIN_ROLE)
- POST /api/v1/admin/treasury/harvest (ADMIN_ROLE)

**Tests**
- Unit test DeployIdleCapital() threshold logic with mock BlockchainClient
- Test harvest calculation and yield routing to sWATT vault
- Integration: seed MintEngine idle WATT -> deploy -> harvest -> verify sWATT NAV increased

---

### Phase 13 — Engine 1 (Pre-Delivery PO Financing)

**Goal:** Pre-delivery purchase order financing. Borrower submits a PO for hardware not yet delivered. AI WATT finances the 30% deposit required.

**Prerequisites before building**
- At least 2 confirmed supplier relationships (GPU vendors or resellers)
- SPV structure confirmed with legal for pre-delivery collateral
- KYC/AML integration fully operational (Phase 7)

**Contract: `Engine1Pool.sol`** — separate from Phase 4 LendingPool to isolate Engine 1 risk
- UUPS upgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable
- Roles: CURATOR_ROLE, SUPPLIER_ROLE (whitelisted supplier wallets), ADMIN_ROLE
- Key functions:
  - `submitPurchaseOrder(bytes32 poHash, address supplier, uint256 depositAmount)` — borrower submits PO
  - `certifyPO(bytes32 poHash)` — CURATOR_ROLE certifies PO is valid and supplier is whitelisted
  - `disburseToEscrow(bytes32 poHash)` — sends WATT to supplier escrow on-chain
  - `confirmDelivery(bytes32 poHash, bytes32 assetId)` — links PO to AssetRegistry entry, releases escrow to supplier
  - `refundOnFail(bytes32 poHash)` — if delivery fails/cancelled, refunds escrow back to protocol
  - Engine 3 integration: while PO is in ESCROWED state, 70% of idle escrow auto-deployed to T-bills via TreasuryService
- Enum: `POStatus { SUBMITTED, CERTIFIED, ESCROWED, DELIVERED, REFUNDED }`
- Events: PurchaseOrderSubmitted, POCertified, EscrowDisbursed, DeliveryConfirmed, EscrowRefunded

**Backend additions**
- `po_service.go`: SubmitPO(borrowerID, supplierName, poDocHash, depositAmount), ApprovePO(poID), ConfirmDelivery(poID, assetID), GetPOStatus(poID)
- Migration `000009_create_purchase_orders.up.sql`: id, po_id, borrower_id (FK), supplier_name, po_hash, deposit_amount, status, submitted_at, certified_at, delivered_at
- `po_dto.go`: SubmitPORequest { supplier_name, po_document_hash, deposit_amount }, POResponse { po_id, status, supplier, deposit_amount, estimated_delivery }
- `po_handler.go`: POST /api/v1/purchase-orders, GET /api/v1/purchase-orders/:id, POST /api/v1/purchase-orders/:id/approve (CURATOR_ROLE), POST /api/v1/purchase-orders/:id/confirm-delivery

**Frontend additions**
- Update Borrow.tsx Engine 1 section: PO document hash input, supplier name field, PO status tracker
- Add POTracker.tsx component: timeline — PO Submitted -> Certified -> Escrowed -> Delivered -> Loan Active

**Tests**
- Full Engine 1 lifecycle: submit -> certify -> disburse -> confirm delivery -> asset registered -> Engine 2 loan active
- Refund flow: certify -> disburse -> delivery fails -> escrow refunded
- Engine 3 integration: verify idle escrow is deployed to T-bills while in ESCROWED state

---

### Phase Future — Multi-Chain Bridge (LayerZero OFT)

**Goal:** WATT and sWATT bridgeable across chains via LayerZero V2 OFT standard. Target chains: Ethereum, Arbitrum One, Base.

**Prerequisites**
- LayerZero V2 deployed and confirmed on XDC Network
- Sufficient TVL on XDC to support cross-chain demand
- Legal/compliance review for multi-chain token presence

**Contract additions**
- Upgrade WattUSD.sol to also implement IOFT (LayerZero): send(), sendFrom(), estimateSendFee()
- Deploy OFT lockbox adapters on Ethereum, Arbitrum, Base — hold canonical WATT while it circulates on XDC
- Upgrade sWattUSD.sol similarly — bridged sWATT is read-only on spoke chains, yield accrues only on XDC

**Backend additions**
- `bridge_service.go`: GetBridgeFee(fromChain, toChain, amount), InitiateBridge(user, fromChain, toChain, amount), GetBridgeStatus(txHash)
- Extend EventIndexer to subscribe to OFT events on Ethereum/Arbitrum/Base RPCs

**Frontend additions**
- Fully implement Bridge.tsx — chain selector, amount input, fee estimate, bridge button
- Multi-chain wallet connector in Wagmi config

---

## Decisions Log

| Date | Decision | Reason |
|---|---|---|
| Apr 2026 | XDC Network as primary chain | Trade finance native, ISO 20022, institutional integrations |
| Apr 2026 | UUPS proxy pattern for all contracts | Gas efficient, current OZ recommendation vs Transparent Proxy |
| Apr 2026 | All contracts are upgradeable | Protocol is early-stage — need ability to fix and iterate safely |
| Apr 2026 | go-migrate over GORM AutoMigrate | Production-safe, explicit SQL, reversible with .down.sql |
| Apr 2026 | Standard Go project layout (/cmd, /internal, /pkg, /api, /configs, /scripts) | Industry standard, enforces clear separation of concerns |
| Apr 2026 | DTO layer in internal/api/dto/ | Decouples API contract from DB models — models never leak out of handlers |
| Apr 2026 | Layer flow: Handler -> DTO -> Service -> Repository -> Model | Enforces single responsibility per layer, makes unit testing clean |
| Apr 2026 | No cross-chain bridge in v1 | Scope control — add LayerZero OFT in future phase |
| Apr 2026 | Rules-based scoring in Veriflow v1 | ML needs training data first — graduate to ML in v3+ |
| Apr 2026 | Single Go binary in v1 (not microservices) | Deployment simplicity — extract services when scale demands it |
| Apr 2026 | Gin framework + GORM | Standard, well-supported, good ecosystem for this use case |
| Apr 2026 | swaggo/swag for Swagger | Annotation-based, integrates cleanly with Gin handlers |
| Apr 2026 | contracts/src/ for Solidity sources | Keeps node_modules out of Hardhat compilation scope (avoids HH1006) |
| Apr 2026 | WattUSD uses 6 decimals | Matches USDC/USDT — no decimal conversion needed in MintEngine |
| Apr 2026 | MockStablecoin deployed on Apothem | Avoids dependency on third-party testnet tokens for integration testing |
| Apr 2026 | Verification separate from deploy | Transient explorer failures should not look like failed deploys; re-runnable |
| Apr 2026 | evmVersion: cancun in Hardhat config | OZ v5 Memory.sol uses mcopy opcode (EIP-5656), unavailable in earlier EVM targets |
| Apr 2026 | sWattUSD seed deposit on deploy | Anchors exchange rate to prevent ERC-4626 inflation attack before first real depositor |
| Apr 2026 | maxWithdraw/maxRedeem cap enforces WEV threshold | ERC-4626 checks maxWithdraw before _withdraw, so capping there is the right guard point |
| Apr 2026 | Frontend moved to Phase 10 | Backend + Veriflow must be functional before frontend can integrate meaningfully |
| Apr 2026 | Engine 1 pool separate from Engine 2 LendingPool | Pre-delivery escrow logic is complex — isolate to protect Engine 2 TVL from bugs |
| Apr 2026 | OCNFT is soulbound by default | Hardware title NFTs should not be freely transferable — MINTER_ROLE transfers on settlement only |

---

## Blockers

- [ ] None currently

---

## Notes for Next Session

- Contract sources: `contracts/src/` — Hardhat `paths.sources = "./src"`, `evmVersion: cancun`
- XDC RPC: testnet `erpc.apothem.network`, mainnet `rpc.xdcrpc.com`
- Block explorer: `testnet.xdcscan.com` (Apothem), `xdcscan.com` (mainnet)
- Verification: `ETHERSCAN_API_KEY` + proxy addresses in `contracts/.env`, run `verify.ts`
- Phase 4: `LendingPool.sol` goes in `contracts/src/credit/`, tests in `test/credit/`
- LendingPool reads `AssetRegistry.isActive()` and `HealthAttestation.getLatestAttestation()` before originating
- LendingPool routes interest to `sWattUSD.receiveYield()` on repayment
- Phase 4 deploy script must grant `LENDINGPOOL_ROLE` on AssetRegistry to LendingPool proxy address
- 167 tests passing across Phases 1–3
- Phase 4 LendingPool imports AssetRegistry, HealthAttestation, sWattUSD, WattUSD addresses from `contracts/.env`
- Phase 7: run `swag init` after adding every new handler to regenerate Swagger docs
- Phase 8 abigen: `abigen --abi=abis/WattUSD.json --pkg=contracts --out=internal/blockchain/contracts/watt_usd.go`
- Phase 9 Veriflow agent: mock nvidia-smi in unit tests; test on real GPU machine before deploying to borrower
- Phase 10 frontend: use Vite not CRA; configure XDC as custom chain in Wagmi (see Phase 10 chain config)
- Phase 11 AttestationWriter: keccak256 encoding must use abi.encodePacked to match HealthAttestation.sol exactly
