import { useQuery } from '@tanstack/react-query'

const BASE = import.meta.env.VITE_BACKEND_URL ?? ''

async function fetchActivity(page: number) {
  const res = await fetch(`${BASE}/api/v1/events?page=${page}&limit=20`)
  if (!res.ok) throw new Error('Failed to fetch activity')
  return res.json()
}

export function useActivity(page = 1) {
  return useQuery({
    queryKey: ['activity', page],
    queryFn: () => fetchActivity(page),
    refetchInterval: 30_000,
  })
}
