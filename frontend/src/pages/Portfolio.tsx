import { PortfolioCard } from '@/components/portfolio/PortfolioCard'
import { YieldBanner }   from '@/components/portfolio/YieldBanner'
import { WEVOptions }    from '@/components/portfolio/WEVOptions'
import { ProgressBar }   from '@/components/shared/ProgressBar'
import { ActionButton }  from '@/components/shared/ActionButton'
import { DataTable }     from '@/components/data/DataTable'
import { StatusDot }     from '@/components/data/StatusDot'
import { useUIStore }    from '@/stores/uiStore'
import { useNavigate }   from 'react-router-dom'

function PosRow({ icon, name, sub, value, valueColor, usd }: { icon: React.ReactNode; name: string; sub: string; value: string; valueColor?: string; usd: string }) {
  return (
    <div className="flex items-center justify-between py-2 border-b border-border last:border-b-0">
      <div className="flex items-center gap-2">
        {icon}
        <div>
          <div className="font-semibold" style={{ fontSize: 13 }}>{name}</div>
          <div className="text-text-3" style={{ fontSize: 10 }}>{sub}</div>
        </div>
      </div>
      <div className="text-right">
        <div className="font-mono font-medium" style={{ fontSize: 13, color: valueColor ?? '#1C1A14' }}>{value}</div>
        <div className="font-mono text-text-3 text-right" style={{ fontSize: 10 }}>{usd}</div>
      </div>
    </div>
  )
}

const ACTIVITY_ROWS = [
  { type: 'Loan Repayment',    amount: '+$42,800',    notes: 'ASSET-001 H200 ×64',          engine: 'E2', status: 'Settled',   dot: '#0A7068', time: '2 min ago'  },
  { type: 'Mint WATT',         amount: '+500,000',    notes: 'Institutional deposit',        engine: '—',  status: 'Confirmed', dot: '#0A7068', time: '8 min ago'  },
  { type: 'Stake sWATT',       amount: '+480,307',    notes: 'NAV $1.0411',                 engine: '—',  status: 'Confirmed', dot: '#0A7068', time: '12 min ago' },
  { type: 'OC-NFT Minted',     amount: 'NFT #005',    notes: 'ASSET-005 B200 ×16 Bangkok',  engine: 'E1', status: 'Active',    dot: '#3B82F6', time: '34 min ago' },
  { type: 'Health Attestation',amount: '64 / 100',    notes: 'ASSET-003 — warning',         engine: '—',  status: 'Warning',   dot: '#B45309', time: '41 min ago' },
  { type: 'T-Bill Deploy',     amount: '$18,400,000', notes: 'Engine 3 idle sweep',         engine: 'E3', status: 'Deployed',  dot: '#0A7068', time: '2h ago'     },
]

export function Portfolio() {
  const { showToast } = useUIStore()
  const navigate = useNavigate()

  return (
    <div className="animate-fadeup">
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 16, marginBottom: 16 }}>
        {/* My Holdings */}
        <PortfolioCard title="My Holdings">
          <PosRow
            icon={<div className="rounded-full bg-gold-bg border border-gold-border flex items-center justify-center font-mono font-bold text-gold flex-shrink-0" style={{ width: 28, height: 28, fontSize: 8 }}>W</div>}
            name="WATT" sub="WattUSD"
            value="3,200.00" valueColor="#9A6B0A" usd="$3,200.00"
          />
          <PosRow
            icon={<div className="rounded-full bg-teal-bg border border-teal-border flex items-center justify-center font-mono font-bold text-teal flex-shrink-0" style={{ width: 28, height: 28, fontSize: 8 }}>sW</div>}
            name="sWATT" sub="sWattUSD vault"
            value="3,068.42" valueColor="#0A7068" usd="$3,196.24"
          />
          <YieldBanner amount="+$128.74" onClaim={() => showToast('Yield auto-compounds into sWATT — no claim needed')} />
        </PortfolioCard>

        {/* sWATT Vault */}
        <PortfolioCard title="sWATT Vault">
          {[
            { k: 'NAV / sWATT',      v: '$1.0418', color: '#9A6B0A' },
            { k: 'Total deposited',  v: '$84.2M'   },
            { k: 'Deployed',         v: '76.4%',   color: '#0A7068' },
            { k: 'T-Bill reserve',   v: '23.6%'    },
          ].map((item, i) => (
            <div key={i} className="flex items-center justify-between py-2 border-b border-border last:border-b-0" style={{ fontSize: 11 }}>
              <span className="text-text-3">{item.k}</span>
              <span className="font-mono font-medium" style={{ fontSize: 13, color: item.color ?? '#1C1A14' }}>{item.v}</span>
            </div>
          ))}
          <div className="mt-2">
            <div className="flex justify-between text-text-3 mb-1" style={{ fontSize: 10 }}>
              <span>Deployed</span><span>76%</span>
            </div>
            <ProgressBar value={76} variant="teal" />
          </div>
        </PortfolioCard>

        {/* WEV Redemption Queue */}
        <PortfolioCard title="WEV Redemption Queue">
          {[
            { k: 'Queue depth',     v: '$2.14M' },
            { k: 'Next processing', v: '6d 14h'  },
          ].map((item, i) => (
            <div key={i} className="flex items-center justify-between py-2 border-b border-border last:border-b-0" style={{ fontSize: 11 }}>
              <span className="text-text-3">{item.k}</span>
              <span className="font-mono" style={{ fontSize: 13 }}>{item.v}</span>
            </div>
          ))}
          <ProgressBar value={34} variant="gold" />
          <WEVOptions
            options={[
              { time: '~30d', label: 'Standard',  onClick: () => showToast('Standard queue — ~30 days, no fee') },
              { time: '~3d',  label: 'Priority',  isPriority: true, onClick: () => showToast('Priority exit — 0.5% fee, ~3 days') },
              { time: 'Now',  label: 'DEX swap',  onClick: () => showToast('Swap sWATT on DEX for instant exit') },
            ]}
          />
        </PortfolioCard>

        {/* Protocol Stats */}
        <PortfolioCard title="Protocol Stats">
          {[
            { k: 'Protocol TVL',  v: '$315.2M', color: '#9A6B0A' },
            { k: 'Active loans',  v: '14'       },
            { k: 'T-Bill reserve',v: '$19.8M',  color: '#0A7068' },
          ].map((item, i) => (
            <div key={i} className="flex items-center justify-between py-2 border-b border-border last:border-b-0" style={{ fontSize: 11 }}>
              <span className="text-text-3">{item.k}</span>
              <span className="font-mono font-medium" style={{ fontSize: 13, color: item.color ?? '#1C1A14' }}>{item.v}</span>
            </div>
          ))}
          <div className="mt-3">
            <ActionButton variant="outline" onClick={() => navigate('/dashboard')} className="text-[11px] py-2">
              View Dashboard →
            </ActionButton>
          </div>
        </PortfolioCard>
      </div>

      {/* Activity table */}
      <div className="bg-white border border-border rounded mb-4">
        <div className="border-b border-border font-semibold" style={{ padding: '14px 18px', fontSize: 13 }}>
          Recent Activity
        </div>
        <DataTable
          columns={[
            { key: 'type',   header: 'Type',
              render: (r) => <span className="font-sans text-text-1 font-medium text-xs">{String(r.type)}</span> },
            { key: 'amount', header: 'Amount',
              render: (r) => <span style={{ color: String(r.amount).startsWith('+') ? '#0A7068' : '#1C1A14' }}>{String(r.amount)}</span> },
            { key: 'notes',  header: 'Notes' },
            { key: 'engine', header: 'Engine',
              render: (r) => r.engine === '—' ? <>—</> : (
                <span className="font-sans font-semibold rounded" style={{ background: '#EAF5F3', color: '#0A7068', padding: '2px 6px', fontSize: 10 }}>{String(r.engine)}</span>
              ) },
            { key: 'status', header: 'Status',
              render: (r) => <StatusDot color={String(r.dot)} label={String(r.status)} /> },
            { key: 'time',   header: 'Time',
              render: (r) => <span className="text-text-3">{String(r.time)}</span> },
          ]}
          rows={ACTIVITY_ROWS as unknown as Record<string, unknown>[]}
          onRowClick={(r) => showToast(`${r.type} — ${r.notes}`)}
        />
      </div>
    </div>
  )
}
