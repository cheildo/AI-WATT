import { useState } from 'react'
import { fmtAddr } from '@/lib/formatters'
import { useUIStore } from '@/stores/uiStore'

// ── Contract registry (Apothem testnet) ─────────────────────────────────────
const CONTRACTS = [
  { name: 'WattUSD',           address: '0x974B2bd650c88D290469a471DeB1Ee6aC55AD2d9', phase: 1, status: 'live' },
  { name: 'MintEngine',        address: '0x37a1d3b072af3055993215DB82e3832dA2d10AEd', phase: 1, status: 'live' },
  { name: 'sWattUSD',          address: '0x08c14aCc5547f03fC523453e005A243C50A7aa94', phase: 2, status: 'live' },
  { name: 'AssetRegistry',     address: '',                                            phase: 3, status: 'pending' },
  { name: 'OCNFT',             address: '',                                            phase: 3, status: 'pending' },
  { name: 'HealthAttestation', address: '',                                            phase: 3, status: 'pending' },
  { name: 'LendingPool',       address: '',                                            phase: 4, status: 'pending' },
  { name: 'WEVQueue',          address: '',                                            phase: 5, status: 'pending' },
]

const ENGINES = [
  {
    id: 'E1', title: 'Pre-Delivery',
    desc: 'Finances 30% of a purchase order deposit before hardware ships. Veriflow certifies the PO. Repaid from deployment proceeds over term.',
    fee: '2–3% origination', ltv: 'N/A (PO-backed)', term: '3–12 months',
  },
  {
    id: 'E2', title: 'Post-Delivery',
    desc: 'Productivity-backed loan originated once hardware is deployed and live. Veriflow monitors real-time telemetry. Health score must be ≥ 60 to originate.',
    fee: '1.5–2.5% AUM/yr', ltv: '70% of asset value', term: '12–36 months',
  },
  {
    id: 'E3', title: 'Capital Reactivation',
    desc: 'Idle escrow capital during the hardware delivery window is auto-deployed into US T-bills via TreasuryService. Generates base yield while capital waits.',
    fee: '~5% APY', ltv: 'N/A (T-bill backed)', term: 'Rolling',
  },
]

const TOKENS = [
  {
    symbol: 'WATT', full: 'WattUSD', color: '#9A6B0A', bg: '#FBF4E4', border: '#E2C46A',
    desc: 'USD-pegged synthetic dollar. Minted 1:1 from USDC/USDT via MintEngine. Does not accrue yield itself — stake into sWATT for yield exposure.',
    stats: [
      { k: 'Peg',      v: '$1.00 USD' },
      { k: 'Backing',  v: 'USDC/USDT + T-Bills' },
      { k: 'Decimals', v: '6' },
      { k: 'Chain',    v: 'XDC Network' },
    ],
  },
  {
    symbol: 'sWATT', full: 'sWattUSD', color: '#0A7068', bg: '#EAF5F3', border: '#80C9C0',
    desc: 'ERC-4626 yield vault. Deposit WATT, receive sWATT shares that appreciate in NAV as GPU loan interest and origination fees flow in via receiveYield().',
    stats: [
      { k: 'Current NAV',  v: '$1.0418 / sWATT' },
      { k: '7d APR',       v: '12.81%' },
      { k: 'Unstake',      v: 'WEV queue (~14–30d)' },
      { k: 'Priority exit',v: '0.5% fee, ~3 days' },
    ],
  },
  {
    symbol: '$WATT', full: 'WattToken (governance)', color: '#2563EB', bg: '#EEF4FF', border: '#C7D7F8',
    desc: 'Governance token. Holders vote on protocol parameters (mint fee, LTV limits, curator whitelisting) via Governor + 48h Timelock. Phase 6 — coming soon.',
    stats: [
      { k: 'Role',       v: 'On-chain governance' },
      { k: 'Quorum',     v: '4% of supply' },
      { k: 'Timelock',   v: '48h delay after vote' },
      { k: 'Status',     v: 'Phase 6 — in progress' },
    ],
  },
]

function SectionHeader({ title, sub }: { title: string; sub?: string }) {
  return (
    <div className="mb-4">
      <h2 className="font-serif font-semibold text-text-1" style={{ fontSize: 18 }}>{title}</h2>
      {sub && <p className="text-text-3 mt-1" style={{ fontSize: 12 }}>{sub}</p>}
    </div>
  )
}

function Card({ children, className }: { children: React.ReactNode; className?: string }) {
  return (
    <div className={`bg-white border border-border ${className ?? ''}`} style={{ borderRadius: 4 }}>
      {children}
    </div>
  )
}

function CopyButton({ text }: { text: string }) {
  const { showToast } = useUIStore()
  const [copied, setCopied] = useState(false)

  const copy = () => {
    navigator.clipboard.writeText(text).then(() => {
      setCopied(true)
      showToast('Address copied to clipboard', 'success')
      setTimeout(() => setCopied(false), 2000)
    })
  }

  return (
    <button
      onClick={copy}
      className="text-text-3 hover:text-text-1 cursor-pointer transition-colors flex-shrink-0"
      style={{ background: 'none', border: 'none', fontSize: 11, padding: '0 0 0 6px' }}
      title="Copy full address"
    >
      {copied ? '✓' : '⧉'}
    </button>
  )
}

export function Docs() {
  return (
    <div className="animate-fadeup" style={{ maxWidth: 900, margin: '0 auto' }}>

      {/* Header banner */}
      <div className="bg-green text-white mb-8" style={{ borderRadius: 4, padding: '24px 28px' }}>
        <div className="flex items-center justify-between">
          <div>
            <div className="font-serif" style={{ fontSize: 22 }}>AI WATT Protocol</div>
            <div className="mt-1 opacity-75" style={{ fontSize: 12 }}>
              Decentralised credit protocol for AI and automation assets on XDC Network
            </div>
          </div>
          <div className="flex flex-col items-end gap-2">
            <span className="font-mono bg-white text-green font-semibold" style={{ padding: '3px 10px', borderRadius: 3, fontSize: 11 }}>
              XDC Apothem · chainId 51
            </span>
            <span className="text-white opacity-60 font-mono" style={{ fontSize: 10 }}>
              Parent: Neurowatt Pte. Ltd., Singapore
            </span>
          </div>
        </div>
      </div>

      {/* Token System */}
      <div className="mb-8">
        <SectionHeader title="Token System" sub="Three tokens — two for protocol economics, one for governance" />
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: 12 }}>
          {TOKENS.map((t) => (
            <Card key={t.symbol}>
              <div className="border-b border-border" style={{ padding: '14px 16px' }}>
                <div className="flex items-center gap-2 mb-2">
                  <span
                    className="font-mono font-bold rounded"
                    style={{ padding: '2px 8px', fontSize: 11, background: t.bg, color: t.color, border: `1.5px solid ${t.border}` }}
                  >{t.symbol}</span>
                  <span className="text-text-3" style={{ fontSize: 11 }}>{t.full}</span>
                </div>
                <p className="text-text-2 leading-snug" style={{ fontSize: 11 }}>{t.desc}</p>
              </div>
              <div style={{ padding: '4px 16px 10px' }}>
                {t.stats.map((s) => (
                  <div key={s.k} className="flex items-center justify-between py-[7px] border-b border-border last:border-b-0">
                    <span className="text-text-3" style={{ fontSize: 10 }}>{s.k}</span>
                    <span className="font-mono font-medium" style={{ fontSize: 11, color: t.color }}>{s.v}</span>
                  </div>
                ))}
              </div>
            </Card>
          ))}
        </div>
      </div>

      {/* Credit Engines */}
      <div className="mb-8">
        <SectionHeader title="Credit Engines" sub="Three origination pathways — each matched to a hardware lifecycle stage" />
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: 12 }}>
          {ENGINES.map((e) => (
            <Card key={e.id}>
              <div style={{ padding: '14px 16px' }}>
                <div className="flex items-center gap-2 mb-3">
                  <span
                    className="font-mono font-bold text-white bg-green"
                    style={{ padding: '2px 7px', borderRadius: 3, fontSize: 10 }}
                  >{e.id}</span>
                  <span className="font-serif font-semibold text-text-1" style={{ fontSize: 14 }}>{e.title}</span>
                </div>
                <p className="text-text-2 leading-snug mb-3" style={{ fontSize: 11 }}>{e.desc}</p>
                <div className="border-t border-border" style={{ paddingTop: 10 }}>
                  {[
                    { k: 'Fee',  v: e.fee },
                    { k: 'LTV',  v: e.ltv },
                    { k: 'Term', v: e.term },
                  ].map((row) => (
                    <div key={row.k} className="flex items-center justify-between py-[5px] border-b border-border last:border-b-0">
                      <span className="text-text-3" style={{ fontSize: 10 }}>{row.k}</span>
                      <span className="font-mono text-gold font-medium" style={{ fontSize: 11 }}>{row.v}</span>
                    </div>
                  ))}
                </div>
              </div>
            </Card>
          ))}
        </div>
      </div>

      {/* Veriflow */}
      <div className="mb-8">
        <SectionHeader title="Veriflow — Asset Health Attestation" sub="Real-time hardware telemetry written on-chain via the Veriflow Agent" />
        <Card>
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', padding: '20px 24px', gap: 24 }}>
            <div>
              <div className="font-semibold text-text-1 mb-2" style={{ fontSize: 13 }}>How it works</div>
              {[
                { n: 1, t: 'Agent deployed on borrower hardware', d: 'Reads nvidia-smi + ipmitool every 5 minutes' },
                { n: 2, t: 'Signs telemetry with HMAC-SHA256',   d: 'Using a secret provisioned at loan origination' },
                { n: 3, t: 'POSTs to backend ingestion endpoint', d: 'Verified, scored, stored in MySQL telemetry table' },
                { n: 4, t: 'Score written on-chain daily',        d: 'HealthAttestation.sol · keccak256 hash + score 0–100' },
              ].map((step) => (
                <div key={step.n} className="flex items-start gap-3 mb-3">
                  <div
                    className="bg-green text-white font-mono font-bold flex-shrink-0 flex items-center justify-center"
                    style={{ width: 20, height: 20, borderRadius: '50%', fontSize: 9, marginTop: 1 }}
                  >{step.n}</div>
                  <div>
                    <div className="font-semibold text-text-1" style={{ fontSize: 12 }}>{step.t}</div>
                    <div className="text-text-3" style={{ fontSize: 11 }}>{step.d}</div>
                  </div>
                </div>
              ))}
            </div>
            <div>
              <div className="font-semibold text-text-1 mb-2" style={{ fontSize: 13 }}>Scoring Model</div>
              {[
                { metric: 'GPU Utilisation',   weight: '25 pts', threshold: '≥ 70% full, ≥ 40% partial' },
                { metric: 'Temperature',        weight: '25 pts', threshold: '≤ 75°C full, ≤ 85°C partial' },
                { metric: 'ECC Error Rate',     weight: '25 pts', threshold: '< 0.01% full, < 0.1% partial' },
                { metric: 'Uptime %',           weight: '25 pts', threshold: '≥ 99.5% full, ≥ 95% partial' },
              ].map((row) => (
                <div key={row.metric} className="py-[7px] border-b border-border last:border-b-0">
                  <div className="flex items-center justify-between mb-[2px]">
                    <span className="font-medium text-text-1" style={{ fontSize: 11 }}>{row.metric}</span>
                    <span className="font-mono text-teal font-semibold" style={{ fontSize: 11 }}>{row.weight}</span>
                  </div>
                  <div className="text-text-3" style={{ fontSize: 10 }}>{row.threshold}</div>
                </div>
              ))}
              <div className="mt-3 border border-border bg-bg rounded" style={{ padding: '8px 12px', borderRadius: 4 }}>
                <div className="text-text-3" style={{ fontSize: 10 }}>
                  Score &lt; 60 → alert sent · Score &lt; 40 → asset flagged on-chain
                </div>
                <div className="text-text-3 mt-1" style={{ fontSize: 10 }}>
                  Heartbeat &gt; 15 min → score = 0 (offline)
                </div>
              </div>
            </div>
          </div>
        </Card>
      </div>

      {/* Smart Contracts */}
      <div className="mb-8">
        <SectionHeader title="Smart Contracts" sub="All contracts are UUPS upgradeable proxies · XDC Apothem testnet (chainId 51)" />
        <Card>
          <div className="border-b border-border" style={{ padding: '12px 18px' }}>
            <div style={{ display: 'grid', gridTemplateColumns: '140px 80px 1fr 80px', gap: 12, fontSize: 10 }}
              className="text-text-3 uppercase tracking-[.08em] font-medium">
              <span>Contract</span><span>Phase</span><span>Address</span><span>Status</span>
            </div>
          </div>
          {CONTRACTS.map((c) => (
            <div
              key={c.name}
              className="flex items-center border-b border-border last:border-b-0"
              style={{ padding: '11px 18px', display: 'grid', gridTemplateColumns: '140px 80px 1fr 80px', gap: 12 }}
            >
              <span className="font-mono font-medium text-text-1" style={{ fontSize: 12 }}>{c.name}</span>
              <span className="text-text-3" style={{ fontSize: 11 }}>Phase {c.phase}</span>
              <span className="flex items-center gap-1">
                {c.address ? (
                  <>
                    <span className="font-mono text-text-2" style={{ fontSize: 11 }}>
                      <span className="hidden-full">{c.address}</span>
                      <a
                        href={`https://explorer.apothem.network/address/${c.address}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="font-mono text-text-2 hover:text-green transition-colors"
                        style={{ fontSize: 11 }}
                        title={c.address}
                      >
                        {fmtAddr(c.address)}
                      </a>
                    </span>
                    <CopyButton text={c.address} />
                  </>
                ) : (
                  <span className="text-text-3" style={{ fontSize: 11 }}>Not deployed</span>
                )}
              </span>
              <span>
                <span
                  className="font-mono font-semibold"
                  style={{
                    fontSize: 10, padding: '2px 7px', borderRadius: 3,
                    background: c.status === 'live' ? '#EAF5F3' : '#EAECF3',
                    color: c.status === 'live' ? '#0A7068' : '#9A9484',
                    border: `1px solid ${c.status === 'live' ? '#80C9C0' : '#D8D3C6'}`,
                  }}
                >
                  {c.status === 'live' ? 'Live' : 'Pending deploy'}
                </span>
              </span>
            </div>
          ))}
        </Card>
      </div>

      {/* Backend API */}
      <div className="mb-8">
        <SectionHeader title="Backend API" sub="Golang REST API · /api/v1/ · JWT-authenticated" />
        <Card>
          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', padding: '20px 24px', gap: 24 }}>
            <div>
              <div className="font-semibold text-text-1 mb-3" style={{ fontSize: 13 }}>Endpoints</div>
              {[
                { method: 'POST', path: '/auth/login',             note: 'EIP-191 wallet or email login → JWT' },
                { method: 'GET',  path: '/assets',                 note: 'List registered assets + health scores' },
                { method: 'POST', path: '/assets',                 note: 'Register new asset (REGISTRAR role)' },
                { method: 'GET',  path: '/assets/:id/health',      note: 'Latest Veriflow health score' },
                { method: 'POST', path: '/loans',                  note: 'Originate loan (CURATOR role)' },
                { method: 'POST', path: '/loans/:id/repay',        note: 'Submit repayment' },
                { method: 'POST', path: '/mint',                   note: 'Mint WATT from USDC/USDT' },
                { method: 'GET',  path: '/vault/stats',            note: 'NAV, APR, TVL, deployed %' },
                { method: 'POST', path: '/wev/queue',              note: 'Enter WEV redemption queue' },
                { method: 'POST', path: '/veriflow/telemetry',     note: 'Agent telemetry ingestion (HMAC)' },
              ].map((ep) => (
                <div key={ep.path} className="flex items-start gap-2 py-[6px] border-b border-border last:border-b-0">
                  <span
                    className="font-mono font-bold flex-shrink-0"
                    style={{
                      fontSize: 9, padding: '2px 5px', borderRadius: 2, width: 36, textAlign: 'center',
                      background: ep.method === 'GET' ? '#EEF4FF' : '#EAF5F3',
                      color: ep.method === 'GET' ? '#2563EB' : '#0A7068',
                    }}
                  >{ep.method}</span>
                  <div>
                    <div className="font-mono text-text-1" style={{ fontSize: 11 }}>/api/v1{ep.path}</div>
                    <div className="text-text-3" style={{ fontSize: 10 }}>{ep.note}</div>
                  </div>
                </div>
              ))}
            </div>
            <div>
              <div className="font-semibold text-text-1 mb-3" style={{ fontSize: 13 }}>Stack</div>
              {[
                { k: 'Language',   v: 'Go 1.22+' },
                { k: 'Framework',  v: 'Gin HTTP' },
                { k: 'ORM',        v: 'GORM + MySQL 8' },
                { k: 'Migrations', v: 'go-migrate' },
                { k: 'Auth',       v: 'JWT + EIP-191 sig verify' },
                { k: 'Cache',      v: 'Redis 7' },
                { k: 'Blockchain', v: 'go-ethereum + abigen' },
                { k: 'Logging',    v: 'zap structured logs' },
                { k: 'Docs',       v: 'Swagger / OpenAPI 3.0' },
              ].map((row) => (
                <div key={row.k} className="flex items-center justify-between py-[7px] border-b border-border last:border-b-0">
                  <span className="text-text-3" style={{ fontSize: 11 }}>{row.k}</span>
                  <span className="font-mono font-medium text-text-1" style={{ fontSize: 11 }}>{row.v}</span>
                </div>
              ))}
              <div className="mt-4 border border-border bg-bg rounded" style={{ padding: '10px 12px', borderRadius: 4 }}>
                <div className="font-semibold text-text-2 mb-1" style={{ fontSize: 11 }}>Response envelope</div>
                <pre className="font-mono text-text-3 leading-snug" style={{ fontSize: 10 }}>
{`{
  "success": true,
  "data": { ... },
  "error": null
}`}
                </pre>
              </div>
            </div>
          </div>
        </Card>
      </div>

      {/* Security */}
      <div className="mb-4">
        <SectionHeader title="Security Model" />
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: 12 }}>
          {[
            {
              title: 'Smart Contract',
              items: [
                'UUPS upgradeable proxies',
                'AccessControl roles (no raw onlyOwner)',
                '48h Timelock on all upgrades',
                'Governor vote required before Timelock',
                'Inline reentrancy guard on LendingPool + WEVQueue',
                'Custom errors (gas efficient)',
              ],
            },
            {
              title: 'Backend',
              items: [
                'JWT tokens, 1h expiry',
                'EIP-191 signature verification for wallet login',
                'HMAC-SHA256 per-asset secrets for Veriflow',
                'Redis rate limiter (60 req/60s per IP)',
                'CORS allowlist — no wildcard origins',
                'go-migrate only — no AutoMigrate in production',
              ],
            },
            {
              title: 'Veriflow Agent',
              items: [
                'Static binary (CGO_ENABLED=0)',
                'HMAC-SHA256 over raw JSON payload bytes',
                'Systemd hardening (NoNewPrivileges, PrivateTmp)',
                'Retry with exponential backoff on failure',
                '3x retries before marking offline',
                'No private keys on borrower hardware',
              ],
            },
          ].map((col) => (
            <Card key={col.title}>
              <div style={{ padding: '14px 16px' }}>
                <div className="font-semibold text-text-1 mb-3" style={{ fontSize: 12 }}>{col.title}</div>
                <ul className="flex flex-col gap-[6px]">
                  {col.items.map((item) => (
                    <li key={item} className="flex items-start gap-2">
                      <span className="text-teal flex-shrink-0" style={{ fontSize: 10, marginTop: 2 }}>✓</span>
                      <span className="text-text-2 leading-snug" style={{ fontSize: 11 }}>{item}</span>
                    </li>
                  ))}
                </ul>
              </div>
            </Card>
          ))}
        </div>
      </div>

    </div>
  )
}
