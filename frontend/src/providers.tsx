import { WagmiProvider, createConfig, http } from 'wagmi'
import { injected, walletConnect } from 'wagmi/connectors'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { xdcMainnet, xdcApothem } from '@/chains'

const config = createConfig({
  chains: [xdcMainnet, xdcApothem],
  connectors: [
    injected(),
    walletConnect({
      projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID ?? 'demo',
    }),
  ],
  transports: {
    [xdcMainnet.id]: http(),
    [xdcApothem.id]: http(),
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
