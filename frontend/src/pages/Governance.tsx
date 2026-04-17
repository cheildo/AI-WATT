import { PipelineCard } from '@/components/data/PipelineCard'
import { DataTable }    from '@/components/data/DataTable'
import { StatusDot }    from '@/components/data/StatusDot'
import { useUIStore }   from '@/stores/uiStore'

const PROPOSALS = [
  { proposal: 'Update mint fee to 0.08%',             type: 'Parameter', for: '12.4M WATT', against: '2.1M WATT', status: 'Active', dot: '#0A7068', deadline: '2d 14h' },
  { proposal: 'Whitelist curator: DataCenterAsia.io', type: 'Curator',   for: '8.7M WATT',  against: '0.9M WATT', status: 'Active', dot: '#0A7068', deadline: '4d 8h'  },
]

export function Governance() {
  const { showToast } = useUIStore()

  return (
    <div className="animate-fadeup">
      <PipelineCard
        title="$WATT Governance"
        cells={[
          { label: 'Your Voting Power',  value: '0',    sub: 'Connect wallet' },
          { label: 'Active Proposals',   value: '2',    sub: <span className="text-teal">Voting open</span> },
          { label: 'Quorum',             value: '4%',   sub: 'Of circulating supply' },
          { label: 'Timelock Delay',     value: '48h',  sub: 'After vote passes' },
        ]}
      />

      <div className="bg-white border border-border rounded">
        <div className="border-b border-border font-semibold" style={{ padding: '14px 18px', fontSize: 13 }}>
          Active Proposals
        </div>
        <DataTable
          columns={[
            { key: 'proposal', header: 'Proposal',
              render: (r) => <span className="font-sans text-text-1 font-medium text-xs">{String(r.proposal)}</span> },
            { key: 'type',    header: 'Type' },
            { key: 'for',     header: 'Votes For',
              render: (r) => <span className="text-teal">{String(r.for)}</span> },
            { key: 'against', header: 'Votes Against',
              render: (r) => <span className="text-red">{String(r.against)}</span> },
            { key: 'status',  header: 'Status',
              render: (r) => <StatusDot color={String(r.dot)} label={String(r.status)} /> },
            { key: 'deadline', header: 'Deadline',
              render: (r) => <span className="text-text-3">{String(r.deadline)}</span> },
          ]}
          rows={PROPOSALS as unknown as Record<string, unknown>[]}
          onRowClick={(r) => showToast(`${r.proposal}`)}
        />
      </div>
    </div>
  )
}
