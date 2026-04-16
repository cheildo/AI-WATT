import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { WagmiProvider, createConfig, http } from 'wagmi'
import { injected, walletConnect } from 'wagmi/connectors'
import { xdcMainnet, xdcApothem } from './chains'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: { staleTime: 30_000, retry: 2 },
  },
})

const wagmiConfig = createConfig({
  chains: [xdcApothem, xdcMainnet],
  connectors: [
    injected(),
    walletConnect({ projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID ?? 'aiwatt' }),
  ],
  transports: {
    [xdcApothem.id]: http(import.meta.env.VITE_XDC_RPC_URL ?? 'https://erpc.apothem.network'),
    [xdcMainnet.id]: http('https://erpc.xinfin.network'),
  },
})

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        {children}
      </QueryClientProvider>
    </WagmiProvider>
  )
}
