/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string
  readonly VITE_XDC_RPC_URL: string
  readonly VITE_CHAIN_ID: string
  readonly VITE_WATT_ADDRESS: string
  readonly VITE_SWATT_ADDRESS: string
  readonly VITE_MINT_ENGINE_ADDRESS: string
  readonly VITE_LENDING_POOL_ADDRESS: string
  readonly VITE_WEV_QUEUE_ADDRESS: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
