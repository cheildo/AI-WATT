import { useEffect, useState } from 'react'
import { useAccount, useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import { parseUnits, formatUnits } from 'viem'
import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TokenRow }           from '@/components/swap/TokenRow'
import { SwapDivider }        from '@/components/swap/SwapDivider'
import { ExchangeRow }        from '@/components/swap/ExchangeRow'
import { TransactionDetails } from '@/components/swap/TransactionDetails'
import { ProgressBar }        from '@/components/shared/ProgressBar'
import { ConfirmModal }       from '@/components/shared/ConfirmModal'
import { useUIStore }         from '@/stores/uiStore'
import { cn }                 from '@/lib/formatters'
import { useStakeSWatt, useSWattBalance, useNAVPerShare } from '@/hooks/contracts/useSWattUSD'
import { useWattBalance }     from '@/hooks/contracts/useWattUSD'
import { useRequestRedeem }   from '@/hooks/contracts/useWEVQueue'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { erc20Abi }           from '@/contracts/abis'

const WATT_DECIMALS  = 18
const SWATT_DECIMALS = 18
const ZERO_ADDR      = '0x0000000000000000000000000000000000000000'

const YIELD_SOURCES = [
  { label: 'GPU loan interest',  pct: '~8.5% APR' },
  { label: 'Origination fees',   pct: '~2.1% APR' },
  { label: 'T-Bill sweep yield', pct: '~2.3% APR' },
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

export function Stake() {
  const { showToast } = useUIStore()
  const { address }   = useAccount()
  const [mode, setMode]       = useState<'stake' | 'unstake'>('stake')
  const [amount, setAmount]   = useState('')
  const [modalOpen, setModal] = useState(false)

  const isStake   = mode === 'stake'
  const parsed    = parseFloat(amount) || 0
  const wevDeployed = CONTRACT_ADDRESSES.wevQueue !== ZERO_ADDR

  // On-chain data
  const { data: wattRaw,  refetch: refetchWatt  } = useWattBalance(address)
  const { data: swattRaw, refetch: refetchSwatt } = useSWattBalance(address)
  const { nav, isLoading: navLoading }             = useNAVPerShare()

  const wattBal  = wattRaw  !== undefined ? Number(formatUnits(wattRaw,  WATT_DECIMALS))  : 0
  const swattBal = swattRaw !== undefined ? Number(formatUnits(swattRaw, SWATT_DECIMALS)) : 0

  const wattDisp  = address ? wattBal.toLocaleString('en-US',  { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : '—'
  const swattDisp = address ? swattBal.toLocaleString('en-US', { minimumFractionDigits: 3, maximumFractionDigits: 3 }) : '—'

  const stakeOut   = parsed > 0 ? (parsed / nav).toFixed(3) : ''
  const unstakeOut = parsed > 0 ? (parsed * nav).toFixed(2)  : ''

  // Stake: need WATT allowance for sWattUSD
  const amountWattBig  = isStake  && parsed > 0 ? parseUnits(amount, WATT_DECIMALS)  : 0n
  const amountSwattBig = !isStake && parsed > 0 ? parseUnits(amount, SWATT_DECIMALS) : 0n

  const { data: wattAllowance, refetch: refetchWattAllowance } = useReadContract({
    address: CONTRACT_ADDRESSES.wattUSD,
    abi: erc20Abi,
    functionName: 'allowance',
    args: address ? [address, CONTRACT_ADDRESSES.sWattUSD] : undefined,
    query: { enabled: !!address && isStake },
  })

  const { data: swattAllowance, refetch: refetchSwattAllowance } = useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: erc20Abi,
    functionName: 'allowance',
    args: address ? [address, CONTRACT_ADDRESSES.wevQueue] : undefined,
    query: { enabled: !!address && !isStake && wevDeployed },
  })

  const needsApproval = isStake
    ? (wattAllowance !== undefined && amountWattBig > 0n && wattAllowance < amountWattBig)
    : (swattAllowance !== undefined && amountSwattBig > 0n && swattAllowance < amountSwattBig)

  // Write hooks
  const { stake, isPending: isStaking, isSuccess: stakeSuccess, error: stakeError } = useStakeSWatt()
  const { requestRedeem, isPending: isRedeeming, isSuccess: redeemSuccess, error: redeemError } = useRequestRedeem()
  const { writeContract: sendApprove, data: approveHash, isPending: isApproving, error: approveError } = useWriteContract()
  const { isSuccess: approveSuccess } = useWaitForTransactionReceipt({ hash: approveHash })

  // Auto-continue after approval
  useEffect(() => {
    if (!approveSuccess || !address) return
    if (isStake) {
      refetchWattAllowance()
      stake(amountWattBig, address)
    } else {
      refetchSwattAllowance()
      requestRedeem(amountSwattBig)
    }
  }, [approveSuccess])

  useEffect(() => {
    if (stakeSuccess) {
      showToast('Staking confirmed — sWATT arriving shortly', 'success')
      refetchWatt()
      refetchSwatt()
      setAmount('')
    }
  }, [stakeSuccess])

  useEffect(() => {
    if (redeemSuccess) {
      showToast('Unstake queued via WEV redemption queue', 'success')
      refetchSwatt()
      setAmount('')
    }
  }, [redeemSuccess])

  useEffect(() => {
    if (approveError) showToast(`Approval failed: ${approveError.message.slice(0, 80)}`, 'error')
  }, [approveError])
  useEffect(() => {
    if (stakeError) showToast(`Stake failed: ${stakeError.message.slice(0, 80)}`, 'error')
  }, [stakeError])
  useEffect(() => {
    if (redeemError) showToast(`Redeem failed: ${redeemError.message.slice(0, 80)}`, 'error')
  }, [redeemError])

  const balLimit = isStake ? wattBal : swattBal
  const token    = isStake ? 'WATT'  : 'sWATT'
  const isPending = isApproving || isStaking || isRedeeming

  const error =
    !address            ? `Connect your wallet to ${isStake ? 'stake' : 'unstake'}` :
    !isStake && !wevDeployed ? 'Redemption queue not yet deployed on this network' :
    amount && parsed <= 0 ? 'Enter a valid amount' :
    parsed > balLimit   ? `Insufficient ${token} balance` :
    undefined

  const canSubmit = !!address && parsed > 0 && !error && !isPending && (isStake || wevDeployed)

  const steps = isStake
    ? [
        { label: 'Approve WATT', status: (approveSuccess || !needsApproval ? 'done' : 'active') as 'done' | 'active' | 'pending' },
        { label: 'Stake WATT',   status: (stakeSuccess ? 'done' : isStaking ? 'active' : (!needsApproval || approveSuccess) ? 'active' : 'pending') as 'done' | 'active' | 'pending' },
      ]
    : [
        { label: 'Approve sWATT',  status: (approveSuccess || !needsApproval ? 'done' : 'active') as 'done' | 'active' | 'pending' },
        { label: 'Queue redeem',   status: (redeemSuccess ? 'done' : isRedeeming ? 'active' : (!needsApproval || approveSuccess) ? 'active' : 'pending') as 'done' | 'active' | 'pending' },
      ]

  const actionLabel =
    isApproving ? `Approving ${token}…` :
    isStaking   ? 'Staking WATT…' :
    isRedeeming ? 'Queuing redeem…' :
    needsApproval ? `Approve ${token}` :
    isStake ? 'Stake WATT' : 'Unstake sWATT'

  const navStr = navLoading ? '…' : `$${nav.toFixed(4)}`

  const info = [
    { key: 'NAV / sWATT',  value: navStr,    variant: 'gold' as const },
    { key: '7d APR',       value: '12.81%',  variant: 'teal' as const },
    { key: '30d APR',      value: '11.94%',  variant: 'teal' as const },
    { key: 'Yield source', value: 'GPU loans' },
    { key: 'Unstake est.', value: wevDeployed ? '14–30 days' : 'Queue not deployed' },
  ]

  const confirmRows = isStake
    ? [
        { label: 'You are staking',   value: `${parsed.toLocaleString('en-US', { minimumFractionDigits: 2 })} WATT` },
        { label: 'You will receive',  value: `${stakeOut || '0.000'} sWATT`, valueColor: '#0A7068' },
        { label: 'NAV',               value: navStr },
        { label: 'Estimated APR',     value: '12.81% (7d)', valueColor: '#0A7068' },
        { label: 'Unstake period',    value: '14–30 days' },
      ]
    : [
        { label: 'You are unstaking', value: `${parsed.toLocaleString('en-US', { minimumFractionDigits: 3 })} sWATT` },
        { label: 'You will receive',  value: `${unstakeOut || '0.00'} WATT`, valueColor: '#9A6B0A' },
        { label: 'NAV',               value: navStr },
        { label: 'Queue wait',        value: '~14–30 days' },
        { label: 'Priority exit',     value: '~3 days · 0.5% fee' },
      ]

  const onConfirm = () => {
    if (!address) return
    const token_addr  = isStake ? CONTRACT_ADDRESSES.wattUSD  : CONTRACT_ADDRESSES.sWattUSD
    const spender     = isStake ? CONTRACT_ADDRESSES.sWattUSD : CONTRACT_ADDRESSES.wevQueue
    const bigAmt      = isStake ? amountWattBig : amountSwattBig
    if (needsApproval) {
      sendApprove({ address: token_addr, abi: erc20Abi, functionName: 'approve', args: [spender, bigAmt] })
    } else if (isStake) {
      stake(bigAmt, address)
    } else {
      requestRedeem(bigAmt)
    }
  }

  return (
    <div className="animate-fadeup" style={{ maxWidth: 920, margin: '0 auto' }}>
      <ConfirmModal
        open={modalOpen}
        title={isStake ? 'Confirm Stake' : 'Confirm Unstake'}
        subtitle={needsApproval
          ? `Step 1 of 2: Approve ${token}, then ${isStake ? 'stake' : 'queue redeem'}`
          : isStake
            ? 'Your WATT will be deposited into the sWATT vault'
            : 'Your sWATT will enter the WEV redemption queue'}
        rows={confirmRows}
        actionLabel={needsApproval ? `Approve ${token}` : isStake ? 'Stake WATT' : 'Unstake sWATT'}
        onConfirm={onConfirm}
        onCancel={() => setModal(false)}
      />

      <div style={{ display: 'grid', gridTemplateColumns: '1fr 288px', gap: 16, alignItems: 'start' }}>
        <div>
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
                <div className="border-b border-border" style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', margin: '12px 18px 0' }}>
                  {(['stake', 'unstake'] as const).map((m) => (
                    <button
                      key={m}
                      onClick={() => { setMode(m); setAmount('') }}
                      className={cn(
                        'py-2 text-center font-semibold uppercase tracking-[.07em] cursor-pointer transition-colors border-b-2',
                        mode === m
                          ? 'border-green bg-green-bg text-green'
                          : 'border-transparent text-text-3 hover:text-text-2'
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
                      balance={wattDisp}
                      balanceNum={wattBal}
                      amount={amount}
                      onAmountChange={setAmount}
                      showMax
                      onMax={() => setAmount(wattBal.toFixed(2))}
                      subValue={parsed > 0 ? `${stakeOut} sWATT` : '0.000 sWATT'}
                      error={error}
                    />
                    <SwapDivider />
                    <TokenRow
                      token="sWATT"
                      balance={swattDisp}
                      amount={stakeOut}
                      readOnly
                      subValue="sWattUSD"
                    />
                    <ExchangeRow left="Exchange" right={`1 WATT = ${(1 / nav).toFixed(4)} sWATT · NAV ${navStr}`} />
                  </>
                ) : (
                  <>
                    <TokenRow
                      token="sWATT"
                      balance={swattDisp}
                      balanceNum={swattBal}
                      amount={amount}
                      onAmountChange={setAmount}
                      showMax
                      onMax={() => setAmount(swattBal.toFixed(3))}
                      subValue={parsed > 0 ? `${unstakeOut} WATT` : '0.00 WATT'}
                      error={error}
                    />
                    <SwapDivider />
                    <TokenRow
                      token="WATT"
                      balance={wattDisp}
                      amount={unstakeOut}
                      readOnly
                    />
                    <ExchangeRow left="Exchange" right={`1 sWATT = ${nav.toFixed(4)} WATT`} />
                  </>
                )}
              </>
            }
            right={
              <TransactionDetails
                steps={steps}
                info={info}
                actionLabel={actionLabel}
                disabled={!canSubmit}
                onAction={() => setModal(true)}
              />
            }
          />
        </div>

        <div className="flex flex-col gap-3" style={{ position: 'sticky', top: 0 }}>
          <Panel title="sWATT Vault">
            {[
              { k: 'NAV / sWATT',     v: navStr,      color: '#9A6B0A' },
              { k: '7d APR',          v: '12.81%',    color: '#0A7068' },
              { k: '30d APR',         v: '11.94%',    color: '#0A7068' },
              { k: 'Total deposited', v: '$84.2M' },
              { k: 'T-Bill reserve',  v: '23.6%' },
            ].map((s) => (
              <div key={s.k} className="flex items-center justify-between py-[9px] border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{s.k}</span>
                <span className="font-mono font-medium" style={{ fontSize: 12, color: s.color ?? '#1C1A14' }}>{s.v}</span>
              </div>
            ))}
            <div className="mt-3">
              <div className="flex justify-between text-text-3 mb-1" style={{ fontSize: 10 }}>
                <span>Capital deployed</span><span>76%</span>
              </div>
              <ProgressBar value={76} variant="teal" />
            </div>
          </Panel>

          <Panel title="Yield Sources">
            {YIELD_SOURCES.map((s) => (
              <div key={s.label} className="flex items-center justify-between py-[9px] border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{s.label}</span>
                <span className="font-mono text-teal font-medium" style={{ fontSize: 11 }}>{s.pct}</span>
              </div>
            ))}
          </Panel>

          <Panel title="WEV Redemption Queue">
            {[
              { k: 'Queue depth',     v: '$2.14M' },
              { k: 'Next processing', v: '6d 14h' },
              { k: 'Priority exit',   v: '~3 days · 0.5%' },
              { k: 'Standard exit',   v: '~30 days · free' },
            ].map((s) => (
              <div key={s.k} className="flex items-center justify-between py-[9px] border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{s.k}</span>
                <span className="font-mono font-medium text-text-1" style={{ fontSize: 12 }}>{s.v}</span>
              </div>
            ))}
            {!wevDeployed ? (
              <p className="text-text-3 mt-3" style={{ fontSize: 10 }}>
                WEV queue not yet deployed on this network.
              </p>
            ) : (
              <button
                onClick={() => showToast('Priority exit — 0.5% fee applied, processing in ~3 days', 'info')}
                className="mt-3 w-full border border-gold-border bg-gold-bg rounded text-gold font-semibold cursor-pointer hover:bg-[#f7ebc0] transition-colors"
                style={{ padding: '8px 12px', fontSize: 11, borderRadius: 4 }}
              >
                Priority exit (0.5% fee)
              </button>
            )}
          </Panel>
        </div>
      </div>
    </div>
  )
}
