import { useEffect, useRef, useState } from 'react'
import { cn } from '@/lib/formatters'

type TokenType = 'USDC' | 'USDT' | 'WATT' | 'sWATT'

const TOKEN_STYLES: Record<TokenType, string> = {
  USDC:  'bg-[#EEF4FF] text-[#2563EB] border-[#C7D7F8]',
  USDT:  'bg-[#F0FDF4] text-[#16A34A] border-[#BBF7D0]',
  WATT:  'bg-gold-bg text-gold border-gold-border',
  sWATT: 'bg-teal-bg text-teal border-teal-border',
}

const TOKEN_ABBR: Record<TokenType, string> = {
  USDC: 'USDC', USDT: 'USDT', WATT: 'WATT', sWATT: 'sW',
}

interface TokenRowProps {
  token: TokenType
  nameOverride?: string
  chainOverride?: string
  chain?: string
  balance: string
  /** Raw numeric balance for percentage buttons */
  balanceNum?: number
  amount: string
  onAmountChange?: (v: string) => void
  readOnly?: boolean
  showMax?: boolean
  onMax?: () => void
  subValue?: string
  /** Validation error message — shows red border + message */
  error?: string
}

export function TokenRow({
  token,
  nameOverride,
  chainOverride,
  chain = 'XDC Network',
  balance,
  balanceNum,
  amount,
  onAmountChange,
  readOnly,
  showMax,
  onMax,
  subValue,
  error,
}: TokenRowProps) {
  const handlePct = (pct: number) => {
    if (balanceNum === undefined || !onAmountChange) return
    onAmountChange((balanceNum * pct).toFixed(2))
  }

  // Fade-in animation when readOnly amount changes
  const [faded, setFaded] = useState(false)
  const prevAmount = useRef(amount)
  useEffect(() => {
    if (!readOnly) return
    if (amount !== prevAmount.current) {
      setFaded(true)
      const t = setTimeout(() => setFaded(false), 80)
      prevAmount.current = amount
      return () => clearTimeout(t)
    }
  }, [amount, readOnly])

  const hasError = Boolean(error)

  return (
    <div
      className={cn('border-b border-border last:border-b-0 transition-colors')}
      style={{ padding: '14px 18px' }}
    >
      {/* Top row: token + amount */}
      <div className="flex items-center justify-between mb-2">
        <div className="flex items-center gap-[10px]">
          <div
            className={cn(
              'rounded-full flex items-center justify-center font-mono font-bold flex-shrink-0 border-[1.5px]',
              TOKEN_STYLES[token]
            )}
            style={{ width: 36, height: 36, fontSize: 9, letterSpacing: '.01em' }}
          >
            {TOKEN_ABBR[token]}
          </div>
          <div>
            <div className="font-serif font-semibold leading-none" style={{ fontSize: 15 }}>
              {nameOverride ?? token}
            </div>
            <div className="flex items-center gap-1 text-text-3 mt-1" style={{ fontSize: 10 }}>
              {!chainOverride && (
                <span
                  className="rounded-sm bg-bg-3 text-text-2 inline-flex items-center justify-center font-mono font-bold"
                  style={{ width: 9, height: 9, fontSize: 7 }}
                >X</span>
              )}
              <span>{chainOverride ?? chain}</span>
            </div>
          </div>
        </div>

        {/* Amount input */}
        <div className="flex flex-col items-end gap-1 min-w-0">
          <input
            type="number"
            inputMode="decimal"
            placeholder="0"
            value={amount}
            readOnly={readOnly}
            onChange={(e) => onAmountChange?.(e.target.value)}
            className={cn(
              'font-mono text-right bg-transparent border-none outline-none w-full transition-opacity duration-75',
              readOnly   ? 'text-text-3 cursor-default' : 'text-text-1',
              !amount && !readOnly && 'text-text-3',
              hasError   && !readOnly && 'text-red',
            )}
            style={{
              fontSize: 26,
              maxWidth: 180,
              opacity: faded ? 0 : 1,
              // Remove browser number spinners
              MozAppearance: 'textfield',
            } as React.CSSProperties}
          />
          {subValue && (
            <span
              className={cn('font-mono', hasError ? 'text-red' : 'text-text-3')}
              style={{ fontSize: 11 }}
            >{subValue}</span>
          )}
        </div>
      </div>

      {/* Bottom row: balance + percentage buttons + error */}
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2 flex-wrap">
          <span
            className={cn('transition-colors', hasError ? 'text-red' : 'text-text-3')}
            style={{ fontSize: 11 }}
          >
            Balance: {balance}
          </span>
          {showMax && !readOnly && (
            <div className="flex items-center gap-1">
              {balanceNum !== undefined && [0.25, 0.5, 0.75].map((pct) => (
                <button
                  key={pct}
                  onClick={() => handlePct(pct)}
                  className="border border-border rounded text-text-3 font-mono cursor-pointer transition-colors hover:border-border-strong hover:text-text-2"
                  style={{ padding: '1px 5px', fontSize: 9 }}
                >
                  {pct * 100}%
                </button>
              ))}
              <button
                onClick={onMax}
                className="bg-gold-bg border border-gold-border rounded text-gold font-mono cursor-pointer transition-colors hover:bg-[#f7ebc0]"
                style={{ padding: '1px 6px', fontSize: 9 }}
              >
                MAX
              </button>
            </div>
          )}
        </div>
      </div>

      {/* Validation error */}
      {error && (
        <div className="text-red mt-1 flex items-center gap-1" style={{ fontSize: 11 }}>
          <span style={{ fontSize: 12 }}>⚠</span> {error}
        </div>
      )}
    </div>
  )
}
