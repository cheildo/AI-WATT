import { useState } from 'react'
import { useAccount } from 'wagmi'
import { formatUnits, parseUnits } from 'viem'
import { AmountInput } from '@/components/AmountInput'
import { DetailsPanel } from '@/components/DetailsPanel'
import { ActionButton } from '@/components/ActionButton'
import { WalletButton } from '@/components/WalletButton'
import { useWattBalance } from '@/hooks/contracts/useWattUSD'
import { useSWattBalance, useStakeWatt, useNAVPerShare, useVaultStats, useRequestUnstake } from '@/hooks/contracts/useSWattUSD'
import { useQueueStatus } from '@/hooks/contracts/useWEVQueue'
import { useTxStore } from '@/stores/txStore'

type Tab = 'stake' | 'unstake'

function fmt(v: bigint | undefined, dec = 18) {
  if (!v) return '0.00'
  return parseFloat(formatUnits(v, dec)).toLocaleString(undefined, { maximumFractionDigits: 4 })
}

export function Stake() {
  const { address, isConnected } = useAccount()
  const [tab, setTab] = useState<Tab>('stake')
  const [amount, setAmount] = useState('')

  const { data: wattBal }  = useWattBalance(address)
  const { data: sWattBal } = useSWattBalance(address)
  const { data: nav }      = useNAVPerShare()
  const { totalAssets }    = useVaultStats()
  const { data: queueDepth } = useQueueStatus()
  const { stake, isPending: staking }           = useStakeWatt()
  const { requestUnstake, isPending: unstaking } = useRequestUnstake()
  const add = useTxStore((s) => s.add)

  const parsed = amount ? parseUnits(amount, 18) : 0n
  const sharesOut = nav && parsed ? (parsed * parseUnits('1', 18)) / nav : 0n

  const handleSubmit = () => {
    const id = crypto.randomUUID()
    if (tab === 'stake' && address) {
      add({ id, description: `Stake ${amount} WATT`, status: 'pending' })
      stake(parsed, address)
    } else {
      add({ id, description: `Request unstake ${amount} sWATT`, status: 'pending' })
      requestUnstake(parsed)
    }
  }

  return (
    <div className="mx-auto max-w-lg px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Stake</h1>

      {/* Vault stats */}
      <div className="mb-6 grid grid-cols-3 gap-3">
        {[
          { label: 'TVL',           value: `${fmt(totalAssets.data)} WATT` },
          { label: 'NAV / sWATT',   value: `${fmt(nav)} WATT` },
          { label: 'Queue depth',   value: queueDepth?.toString() ?? '—' },
        ].map((s) => (
          <div key={s.label} className="rounded-xl border border-surface-border bg-surface-card p-3 text-center">
            <p className="text-xs text-text-secondary">{s.label}</p>
            <p className="mt-1 text-sm font-semibold text-text-primary">{s.value}</p>
          </div>
        ))}
      </div>

      {/* Tabs */}
      <div className="mb-4 flex rounded-xl border border-surface-border bg-surface-card p-1">
        {(['stake', 'unstake'] as Tab[]).map((t) => (
          <button
            key={t}
            onClick={() => setTab(t)}
            className={`flex-1 rounded-lg py-2 text-sm font-medium capitalize transition-colors ${
              tab === t ? 'bg-brand text-white' : 'text-text-secondary hover:text-text-primary'
            }`}
          >
            {t}
          </button>
        ))}
      </div>

      <div className="space-y-3">
        <AmountInput
          label={tab === 'stake' ? 'WATT to stake' : 'sWATT to unstake'}
          value={amount}
          onChange={setAmount}
          symbol={tab === 'stake' ? 'WATT' : 'sWATT'}
          max={tab === 'stake' ? fmt(wattBal) : fmt(sWattBal)}
        />

        <DetailsPanel
          rows={
            tab === 'stake'
              ? [
                  { label: 'You receive (est.)', value: amount ? `${fmt(sharesOut)} sWATT` : '—', highlight: true },
                  { label: 'WATT balance',       value: `${fmt(wattBal)} WATT` },
                ]
              : [
                  { label: 'Est. wait',          value: '~30 days (standard)' },
                  { label: 'sWATT balance',      value: `${fmt(sWattBal)} sWATT` },
                  { label: 'Queue depth',        value: queueDepth?.toString() ?? '—' },
                ]
          }
        />

        {tab === 'unstake' && (
          <p className="rounded-xl border border-warn/30 bg-warn/10 px-4 py-3 text-xs text-warn">
            Unstaking enters the WEV redemption queue (~30 days). Priority exit available for 0.5% fee.
          </p>
        )}

        {!isConnected ? (
          <div className="flex justify-center pt-2"><WalletButton /></div>
        ) : (
          <ActionButton
            onClick={handleSubmit}
            isLoading={staking || unstaking}
            disabled={!amount || parsed === 0n}
            className="w-full"
          >
            {tab === 'stake' ? 'Stake WATT' : 'Enter Unstake Queue'}
          </ActionButton>
        )}
      </div>
    </div>
  )
}
