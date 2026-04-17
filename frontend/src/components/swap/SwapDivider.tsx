interface SwapDividerProps {
  onSwap?: () => void
}

export function SwapDivider({ onSwap }: SwapDividerProps) {
  return (
    <div
      className="flex items-center justify-center border-b border-border"
      style={{ padding: '10px 18px' }}
    >
      <div className="flex-1 border-t border-border" />
      <button
        onClick={onSwap}
        className="rounded-full border border-border bg-white flex items-center justify-center cursor-pointer text-text-2 transition-colors hover:bg-bg-2 hover:border-border-strong flex-shrink-0"
        style={{ width: 28, height: 28, margin: '0 10px', fontSize: 13 }}
      >
        ↕
      </button>
      <div className="flex-1 border-t border-border" />
    </div>
  )
}
