import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { parseUnits } from 'viem'
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
  return useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'convertToAssets',
    args: [parseUnits('1', 18)],
  })
}

export function useVaultStats() {
  const totalAssets = useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'totalAssets',
  })
  const totalSupply = useReadContract({
    address: CONTRACT_ADDRESSES.sWattUSD,
    abi: sWattUSDabi,
    functionName: 'totalSupply',
  })
  const nav = useNAVPerShare()
  return { totalAssets, totalSupply, nav }
}

export function useStakeWatt() {
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

export function useRequestUnstake() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const requestUnstake = (sWattAmount: bigint) =>
    writeContract({
      address: CONTRACT_ADDRESSES.wevQueue,
      abi: [
        { name: 'requestRedeem', type: 'function', stateMutability: 'nonpayable',
          inputs: [{ name: 'sWattAmount', type: 'uint256' }], outputs: [{ name: 'requestId', type: 'bytes32' }] },
      ] as const,
      functionName: 'requestRedeem',
      args: [sWattAmount],
    })

  return { requestUnstake, hash, isPending: isPending || isConfirming, isSuccess, error }
}
