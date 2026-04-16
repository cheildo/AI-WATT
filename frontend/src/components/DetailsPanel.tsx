interface Row {
  label: string
  value: string
  highlight?: boolean
}

export function DetailsPanel({ rows }: { rows: Row[] }) {
  return (
    <div className="rounded-xl border border-surface-border bg-surface-card divide-y divide-surface-border">
      {rows.map((r) => (
        <div key={r.label} className="flex items-center justify-between px-4 py-3 text-sm">
          <span className="text-text-secondary">{r.label}</span>
          <span className={r.highlight ? 'font-medium text-yield' : 'text-text-primary'}>{r.value}</span>
        </div>
      ))}
    </div>
  )
}
