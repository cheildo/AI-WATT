interface YieldBannerProps {
  amount: string
  onClaim?: () => void
}

export function YieldBanner({ amount, onClaim }: YieldBannerProps) {
  return (
    <div
      className="bg-teal-bg border border-teal-border rounded flex items-center justify-between mt-3"
      style={{ padding: '12px 14px', borderRadius: 3 }}
    >
      <div>
        <div className="uppercase tracking-[.08em] text-teal mb-[3px]" style={{ fontSize: 10 }}>Yield Earned</div>
        <div className="font-serif" style={{ fontSize: 22 }}>{amount}</div>
      </div>
      <button
        onClick={onClaim}
        className="bg-teal text-white border-none rounded font-semibold cursor-pointer transition-opacity hover:opacity-85"
        style={{ padding: '7px 14px', fontSize: 11, borderRadius: 3 }}
      >
        Auto ✓
      </button>
    </div>
  )
}
