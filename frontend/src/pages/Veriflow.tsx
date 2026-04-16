import { useState } from 'react'
import { useAccount } from 'wagmi'
import { LineChart, Line, XAxis, YAxis, Tooltip, ResponsiveContainer, CartesianGrid } from 'recharts'
import { WalletButton } from '@/components/WalletButton'
import { HealthCard } from '@/components/HealthCard'
import { HealthBadge } from '@/components/HealthBadge'
import { useAssets, useHealthScore, useAttestation } from '@/hooks/api/useVeriflow'

function shortHash(h: string) { return `${h.slice(0, 10)}…${h.slice(-8)}` }

// Placeholder trend data — in production fetched via GET /api/v1/telemetry/history
const mockTrend = Array.from({ length: 12 }, (_, i) => ({
  t: `${55 - i * 5}m`,
  score: 70 + Math.random() * 25,
})).reverse()

function AssetDetail({ assetId }: { assetId: string }) {
  const { data: score, isLoading: scoreLoading } = useHealthScore(assetId)
  const { data: att }  = useAttestation(assetId)

  return (
    <div className="space-y-4">
      <div className="flex items-center gap-3">
        {score && <HealthBadge score={score.healthScore} />}
        {scoreLoading && <span className="text-sm text-text-secondary">Loading…</span>}
        <span className="text-xs text-text-secondary">Computed {score ? new Date(score.computedAt).toLocaleTimeString() : '—'}</span>
      </div>

      {/* Score trend chart */}
      <div className="rounded-xl border border-surface-border bg-surface-card p-4">
        <p className="mb-3 text-xs font-medium text-text-secondary">Health Score — Last 60 min</p>
        <ResponsiveContainer width="100%" height={160}>
          <LineChart data={mockTrend}>
            <CartesianGrid strokeDasharray="3 3" stroke="#30363D" />
            <XAxis dataKey="t" tick={{ fontSize: 10, fill: '#8B949E' }} />
            <YAxis domain={[0, 100]} tick={{ fontSize: 10, fill: '#8B949E' }} />
            <Tooltip
              contentStyle={{ background: '#161B22', border: '1px solid #30363D', borderRadius: 8 }}
              labelStyle={{ color: '#8B949E' }}
              itemStyle={{ color: '#58A6FF' }}
            />
            <Line type="monotone" dataKey="score" stroke="#58A6FF" strokeWidth={2} dot={false} />
          </LineChart>
        </ResponsiveContainer>
      </div>

      {/* Latest attestation */}
      {att && (
        <div className="rounded-xl border border-surface-border bg-surface-card p-4">
          <p className="mb-2 text-xs font-medium text-text-secondary">Latest On-chain Attestation</p>
          <div className="space-y-1.5 text-xs">
            <div className="flex justify-between">
              <span className="text-text-secondary">Score</span>
              <span className="font-mono text-text-primary">{att.healthScore}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-text-secondary">Health hash</span>
              <span className="font-mono text-brand">{shortHash(att.healthHash)}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-text-secondary">XDC tx</span>
              <span className="font-mono text-brand">{shortHash(att.xdcTxHash)}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-text-secondary">Attested at</span>
              <span className="text-text-primary">{new Date(att.timestamp).toLocaleString()}</span>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export function Veriflow() {
  const { address, isConnected } = useAccount()
  const { data: assetsResp } = useAssets(address)
  const [selected, setSelected] = useState<string | null>(null)

  if (!isConnected) {
    return (
      <div className="flex flex-col items-center justify-center gap-4 py-24">
        <p className="text-text-secondary">Connect your wallet to view Veriflow dashboard.</p>
        <WalletButton />
      </div>
    )
  }

  const assets = assetsResp?.assets ?? []

  return (
    <div className="px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Veriflow</h1>

      <div className="flex gap-6">
        {/* Asset grid */}
        <div className="w-72 shrink-0 space-y-3">
          {assets.length === 0 && (
            <p className="text-sm text-text-secondary">No assets registered yet.</p>
          )}
          {assets.map((a) => (
            <HealthCard
              key={a.id}
              assetId={a.id}
              assetType={a.assetType}
              status={a.status}
              healthScore={a.healthScore}
              onClick={() => setSelected(a.id)}
            />
          ))}
        </div>

        {/* Detail panel */}
        <div className="flex-1">
          {selected ? (
            <AssetDetail assetId={selected} />
          ) : (
            <div className="flex h-48 items-center justify-center rounded-xl border border-surface-border bg-surface-card text-sm text-text-secondary">
              Select an asset to view details
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
