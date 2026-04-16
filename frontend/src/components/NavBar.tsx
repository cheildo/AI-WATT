import { Link } from 'react-router-dom'
import { WalletButton } from './WalletButton'

export function NavBar() {
  return (
    <header className="flex h-14 shrink-0 items-center justify-between border-b border-surface-border bg-surface px-6">
      <Link to="/" className="flex items-center gap-2">
        <span className="text-lg font-bold text-text-primary">
          AI <span className="text-brand">WATT</span>
        </span>
        <span className="rounded bg-brand-muted px-1.5 py-0.5 text-[10px] font-medium text-brand">BETA</span>
      </Link>
      <WalletButton />
    </header>
  )
}
