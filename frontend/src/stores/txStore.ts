import { create } from 'zustand'

export type TxStatus = 'pending' | 'success' | 'error'

export interface TxRecord {
  id: string
  description: string
  hash?: string
  status: TxStatus
  createdAt: number
}

interface TxState {
  transactions: TxRecord[]
  add: (tx: Omit<TxRecord, 'createdAt'>) => void
  update: (id: string, patch: Partial<TxRecord>) => void
  clear: () => void
}

export const useTxStore = create<TxState>((set) => ({
  transactions: [],
  add: (tx) =>
    set((s) => ({
      transactions: [{ ...tx, createdAt: Date.now() }, ...s.transactions].slice(0, 20),
    })),
  update: (id, patch) =>
    set((s) => ({
      transactions: s.transactions.map((t) => (t.id === id ? { ...t, ...patch } : t)),
    })),
  clear: () => set({ transactions: [] }),
}))
