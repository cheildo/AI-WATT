import { useState } from 'react'
import { TabBar }              from '@/components/swap/TabBar'
import { SwapWidget }          from '@/components/swap/SwapWidget'
import { TokenRow }            from '@/components/swap/TokenRow'
import { SwapDivider }         from '@/components/swap/SwapDivider'
import { ExchangeRow }         from '@/components/swap/ExchangeRow'
import { TransactionDetails }  from '@/components/swap/TransactionDetails'
import { useUIStore }          from '@/stores/uiStore'

export function Buy() {
  const { showToast } = useUIStore()
  const [amount, setAmount] = useState('')

  const parsed   = parseFloat(amount) || 0
  const fee      = parsed * 0.001
  const outAmt   = parsed > 0 ? (parsed - fee).toFixed(2) : ''
  const feeStr   = fee > 0 ? fee.toFixed(4) : '0.00'
  const usdStr   = parsed > 0 ? '$' + parsed.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '$0.00'

  const steps = [
    { label: 'Approve USDC', status: 'done' as const },
    { label: 'Buy WATT',     status: 'active' as const },
  ]

  const info = [
    { key: 'Slippage', value: '0.1%' },
    { key: 'Fee',      value: `${feeStr} USDC` },
    { key: 'Redeem',   value: 'Instant', variant: 'teal' as const },
  ]

  return (
    <div className="animate-fadeup">
      <TabBar />
      <SwapWidget
        notice={
          <>
            WATT does not accrue yield. Stake WATT →{' '}
            <span className="text-teal font-semibold">sWATT</span>{' '}
            to earn GPU infrastructure yield from real hardware loans.
          </>
        }
        left={
          <>
            <TokenRow
              token="USDC"
              balance="12,450.00"
              amount={amount}
              onAmountChange={setAmount}
              showMax
              onMax={() => setAmount('12450')}
              subValue={usdStr}
            />
            <SwapDivider />
            <TokenRow
              token="WATT"
              balance="3,200.00"
              amount={outAmt}
              readOnly
              subValue="WattUSD"
            />
            <ExchangeRow left="Exchange" right={`1 USDC = 1 WATT · ${feeStr} fee`} />
          </>
        }
        right={
          <TransactionDetails
            steps={steps}
            info={info}
            actionLabel="Buy WATT"
            onAction={() => showToast('WATT minted successfully')}
          />
        }
      />
    </div>
  )
}
