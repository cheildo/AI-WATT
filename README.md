# AI WATT

Decentralized credit protocol for AI and automation assets, built on XDC Network.
Finances GPU clusters, robotics, and AI energy infrastructure through a dual-token system (WATT + sWATT).

**Neurowatt Pte. Ltd. — Singapore**

---

## Monorepo structure

```
aiwatt/
├── backend/          # Golang API server (Gin, GORM, go-migrate)
├── contracts/        # Solidity smart contracts (Hardhat, OZ Upgradeable, XDC)
├── veriflow-agent/   # Standalone Go binary deployed on borrower hardware
├── frontend/         # React TypeScript dApp (Wagmi, Viem, React Query)
└── docs/             # Architecture docs and runbooks
```

Each sub-repo is independently runnable. See its own `README.md` for setup instructions.

---

## Quick start (local full stack)

**Prerequisites:** Docker, Docker Compose

```bash
cp .env.example .env          # edit MySQL/Redis passwords if needed
docker-compose up -d          # starts MySQL 8 + Redis 7

cd backend
cp .env.example .env          # fill in JWT_SECRET etc.
go run ./cmd/api              # starts API on :8080
```

---

## Sub-repos

| Repo | README |
|---|---|
| Backend API | [backend/README.md](backend/README.md) |
| Smart contracts | [contracts/README.md](contracts/README.md) |
| Veriflow agent | [veriflow-agent/README.md](veriflow-agent/README.md) |
| Frontend dApp | [frontend/README.md](frontend/README.md) |

---

## Networks

| Network | Chain ID | RPC |
|---|---|---|
| XDC Apothem (testnet) | 51 | https://erpc.apothem.network |
| XDC Mainnet | 50 | https://rpc.xdcrpc.com |
