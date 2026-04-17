import { useQuery } from '@tanstack/react-query'

const BASE = import.meta.env.VITE_BACKEND_URL ?? ''

async function fetchHealthScore(assetId: string) {
  const res = await fetch(`${BASE}/api/v1/veriflow/score/${assetId}`)
  if (!res.ok) throw new Error('Failed to fetch health score')
  return res.json()
}

async function fetchAttestation(assetId: string) {
  const res = await fetch(`${BASE}/api/v1/veriflow/attestation/${assetId}`)
  if (!res.ok) throw new Error('Failed to fetch attestation')
  return res.json()
}

export function useHealthScore(assetId?: string) {
  return useQuery({
    queryKey: ['health-score', assetId],
    queryFn: () => fetchHealthScore(assetId!),
    enabled: !!assetId,
    refetchInterval: 60_000,
  })
}

export function useAttestation(assetId?: string) {
  return useQuery({
    queryKey: ['attestation', assetId],
    queryFn: () => fetchAttestation(assetId!),
    enabled: !!assetId,
  })
}
