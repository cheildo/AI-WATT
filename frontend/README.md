# AI WATT — Frontend

React TypeScript dApp for the AI WATT protocol.

**Stack:** React 18 · TypeScript (strict) · Vite · Wagmi · Viem · React Query · Zustand · Tailwind CSS

---

## Prerequisites

- Node.js 18+

---

## Running with Docker (recommended)

From the repo root:

```bash
cp .env.example .env        # fill in VITE_WALLETCONNECT_PROJECT_ID and contract addresses
docker compose up -d --build
# Frontend: http://localhost (port 80, served by nginx)
```

`VITE_*` variables are baked into the static bundle at Docker build time. Changing them requires a rebuild:

```bash
docker compose up -d --build --no-deps frontend
```

## Running natively

```bash
npm install
cp .env.local.example .env.local    # already populated for local dev
npm run dev                          # http://localhost:5173
```

`VITE_API_URL` in `.env.local` points to `http://localhost:8080` (the native backend).

---

## Build

```bash
npm run build      # output → dist/
npm run preview    # preview the production build locally
```

---

## nginx proxy (Docker only)

In Docker, the frontend container runs nginx which:
- Serves the React SPA at `/`
- Proxies `/api/` requests to `http://backend:8080` (internal Docker network)

This means `VITE_API_URL` should be left **empty** in Docker env files so the browser uses relative paths that nginx intercepts. No CORS configuration is needed.

---

## Environment variables

All `VITE_*` vars are baked in at build time (Vite inlines them into the bundle).

| Variable | Description |
|---|---|
| `VITE_API_URL` | Backend base URL. Empty = nginx proxy (Docker). Full URL for native dev. |
| `VITE_CHAIN_ID` | XDC chain ID: `31337` local · `51` Apothem · `50` mainnet |
| `VITE_RPC_URL` | XDC JSON-RPC endpoint |
| `VITE_WALLETCONNECT_PROJECT_ID` | From https://cloud.walletconnect.com |
| `VITE_WATT_ADDRESS` etc. | Contract proxy addresses — update after each deploy |

See [.env.local.example](.env.local.example) for local dev and the root [.env.staging.example](../.env.staging.example) for staging.

---

## Wallet connection

Uses Wagmi + Viem with a custom XDC chain config.
Supports MetaMask and WalletConnect. XDC Apothem testnet (chain ID 51) by default.
