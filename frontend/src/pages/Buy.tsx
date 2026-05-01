import { useEffect, useState } from 'react'
import { useAccount, useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import { parseUnits, formatUnits } from 'viem'
import type { Address } from 'viem'
import { TabBar }              from '@/components/swap/TabBar'
import { SwapWidget }          from '@/components/swap/SwapWidget'
import { TokenRow }            from '@/components/swap/TokenRow'
import { SwapDivider }         from '@/components/swap/SwapDivider'
import { ExchangeRow }         from '@/components/swap/ExchangeRow'
import { TransactionDetails }  from '@/components/swap/TransactionDetails'
import { ConfirmModal }        from '@/components/shared/ConfirmModal'
import { useUIStore }          from '@/stores/uiStore'
import { useMintWatt }         from '@/hooks/contracts/useWattUSD'
import { CONTRACT_ADDRESSES }  from '@/contracts/addresses'
import { erc20Abi }            from '@/contracts/abis'

const USDC_ADDRESS = import.meta.env.VITE_USDC_ADDRESS as Address

const STATS = [
  { k: 'WATT Price',    v: '$1.00',    color: '#9A6B0A' },
  { k: 'Protocol TVL', v: '$315.2M',  color: '#9A6B0A' },
  { k: 'Total Minted', v: '8.4M WATT' },
  { k: '24h Volume',   v: '$2.1M' },
  { k: 'Fee',          v: '0.10%' },
  { k: 'Settlement',   v: 'Instant' },
]

const HOW_IT_WORKS = [
  { n: 1, title: 'Deposit USDC',     sub: 'Send USDC from any EVM wallet' },
  { n: 2, title: 'Receive WATT 1:1', sub: 'Minus 0.1% protocol fee' },
  { n: 3, title: 'Stake for yield',  sub: 'Stake WATT → sWATT to earn GPU loan yield' },
]

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

export function Buy() {
  const { showToast } = useUIStore()
  const { address } = useAccount()
  const [amount, setAmount]       = useState('')
  const [modalOpen, setModalOpen] = useState(false)

  // USDC decimals — mock contracts on testnet may use 18
  const { data: usdcDecimals = 6 } = useReadContract({
    address: USDC_ADDRESS,
    abi: erc20Abi,
    functionName: 'decimals',
  })

  const { data: usdcBalanceRaw, refetch: refetchBalance } = useReadContract({
    address: USDC_ADDRESS,
    abi: erc20Abi,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: !!address },
  })

  const { data: allowanceRaw, refetch: refetchAllowance } = useReadContract({
    address: USDC_ADDRESS,
    abi: erc20Abi,
    functionName: 'allowance',
    args: address ? [address, CONTRACT_ADDRESSES.mintEngine] : undefined,
    query: { enabled: !!address },
  })

  const balanceNum = usdcBalanceRaw !== undefined
    ? Number(formatUnits(usdcBalanceRaw, usdcDecimals))
    : 0
  const balanceDisp = address
    ? balanceNum.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })
    : '—'

  const parsed       = parseFloat(amount) || 0
  const amountBigInt = parsed > 0 ? parseUnits(amount, usdcDecimals) : 0n
  const needsApproval = allowanceRaw !== undefined && amountBigInt > 0n && allowanceRaw < amountBigInt

  const { mint, isPending: isMinting, isSuccess: mintSuccess, error: mintError } = useMintWatt()
  const { writeContract: sendApprove, data: approveHash, isPending: isApproving, error: approveError } = useWriteContract()
  const { isSuccess: approveSuccess } = useWaitForTransactionReceipt({ hash: approveHash })

  // Auto-mint after approval is confirmed on-chain
  useEffect(() => {
    if (approveSuccess) {
      refetchAllowance()
      mint(USDC_ADDRESS, amountBigInt)
    }
  }, [approveSuccess])

  useEffect(() => {
    if (mintSuccess) {
      showToast('WATT minted successfully — check your wallet', 'success')
      refetchBalance()
      setAmount('')
    }
  }, [mintSuccess])

  useEffect(() => {
    if (approveError) showToast(`Approval failed: ${approveError.message.slice(0, 80)}`, 'error')
  }, [approveError])

  useEffect(() => {
    if (mintError) showToast(`Mint failed: ${mintError.message.slice(0, 80)}`, 'error')
  }, [mintError])

  const fee    = parsed * 0.001
  const outAmt = parsed > 0 ? (parsed - fee).toFixed(2) : ''
  const feeStr = fee > 0 ? fee.toFixed(4) : '0.00'
  const usdStr = parsed > 0
    ? '$' + parsed.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })
    : '$0.00'

  const isPending = isApproving || isMinting

  const error =
    !address                ? 'Connect your wallet to buy WATT' :
    amount && parsed <= 0   ? 'Enter a valid amount' :
    parsed > balanceNum     ? 'Insufficient USDC balance' :
    undefined

  const canSubmit = !!address && parsed > 0 && !error && !isPending

  const steps = [
    {
      label: 'Approve USDC',
      status: (approveSuccess || !needsApproval ? 'done' : 'active') as 'done' | 'active' | 'pending',
    },
    {
      label: 'Mint WATT',
      status: (mintSuccess ? 'done' : isMinting ? 'active' : (!needsApproval || approveSuccess) ? 'active' : 'pending') as 'done' | 'active' | 'pending',
    },
  ]

  const actionLabel =
    isApproving ? 'Approving USDC…' :
    isMinting   ? 'Minting WATT…'  :
    needsApproval ? 'Approve & Buy WATT' :
    'Buy WATT'

  const confirmRows = [
    { label: 'You are paying',   value: `${parsed.toLocaleString('en-US', { minimumFractionDigits: 2 })} USDC` },
    { label: 'You will receive', value: `${outAmt || '0.00'} WATT`, valueColor: '#9A6B0A' },
    { label: 'Protocol fee',     value: `${feeStr} USDC (0.1%)` },
    { label: 'Rate',             value: '1 USDC = 1 WATT' },
    { label: 'Settlement',       value: 'Instant' },
  ]

  const onConfirm = () => {
    if (!address) return
    if (needsApproval) {
      sendApprove({
        address: USDC_ADDRESS,
        abi: erc20Abi,
        functionName: 'approve',
        args: [CONTRACT_ADDRESSES.mintEngine, amountBigInt],
      })
    } else {
      mint(USDC_ADDRESS, amountBigInt)
    }
  }

  return (
    <div className="animate-fadeup" style={{ maxWidth: 920, margin: '0 auto' }}>
      <ConfirmModal
        open={modalOpen}
        title="Confirm Purchase"
        subtitle={needsApproval
          ? 'Step 1 of 2: Approve USDC spend, then mint WATT'
          : 'Review your transaction before submitting'}
        rows={confirmRows}
        actionLabel={needsApproval ? 'Approve USDC' : 'Buy WATT'}
        onConfirm={onConfirm}
        onCancel={() => setModalOpen(false)}
      />

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 288px', gap: 16, alignItems: 'start' }}>
        <div>
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
                  balance={balanceDisp}
                  balanceNum={balanceNum}
                  amount={amount}
                  onAmountChange={setAmount}
                  showMax
                  onMax={() => setAmount(balanceNum.toFixed(usdcDecimals))}
                  subValue={usdStr}
                  error={error}
                />
                <SwapDivider />
                <TokenRow
                  token="WATT"
                  balance="—"
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
                info={[
                  { key: 'Slippage', value: '0.1%' },
                  { key: 'Fee',      value: `${feeStr} USDC` },
                  { key: 'Redeem',   value: 'Instant', variant: 'teal' as const },
                ]}
                actionLabel={actionLabel}
                disabled={!canSubmit}
                onAction={() => setModalOpen(true)}
              />
            }
          />
        </div>

        <div className="flex flex-col gap-3" style={{ position: 'sticky', top: 0 }}>
          <Panel title="Protocol Stats">
            {STATS.map((s) => (
              <div key={s.k} className="flex items-center justify-between py-[9px] border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{s.k}</span>
                <span className="font-mono font-medium" style={{ fontSize: 12, color: s.color ?? '#1C1A14' }}>{s.v}</span>
              </div>
            ))}
          </Panel>

          <Panel title="How It Works">
            <div className="flex flex-col gap-3 pt-2">
              {HOW_IT_WORKS.map((step) => (
                <div key={step.n} className="flex items-start gap-3">
                  <div
                    className="bg-green text-white font-mono font-bold flex-shrink-0 flex items-center justify-center"
                    style={{ width: 20, height: 20, borderRadius: '50%', fontSize: 9, marginTop: 1 }}
                  >{step.n}</div>
                  <div>
                    <div className="font-semibold text-text-1" style={{ fontSize: 12 }}>{step.title}</div>
                    <div className="text-text-3 leading-snug" style={{ fontSize: 11 }}>{step.sub}</div>
                  </div>
                </div>
              ))}
            </div>
          </Panel>

          <Panel title="About WATT">
            <p className="text-text-3 leading-relaxed" style={{ fontSize: 11, paddingTop: 6 }}>
              WATT is the protocol stablecoin pegged 1:1 to USD, backed by a portfolio of GPU
              hardware loans and US T-bill reserves. It is fully redeemable at any time.
            </p>
            <button
              onClick={() => showToast('Navigate to the Stake tab to earn sWATT yield', 'info')}
              className="mt-3 w-full border border-border rounded text-text-2 font-semibold cursor-pointer hover:border-border-strong hover:text-text-1 transition-colors"
              style={{ padding: '8px 12px', fontSize: 11, borderRadius: 4 }}
            >
              Stake WATT → Earn yield
            </button>
          </Panel>
        </div>
      </div>
    </div>
  )
}
