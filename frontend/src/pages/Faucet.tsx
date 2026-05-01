import { useEffect } from 'react'
import {
  useAccount,
  useBalance,
  useReadContract,
  useWriteContract,
  useWaitForTransactionReceipt,
  useConnect,
} from 'wagmi'
import { injected } from 'wagmi/connectors'
import { useUIStore } from '@/stores/uiStore'
import { fmtAddr }   from '@/lib/formatters'

// ── Contract addresses from env (set per environment in .env.*) ─────────────
const USDC_ADDRESS = (import.meta.env.VITE_USDC_ADDRESS ?? '0x6404e052B842fB1c82C1A404C7E02BE2A92A7970') as `0x${string}`
const USDT_ADDRESS = (import.meta.env.VITE_USDT_ADDRESS ?? '0x59828572B8ABEDa2A0f963ECC94b80Eee01bbA6c') as `0x${string}`

const CHAIN_ID     = Number(import.meta.env.VITE_CHAIN_ID ?? 51)
const EXPLORER_BASE = CHAIN_ID === 50
  ? 'https://xdcscan.com'
  : 'https://explorer.apothem.network'
const NETWORK_LABEL = CHAIN_ID === 50 ? 'XDC Mainnet · chainId 50' : 'XDC Apothem · chainId 51'

// Minimal ABI — balanceOf + faucet
const MOCK_ABI = [
  {
    name: 'balanceOf',
    type: 'function',
    stateMutability: 'view',
    inputs:  [{ name: 'account', type: 'address' }],
    outputs: [{ name: '',        type: 'uint256' }],
  },
  {
    name: 'faucet',
    type: 'function',
    stateMutability: 'nonpayable',
    inputs:  [],
    outputs: [],
  },
] as const

// Format 6-decimal token amount → "10,000.00"
function fmt6(raw: bigint | undefined): string {
  if (raw === undefined) return '—'
  const n = Number(raw) / 1e6
  return n.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })
}

// Format 18-decimal XDC amount → "1,234.56"
function fmtXdc(raw: bigint | undefined): string {
  if (raw === undefined) return '—'
  const n = Number(raw) / 1e18
  return n.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 4 })
}

// ── Sub-component: one faucet card ──────────────────────────────────────────
interface FaucetCardProps {
  symbol: 'USDC' | 'USDT'
  address: `0x${string}`
  walletAddress: `0x${string}` | undefined
  onSuccess: () => void
}

function FaucetCard({ symbol, address, walletAddress, onSuccess }: FaucetCardProps) {
  const { showToast } = useUIStore()
  const isUSDC = symbol === 'USDC'

  const { data: balance, refetch } = useReadContract({
    address,
    abi: MOCK_ABI,
    functionName: 'balanceOf',
    args: walletAddress ? [walletAddress] : undefined,
    query: { enabled: Boolean(walletAddress) },
  })

  const {
    writeContract,
    data: txHash,
    isPending: waitingWallet,
    reset,
  } = useWriteContract()

  const { isLoading: confirming, isSuccess: confirmed } = useWaitForTransactionReceipt({
    hash: txHash,
  })

  useEffect(() => {
    if (confirmed) {
      refetch()
      onSuccess()
      showToast(`10,000 ${symbol} added to your wallet`, 'success')
      reset()
    }
  }, [confirmed]) // eslint-disable-line react-hooks/exhaustive-deps

  const color  = isUSDC ? '#2563EB' : '#16A34A'
  const bg     = isUSDC ? '#EEF4FF' : '#F0FDF4'
  const border = isUSDC ? '#C7D7F8' : '#BBF7D0'

  const isBusy = waitingWallet || confirming
  const statusLabel = waitingWallet ? 'Check wallet…'
    : confirming ? 'Confirming…'
    : `Get 10,000 ${symbol}`

  return (
    <div className="bg-white border border-border" style={{ borderRadius: 4 }}>
      {/* Header */}
      <div className="flex items-center justify-between border-b border-border" style={{ padding: '14px 18px' }}>
        <div className="flex items-center gap-3">
          <div
            className="rounded-full flex items-center justify-center font-mono font-bold border-[1.5px]"
            style={{ width: 36, height: 36, fontSize: 9, color, background: bg, borderColor: border }}
          >
            {symbol}
          </div>
          <div>
            <div className="font-serif font-semibold text-text-1" style={{ fontSize: 15 }}>{symbol}</div>
            <div className="text-text-3" style={{ fontSize: 10 }}>Mock {symbol} · {NETWORK_LABEL}</div>
          </div>
        </div>
        <div className="text-right">
          <div className="font-mono font-medium text-text-1" style={{ fontSize: 18 }}>{fmt6(balance as bigint | undefined)}</div>
          <div className="text-text-3" style={{ fontSize: 10 }}>{symbol} balance</div>
        </div>
      </div>

      {/* Action */}
      <div style={{ padding: '12px 18px' }}>
        <button
          onClick={() => writeContract({ address, abi: MOCK_ABI, functionName: 'faucet' })}
          disabled={!walletAddress || isBusy}
          className="w-full font-sans font-semibold rounded cursor-pointer transition-all hover:-translate-y-px disabled:opacity-40 disabled:cursor-not-allowed disabled:transform-none"
          style={{
            padding: '10px',
            fontSize: 12,
            border: 'none',
            borderRadius: 4,
            background: color,
            color: '#fff',
            letterSpacing: '.04em',
          }}
        >
          {isBusy ? (
            <span className="flex items-center justify-center gap-2">
              <span
                style={{
                  width: 10, height: 10, border: '1.5px solid rgba(255,255,255,0.4)',
                  borderTopColor: '#fff', borderRadius: '50%',
                  animation: 'spin 0.7s linear infinite', display: 'inline-block',
                }}
              />
              {statusLabel}
            </span>
          ) : statusLabel}
        </button>
        {txHash && (
          <div className="text-center mt-2">
            <a
              href={`${EXPLORER_BASE}/tx/${txHash}`}
              target="_blank"
              rel="noopener noreferrer"
              className="font-mono text-text-3 hover:text-green transition-colors"
              style={{ fontSize: 10 }}
            >
              {fmtAddr(txHash)} ↗
            </a>
          </div>
        )}
      </div>
    </div>
  )
}

// ── Main page ────────────────────────────────────────────────────────────────
export function Faucet() {
  const { address, isConnected } = useAccount()
  const { connect }  = useConnect()
  const { showToast } = useUIStore()

  const { data: xdcBalance, refetch: refetchXdc } = useBalance({
    address,
    query: { enabled: Boolean(address) },
  })

  const { data: usdcBalance, refetch: refetchUsdc } = useReadContract({
    address: USDC_ADDRESS,
    abi: MOCK_ABI,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: Boolean(address) },
  })

  const { data: usdtBalance, refetch: refetchUsdt } = useReadContract({
    address: USDT_ADDRESS,
    abi: MOCK_ABI,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: Boolean(address) },
  })

  const refetchAll = () => {
    refetchXdc()
    refetchUsdc()
    refetchUsdt()
  }

  return (
    <>
      {/* spinner keyframe */}
      <style>{`@keyframes spin { to { transform: rotate(360deg) } }`}</style>

      <div className="animate-fadeup" style={{ maxWidth: 520, margin: '0 auto' }}>

        {/* Page header */}
        <div className="mb-6">
          <div className="flex items-center gap-3 mb-1">
            <h1 className="font-serif font-semibold text-text-1" style={{ fontSize: 20 }}>Testnet Faucet</h1>
            <span
              className="font-mono font-semibold"
              style={{ fontSize: 10, padding: '2px 8px', borderRadius: 3, background: '#FBF4E4', color: '#9A6B0A', border: '1px solid #E2C46A' }}
            >
              {NETWORK_LABEL}
            </span>
          </div>
          <p className="text-text-3" style={{ fontSize: 12 }}>
            Mint free test tokens to your wallet. For development only — never deploy mock contracts to mainnet.
          </p>
        </div>

        {/* Wallet + XDC balance card */}
        <div className="bg-white border border-border mb-4" style={{ borderRadius: 4 }}>
          <div className="border-b border-border font-semibold text-text-1" style={{ padding: '11px 18px', fontSize: 12 }}>
            Your Wallet
          </div>

          {isConnected && address ? (
            <div style={{ padding: '4px 18px 12px' }}>
              {/* Address */}
              <div className="flex items-center justify-between py-3 border-b border-border">
                <span className="text-text-3" style={{ fontSize: 11 }}>Address</span>
                <span className="font-mono text-text-1" style={{ fontSize: 12 }}>{address}</span>
              </div>

              {/* XDC balance */}
              <div className="flex items-center justify-between py-3 border-b border-border">
                <span className="flex items-center gap-2 text-text-3" style={{ fontSize: 11 }}>
                  <div
                    className="rounded-full flex items-center justify-center font-mono font-bold"
                    style={{ width: 18, height: 18, fontSize: 6, background: '#EEF4FF', color: '#2563EB', border: '1.5px solid #C7D7F8' }}
                  >XDC</div>
                  XDC (native)
                </span>
                <span className="font-mono font-medium text-text-1" style={{ fontSize: 13 }}>
                  {fmtXdc(xdcBalance?.value)} XDC
                </span>
              </div>

              {/* USDC balance */}
              <div className="flex items-center justify-between py-3 border-b border-border">
                <span className="flex items-center gap-2 text-text-3" style={{ fontSize: 11 }}>
                  <div
                    className="rounded-full flex items-center justify-center font-mono font-bold"
                    style={{ width: 18, height: 18, fontSize: 6, background: '#EEF4FF', color: '#2563EB', border: '1.5px solid #C7D7F8' }}
                  >USDC</div>
                  USDC (Mock)
                </span>
                <span className="font-mono font-medium text-text-1" style={{ fontSize: 13 }}>
                  {fmt6(usdcBalance as bigint | undefined)} USDC
                </span>
              </div>

              {/* USDT balance */}
              <div className="flex items-center justify-between py-3">
                <span className="flex items-center gap-2 text-text-3" style={{ fontSize: 11 }}>
                  <div
                    className="rounded-full flex items-center justify-center font-mono font-bold"
                    style={{ width: 18, height: 18, fontSize: 6, background: '#F0FDF4', color: '#16A34A', border: '1.5px solid #BBF7D0' }}
                  >USDT</div>
                  USDT (Mock)
                </span>
                <span className="font-mono font-medium text-text-1" style={{ fontSize: 13 }}>
                  {fmt6(usdtBalance as bigint | undefined)} USDT
                </span>
              </div>
            </div>
          ) : (
            <div className="text-center" style={{ padding: '28px 18px' }}>
              <div className="text-text-3 mb-4" style={{ fontSize: 12 }}>
                Connect your wallet to view balances and claim tokens
              </div>
              <button
                onClick={() => connect({ connector: injected() })}
                className="bg-green text-white font-sans font-semibold rounded cursor-pointer transition-opacity hover:opacity-85"
                style={{ padding: '10px 28px', fontSize: 12, border: 'none', borderRadius: 4, letterSpacing: '.03em' }}
              >
                Connect Wallet
              </button>
            </div>
          )}
        </div>

        {/* Faucet cards */}
        <div className="flex flex-col gap-3 mb-4">
          <FaucetCard
            symbol="USDC"
            address={USDC_ADDRESS}
            walletAddress={address}
            onSuccess={refetchAll}
          />
          <FaucetCard
            symbol="USDT"
            address={USDT_ADDRESS}
            walletAddress={address}
            onSuccess={refetchAll}
          />
        </div>

        {/* Contract addresses */}
        <div className="bg-white border border-border" style={{ borderRadius: 4 }}>
          <div className="border-b border-border font-semibold text-text-1" style={{ padding: '11px 18px', fontSize: 12 }}>
            Contract Addresses
          </div>
          <div style={{ padding: '4px 18px 10px' }}>
            {[
              { label: 'Mock USDC', address: USDC_ADDRESS },
              { label: 'Mock USDT', address: USDT_ADDRESS },
            ].map((c) => (
              <div key={c.label} className="flex items-center justify-between py-3 border-b border-border last:border-b-0">
                <span className="text-text-3" style={{ fontSize: 11 }}>{c.label}</span>
                <div className="flex items-center gap-2">
                  <a
                    href={`${EXPLORER_BASE}/address/${c.address}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="font-mono text-text-2 hover:text-green transition-colors"
                    style={{ fontSize: 11 }}
                    title={c.address}
                  >
                    {c.address}
                  </a>
                  <button
                    onClick={() => {
                      navigator.clipboard.writeText(c.address)
                      showToast(`${c.label} address copied`, 'success')
                    }}
                    className="text-text-3 hover:text-text-1 cursor-pointer transition-colors"
                    style={{ background: 'none', border: 'none', fontSize: 12, padding: 0 }}
                    title="Copy address"
                  >⧉</button>
                </div>
              </div>
            ))}
            <div className="mt-2 text-text-3" style={{ fontSize: 10 }}>
              Get free XDC for gas at{' '}
              <a
                href="https://faucet.apothem.network"
                target="_blank"
                rel="noopener noreferrer"
                className="text-green hover:underline"
              >
                faucet.apothem.network
              </a>
            </div>
          </div>
        </div>

      </div>
    </>
  )
}
