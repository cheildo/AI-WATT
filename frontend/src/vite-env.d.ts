/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_ENV: string
  readonly VITE_CHAIN_ID: string
  readonly VITE_RPC_URL: string
  readonly VITE_API_URL: string
  readonly VITE_WALLETCONNECT_PROJECT_ID: string
  // Stablecoin addresses (mock on testnet, real on mainnet)
  readonly VITE_USDC_ADDRESS: string
  readonly VITE_USDT_ADDRESS: string
  // Protocol contract proxy addresses
  readonly VITE_WATT_ADDRESS: string
  readonly VITE_SWATT_ADDRESS: string
  readonly VITE_MINT_ENGINE_ADDRESS: string
  readonly VITE_LENDING_POOL_ADDRESS: string
  readonly VITE_WEV_QUEUE_ADDRESS: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
