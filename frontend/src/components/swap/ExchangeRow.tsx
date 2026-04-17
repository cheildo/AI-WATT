interface ExchangeRowProps {
  left: string
  right: string
}

export function ExchangeRow({ left, right }: ExchangeRowProps) {
  return (
    <div
      className="flex items-center justify-between border-t border-border text-text-3"
      style={{ padding: '10px 18px', fontSize: 11 }}
    >
      <span>{left}</span>
      <span className="font-mono text-text-2">{right}</span>
    </div>
  )
}
