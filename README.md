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

---

## Local development (Docker)

**Prerequisites:** Docker Desktop

```bash
cp .env.example .env        # passwords are pre-filled for local use
docker compose up -d --build
```

This starts five containers in dependency order:

| Container | Role | Exposed |
|---|---|---|
| `aiwatt-mysql` | MySQL 8 | internal only |
| `aiwatt-redis` | Redis 7 | internal only |
| `aiwatt-migrator` | Runs all migrations then exits | — |
| `aiwatt-backend` | Go API | internal only |
| `aiwatt-frontend` | nginx + React SPA | `http://localhost` (port 80) |

Migrations run automatically before the backend starts. On subsequent `up --build` runs, the migrator is a no-op (already at latest version).

## Local development (without Docker)

**Prerequisites:** Go 1.22+, Node 18+, MySQL 8, Redis 7 running locally

```bash
# Backend
cd backend
cp .env.example .env.local  # already populated for local
make migrate-up ENV=local
make dev ENV=local           # hot-reload via air, falls back to go run

# Frontend (separate terminal)
cd frontend
npm install
npm run dev                  # http://localhost:5173
```

---

## Staging deployment (GCP VM)

The staging VM (`34.28.184.105`) runs MySQL and Redis as pre-existing Docker services.
The app joins their networks — no new database containers are spun up.

```bash
# On the VM
git clone <repo> ~/aiwatt && cd ~/aiwatt
cp .env.staging.example .env

# Fill in:
#   MYSQL_PASSWORD   → from ~/mysql/.env
#   REDIS_PASSWORD   → from ~/redis/.env
#   JWT_SECRET       → openssl rand -hex 32
#   VERIFLOW_SIGNER_PRIVATE_KEY → your hot-wallet key (never commit)
nano .env

docker compose -f docker-compose.staging.yml up -d --build
```

See [.env.staging.example](.env.staging.example) for the full variable list.

### Cloudflare tunnel (one-time)

Add to `/etc/cloudflared/config.yml` on the VM before the catch-all:

```yaml
- hostname: staging.neurowatt.services
  service: http://localhost:3000
  originRequest:
    httpHostHeader: staging.neurowatt.services
    noTLSVerify: true
    connectTimeout: 30s
```

```bash
cloudflared tunnel route dns ai-watt staging.neurowatt.services
sudo systemctl restart cloudflared
```

### Redeploy after code changes

```bash
git pull
docker compose -f docker-compose.staging.yml up -d --build --no-deps frontend backend
```

---

## Environment summary

| | Local (Docker) | Local (native) | Staging (VM) |
|---|---|---|---|
| Compose file | `docker-compose.yml` | — | `docker-compose.staging.yml` |
| Env template | `.env.example` | `backend/.env.example` | `.env.staging.example` |
| MySQL | Spun up by compose | Homebrew | Pre-existing on VM (`mysql-net`) |
| Redis | Spun up by compose | Homebrew | Pre-existing on VM (`redis-net`) |
| Frontend URL | `http://localhost` | `http://localhost:5173` | `https://staging.neurowatt.services` |

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
