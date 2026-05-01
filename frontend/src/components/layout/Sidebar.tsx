import { useAccount, useDisconnect } from 'wagmi'
import { NavIcon } from './NavIcon'
import { fmtAddr } from '@/lib/formatters'

function Sep() {
  return <div className="bg-border my-[6px]" style={{ width: 32, height: 1 }} />
}

function WalletFooter() {
  const { address, isConnected } = useAccount()
  const { disconnect } = useDisconnect()

  if (!isConnected || !address) return null

  return (
    <>
      <Sep />
      <div className="flex flex-col items-center" style={{ padding: '4px 0 2px' }}>
        {/* Green connected dot */}
        <div
          style={{
            width: 6, height: 6, borderRadius: '50%',
            background: '#0A7068',
            boxShadow: '0 0 0 2px #EAF5F3',
            marginBottom: 4,
          }}
        />
        {/* Truncated address */}
        <button
          onClick={() => disconnect()}
          title={`Connected: ${address}\nClick to disconnect`}
          className="font-mono text-text-3 hover:text-text-1 cursor-pointer transition-colors text-center leading-tight"
          style={{ fontSize: 7, background: 'none', border: 'none', padding: 0, wordBreak: 'break-all', maxWidth: 44 }}
        >
          {fmtAddr(address)}
        </button>
      </div>
    </>
  )
}

export function Sidebar() {
  return (
    <div
      className="bg-white border-r border-border flex flex-col items-center flex-shrink-0 overflow-y-auto"
      style={{ width: 56, padding: '8px 0' }}
    >
      {/* Buy */}
      <NavIcon to="/" tip="Buy WATT">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path d="M8 2v12M4 6l4-4 4 4M3 12h10" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
      </NavIcon>

      {/* Stake */}
      <NavIcon to="/stake" tip="Stake sWATT">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path d="M8 14V2M5 5l3-3 3 3M3 10l2 2 2-2M9 10l2 2 2-2" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
      </NavIcon>

      {/* Borrow */}
      <NavIcon to="/borrow" tip="Borrow">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect x="2" y="10" width="12" height="4" rx="1" stroke="currentColor" strokeWidth="1.5"/>
          <path d="M5 10V6a3 3 0 016 0v4" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
          <path d="M8 2v1M11.5 3.5l-.7.7M4.5 3.5l.7.7" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
        </svg>
      </NavIcon>

      {/* Bridge */}
      <NavIcon to="/bridge" tip="Bridge (soon)">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path d="M2 10c0-2.2 2.7-4 6-4s6 1.8 6 4M4 10V7M12 10V7M2 10h12" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
        </svg>
      </NavIcon>

      <Sep />

      {/* Dashboard */}
      <NavIcon to="/dashboard" tip="Dashboard">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect x="2" y="2" width="5" height="5" rx="1" stroke="currentColor" strokeWidth="1.5"/>
          <rect x="9" y="2" width="5" height="5" rx="1" stroke="currentColor" strokeWidth="1.5"/>
          <rect x="2" y="9" width="5" height="5" rx="1" stroke="currentColor" strokeWidth="1.5"/>
          <rect x="9" y="9" width="5" height="5" rx="1" stroke="currentColor" strokeWidth="1.5"/>
        </svg>
      </NavIcon>

      {/* Veriflow */}
      <NavIcon to="/veriflow" tip="Veriflow">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path d="M8 2L14 5v4c0 3-2.5 5-6 6-3.5-1-6-3-6-6V5L8 2z" stroke="currentColor" strokeWidth="1.5" strokeLinejoin="round"/>
          <path d="M5.5 8l2 2 3-3" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
      </NavIcon>

      {/* Portfolio */}
      <NavIcon to="/portfolio" tip="Portfolio">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect x="2" y="4" width="12" height="10" rx="1.5" stroke="currentColor" strokeWidth="1.5"/>
          <path d="M5 4V3a1 1 0 011-1h4a1 1 0 011 1v1M2 8h12" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
        </svg>
      </NavIcon>

      <Sep />

      {/* Governance */}
      <NavIcon to="/governance" tip="Governance">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <circle cx="8" cy="8" r="6" stroke="currentColor" strokeWidth="1.5"/>
          <path d="M8 5v3l2 2" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
      </NavIcon>

      {/* Docs */}
      <NavIcon to="/docs" tip="Documentation">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <rect x="3" y="2" width="10" height="12" rx="1.5" stroke="currentColor" strokeWidth="1.5"/>
          <path d="M6 6h4M6 9h4M6 12h2" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
        </svg>
      </NavIcon>

      {/* Faucet — testnet only */}
      <NavIcon to="/faucet" tip="Testnet Faucet">
        <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
          <path d="M6 2h4M7 2v3L5 8c-.6.8-1 1.8-1 3a4 4 0 008 0c0-1.2-.4-2.2-1-3L9 5V2" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
          <circle cx="8" cy="11" r="1.5" fill="currentColor" opacity=".4"/>
        </svg>
      </NavIcon>

      {/* Wallet state — only shown when connected */}
      <div className="flex-1" />
      <WalletFooter />
    </div>
  )
}
