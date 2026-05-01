import { useState, useEffect }   from 'react'
import { AttestationStrip }      from '@/components/veriflow/AttestationStrip'
import { HealthCard }            from '@/components/veriflow/HealthCard'
import { Skeleton }              from '@/components/shared/Skeleton'
import { useUIStore }            from '@/stores/uiStore'

const ASSETS = [
  {
    assetId: 'ASSET-001', engineType: 2, name: 'H200 Cluster × 64',      location: 'Cyberjaya, Malaysia',
    healthScore: 98,
    metrics: [
      { label: 'GPU Util',     value: '94.2%',  variant: 'ok' as const },
      { label: 'Temp',         value: '72°C' },
      { label: 'Health Score', value: '98 / 100', variant: 'ok' as const },
      { label: 'Uptime',       value: '99.98%' },
    ],
    toast: 'ASSET-001 · H200 ×64 · Health 98/100 · Engine 2 · Cyberjaya',
  },
  {
    assetId: 'ASSET-002', engineType: 2, name: 'H100 Cluster × 128',     location: 'Kumamoto, Japan',
    healthScore: 95,
    metrics: [
      { label: 'GPU Util',     value: '87.1%',  variant: 'ok' as const },
      { label: 'Temp',         value: '68°C' },
      { label: 'Health Score', value: '95 / 100', variant: 'ok' as const },
      { label: 'Uptime',       value: '99.91%' },
    ],
    toast: 'ASSET-002 · H100 ×128 · Health 95/100 · Engine 2 · Kumamoto',
  },
  {
    assetId: 'ASSET-003', engineType: 1, name: 'B200 Cluster × 32',      location: 'Ho Chi Minh, Vietnam',
    healthScore: 64,
    metrics: [
      { label: 'GPU Util',     value: '61.4%',  variant: 'warn' as const },
      { label: 'Temp',         value: '84°C',   variant: 'warn' as const },
      { label: 'Health Score', value: '64 / 100', variant: 'warn' as const },
      { label: 'Uptime',       value: '97.42%' },
    ],
    toast: 'ASSET-003 · B200 ×32 · WARNING — Health 64/100 · High temp',
  },
  {
    assetId: 'ASSET-004', engineType: 2, name: 'Industrial Robot × 24', location: 'Singapore',
    healthScore: 93,
    metrics: [
      { label: 'Op. Rate',     value: '91.0%',  variant: 'ok' as const },
      { label: 'Errors',       value: '0.002%' },
      { label: 'Health Score', value: '93 / 100', variant: 'ok' as const },
      { label: 'Uptime',       value: '99.60%' },
    ],
    toast: 'ASSET-004 · Robot ×24 · Health 93/100 · Engine 2 · Singapore',
  },
  {
    assetId: 'ASSET-005', engineType: 1, name: 'B200 Cluster × 16',      location: 'Bangkok, Thailand',
    healthScore: 0, isPending: true,
    metrics: [
      { label: 'GPU Util', value: '—' },
      { label: 'Temp',     value: '—' },
      { label: 'Status',   value: 'PO Submitted' },
      { label: 'Engine',   value: 'Engine 1' },
    ],
    toast: 'ASSET-005 · B200 ×16 · Bangkok · Engine 1 — PO submitted',
  },
]

function CardSkeleton() {
  return (
    <div className="bg-white border border-border" style={{ borderRadius: 4, padding: '16px' }}>
      <div className="flex items-center justify-between mb-4">
        <div className="flex flex-col gap-2">
          <Skeleton width={120} height={11} />
          <Skeleton width={80}  height={9} />
        </div>
        <Skeleton width={36} height={36} radius={18} />
      </div>
      <Skeleton width="100%" height={3} radius={2} className="mb-4" />
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 8 }}>
        {Array.from({ length: 4 }).map((_, i) => (
          <div key={i} className="border border-border rounded" style={{ padding: '8px', borderRadius: 3 }}>
            <Skeleton width={50}  height={9}  className="mb-1" />
            <Skeleton width={65}  height={11} />
          </div>
        ))}
      </div>
    </div>
  )
}

export function Veriflow() {
  const { showToast } = useUIStore()
  const [loaded, setLoaded] = useState(false)

  useEffect(() => {
    const t = setTimeout(() => setLoaded(true), 700)
    return () => clearTimeout(t)
  }, [])

  return (
    <div className="animate-fadeup">
      <AttestationStrip
        assetCount={4}
        lastAttestation="4 min ago"
        blockNumber="82,441,209"
      />

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3,1fr)', gap: 12 }}>
        {loaded ? (
          <>
            {ASSETS.map((a) => (
              <HealthCard
                key={a.assetId}
                assetId={a.assetId}
                engineType={a.engineType}
                name={a.name}
                location={a.location}
                healthScore={a.healthScore}
                metrics={a.metrics}
                isPending={a.isPending}
                onClick={() => showToast(a.toast)}
              />
            ))}
            <HealthCard
              assetId=""
              engineType={0}
              name=""
              location=""
              healthScore={0}
              metrics={[]}
              isPlaceholder
            />
          </>
        ) : (
          Array.from({ length: 6 }).map((_, i) => <CardSkeleton key={i} />)
        )}
      </div>
    </div>
  )
}
