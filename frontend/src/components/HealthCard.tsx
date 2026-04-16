import { HealthBadge } from './HealthBadge'

interface Props {
  assetId: string
  assetType: string
  status: string
  healthScore: number
  onClick?: () => void
}

export function HealthCard({ assetId, assetType, status, healthScore, onClick }: Props) {
  return (
    <button
      onClick={onClick}
      className="w-full rounded-xl border border-surface-border bg-surface-card p-4 text-left hover:border-brand/50 transition-colors"
    >
      <div className="mb-3 flex items-start justify-between">
        <div>
          <p className="text-xs text-text-secondary font-mono">{assetId.slice(0, 8)}…</p>
          <p className="mt-0.5 text-sm font-medium text-text-primary capitalize">{assetType.replace('_', ' ')}</p>
        </div>
        <HealthBadge score={healthScore} size="sm" />
      </div>
      <div className="h-1.5 w-full rounded-full bg-surface-hover">
        <div
          className="h-full rounded-full bg-gradient-to-r from-danger via-warn to-yield transition-all"
          style={{ width: `${healthScore}%` }}
        />
      </div>
      <p className="mt-2 text-xs text-text-secondary capitalize">{status}</p>
    </button>
  )
}
