import { cn } from '@/lib/formatters'

type TokenType = 'USDC' | 'USDT' | 'WATT' | 'sWATT'

const TOKEN_STYLES: Record<TokenType, string> = {
  USDC:  'bg-[#EEF4FF] text-[#2563EB] border-[#C7D7F8]',
  USDT:  'bg-[#F0FDF4] text-[#16A34A] border-[#BBF7D0]',
  WATT:  'bg-gold-bg text-gold border-gold-border',
  sWATT: 'bg-teal-bg text-teal border-teal-border',
}

const TOKEN_LABELS: Record<TokenType, string> = {
  USDC: 'USDC', USDT: 'USDT', WATT: 'WATT', sWATT: 'sWATT',
}

interface TokenRowProps {
  token: TokenType
  /** Override displayed name (e.g. "Loan Amount") */
  nameOverride?: string
  /** Override chain label */
  chainOverride?: string
  chain?: string
  balance: string
  amount: string
  onAmountChange?: (v: string) => void
  readOnly?: boolean
  showMax?: boolean
  onMax?: () => void
  /** USD or descriptive value shown bottom-right */
  subValue?: string
}

export function TokenRow({
  token,
  nameOverride,
  chainOverride,
  chain = 'XDC Network',
  balance,
  amount,
  onAmountChange,
  readOnly,
  showMax,
  onMax,
  subValue,
}: TokenRowProps) {
  return (
    <div
      className="border-b border-border"
      style={{ padding: '16px 18px' }}
    >
      {/* top row */}
      <div className="flex items-center justify-between mb-[10px]">
        <div className="flex items-center gap-[10px]">
          {/* token icon */}
          <div
            className={cn(
              'rounded-full flex items-center justify-center font-mono font-bold flex-shrink-0 border-[1.5px]',
              TOKEN_STYLES[token]
            )}
            style={{ width: 36, height: 36, fontSize: 9, letterSpacing: '.01em' }}
          >
            {token === 'sWATT' ? 'sW' : token.slice(0, 4)}
          </div>
          {/* token name */}
          <div>
            <div className="font-serif font-semibold" style={{ fontSize: 15 }}>
              {nameOverride ?? TOKEN_LABELS[token]}
            </div>
            <div className="flex items-center gap-1 text-text-3 mt-px" style={{ fontSize: 10 }}>
              {!chainOverride && (
                <span
                  className="rounded-sm bg-bg-3 text-text-2 flex items-center justify-center font-mono font-bold"
                  style={{ width: 8, height: 8, fontSize: 7 }}
                >X</span>
              )}
              {chainOverride ?? chain}
            </div>
          </div>
        </div>
        {/* amount input */}
        <input
          type="number"
          placeholder="0"
          value={amount}
          readOnly={readOnly}
          onChange={(e) => onAmountChange?.(e.target.value)}
          className={cn(
            'font-mono text-right bg-transparent border-none outline-none',
            readOnly ? 'text-text-3' : 'text-text-3 focus:text-text-1'
          )}
          style={{ fontSize: 26, width: 160 }}
        />
      </div>
      {/* bottom row */}
      <div className="flex items-center justify-between">
        <div className="text-text-3" style={{ fontSize: 11 }}>
          Balance: {balance}
          {showMax && (
            <button
              onClick={onMax}
              className="ml-2 bg-gold-bg border border-gold-border rounded text-gold font-mono cursor-pointer transition-colors hover:bg-[#f7ebc0]"
              style={{ padding: '2px 7px', fontSize: 10 }}
            >
              MAX
            </button>
          )}
        </div>
        {subValue && (
          <div className="font-mono text-text-3" style={{ fontSize: 11 }}>
            {subValue}
          </div>
        )}
      </div>
    </div>
  )
}
