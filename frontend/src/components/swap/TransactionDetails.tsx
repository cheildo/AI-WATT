import { cn } from '@/lib/formatters'
import { ActionButton } from '@/components/shared/ActionButton'

export interface TxStep {
  label: string
  status: 'pending' | 'active' | 'done'
}

export interface InfoItem {
  key: string
  value: string
  variant?: 'default' | 'teal' | 'gold'
}

interface TransactionDetailsProps {
  steps: TxStep[]
  info: InfoItem[]
  actionLabel: string
  onAction?: () => void
  disabled?: boolean
  actionVariant?: 'green' | 'gold'
}

export function TransactionDetails({
  steps,
  info,
  actionLabel,
  onAction,
  disabled,
  actionVariant = 'green',
}: TransactionDetailsProps) {
  return (
    <div className="bg-bg flex flex-col">
      <div className="border-b border-border text-text-2 font-semibold tracking-[.02em]" style={{ padding: '14px 18px', fontSize: 12 }}>
        Transaction Details
      </div>

      {/* Steps */}
      <div className="flex flex-col" style={{ padding: '14px 18px' }}>
        {steps.map((step, i) => {
          const isLast = i === steps.length - 1
          return (
            <div key={i} className="relative flex items-start gap-[10px] py-[10px]">
              {!isLast && (
                <div
                  className="absolute bg-border"
                  style={{ left: 8, top: 26, width: 1, height: 'calc(100% - 6px)' }}
                />
              )}
              {/* dot */}
              <div
                className={cn(
                  'rounded-full border-[1.5px] flex items-center justify-center flex-shrink-0 mt-px',
                  step.status === 'done'
                    ? 'bg-green-bg border-green-border'
                    : 'bg-white border-border-strong'
                )}
                style={{ width: 18, height: 18 }}
              >
                {step.status === 'done' && (
                  <svg width="8" height="8" viewBox="0 0 8 8" fill="none">
                    <path d="M1.5 4l2 2 3-3" stroke="#0A7068" strokeWidth="1.3" strokeLinecap="round" strokeLinejoin="round"/>
                  </svg>
                )}
                {step.status !== 'done' && (
                  <svg width="8" height="8" viewBox="0 0 8 8">
                    <circle cx="4" cy="4" r="1.5" fill="#9A9484"/>
                  </svg>
                )}
              </div>
              {/* label */}
              <span
                className={cn(
                  '',
                  step.status === 'done'   && 'text-teal',
                  step.status === 'active' && 'text-text-1 font-medium',
                  step.status === 'pending' && 'text-text-3'
                )}
                style={{ fontSize: 12 }}
              >
                {step.label}
              </span>
            </div>
          )
        })}
      </div>

      {/* Info grid */}
      <div
        className="grid border-t border-border"
        style={{ gridTemplateColumns: 'repeat(2,1fr)', padding: '10px 18px 0' }}
      >
        {info.map((item, i) => (
          <div
            key={i}
            className="flex justify-between py-2"
            style={{
              fontSize: 11,
              borderRight: i % 2 === 0 ? '1px solid #D8D3C6' : 'none',
              paddingRight: i % 2 === 0 ? 12 : 0,
              paddingLeft: i % 2 === 1 ? 12 : 0,
            }}
          >
            <span className="text-text-3">{item.key}</span>
            <span
              className={cn(
                'font-mono',
                item.variant === 'teal' && 'text-teal',
                item.variant === 'gold' && 'text-gold',
                !item.variant && 'text-text-1',
              )}
            >
              {item.value}
            </span>
          </div>
        ))}
      </div>

      {/* Action button */}
      <div className="border-t border-border" style={{ padding: 14 }}>
        <ActionButton onClick={onAction} disabled={disabled} variant={actionVariant}>
          {actionLabel}
        </ActionButton>
      </div>
    </div>
  )
}
