import { ReactNode } from 'react'

interface PortfolioCardProps {
  title: string
  children: ReactNode
}

export function PortfolioCard({ title, children }: PortfolioCardProps) {
  return (
    <div className="bg-white border border-border rounded overflow-hidden" style={{ borderRadius: 4 }}>
      <div
        className="border-b border-border font-semibold uppercase tracking-[.08em] text-text-3 bg-bg-2"
        style={{ padding: '12px 16px', fontSize: 11 }}
      >
        {title}
      </div>
      <div style={{ padding: '14px 16px' }}>
        {children}
      </div>
    </div>
  )
}
