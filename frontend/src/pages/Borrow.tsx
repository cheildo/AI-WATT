import { useState } from 'react'
import { useAccount } from 'wagmi'
import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TokenRow }           from '@/components/swap/TokenRow'
import { ExchangeRow }        from '@/components/swap/ExchangeRow'
import { TransactionDetails } from '@/components/swap/TransactionDetails'
import { ConfirmModal }       from '@/components/shared/ConfirmModal'
import { useUIStore }         from '@/stores/uiStore'
import { cn }                 from '@/lib/formatters'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'

const MAX_LOAN = 5_200_000
const ZERO_ADDR = '0x0000000000000000000000000000000000000000'

const ENGINES = [
  {
    id: 1, label: 'Engine 1', title: 'Pre-Delivery',
    sub: 'Finance your 30% purchase order deposit before hardware ships',
    fee: '2–3% origination', feeRate: 0.025,
  },
  {
    id: 2, label: 'Engine 2', title: 'Post-Delivery',
    sub: 'Productivity-backed loan once hardware is deployed and running',
    fee: '1.5–2.5% AUM/yr', feeRate: 0.020,
  },
  {
    id: 3, label: 'Engine 3', title: 'Reactivation',
    sub: 'Idle escrow capital auto-deployed into T-bills during delivery window',
    fee: '~5% APY', feeRate: 0,
  },
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

function Panel({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <div className="bg-white border border-border" style={{ borderRadius: 4 }}>
      <div className="border-b border-border font-semibold text-text-1" style={{ padding: '11px 16px', fontSize: 12 }}>
        {title}
      </div>
      <div style={{ padding: '4px 16px 12px' }}>{children}</div>
    </div>
  )
}

function CheckItem({ label, met }: { label: string; met: boolean }) {
  return (
    <div className="flex items-start gap-2 py-[7px] border-b border-border last:border-b-0">
      <span
        className="flex-shrink-0 flex items-center justify-center rounded-full font-bold"
        style={{
          width: 16, height: 16, fontSize: 8, marginTop: 1,
          background: met ? '#EAF5F3' : '#FBF0F0',
          color: met ? '#0A7068' : '#8B2020',
          border: `1.5px solid ${met ? '#80C9C0' : '#EBADAD'}`,
        }}
      >{met ? '✓' : '✗'}</span>
      <span className="text-text-2 leading-snug" style={{ fontSize: 11 }}>{label}</span>
    </div>
  )
}

export function Borrow() {
  const { showToast } = useUIStore()
  const { address }   = useAccount()
  const [engine, setEngine]   = useState(1)
  const [amount, setAmount]   = useState('')
  const [modalOpen, setModal] = useState(false)
  const [isSubmitting, setSubmitting] = useState(false)

  const lendingDeployed = CONTRACT_ADDRESSES.lendingPool !== ZERO_ADDR

  const parsed  = parseFloat(amount) || 0
  const engData = ENGINES.find((e) => e.id === engine)!
  const feeAmt  = (parsed * engData.feeRate).toFixed(2)
  const usdStr  = parsed > 0
    ? '$' + parsed.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })
    : '$0.00'

  const error =
    !address               ? 'Connect your wallet to apply for a loan' :
    !lendingDeployed       ? 'Lending pool not yet deployed on this network' :
    amount && parsed <= 0  ? 'Enter a valid loan amount' :
    parsed > MAX_LOAN      ? `Exceeds maximum available ($${MAX_LOAN.toLocaleString()})` :
    undefined

  const canSubmit = !!address && lendingDeployed && parsed > 0 && !error && !isSubmitting

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

  const confirmRows = [
    { label: 'Engine',          value: `Engine ${engine} — ${engData.title}` },
    { label: 'Loan amount',     value: `${parsed.toLocaleString('en-US', { minimumFractionDigits: 2 })} WATT`, valueColor: '#9A6B0A' },
    { label: 'Origination fee', value: `${feeAmt} WATT (${(engData.feeRate * 100).toFixed(1)}%)` },
    { label: 'Collateral',      value: 'OC-NFT + SPV entity' },
    { label: 'Term',            value: '3–36 months' },
    { label: 'Repayment',       value: 'Monthly' },
  ]

  const onConfirm = async () => {
    if (!address || !lendingDeployed) return
    setSubmitting(true)
    try {
      const token = localStorage.getItem('aiwatt_jwt') ?? ''
      const res = await fetch(`${import.meta.env.VITE_API_URL}/api/v1/loans`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          ...(token ? { Authorization: `Bearer ${token}` } : {}),
        },
        body: JSON.stringify({
          amount: parsed,
          engine_type: engine,
          borrower_id: address,
        }),
      })
      if (!res.ok) {
        const body = await res.json().catch(() => ({}))
        throw new Error(body?.error ?? `HTTP ${res.status}`)
      }
      showToast('Application submitted — Veriflow certification begins', 'success')
      setAmount('')
    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : 'Unknown error'
      showToast(`Submission failed: ${msg}`, 'error')
    } finally {
      setSubmitting(false)
    }
  }

  const REQUIREMENTS = [
    { label: 'Veriflow health score ≥ 60',    met: true  },
    { label: 'OC-NFT as collateral',           met: true  },
    { label: 'KYC / AML verified',             met: true  },
    { label: 'No double-encumbrance on asset', met: false },
  ]

  const LOAN_STATS = [
    { k: 'Active loans',  v: '14' },
    { k: 'Total loaned',  v: '$42.8M', color: '#9A6B0A' },
    { k: 'Avg. health',   v: '78 / 100', color: '#0A7068' },
    { k: 'Avg. term',     v: '18 months' },
    { k: 'Max LTV',       v: '70%' },
  ]

  return (
    <div className="animate-fadeup" style={{ maxWidth: 920, margin: '0 auto' }}>
      <ConfirmModal
        open={modalOpen}
        title="Confirm Loan Application"
        subtitle="Veriflow will begin certification after submission"
        rows={confirmRows}
        actionLabel="Submit Application"
        actionVariant="gold"
        onConfirm={onConfirm}
        onCancel={() => setModal(false)}
      />

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 288px', gap: 16, alignItems: 'start' }}>
        <div>
          <TabBar />

          <div className="mb-4" style={{ display: 'grid', gridTemplateColumns: 'repeat(3,1fr)', gap: 8 }}>
            {ENGINES.map((e) => (
              <button
                key={e.id}
                onClick={() => setEngine(e.id)}
                className={cn(
                  'border-[1.5px] text-left cursor-pointer transition-colors',
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
                  balance={lendingDeployed ? `Max available: $${MAX_LOAN.toLocaleString()}` : 'Lending pool not deployed'}
                  amount={amount}
                  onAmountChange={setAmount}
                  subValue={usdStr}
                  error={error}
                />
                <ExchangeRow left="Engine" right={ENGINE_LABELS[engine]} />
              </>
            }
            right={
              <TransactionDetails
                steps={steps}
                info={info}
                actionLabel={isSubmitting ? 'Submitting…' : 'Submit Application'}
                actionVariant="gold"
                disabled={!canSubmit}
                onAction={() => setModal(true)}
              />
            }
          />
        </div>

        <div className="flex flex-col gap-3" style={{ position: 'sticky', top: 0 }}>
          <Panel title="Requirements">
            {REQUIREMENTS.map((r) => <CheckItem key={r.label} label={r.label} met={r.met} />)}
            {!lendingDeployed && (
              <p className="text-text-3 mt-2" style={{ fontSize: 10 }}>
                Lending pool not yet deployed on this network.
              </p>
            )}
          </Panel>

          <Panel title="Protocol Loans">
            {LOAN_STATS.map((s) => (
              <div key={s.k} className="flex items-center justify-between py-[9px] border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{s.k}</span>
                <span className="font-mono font-medium" style={{ fontSize: 12, color: s.color ?? '#1C1A14' }}>{s.v}</span>
              </div>
            ))}
          </Panel>

          <Panel title="Engine Guide">
            <div className="flex flex-col gap-2 pt-1">
              {ENGINES.map((e) => (
                <div
                  key={e.id}
                  className={cn(
                    'border rounded cursor-pointer transition-colors',
                    engine === e.id ? 'border-green-mid bg-green-bg' : 'border-border hover:border-border-strong'
                  )}
                  style={{ padding: '8px 10px', borderRadius: 4 }}
                  onClick={() => setEngine(e.id)}
                >
                  <div className="flex items-center justify-between mb-[2px]">
                    <span className="font-semibold text-text-2" style={{ fontSize: 11 }}>{e.title}</span>
                    <span className="font-mono text-gold" style={{ fontSize: 10 }}>{e.fee}</span>
                  </div>
                  <div className="text-text-3 leading-snug" style={{ fontSize: 10 }}>{e.sub}</div>
                </div>
              ))}
            </div>
          </Panel>
        </div>
      </div>
    </div>
  )
}
