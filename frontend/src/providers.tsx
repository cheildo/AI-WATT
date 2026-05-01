import { WagmiProvider, createConfig, http } from 'wagmi'
import { injected, walletConnect } from 'wagmi/connectors'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { xdcMainnet, xdcApothem } from '@/chains'

const chainId  = Number(import.meta.env.VITE_CHAIN_ID ?? 51)
const rpcUrl   = import.meta.env.VITE_RPC_URL as string | undefined

// Build env-specific config — always include both chains so the transport
// type satisfies wagmi's Record<50 | 51, Transport>, but only connect to
// the active chain at runtime.
const config = createConfig({
  chains: [xdcMainnet, xdcApothem],
  connectors: [
    injected(),
    walletConnect({
      projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID ?? 'demo',
    }),
  ],
  transports: {
    // Use VITE_RPC_URL for the active chain; fall back to chain defaults for the other
    [xdcMainnet.id]: chainId === 50 ? http(rpcUrl) : http(),
    [xdcApothem.id]: chainId === 51 ? http(rpcUrl) : http(),
  },
})

const queryClient = new QueryClient({
  defaultOptions: { queries: { staleTime: 30_000 } },
})

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        {children}
      </QueryClientProvider>
    </WagmiProvider>
  )
}
