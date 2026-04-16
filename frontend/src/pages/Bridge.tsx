export function Bridge() {
  return (
    <div className="px-4 py-8">
      <h1 className="mb-2 text-2xl font-bold text-text-primary">Bridge</h1>
      <p className="mb-8 text-sm text-text-secondary">
        Cross-chain transfers via LayerZero OFT standard.
      </p>

      <div className="rounded-xl border border-surface-border bg-surface-card p-8 text-center">
        <p className="text-4xl">🌉</p>
        <p className="mt-3 text-lg font-medium text-text-primary">Multi-chain Bridge — Coming Soon</p>
        <p className="mt-2 text-sm text-text-secondary">
          The AI WATT bridge to Ethereum, Arbitrum, and other chains via LayerZero OFT
          is planned for a future phase. Protocol is currently XDC Network only.
        </p>
        <div className="mt-6 rounded-xl border border-surface-border bg-surface p-4 text-left text-sm">
          <p className="font-medium text-text-primary">Planned chains</p>
          <ul className="mt-2 space-y-1 text-text-secondary">
            <li>• XDC Network (live)</li>
            <li>• Ethereum mainnet (future)</li>
            <li>• Arbitrum One (future)</li>
            <li>• Base (future)</li>
          </ul>
        </div>
      </div>
    </div>
  )
}
