import { useAccount, useConnect, useDisconnect } from 'wagmi'
import { injected } from 'wagmi/connectors'

function truncate(addr: string) {
  return `${addr.slice(0, 6)}…${addr.slice(-4)}`
}

export function WalletButton() {
  const { address, isConnected } = useAccount()
  const { connect } = useConnect()
  const { disconnect } = useDisconnect()

  if (isConnected && address) {
    return (
      <button
        onClick={() => disconnect()}
        className="flex items-center gap-2 rounded-lg border border-surface-border bg-surface-card px-3 py-1.5 text-sm text-text-primary hover:border-brand transition-colors"
      >
        <span className="h-2 w-2 rounded-full bg-yield" />
        {truncate(address)}
      </button>
    )
  }

  return (
    <button
      onClick={() => connect({ connector: injected() })}
      className="rounded-lg bg-brand px-4 py-1.5 text-sm font-medium text-white hover:bg-brand-dim transition-colors"
    >
      Connect Wallet
    </button>
  )
}
