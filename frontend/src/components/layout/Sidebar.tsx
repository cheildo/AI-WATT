import { NavIcon } from './NavIcon'

function Sep() {
  return <div className="bg-border my-[6px]" style={{ width: 32, height: 1 }} />
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
    </div>
  )
}
