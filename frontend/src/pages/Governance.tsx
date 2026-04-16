export function Governance() {
  return (
    <div className="px-4 py-8">
      <h1 className="mb-2 text-2xl font-bold text-text-primary">Governance</h1>
      <p className="mb-8 text-sm text-text-secondary">
        $WATT holders vote on protocol parameters, curator approvals, and contract upgrades.
        All changes are queued through a 48-hour timelock.
      </p>

      <div className="rounded-xl border border-surface-border bg-surface-card p-8 text-center">
        <p className="text-4xl">🗳️</p>
        <p className="mt-3 text-lg font-medium text-text-primary">Governance — Phase 6</p>
        <p className="mt-2 text-sm text-text-secondary">
          $WATT token and Governor contract deploy in Phase 6.
          On-chain governance will be available after TGE.
        </p>
        <div className="mt-6 rounded-xl border border-brand/30 bg-brand-muted p-4 text-left text-sm">
          <p className="font-medium text-brand">Coming in Phase 6</p>
          <ul className="mt-2 space-y-1 text-text-secondary">
            <li>• Propose and vote on protocol parameter changes</li>
            <li>• Curator whitelist governance</li>
            <li>• Contract upgrade proposals (48h timelock)</li>
            <li>• 4% quorum, 3-day voting period</li>
          </ul>
        </div>
      </div>
    </div>
  )
}
