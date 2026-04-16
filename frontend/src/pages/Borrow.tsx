import { useState } from 'react'
import { useAccount } from 'wagmi'
import { DetailsPanel } from '@/components/DetailsPanel'
import { ActionButton } from '@/components/ActionButton'
import { WalletButton } from '@/components/WalletButton'
import { AmountInput } from '@/components/AmountInput'

const ENGINES = [
  { id: 1, label: 'Engine 1 — Pre-delivery PO Financing', description: 'Finance hardware before delivery against purchase orders.' },
  { id: 2, label: 'Engine 2 — Productivity-backed Loans', description: 'Post-delivery loans backed by on-chain GPU productivity attestations.' },
  { id: 3, label: 'Engine 3 — Idle Capital Sweep',        description: 'Short-term yield on treasury idle capital.' },
]

export function Borrow() {
  const { isConnected } = useAccount()
  const [engine, setEngine]   = useState(2)
  const [assetId, setAssetId] = useState('')
  const [amount, setAmount]   = useState('')
  const [term, setTerm]       = useState('90')

  const interestRate = 0.12 // 12% APR (placeholder)
  const totalInterest = amount ? parseFloat(amount) * interestRate * (parseInt(term) / 365) : 0

  return (
    <div className="mx-auto max-w-lg px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Borrow</h1>

      {/* Engine selector */}
      <div className="mb-6 space-y-2">
        <p className="text-xs font-medium text-text-secondary">Select Engine</p>
        {ENGINES.map((e) => (
          <button
            key={e.id}
            onClick={() => setEngine(e.id)}
            className={`w-full rounded-xl border p-4 text-left transition-colors ${
              engine === e.id ? 'border-brand bg-brand-muted' : 'border-surface-border bg-surface-card hover:border-brand/40'
            }`}
          >
            <p className={`text-sm font-medium ${engine === e.id ? 'text-brand' : 'text-text-primary'}`}>
              {e.label}
            </p>
            <p className="mt-0.5 text-xs text-text-secondary">{e.description}</p>
          </button>
        ))}
      </div>

      <div className="space-y-3">
        {engine === 2 && (
          <>
            <div className="rounded-xl border border-surface-border bg-surface-card p-4">
              <p className="mb-2 text-xs text-text-secondary">Asset ID (from Veriflow dashboard)</p>
              <input
                value={assetId}
                onChange={(e) => setAssetId(e.target.value)}
                placeholder="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
                className="w-full bg-transparent font-mono text-sm text-text-primary outline-none placeholder:text-text-muted"
              />
            </div>
            <p className="rounded-xl border border-brand/30 bg-brand-muted px-4 py-3 text-xs text-brand">
              Engine 2 requires a Veriflow health attestation ≥ 60 and asset registered in AssetRegistry.
            </p>
          </>
        )}

        <AmountInput
          label="Loan amount"
          value={amount}
          onChange={setAmount}
          symbol="WATT"
        />

        <div className="rounded-xl border border-surface-border bg-surface-card p-4">
          <p className="mb-2 text-xs text-text-secondary">Loan term (days)</p>
          <div className="flex gap-2">
            {['30', '60', '90', '180', '365'].map((d) => (
              <button
                key={d}
                onClick={() => setTerm(d)}
                className={`rounded-lg px-3 py-1.5 text-sm transition-colors ${
                  term === d ? 'bg-brand text-white' : 'bg-surface-hover text-text-secondary hover:text-text-primary'
                }`}
              >
                {d}d
              </button>
            ))}
          </div>
        </div>

        <DetailsPanel
          rows={[
            { label: 'Interest rate (APR)', value: '12%' },
            { label: 'Term',               value: `${term} days` },
            { label: 'Total interest',     value: amount ? `${totalInterest.toFixed(2)} WATT` : '—' },
            { label: 'Total repayment',    value: amount ? `${(parseFloat(amount) + totalInterest).toFixed(2)} WATT` : '—', highlight: true },
          ]}
        />

        {!isConnected ? (
          <div className="flex justify-center pt-2"><WalletButton /></div>
        ) : (
          <ActionButton disabled={!amount || (engine === 2 && !assetId)} className="w-full">
            Apply for Loan
          </ActionButton>
        )}
      </div>
    </div>
  )
}
