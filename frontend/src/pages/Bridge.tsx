import { TabBar }             from '@/components/swap/TabBar'
import { SwapWidget }         from '@/components/swap/SwapWidget'
import { TransactionDetails } from '@/components/swap/TransactionDetails'

export function Bridge() {
  return (
    <div className="animate-fadeup">
      <TabBar />
      <SwapWidget
        notice="Cross-chain bridging via LayerZero OFT is planned for Phase 3. WATT and sWATT will bridge to Ethereum, Arbitrum, and Base."
        left={
          <div className="text-center text-text-3" style={{ padding: '40px 24px' }}>
            <div style={{ fontSize: 32, marginBottom: 12 }}>🔗</div>
            <div className="font-serif text-text-1 mb-2" style={{ fontSize: 18 }}>Coming in Phase 3</div>
            <div className="leading-relaxed mx-auto" style={{ fontSize: 12, maxWidth: 260 }}>
              Bridge WATT and sWATT across Ethereum, Arbitrum, and Base using the LayerZero OFT framework.
            </div>
            <div
              className="border-t border-border mx-auto mt-5"
              style={{ display: 'grid', gridTemplateColumns: 'repeat(2,1fr)', maxWidth: 280 }}
            >
              {[
                { k: 'Target chains', v: 'ETH · ARB · Base' },
                { k: 'Protocol',      v: 'LayerZero OFT'     },
              ].map((item, i) => (
                <div
                  key={i}
                  className="flex justify-between py-2"
                  style={{
                    fontSize: 11,
                    borderRight: i % 2 === 0 ? '1px solid #D8D3C6' : 'none',
                    paddingRight: i % 2 === 0 ? 12 : 0,
                    paddingLeft:  i % 2 === 1 ? 12 : 0,
                  }}
                >
                  <span className="text-text-3">{item.k}</span>
                  <span className="font-mono text-text-1">{item.v}</span>
                </div>
              ))}
            </div>
          </div>
        }
        right={
          <TransactionDetails
            steps={[{ label: 'Bridge WATT', status: 'pending' }]}
            info={[]}
            actionLabel="Bridge (Coming Soon)"
            disabled
          />
        }
      />
    </div>
  )
}
