import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'
import { useWalletStore } from '@/stores/walletStore'

export interface HealthScore {
  assetId: string
  healthScore: number
  status: string
  computedAt: string
}

export interface Attestation {
  assetId: string
  healthScore: number
  healthHash: string
  xdcTxHash: string
  timestamp: string
}

export function useHealthScore(assetId?: string) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['healthScore', assetId],
    queryFn: () => api.get<HealthScore>(`/api/v1/veriflow/assets/${assetId}/score`, jwt),
    enabled: !!assetId,
    refetchInterval: 60_000,
  })
}

export function useAttestation(assetId?: string) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['attestation', assetId],
    queryFn: () => api.get<Attestation>(`/api/v1/veriflow/assets/${assetId}/attestation`, jwt),
    enabled: !!assetId,
  })
}

export function useAssets(ownerId?: string) {
  const jwt = useWalletStore((s) => s.jwt)
  return useQuery({
    queryKey: ['assets', ownerId],
    queryFn: () =>
      api.get<{ assets: { id: string; assetType: string; status: string; healthScore: number }[] }>(
        `/api/v1/assets?owner_id=${ownerId}&page=1&page_size=50`,
        jwt,
      ),
    enabled: !!ownerId,
  })
}
