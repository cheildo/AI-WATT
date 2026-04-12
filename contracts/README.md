# AI WATT — Smart Contracts

Solidity smart contracts for the AI WATT protocol, deployed on XDC Network.

**Stack:** Solidity 0.8.24 · Hardhat · OpenZeppelin Upgradeable v5 · TypeScript tests

All contracts use the **UUPS upgradeable proxy pattern**. No constructors — initialization via `initialize()`.

---

## Prerequisites

- Node.js 18+
- npm

---

## Setup

```bash
npm install
cp .env.example .env
# Fill in DEPLOYER_PRIVATE_KEY for testnet deploy
```

---

## Commands

```bash
# Compile
npm run compile

# Run tests
npm test

# Test coverage
npm run coverage

# Start local Hardhat node
npm run node

# Deploy to Apothem testnet
npm run deploy:apothem

# Deploy to XDC mainnet
npm run deploy:mainnet
```

---

## Contract layout

```
src/
├── tokens/        # WattUSD.sol, sWattUSD.sol, WattToken.sol
├── credit/        # MintEngine.sol, LendingPool.sol, WEVQueue.sol, Treasury.sol
├── assets/        # OCNFT.sol, AssetRegistry.sol, HealthAttestation.sol
├── governance/    # Governor.sol, Timelock.sol, AccessControl.sol
├── interfaces/    # IWattUSD.sol, IMintEngine.sol, ...
└── mocks/         # Test-only mock contracts
test/              # Hardhat TypeScript test files
scripts/           # deploy-proxy.ts, upgrade.ts, verify.ts
```

---

## Networks

| Network | Chain ID | RPC |
|---|---|---|
| Hardhat local | 31337 | http://127.0.0.1:8545 |
| XDC Apothem (testnet) | 51 | https://erpc.apothem.network |
| XDC Mainnet | 50 | https://erpc.xinfin.network |

Testnet XDC faucet: https://faucet.apothem.network

---

## Upgrade policy

All upgrades require:
1. Governor proposal + vote (3-day voting period, 4% quorum)
2. Timelock execution (48h delay)
3. `hardhat-upgrades` storage layout safety check before every deploy

Never upgrade a proxy without running `upgrades.validateUpgrade()` first.

---

## Environment variables

See [.env.example](.env.example) for the full list.
