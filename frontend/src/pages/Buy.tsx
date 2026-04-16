import { useState } from 'react'
import { useAccount } from 'wagmi'
import { parseUnits, formatUnits } from 'viem'
import type { Address } from 'viem'
import { AmountInput } from '@/components/AmountInput'
import { DetailsPanel } from '@/components/DetailsPanel'
import { ActionButton } from '@/components/ActionButton'
import { WalletButton } from '@/components/WalletButton'
import { useMintWatt, useRedeemWatt, useWattBalance } from '@/hooks/contracts/useWattUSD'
import { useTxStore } from '@/stores/txStore'

type Mode = 'mint' | 'redeem'

// Testnet USDC/USDT addresses on Apothem — update after MockStablecoin deploy
const STABLECOINS: { label: string; address: Address }[] = [
  { label: 'USDC', address: '0x0000000000000000000000000000000000000000' },
  { label: 'USDT', address: '0x0000000000000000000000000000000000000000' },
]

const FEE_BPS = 10 // 0.1%

export function Buy() {
  const { address, isConnected } = useAccount()
  const [mode, setMode] = useState<Mode>('mint')
  const [amount, setAmount] = useState('')
  const [stablecoin, setStablecoin] = useState(STABLECOINS[0])

  const { data: wattBalance } = useWattBalance(address)
  const { mint, isPending: minting } = useMintWatt()
  const { redeem, isPending: redeeming } = useRedeemWatt()
  const add = useTxStore((s) => s.add)

  const parsedAmount = amount ? parseUnits(amount, 6) : 0n
  const fee = (parsedAmount * BigInt(FEE_BPS)) / 10_000n
  const received = parsedAmount - fee

  const handleSubmit = () => {
    const id = crypto.randomUUID()
    if (mode === 'mint') {
      add({ id, description: `Mint ${amount} WATT`, status: 'pending' })
      mint(stablecoin.address, parseUnits(amount, 6))
    } else {
      add({ id, description: `Redeem ${amount} WATT`, status: 'pending' })
      redeem(parseUnits(amount, 18))
    }
  }

  return (
    <div className="mx-auto max-w-lg px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Buy WATT</h1>

      {/* Mode tabs */}
      <div className="mb-6 flex rounded-xl border border-surface-border bg-surface-card p-1">
        {(['mint', 'redeem'] as Mode[]).map((m) => (
          <button
            key={m}
            onClick={() => setMode(m)}
            className={`flex-1 rounded-lg py-2 text-sm font-medium capitalize transition-colors ${
              mode === m ? 'bg-brand text-white' : 'text-text-secondary hover:text-text-primary'
            }`}
          >
            {m === 'mint' ? 'Mint WATT' : 'Redeem WATT'}
          </button>
        ))}
      </div>

      <div className="space-y-3">
        {mode === 'mint' && (
          <div className="flex gap-2">
            {STABLECOINS.map((s) => (
              <button
                key={s.label}
                onClick={() => setStablecoin(s)}
                className={`rounded-lg border px-4 py-2 text-sm font-medium transition-colors ${
                  stablecoin.label === s.label
                    ? 'border-brand bg-brand-muted text-brand'
                    : 'border-surface-border text-text-secondary hover:border-brand/50'
                }`}
              >
                {s.label}
              </button>
            ))}
          </div>
        )}

        <AmountInput
          label={mode === 'mint' ? `${stablecoin.label} to deposit` : 'WATT to redeem'}
          value={amount}
          onChange={setAmount}
          symbol={mode === 'mint' ? stablecoin.label : 'WATT'}
          max={wattBalance ? formatUnits(wattBalance, 18) : undefined}
        />

        <DetailsPanel
          rows={[
            { label: 'Protocol fee (0.1%)', value: amount ? `${formatUnits(fee, 6)} ${stablecoin.label}` : '—' },
            { label: 'You receive',          value: amount ? `${formatUnits(received, 6)} ${mode === 'mint' ? 'WATT' : stablecoin.label}` : '—', highlight: true },
            { label: 'WATT balance',         value: wattBalance ? `${parseFloat(formatUnits(wattBalance, 18)).toFixed(2)} WATT` : '—' },
          ]}
        />

        {!isConnected ? (
          <div className="flex justify-center pt-2">
            <WalletButton />
          </div>
        ) : (
          <ActionButton
            onClick={handleSubmit}
            isLoading={minting || redeeming}
            disabled={!amount || parsedAmount === 0n}
            className="w-full"
          >
            {mode === 'mint' ? 'Mint WATT' : 'Redeem WATT'}
          </ActionButton>
        )}
      </div>
    </div>
  )
}
