import { ReactNode } from 'react'
import { cn } from '@/lib/formatters'

interface ActionButtonProps {
  children: ReactNode
  onClick?: () => void
  disabled?: boolean
  variant?: 'green' | 'gold' | 'outline'
  className?: string
}

export function ActionButton({
  children,
  onClick,
  disabled,
  variant = 'green',
  className,
}: ActionButtonProps) {
  return (
    <button
      onClick={onClick}
      disabled={disabled}
      className={cn(
        'w-full rounded font-sans font-semibold uppercase tracking-[.05em] cursor-pointer transition-all',
        'hover:-translate-y-px disabled:opacity-35 disabled:cursor-not-allowed disabled:transform-none',
        variant === 'green'   && 'bg-green text-white border-none',
        variant === 'gold'    && 'bg-gold text-white border-none',
        variant === 'outline' && 'bg-transparent text-text-2 border border-border-strong hover:bg-bg-2 hover:-translate-y-0',
        className
      )}
      style={{ padding: '12px', fontSize: 12 }}
    >
      {children}
    </button>
  )
}
