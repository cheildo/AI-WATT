interface Props {
  value: string
  onChange: (v: string) => void
  label: string
  symbol: string
  max?: string
  disabled?: boolean
}

export function AmountInput({ value, onChange, label, symbol, max, disabled }: Props) {
  return (
    <div className="rounded-xl border border-surface-border bg-surface-card p-4">
      <div className="mb-2 flex items-center justify-between text-xs text-text-secondary">
        <span>{label}</span>
        {max && (
          <button
            type="button"
            onClick={() => onChange(max)}
            className="text-brand hover:underline"
          >
            MAX {max} {symbol}
          </button>
        )}
      </div>
      <div className="flex items-center gap-3">
        <input
          type="number"
          min="0"
          step="any"
          value={value}
          onChange={(e) => onChange(e.target.value)}
          disabled={disabled}
          placeholder="0.00"
          className="w-full bg-transparent text-2xl font-medium text-text-primary outline-none placeholder:text-text-muted disabled:opacity-50"
        />
        <span className="shrink-0 rounded-lg bg-surface-hover px-3 py-1.5 text-sm font-medium text-text-primary">
          {symbol}
        </span>
      </div>
    </div>
  )
}
