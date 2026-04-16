import { useAccount } from 'wagmi'
import { formatUnits } from 'viem'
import { WalletButton } from '@/components/WalletButton'
import { useWattBalance } from '@/hooks/contracts/useWattUSD'
import { useSWattBalance } from '@/hooks/contracts/useSWattUSD'
import { useQueueStatus, useUserQueue } from '@/hooks/contracts/useWEVQueue'
import { useLoans } from '@/hooks/api/usePortfolio'

const STATUS_COLORS: Record<string, string> = {
  pending:    'text-warn',
  active:     'text-yield',
  repaying:   'text-brand',
  settled:    'text-text-secondary',
  defaulted:  'text-danger',
  liquidated: 'text-danger',
}

export function Portfolio() {
  const { address, isConnected } = useAccount()

  const { data: wattBal }    = useWattBalance(address)
  const { data: sWattBal }   = useSWattBalance(address)
  const { data: queueDepth } = useQueueStatus()
  const { data: wevRequests } = useUserQueue(address)
  const { data: loansResp }  = useLoans(address)

  if (!isConnected) {
    return (
      <div className="flex flex-col items-center justify-center gap-4 py-24">
        <p className="text-text-secondary">Connect your wallet to view your portfolio.</p>
        <WalletButton />
      </div>
    )
  }

  const fmt = (v: bigint | undefined, dec = 18) =>
    v !== undefined ? parseFloat(formatUnits(v, dec)).toLocaleString(undefined, { maximumFractionDigits: 4 }) : '—'

  const balances = [
    { label: 'WATT Balance',   value: `${fmt(wattBal)} WATT`,   color: 'text-text-primary' },
    { label: 'sWATT Balance',  value: `${fmt(sWattBal)} sWATT`, color: 'text-yield' },
    { label: 'WEV Queue',      value: `${wevRequests?.length ?? 0} requests`, color: 'text-brand' },
    { label: 'Queue Depth',    value: queueDepth?.toString() ?? '—', color: 'text-text-secondary' },
  ]

  return (
    <div className="mx-auto max-w-2xl px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Portfolio</h1>

      {/* Balance cards */}
      <div className="mb-8 grid grid-cols-2 gap-4">
        {balances.map((b) => (
          <div key={b.label} className="rounded-xl border border-surface-border bg-surface-card p-4">
            <p className="text-xs text-text-secondary">{b.label}</p>
            <p className={`mt-1 text-xl font-semibold ${b.color}`}>{b.value}</p>
          </div>
        ))}
      </div>

      {/* Open loans */}
      <div>
        <h2 className="mb-3 text-sm font-medium text-text-secondary">Open Loans</h2>
        {loansResp?.loans?.length ? (
          <div className="space-y-2">
            {loansResp.loans.map((l) => (
              <div key={l.id} className="flex items-center justify-between rounded-xl border border-surface-border bg-surface-card px-4 py-3">
                <div>
                  <p className="text-xs font-mono text-text-secondary">{l.id.slice(0, 8)}…</p>
                  <p className="text-sm text-text-primary">Engine {l.engineType}</p>
                </div>
                <span className={`text-sm font-medium capitalize ${STATUS_COLORS[l.status] ?? 'text-text-secondary'}`}>
                  {l.status}
                </span>
              </div>
            ))}
          </div>
        ) : (
          <p className="rounded-xl border border-surface-border bg-surface-card px-4 py-8 text-center text-sm text-text-secondary">
            No open loans.
          </p>
        )}
      </div>
    </div>
  )
}
