import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TransactionDetails } from '@/components/swap/TransactionDetails'

const ROADMAP = [
  { phase: 'Phase 3', when: 'Q3 2025', label: 'LayerZero OFT deployment on XDC', done: false },
  { phase: 'Phase 4', when: 'Q4 2025', label: 'Ethereum mainnet bridge',         done: false },
  { phase: 'Phase 5', when: '2026',    label: 'Arbitrum + Base support',          done: false },
]

const CHAINS = [
  { name: 'XDC Network',  status: 'Live',   color: '#0A7068', bg: '#EAF5F3', border: '#80C9C0' },
  { name: 'Ethereum',     status: 'Soon',   color: '#5A5646', bg: '#EAECF3', border: '#D8D3C6' },
  { name: 'Arbitrum',     status: 'Planned', color: '#5A5646', bg: '#EAECF3', border: '#D8D3C6' },
  { name: 'Base',         status: 'Planned', color: '#5A5646', bg: '#EAECF3', border: '#D8D3C6' },
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

export function Bridge() {
  return (
    <div className="animate-fadeup" style={{ maxWidth: 920, margin: '0 auto' }}>
    <div style={{ display: 'grid', gridTemplateColumns: '1fr 288px', gap: 16, alignItems: 'start' }}>
      {/* Left: tab bar + swap widget */}
      <div>
        <TabBar />
        <SwapWidget
          notice="Cross-chain bridging via LayerZero OFT is planned for Phase 3. WATT and sWATT will bridge to Ethereum, Arbitrum, and Base."
          left={
            <div className="text-center" style={{ padding: '48px 32px' }}>
              {/* Chain icons */}
              <div className="flex items-center justify-center gap-3 mb-6">
                {['XDC', 'ETH', 'ARB', 'BASE'].map((c, i) => (
                  <div key={c} className="flex items-center gap-3">
                    <div
                      className="flex items-center justify-center font-mono font-bold flex-shrink-0"
                      style={{
                        width: 40, height: 40, borderRadius: '50%', fontSize: 9, letterSpacing: '.01em',
                        background: i === 0 ? '#EAF5F3' : '#EAECF3',
                        color: i === 0 ? '#0A7068' : '#9A9484',
                        border: `1.5px solid ${i === 0 ? '#80C9C0' : '#D8D3C6'}`,
                        opacity: i === 0 ? 1 : 0.6,
                      }}
                    >{c}</div>
                    {i < 3 && (
                      <div className="text-text-3" style={{ fontSize: 14 }}>→</div>
                    )}
                  </div>
                ))}
              </div>

              <div className="font-serif text-text-1 mb-2" style={{ fontSize: 20 }}>Coming in Phase 3</div>
              <div className="text-text-3 leading-relaxed mx-auto" style={{ fontSize: 12, maxWidth: 300 }}>
                Bridge WATT and sWATT across Ethereum, Arbitrum, and Base using the
                LayerZero OFT framework. Currently live on XDC Network only.
              </div>

              <div className="border border-border bg-bg rounded mx-auto mt-6" style={{ maxWidth: 320, padding: '12px 16px', borderRadius: 4 }}>
                <div className="flex items-center justify-between text-text-3 mb-3" style={{ fontSize: 10 }}>
                  <span>TECHNOLOGY</span><span>STATUS</span>
                </div>
                {[
                  { k: 'Protocol',        v: 'LayerZero OFT v2' },
                  { k: 'Token standard',  v: 'OFT (Omnichain Fungible Token)' },
                  { k: 'Security',        v: 'DVN — Decentralised Verification' },
                ].map((row) => (
                  <div key={row.k} className="flex items-start justify-between py-2 border-t border-border" style={{ fontSize: 11 }}>
                    <span className="text-text-3">{row.k}</span>
                    <span className="font-mono text-text-1 text-right ml-4" style={{ maxWidth: 160 }}>{row.v}</span>
                  </div>
                ))}
              </div>
            </div>
          }
          right={
            <TransactionDetails
              steps={[{ label: 'Bridge WATT', status: 'pending' }]}
              info={[
                { key: 'Source',      value: 'XDC Network' },
                { key: 'Destination', value: 'Coming soon' },
              ]}
              actionLabel="Bridge (Coming Soon)"
              disabled
            />
          }
        />
      </div>

      {/* Right: info sidebar */}
      <div className="flex flex-col gap-3" style={{ position: 'sticky', top: 0 }}>
        <Panel title="Supported Networks">
          <div className="flex flex-col gap-2 pt-1">
            {CHAINS.map((c) => (
              <div
                key={c.name}
                className="flex items-center justify-between"
                style={{ padding: '7px 0', borderBottom: '1px solid #D8D3C6' }}
              >
                <span className="text-text-2" style={{ fontSize: 11 }}>{c.name}</span>
                <span
                  className="font-mono font-semibold"
                  style={{
                    fontSize: 10, padding: '2px 7px', borderRadius: 3,
                    background: c.bg, color: c.color, border: `1px solid ${c.border}`,
                  }}
                >{c.status}</span>
              </div>
            ))}
          </div>
        </Panel>

        <Panel title="Phase Roadmap">
          <div className="flex flex-col gap-0 pt-1">
            {ROADMAP.map((r, i) => (
              <div key={r.phase} className="flex items-start gap-3 py-3" style={{ borderBottom: i < ROADMAP.length - 1 ? '1px solid #D8D3C6' : 'none' }}>
                <div className="flex flex-col items-center flex-shrink-0" style={{ marginTop: 2 }}>
                  <div
                    className="font-mono font-bold"
                    style={{
                      width: 8, height: 8, borderRadius: '50%', marginBottom: 2,
                      background: r.done ? '#0A7068' : '#D8D3C6',
                    }}
                  />
                  {i < ROADMAP.length - 1 && (
                    <div style={{ width: 1, height: 20, background: '#D8D3C6' }} />
                  )}
                </div>
                <div>
                  <div className="flex items-center gap-2 mb-[2px]">
                    <span className="font-semibold text-text-2" style={{ fontSize: 11 }}>{r.phase}</span>
                    <span className="text-text-3 font-mono" style={{ fontSize: 10 }}>{r.when}</span>
                  </div>
                  <div className="text-text-3 leading-snug" style={{ fontSize: 11 }}>{r.label}</div>
                </div>
              </div>
            ))}
          </div>
        </Panel>

        <Panel title="LayerZero OFT">
          <p className="text-text-3 leading-relaxed" style={{ fontSize: 11, paddingTop: 6 }}>
            The Omnichain Fungible Token (OFT) standard allows WATT and sWATT to exist natively
            on multiple chains without wrapped tokens or liquidity fragmentation. Assets are
            locked on the source chain and minted on the destination via decentralised verification.
          </p>
        </Panel>
      </div>
    </div>
    </div>
  )
}
