# AI WATT — Frontend

React TypeScript dApp for the AI WATT protocol.

**Stack:** React · TypeScript (strict) · Wagmi · Viem · React Query · Zustand · Tailwind CSS

---

## Prerequisites

- Node.js 18+
- npm or pnpm

---

## Setup

```bash
npm install
cp .env.example .env.local
# Fill in VITE_API_URL and contract addresses after deploy
```

---

## Development

```bash
npm run dev       # starts dev server on http://localhost:5173
npm run build     # production build
npm run preview   # preview production build locally
```

---

## Environment variables

See [.env.example](.env.example) for the full list.

Contract addresses (`VITE_WATT_USD_ADDRESS`, `VITE_MINT_ENGINE_ADDRESS`) are updated after each deploy.

---

## Wallet connection

Uses Wagmi + Viem with a custom XDC chain config.
Supports MetaMask and WalletConnect. XDC Apothem testnet by default in development.

> Frontend scaffold is pending — implementation starts in Phase 12.
