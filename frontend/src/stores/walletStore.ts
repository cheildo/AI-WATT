import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface WalletStore {
  jwt: string | null
  setJwt: (jwt: string | null) => void
}

export const useWalletStore = create<WalletStore>()(
  persist(
    (set) => ({
      jwt: null,
      setJwt: (jwt) => set({ jwt }),
    }),
    { name: 'aiwatt-wallet' }
  )
)
