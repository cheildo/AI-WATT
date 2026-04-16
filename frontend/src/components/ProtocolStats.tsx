import { formatUnits } from 'viem'
import { useVaultStats } from '@/hooks/contracts/useSWattUSD'

function fmt(v: bigint | undefined, decimals = 18) {
  if (v === undefined) return '—'
  const n = parseFloat(formatUnits(v, decimals))
  return n.toLocaleString(undefined, { maximumFractionDigits: 2 })
}

export function ProtocolStats() {
  const { totalAssets, totalSupply, nav } = useVaultStats()

  const stats = [
    { label: 'TVL (WATT)',       value: fmt(totalAssets.data) },
    { label: 'sWATT Supply',     value: fmt(totalSupply.data) },
    { label: 'NAV / sWATT',      value: fmt(nav.data) },
  ]

  return (
    <div className="grid grid-cols-3 gap-4">
      {stats.map((s) => (
        <div key={s.label} className="rounded-xl border border-surface-border bg-surface-card p-4">
          <p className="text-xs text-text-secondary">{s.label}</p>
          <p className="mt-1 text-xl font-semibold text-text-primary">{s.value}</p>
        </div>
      ))}
    </div>
  )
}
