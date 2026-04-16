import type { ButtonHTMLAttributes } from 'react'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  isLoading?: boolean
  variant?: 'primary' | 'secondary' | 'danger'
}

const variants = {
  primary:   'bg-brand hover:bg-brand-dim text-white',
  secondary: 'border border-surface-border bg-surface-card hover:bg-surface-hover text-text-primary',
  danger:    'bg-danger hover:opacity-90 text-white',
}

export function ActionButton({ isLoading, variant = 'primary', className = '', children, disabled, ...props }: Props) {
  return (
    <button
      disabled={disabled || isLoading}
      className={`relative flex items-center justify-center rounded-lg px-4 py-2.5 text-sm font-medium transition-all disabled:cursor-not-allowed disabled:opacity-50 ${variants[variant]} ${className}`}
      {...props}
    >
      {isLoading && (
        <span className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
          <span className="h-4 w-4 animate-spin rounded-full border-2 border-current border-t-transparent" />
        </span>
      )}
      <span className={isLoading ? 'opacity-0' : ''}>{children}</span>
    </button>
  )
}
