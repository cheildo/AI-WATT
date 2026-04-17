import { useQuery } from '@tanstack/react-query'
import { useAccount } from 'wagmi'

const BASE = import.meta.env.VITE_BACKEND_URL ?? ''

async function fetchPortfolio(address: string) {
  const res = await fetch(`${BASE}/api/v1/loans?borrower=${address}`)
  if (!res.ok) throw new Error('Failed to fetch portfolio')
  return res.json()
}

export function usePortfolio() {
  const { address } = useAccount()
  return useQuery({
    queryKey: ['portfolio', address],
    queryFn: () => fetchPortfolio(address!),
    enabled: !!address,
    staleTime: 30_000,
  })
}
