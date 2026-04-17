/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_WATT_ADDRESS: string
  readonly VITE_SWATT_ADDRESS: string
  readonly VITE_MINT_ENGINE_ADDRESS: string
  readonly VITE_LENDING_POOL_ADDRESS: string
  readonly VITE_WEV_QUEUE_ADDRESS: string
  readonly VITE_WALLETCONNECT_PROJECT_ID: string
  readonly VITE_BACKEND_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
