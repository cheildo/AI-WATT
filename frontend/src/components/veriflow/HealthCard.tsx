import { cn } from '@/lib/formatters'
import { HealthBadge, scoreToVariant } from './HealthBadge'

interface Metric {
  label: string
  value: string
  variant?: 'ok' | 'warn' | 'crit'
}

interface HealthCardProps {
  assetId: string
  engineType: number
  name: string
  location: string
  healthScore: number
  metrics: Metric[]
  onClick?: () => void
  isPending?: boolean
  isPlaceholder?: boolean
}

export function HealthCard({
  assetId, engineType, name, location, healthScore, metrics, onClick, isPending, isPlaceholder,
}: HealthCardProps) {
  if (isPlaceholder) {
    return (
      <div
        className="bg-bg border border-border rounded flex items-center justify-center flex-col gap-[6px] cursor-default"
        style={{ borderStyle: 'dashed', minHeight: 180, padding: 16, borderRadius: 4 }}
      >
        <div className="text-border-strong" style={{ fontSize: 22 }}>+</div>
        <div className="text-text-3" style={{ fontSize: 12 }}>Onboard new asset</div>
      </div>
    )
  }

  const variant = isPending ? 'pending' : scoreToVariant(healthScore)
  const barColor = variant === 'ok' ? 'bg-teal' : variant === 'warn' ? 'bg-[#F59E0B]' : variant === 'crit' ? 'bg-red' : 'bg-border'

  return (
    <div
      className="bg-white border border-border rounded cursor-pointer transition-all hover:border-border-strong hover:shadow-sm"
      style={{ padding: 16, borderRadius: 4 }}
      onClick={onClick}
    >
      {/* header */}
      <div className="flex items-center justify-between mb-[10px]">
        <span className="font-mono text-text-3" style={{ fontSize: 10 }}>
          {assetId} · Engine {engineType}
        </span>
        <HealthBadge
          variant={variant}
          label={isPending ? 'Pending' : variant === 'ok' ? 'Healthy' : variant === 'warn' ? 'Warning' : 'Critical'}
        />
      </div>

      {/* name */}
      <div className="font-serif mb-0.5" style={{ fontSize: 16 }}>{name}</div>
      <div className="text-text-3 mb-3" style={{ fontSize: 11 }}>📍 {location}</div>

      {/* health bar */}
      <div className="bg-bg-3 rounded-sm overflow-hidden mb-[10px]" style={{ height: 3, borderRadius: 2 }}>
        <div
          className={cn('h-full rounded-sm transition-all duration-1000 ease-out', barColor)}
          style={{ width: isPending ? '0%' : `${healthScore}%` }}
        />
      </div>

      {/* metrics */}
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 8 }}>
        {metrics.map((m, i) => (
          <div key={i}>
            <div className="text-text-3 uppercase tracking-[.07em]" style={{ fontSize: 9 }}>{m.label}</div>
            <div
              className={cn(
                'font-mono mt-px',
                m.variant === 'ok'   && 'text-teal',
                m.variant === 'warn' && 'text-[#B45309]',
                m.variant === 'crit' && 'text-red',
                !m.variant           && 'text-text-1',
              )}
              style={{ fontSize: 12 }}
            >
              {m.value}
            </div>
          </div>
        ))}
      </div>
    </div>
  )
}
