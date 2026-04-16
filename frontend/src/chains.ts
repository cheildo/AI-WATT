import { defineChain } from 'viem'

export const xdcMainnet = defineChain({
  id: 50,
  name: 'XDC Network',
  nativeCurrency: { name: 'XDC', symbol: 'XDC', decimals: 18 },
  rpcUrls: {
    default: { http: ['https://erpc.xinfin.network'] },
  },
  blockExplorers: {
    default: { name: 'XDCScan', url: 'https://xdcscan.com' },
  },
})

export const xdcApothem = defineChain({
  id: 51,
  name: 'XDC Apothem Testnet',
  nativeCurrency: { name: 'TXDC', symbol: 'TXDC', decimals: 18 },
  rpcUrls: {
    default: { http: ['https://erpc.apothem.network'] },
  },
  blockExplorers: {
    default: { name: 'XDCScan Testnet', url: 'https://testnet.xdcscan.com' },
  },
  testnet: true,
})
