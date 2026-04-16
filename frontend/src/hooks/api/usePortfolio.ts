import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'
import { useWalletStore } from '@/stores/walletStore'

export interface LoanSummary {
  id: string
  assetId: string
  engineType: number
  status: string
  createdAt: string
}

export interface PortfolioData {
  wattBalance: string
  sWattBalance: string
  accruedYield: string
  openLoans: LoanSummary[]
  wevQueueCount: number
}

export function usePortfolio(address?: string) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['portfolio', address],
    queryFn: () => api.get<PortfolioData>(`/api/v1/portfolio/${address}`, jwt),
    enabled: !!address && !!jwt,
  })
}

export function useLoans(borrowerId?: string) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['loans', borrowerId],
    queryFn: () =>
      api.get<{ loans: LoanSummary[]; total: number; page: number }>(
        `/api/v1/loans?borrower_id=${borrowerId}&page=1&page_size=20`,
        jwt,
      ),
    enabled: !!borrowerId && !!jwt,
  })
}
