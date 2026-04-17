import { cn } from '@/lib/formatters'

interface ProgressBarProps {
  value: number  // 0-100
  variant?: 'gold' | 'teal' | 'green'
}

export function ProgressBar({ value, variant = 'gold' }: ProgressBarProps) {
  return (
    <div className="h-1 bg-bg-3 rounded-sm overflow-hidden" style={{ margin: '8px 0' }}>
      <div
        className={cn(
          'h-full rounded-sm transition-all duration-1000 ease-out',
          variant === 'gold'  && 'bg-gold',
          variant === 'teal'  && 'bg-teal',
          variant === 'green' && 'bg-green-light',
        )}
        style={{ width: `${Math.min(Math.max(value, 0), 100)}%` }}
      />
    </div>
  )
}
