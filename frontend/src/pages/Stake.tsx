import { useState } from 'react'
import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TokenRow }           from '@/components/swap/TokenRow'
import { SwapDivider }        from '@/components/swap/SwapDivider'
import { ExchangeRow }        from '@/components/swap/ExchangeRow'
import { TransactionDetails } from '@/components/swap/TransactionDetails'
import { useUIStore }         from '@/stores/uiStore'
import { cn }                 from '@/lib/formatters'

const NAV = 1.0418

export function Stake() {
  const { showToast } = useUIStore()
  const [mode, setMode]       = useState<'stake' | 'unstake'>('stake')
  const [amount, setAmount]   = useState('')

  const parsed  = parseFloat(amount) || 0
  const stakeOut   = parsed > 0 ? (parsed / NAV).toFixed(3) : ''
  const unstakeOut = parsed > 0 ? (parsed * NAV).toFixed(2) : ''

  const isStake = mode === 'stake'

  const steps = [
    { label: isStake ? 'Approve WATT'  : 'Approve sWATT', status: 'active' as const },
    { label: isStake ? 'Stake WATT'    : 'Unstake sWATT', status: 'pending' as const },
  ]

  const info = [
    { key: 'NAV / sWATT',  value: '$1.0418',   variant: 'gold' as const },
    { key: '7d APR',       value: '12.81%',     variant: 'teal' as const },
    { key: '30d APR',      value: '11.94%',     variant: 'teal' as const },
    { key: 'Yield source', value: 'GPU loans' },
    { key: 'Unstake est.', value: '14–30 days' },
  ]

  return (
    <div className="animate-fadeup">
      <TabBar />
      <SwapWidget
        notice={
          <>
            The average sWATT unstaking period is{' '}
            <span className="text-teal font-semibold">14–30 days via WEV queue</span>.
            Priority exit available at 0.5% fee.
          </>
        }
        left={
          <>
            {/* Sub-tab bar */}
            <div
              className="border-b border-border overflow-hidden"
              style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', margin: '12px 18px 0' }}
            >
              {(['stake', 'unstake'] as const).map((m) => (
                <button
                  key={m}
                  onClick={() => { setMode(m); setAmount('') }}
                  className={cn(
                    'py-2 text-center font-semibold uppercase tracking-[.07em] border-b-2 cursor-pointer transition-colors',
                    mode === m
                      ? 'border-green text-green bg-green text-white border-green rounded-sm'
                      : 'border-transparent text-text-3 bg-transparent hover:text-text-2'
                  )}
                  style={{ fontSize: 11 }}
                >
                  {m.charAt(0).toUpperCase() + m.slice(1)}
                </button>
              ))}
            </div>

            {isStake ? (
              <>
                <TokenRow
                  token="WATT"
                  balance="3,200.00"
                  amount={amount}
                  onAmountChange={setAmount}
                  showMax
                  onMax={() => setAmount('3200')}
                  subValue={parsed > 0 ? `${stakeOut} sWATT` : '0.000 sWATT'}
                />
                <SwapDivider />
                <TokenRow
                  token="sWATT"
                  balance="3,068.42"
                  amount={stakeOut}
                  readOnly
                  subValue="sWattUSD"
                />
                <ExchangeRow left="Exchange" right="1 WATT = 0.9599 sWATT · NAV $1.0418" />
              </>
            ) : (
              <>
                <TokenRow
                  token="sWATT"
                  balance="3,068.42"
                  amount={amount}
                  onAmountChange={setAmount}
                  showMax
                  onMax={() => setAmount('3068.42')}
                  subValue={parsed > 0 ? `${unstakeOut} WATT` : '0.00 WATT'}
                />
                <SwapDivider />
                <TokenRow
                  token="WATT"
                  balance="3,200.00"
                  amount={unstakeOut}
                  readOnly
                />
                <ExchangeRow left="Exchange" right="1 sWATT = 1.0418 WATT" />
              </>
            )}
          </>
        }
        right={
          <TransactionDetails
            steps={steps}
            info={info}
            actionLabel={isStake ? 'Stake WATT' : 'Unstake sWATT'}
            onAction={() => showToast(isStake ? 'Staking submitted — sWATT arriving shortly' : 'Unstake queued via WEV')}
          />
        }
      />
    </div>
  )
}
