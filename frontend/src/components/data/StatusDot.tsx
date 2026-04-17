interface StatusDotProps {
  color: string
  label: string
}

export function StatusDot({ color, label }: StatusDotProps) {
  return (
    <div className="flex items-center gap-[5px] font-sans">
      <div className="rounded-full flex-shrink-0" style={{ width: 6, height: 6, background: color }} />
      {label}
    </div>
  )
}
