import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'
import { useWalletStore } from '@/stores/walletStore'

export interface ChainEvent {
  id: string
  contractAddress: string
  eventName: string
  parsedArgs: string
  blockNumber: number
  txHash: string
  logIndex: number
  createdAt: string
}

export function useActivity(page = 1) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['activity', page],
    queryFn: () =>
      api.get<{ events: ChainEvent[]; total: number }>(`/api/v1/events?page=${page}&page_size=20`, jwt),
    refetchInterval: 30_000,
  })
}
