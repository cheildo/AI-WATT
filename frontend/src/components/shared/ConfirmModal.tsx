import { useEffect } from 'react'

interface ConfirmRow {
  label: string
  value: string
  valueColor?: string
}

interface ConfirmModalProps {
  open: boolean
  title: string
  subtitle?: string
  rows: ConfirmRow[]
  actionLabel: string
  actionVariant?: 'green' | 'gold'
  onConfirm: () => void
  onCancel: () => void
}

export function ConfirmModal({
  open,
  title,
  subtitle,
  rows,
  actionLabel,
  actionVariant = 'green',
  onConfirm,
  onCancel,
}: ConfirmModalProps) {
  // Close on Escape
  useEffect(() => {
    if (!open) return
    const handler = (e: KeyboardEvent) => { if (e.key === 'Escape') onCancel() }
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [open, onCancel])

  if (!open) return null

  const handleConfirm = () => {
    onConfirm()
    onCancel()
  }

  return (
    <div
      className="fixed inset-0 z-[9998] flex items-center justify-center"
      style={{ background: 'rgba(28,26,20,0.45)', backdropFilter: 'blur(3px)' }}
      onClick={(e) => { if (e.target === e.currentTarget) onCancel() }}
    >
      <div
        className="bg-white border border-border animate-fadeup"
        style={{ borderRadius: 6, width: 380, boxShadow: '0 20px 60px rgba(0,0,0,0.18)' }}
      >
        {/* Header */}
        <div className="flex items-start justify-between border-b border-border" style={{ padding: '16px 20px 14px' }}>
          <div>
            <div className="font-semibold text-text-1" style={{ fontSize: 14 }}>{title}</div>
            {subtitle && (
              <div className="text-text-3 mt-[3px]" style={{ fontSize: 11 }}>{subtitle}</div>
            )}
          </div>
          <button
            onClick={onCancel}
            className="text-text-3 hover:text-text-1 cursor-pointer transition-colors flex-shrink-0"
            style={{ fontSize: 20, background: 'none', border: 'none', lineHeight: 1, marginTop: 1 }}
            aria-label="Close"
          >×</button>
        </div>

        {/* Summary rows */}
        <div style={{ padding: '0 20px' }}>
          {rows.map((r, i) => (
            <div
              key={i}
              className="flex items-center justify-between"
              style={{ padding: '11px 0', borderBottom: i < rows.length - 1 ? '1px solid #D8D3C6' : 'none' }}
            >
              <span className="text-text-3" style={{ fontSize: 12 }}>{r.label}</span>
              <span
                className="font-mono font-medium"
                style={{ fontSize: 13, color: r.valueColor ?? '#1C1A14' }}
              >{r.value}</span>
            </div>
          ))}
        </div>

        {/* Warning note */}
        <div className="border-t border-border" style={{ padding: '10px 20px 0' }}>
          <p className="text-text-3 leading-snug" style={{ fontSize: 10 }}>
            Please review the details above. This transaction cannot be reversed once submitted to the blockchain.
          </p>
        </div>

        {/* Action buttons */}
        <div className="flex items-center gap-2" style={{ padding: '12px 20px 16px' }}>
          <button
            onClick={onCancel}
            className="flex-1 border border-border rounded font-semibold text-text-2 cursor-pointer hover:bg-bg hover:border-border-strong transition-colors"
            style={{ padding: '10px 0', fontSize: 12, borderRadius: 4, background: 'transparent' }}
          >
            Cancel
          </button>
          <button
            onClick={handleConfirm}
            className="flex-1 rounded font-semibold text-white cursor-pointer transition-opacity hover:opacity-85"
            style={{
              padding: '10px 0', fontSize: 12, borderRadius: 4, border: 'none',
              background: actionVariant === 'gold' ? '#9A6B0A' : '#1A3C2A',
            }}
          >
            {actionLabel}
          </button>
        </div>
      </div>
    </div>
  )
}
