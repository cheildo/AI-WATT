import { cn } from '@/lib/formatters'

interface WEVOption {
  time: string
  label: string
  isPriority?: boolean
  onClick?: () => void
}

interface WEVOptionsProps {
  options: WEVOption[]
}

export function WEVOptions({ options }: WEVOptionsProps) {
  return (
    <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3,1fr)', gap: 6, marginTop: 10 }}>
      {options.map((opt, i) => (
        <button
          key={i}
          onClick={opt.onClick}
          className="bg-bg border border-border rounded text-center cursor-pointer transition-colors hover:border-border-strong"
          style={{ padding: 8, borderRadius: 3 }}
        >
          <div className={cn('font-mono font-medium', opt.isPriority ? 'text-gold' : 'text-text-1')} style={{ fontSize: 12 }}>
            {opt.time}
          </div>
          <div className="text-text-3 mt-0.5" style={{ fontSize: 10 }}>{opt.label}</div>
        </button>
      ))}
    </div>
  )
}
