import { useState } from 'react'
import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TokenRow }           from '@/components/swap/TokenRow'
import { ExchangeRow }        from '@/components/swap/ExchangeRow'
import { TransactionDetails } from '@/components/swap/TransactionDetails'
import { useUIStore }         from '@/stores/uiStore'
import { cn }                 from '@/lib/formatters'

const ENGINES = [
  { id: 1, label: 'Engine 1', title: 'Pre-Delivery',  sub: 'Finance your 30% purchase order deposit before hardware ships', fee: '2–3% origination',  feeRate: 0.025 },
  { id: 2, label: 'Engine 2', title: 'Post-Delivery', sub: 'Productivity-backed loan once hardware is deployed and running', fee: '1.5–2.5% AUM/yr',   feeRate: 0.020 },
  { id: 3, label: 'Engine 3', title: 'Reactivation',  sub: 'Idle escrow capital auto-deployed into T-bills during delivery window', fee: '~5% APY',     feeRate: 0 },
]

const ENGINE_NOTICES: Record<number, string> = {
  1: 'Engine 1: Veriflow certifies your PO. AI WATT finances 30% deposit. Hardware delivered, loan repaid over term.',
  2: 'Engine 2: Hardware is deployed. Veriflow monitors real-time telemetry. Productivity-backed loan against running asset.',
  3: 'Engine 3: Idle escrow capital during delivery window auto-deployed to T-bills. Generates additional base yield.',
}

const ENGINE_LABELS: Record<number, string> = {
  1: 'E1 — Pre-delivery · 2–3% origination',
  2: 'E2 — Post-delivery · 1.5–2.5% AUM/yr',
  3: 'E3 — Capital reactivation · ~5% APY',
}

export function Borrow() {
  const { showToast } = useUIStore()
  const [engine, setEngine]   = useState(1)
  const [amount, setAmount]   = useState('')

  const parsed    = parseFloat(amount) || 0
  const engData   = ENGINES.find((e) => e.id === engine)!
  const feeAmt    = (parsed * engData.feeRate).toFixed(2)
  const usdStr    = parsed > 0 ? '$' + parsed.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '$0.00'

  const steps = [
    { label: 'Veriflow certification', status: 'active' as const },
    { label: 'OC-NFT minted',         status: 'pending' as const },
    { label: 'Loan originated',        status: 'pending' as const },
  ]

  const info = [
    { key: 'Origination fee', value: `${feeAmt} WATT` },
    { key: 'Collateral',      value: 'OC-NFT + SPV' },
    { key: 'Term',            value: '3–36 months' },
    { key: 'Repayment',       value: 'Monthly' },
  ]

  return (
    <div className="animate-fadeup">
      <TabBar />

      {/* Engine selector */}
      <div
        className="mb-4 animate-fadeup"
        style={{ display: 'grid', gridTemplateColumns: 'repeat(3,1fr)', gap: 8, maxWidth: 820 }}
      >
        {ENGINES.map((e) => (
          <button
            key={e.id}
            onClick={() => setEngine(e.id)}
            className={cn(
              'border-[1.5px] rounded text-left cursor-pointer transition-colors',
              engine === e.id
                ? 'border-green-mid bg-green-bg'
                : 'border-border bg-white hover:border-border-strong hover:bg-bg'
            )}
            style={{ padding: 14, borderRadius: 4 }}
          >
            <div className="font-semibold uppercase tracking-[.08em] text-text-3 mb-1" style={{ fontSize: 11 }}>
              {e.label}
            </div>
            <div className="font-serif mb-[3px]" style={{ fontSize: 15 }}>{e.title}</div>
            <div className="text-text-3 leading-[1.4] mb-2" style={{ fontSize: 11 }}>{e.sub}</div>
            <div className="font-mono font-medium text-gold" style={{ fontSize: 12 }}>{e.fee}</div>
          </button>
        ))}
      </div>

      <SwapWidget
        notice={ENGINE_NOTICES[engine]}
        left={
          <>
            <TokenRow
              token="WATT"
              nameOverride="Loan Amount"
              chainOverride="Disbursed in WATT · Max LTV 70%"
              balance="Max available: $5,200,000"
              amount={amount}
              onAmountChange={setAmount}
              subValue={usdStr}
            />
            <ExchangeRow left="Engine" right={ENGINE_LABELS[engine]} />
          </>
        }
        right={
          <TransactionDetails
            steps={steps}
            info={info}
            actionLabel="Submit Application"
            actionVariant="gold"
            onAction={() => showToast('Application submitted — Veriflow certification begins')}
          />
        }
      />
    </div>
  )
}
