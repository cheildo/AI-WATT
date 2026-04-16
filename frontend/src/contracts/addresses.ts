import type { Address } from 'viem'

function addr(key: string): Address {
  const v = key as Address
  if (!v || v === '0x') return '0x0000000000000000000000000000000000000000'
  return v
}

export const CONTRACT_ADDRESSES = {
  wattUSD:     addr(import.meta.env.VITE_WATT_ADDRESS),
  sWattUSD:    addr(import.meta.env.VITE_SWATT_ADDRESS),
  mintEngine:  addr(import.meta.env.VITE_MINT_ENGINE_ADDRESS),
  lendingPool: addr(import.meta.env.VITE_LENDING_POOL_ADDRESS),
  wevQueue:    addr(import.meta.env.VITE_WEV_QUEUE_ADDRESS),
} as const
