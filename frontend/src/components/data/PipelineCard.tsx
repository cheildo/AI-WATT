import { ReactNode } from 'react'

interface PipelineCell {
  label: ReactNode
  value: string
  sub: ReactNode
  highlight?: boolean
}

interface PipelineCardProps {
  title: ReactNode
  cells: PipelineCell[]
}

export function PipelineCard({ title, cells }: PipelineCardProps) {
  return (
    <div className="bg-white border border-border rounded mb-4">
      <div className="border-b border-border font-semibold" style={{ padding: '14px 18px', fontSize: 13 }}>
        {title}
      </div>
      <div style={{ display: 'grid', gridTemplateColumns: `repeat(${cells.length}, 1fr)` }}>
        {cells.map((cell, i) => (
          <div
            key={i}
            className={`border-r border-border last:border-r-0 ${cell.highlight ? 'bg-bg-2' : ''}`}
            style={{ padding: '16px 18px' }}
          >
            <div className="uppercase tracking-[.08em] text-text-3 flex items-center gap-[5px] mb-1" style={{ fontSize: 10 }}>
              {cell.label}
            </div>
            <div className="font-serif mb-0.5" style={{ fontSize: 22 }}>{cell.value}</div>
            <div className="text-text-3 font-mono" style={{ fontSize: 11 }}>{cell.sub}</div>
          </div>
        ))}
      </div>
    </div>
  )
}
