import { create } from 'zustand'

interface WalletState {
  jwt: string | null
  setJwt: (token: string | null) => void
  clearJwt: () => void
}

export const useWalletStore = create<WalletState>((set) => ({
  jwt: localStorage.getItem('aiwatt_jwt'),
  setJwt: (token) => {
    if (token) localStorage.setItem('aiwatt_jwt', token)
    else localStorage.removeItem('aiwatt_jwt')
    set({ jwt: token })
  },
  clearJwt: () => {
    localStorage.removeItem('aiwatt_jwt')
    set({ jwt: null })
  },
}))
