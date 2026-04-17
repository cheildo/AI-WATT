import { PipelineCard } from '@/components/data/PipelineCard'
import { DashCard }     from '@/components/data/DashCard'
import { DataTable }    from '@/components/data/DataTable'
import { StatusDot }    from '@/components/data/StatusDot'
import { useUIStore }   from '@/stores/uiStore'

const TVL_DATA  = [0,8,12,18,22,30,45,60,82,110,150,200,240,310,315]
const APR_DATA  = [4,4.5,5,5.8,6.2,7,8,9.5,10.2,11,11.8,12.4,12.8,12.81,12.81]

const LOAN_ROWS = [
  { asset: 'NVIDIA H200 [64]',       amount: '$84.2M',  util: '(8.9%)',  apr: '12.0%', contrib: '+1.07%', status: 'Active',  dot: '#2D5C3E' },
  { asset: 'NVIDIA H100 [128]',      amount: '$153.8M', util: '(16.2%)', apr: '9.0%',  contrib: '+1.12%', status: 'Active',  dot: '#2D5C3E' },
  { asset: 'NVIDIA B200 [32]',       amount: '$31.4M',  util: '(3.3%)',  apr: '15.0%', contrib: '+0.41%', status: 'Pre-PO', dot: '#B45309' },
  { asset: 'Industrial Robot [24]',  amount: '$29.4M',  util: '(3.1%)',  apr: '12.0%', contrib: '+0.30%', status: 'Active',  dot: '#2D5C3E' },
]

export function Dashboard() {
  const { showToast } = useUIStore()

  return (
    <div className="animate-fadeup">
      <PipelineCard
        title="Committed Capital Pipeline"
        cells={[
          {
            label: 'Total Committed',
            value: '$315.2M',
            sub: <span className="text-teal">14 Loans</span>,
            highlight: true,
          },
          {
            label: <><span className="rounded-full inline-block mr-1" style={{ width: 6, height: 6, background: '#2D5C3E', flexShrink: 0 }} />Escrowed</>,
            value: '$163.5M',
            sub: <><span className="text-teal">↑ 1.67%</span> · 4 Loans</>,
          },
          {
            label: <><span className="rounded-full inline-block mr-1" style={{ width: 6, height: 6, background: '#2563EB' }} />PO Placed</>,
            value: '$28.3M',
            sub: <><span className="text-teal">↑ 0.29%</span> · 3 Loans</>,
          },
          {
            label: <><span className="rounded-full inline-block mr-1" style={{ width: 6, height: 6, background: '#9A6B0A' }} />Pre-PO</>,
            value: '$123.4M',
            sub: <><span className="text-teal">↑ 3.39%</span> · 7 Loans</>,
          },
        ]}
      />

      {/* Charts */}
      <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: 16, marginBottom: 16 }}>
        <DashCard
          label="Total Value Locked ⓘ"
          value="$315,200,000"
          chartData={TVL_DATA}
          chartStroke="#3D7A52"
          chartFill="#EBF3EE"
        />
        <DashCard
          label="Staking APR (sWATT Staking Ratio)"
          value="12.81%"
          subValue="(76%)"
          chartData={APR_DATA}
          chartStroke="#0A7068"
          chartFill="#EAF5F3"
        />
      </div>

      {/* Loan table */}
      <div className="bg-white border border-border rounded mb-4">
        <div className="border-b border-border font-semibold flex items-center gap-[6px]" style={{ padding: '14px 18px', fontSize: 13 }}>
          <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
            <rect x="1" y="3" width="12" height="8" rx="1" stroke="#5A5646" strokeWidth="1.2"/>
            <path d="M4 3V2M7 3V2M10 3V2" stroke="#5A5646" strokeWidth="1.2" strokeLinecap="round"/>
          </svg>
          Loan Details
        </div>
        <DataTable
          columns={[
            { key: 'asset',  header: 'Asset',
              render: (r) => <span className="font-sans text-text-1 font-medium text-xs">{String(r.asset)}</span> },
            { key: 'amount', header: 'Amount (Utilization)',
              render: (r) => <>{r.amount} <span className="text-text-3">{String(r.util)}</span></> },
            { key: 'apr',    header: 'APR',
              render: (r) => <span className="text-teal">{String(r.apr)}</span> },
            { key: 'contrib', header: 'Contribution',
              render: (r) => <span className="text-teal">{String(r.contrib)}</span> },
            { key: 'status', header: 'Status',
              render: (r) => <StatusDot color={String(r.dot)} label={String(r.status)} /> },
          ]}
          rows={LOAN_ROWS as unknown as Record<string, unknown>[]}
          onRowClick={(r) => showToast(`${r.asset} — Engine 2`)}
        />
      </div>
    </div>
  )
}
