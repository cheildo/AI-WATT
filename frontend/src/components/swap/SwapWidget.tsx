import { ReactNode } from 'react'

interface SwapWidgetProps {
  notice: ReactNode
  left: ReactNode
  right: ReactNode
}

function InfoIcon() {
  return (
    <svg width="13" height="13" viewBox="0 0 14 14" fill="none" style={{ flexShrink: 0 }}>
      <circle cx="7" cy="7" r="6" stroke="#5A5646" strokeWidth="1.3"/>
      <path d="M7 6v4M7 4.5v.5" stroke="#5A5646" strokeWidth="1.3" strokeLinecap="round"/>
    </svg>
  )
}

export function SwapWidget({ notice, left, right }: SwapWidgetProps) {
  return (
    <div
      className="border border-border overflow-hidden bg-white animate-fadeup"
      style={{
        display: 'grid',
        gridTemplateColumns: '1fr 280px',
        borderRadius: 2,
        maxWidth: 820,
      }}
    >
      {/* Notice bar — full width */}
      <div
        className="bg-bg-2 border-b border-border flex items-center gap-[7px] text-text-2"
        style={{ gridColumn: '1 / -1', padding: '10px 16px', fontSize: 12 }}
      >
        <InfoIcon />
        {notice}
      </div>

      {/* Left panel */}
      <div className="border-r border-border">
        {left}
      </div>

      {/* Right panel */}
      <div className="bg-bg">
        {right}
      </div>
    </div>
  )
}
