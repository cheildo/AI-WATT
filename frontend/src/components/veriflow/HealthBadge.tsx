import { cn } from '@/lib/formatters'

type BadgeVariant = 'ok' | 'warn' | 'crit' | 'pending'

interface HealthBadgeProps {
  variant: BadgeVariant
  label: string
}

export function HealthBadge({ variant, label }: HealthBadgeProps) {
  return (
    <span
      className={cn(
        'inline-flex items-center gap-1 rounded-full font-semibold',
        variant === 'ok'      && 'bg-teal-bg text-teal',
        variant === 'warn'    && 'bg-[#FFF7ED] text-[#B45309]',
        variant === 'crit'    && 'bg-red-bg text-red',
        variant === 'pending' && 'bg-bg-3 text-text-3',
      )}
      style={{ padding: '3px 8px', fontSize: 10 }}
    >
      <span
        className="rounded-full flex-shrink-0"
        style={{ width: 4, height: 4, background: 'currentColor' }}
      />
      {label}
    </span>
  )
}

export function scoreToVariant(score: number): BadgeVariant {
  if (score >= 80) return 'ok'
  if (score >= 60) return 'warn'
  if (score > 0)   return 'crit'
  return 'pending'
}
