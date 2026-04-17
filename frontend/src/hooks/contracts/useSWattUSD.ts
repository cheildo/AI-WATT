import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { sWattUSDabi } from '@/contracts/abis'

export function useSWattBalance(address?: Address) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: !!address },
  })
}

export function useNAVPerShare() {
  const assets = useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'totalAssets',
  })
  const supply = useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'totalSupply',
  })
  const nav =
    assets.data && supply.data && supply.data > 0n
      ? Number(assets.data) / Number(supply.data)
      : 1.0

  return { nav, isLoading: assets.isLoading || supply.isLoading }
}

export function useStakeSWatt() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const stake = (assets: bigint, receiver: Address) =>
    writeContract({
      address: CONTRACT_ADDRESSES.sWattUSD,
      abi: sWattUSDabi,
      functionName: 'deposit',
      args: [assets, receiver],
    })

  return { stake, hash, isPending: isPending || isConfirming, isSuccess, error }
}
