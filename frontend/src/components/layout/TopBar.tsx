import { useState, useEffect } from 'react'
import { useAccount, useConnect, useDisconnect } from 'wagmi'
import { injected } from 'wagmi/connectors'
import { fmtAddr } from '@/lib/formatters'
import { useWalletStore } from '@/stores/walletStore'
import { useUIStore }     from '@/stores/uiStore'

function useCountUp(target: number, duration: number, delay = 300) {
  const [val, setVal] = useState(0)
  useEffect(() => {
    let raf: number
    const start = performance.now() + delay
    const tick = (now: number) => {
      if (now < start) { raf = requestAnimationFrame(tick); return }
      const p = Math.min((now - start) / duration, 1)
      const ease = 1 - Math.pow(1 - p, 3)
      setVal(target * ease)
      if (p < 1) raf = requestAnimationFrame(tick)
    }
    raf = requestAnimationFrame(tick)
    return () => cancelAnimationFrame(raf)
  }, [target, duration, delay])
  return val
}

export function TopBar() {
  const { address, isConnected } = useAccount()
  const { connect }    = useConnect()
  const { disconnect } = useDisconnect()
  const { jwt, setJwt } = useWalletStore()
  const { showToast }   = useUIStore()

  const [signing, setSigning] = useState(false)

  const tvl    = useCountUp(315.2,  1800, 300)
  const apr    = useCountUp(12.81,  1400, 300)
  const expApr = useCountUp(15.40,  1600, 300)
  const wfv    = useCountUp(1.0000, 1200, 300)
  const sfv    = useCountUp(1.0418, 1500, 300)

  // Clear JWT when wallet disconnects
  useEffect(() => {
    if (!isConnected) setJwt(null)
  }, [isConnected, setJwt])

  const handleSignIn = async () => {
    setSigning(true)
    // Simulate EIP-191 sign + backend auth (1s round-trip)
    await new Promise((r) => setTimeout(r, 1000))
    setJwt('mock.jwt.token')
    setSigning(false)
    showToast('Wallet authenticated — session active', 'success')
  }

  const handleDisconnect = () => {
    setJwt(null)
    disconnect()
  }

  const isAuthenticated = isConnected && Boolean(jwt)

  return (
    <div
      className="bg-white border-b border-border flex items-center flex-shrink-0 z-50"
      style={{ height: 48, paddingRight: 16 }}
    >
      {/* Logo hex */}
      <div
        className="flex items-center justify-center flex-shrink-0 border-r border-border h-full"
        style={{ width: 56 }}
      >
        <div
          className="bg-green flex items-center justify-center cursor-pointer"
          style={{
            width: 28, height: 28,
            clipPath: 'polygon(50% 0%,93% 25%,93% 75%,50% 100%,7% 75%,7% 25%)',
          }}
        >
          <svg width="13" height="13" viewBox="0 0 14 14" fill="none">
            <path d="M7 2L11 4.5V9.5L7 12L3 9.5V4.5L7 2Z" fill="white" opacity=".35"/>
            <path d="M7 3L6 7.5H7.5L6.5 11L10 6.5H8.5L9.5 3H7Z" fill="white"/>
          </svg>
        </div>
      </div>

      {/* Stats bar */}
      <div className="flex items-center flex-1 h-full border-r border-border" style={{ padding: '0 20px' }}>
        <div className="flex items-center gap-[10px] h-full border-r border-border" style={{ padding: '0 18px 0 0' }}>
          <div className="text-[10px] text-text-3 uppercase tracking-[.08em] font-medium">TVL</div>
          <div className="font-mono text-[13px] font-medium text-gold">${tvl.toFixed(1)}M</div>
        </div>
        <div className="flex items-center gap-[10px] h-full border-r border-border" style={{ padding: '0 18px' }}>
          <div className="text-[10px] text-text-3 uppercase tracking-[.08em] font-medium">APR</div>
          <div className="font-mono text-[13px] font-medium text-teal">{apr.toFixed(2)}%</div>
        </div>
        <div className="flex items-center gap-[10px] h-full border-r border-border" style={{ padding: '0 18px' }}>
          <div className="text-[10px] text-text-3 uppercase tracking-[.08em] font-medium">Exp. APR</div>
          <div className="font-mono text-[13px] font-medium text-teal">{expApr.toFixed(2)}%</div>
        </div>
        <div className="flex items-center gap-[10px] h-full" style={{ padding: '0 18px' }}>
          <div className="text-[10px] text-text-3 uppercase tracking-[.08em] font-medium">Fair Value · WATT</div>
          <div className="flex items-center gap-2">
            <div className="flex items-center gap-1">
              <div className="rounded-full bg-gold-bg border border-gold-border flex items-center justify-center font-mono font-bold text-gold" style={{ width: 16, height: 16, fontSize: 5 }}>W</div>
              <span className="font-mono text-[13px] font-medium">${wfv.toFixed(4)}</span>
            </div>
            <div className="flex items-center gap-1">
              <div className="rounded-full bg-teal-bg border border-teal-border flex items-center justify-center font-mono font-bold text-teal" style={{ width: 16, height: 16, fontSize: 5 }}>sW</div>
              <span className="font-mono text-[13px] font-medium text-teal">${sfv.toFixed(4)}</span>
            </div>
          </div>
        </div>
      </div>

      {/* Right: chain badge + wallet */}
      <div className="flex items-center gap-[8px]" style={{ paddingLeft: 16 }}>
        {/* Chain badge */}
        <div className="flex items-center gap-[5px] border border-border rounded bg-bg font-mono text-[11px] text-text-2" style={{ padding: '5px 10px' }}>
          <div className="rounded-full" style={{ width: 5, height: 5, background: '#4CAF50', boxShadow: '0 0 0 2px #E8F5E9' }} />
          XDC Network
        </div>

        {isConnected && address ? (
          <div className="flex items-center gap-[6px]">
            {/* Address pill */}
            <button
              onClick={handleDisconnect}
              title="Click to disconnect"
              className="flex items-center gap-[7px] border border-border rounded bg-bg font-mono text-[12px] text-text-2 cursor-pointer hover:border-border-strong transition-colors"
              style={{ padding: '5px 12px' }}
            >
              {/* Auth status dot */}
              <div
                style={{
                  width: 6, height: 6, borderRadius: '50%', flexShrink: 0,
                  background: isAuthenticated ? '#0A7068' : '#D8D3C6',
                  boxShadow: isAuthenticated ? '0 0 0 2px #EAF5F3' : 'none',
                  transition: 'background 0.3s, box-shadow 0.3s',
                }}
              />
              {fmtAddr(address)}
            </button>

            {/* Sign-in chip — only shown when not yet authenticated */}
            {!isAuthenticated && (
              <button
                onClick={handleSignIn}
                disabled={signing}
                className="border border-border-strong rounded font-sans font-semibold text-text-2 cursor-pointer hover:bg-bg-2 hover:text-text-1 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
                style={{ padding: '5px 11px', fontSize: 11 }}
              >
                {signing ? 'Signing…' : 'Sign In'}
              </button>
            )}
          </div>
        ) : (
          <button
            onClick={() => connect({ connector: injected() })}
            className="bg-green text-white border-none rounded font-sans font-semibold cursor-pointer transition-opacity hover:opacity-85"
            style={{ padding: '7px 18px', fontSize: 12, letterSpacing: '.02em' }}
          >
            Connect Wallet
          </button>
        )}
      </div>
    </div>
  )
}
