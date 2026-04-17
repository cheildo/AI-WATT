import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { wevQueueAbi } from '@/contracts/abis'

export function useQueueDepth() {
  return useReadContract({
    address: CONTRACT_ADDRESSES.wevQueue,
    abi: wevQueueAbi,
    functionName: 'getQueueDepth',
  })
}

export function useUserQueue(user?: Address) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.wevQueue,
    abi: wevQueueAbi,
    functionName: 'getUserRequests',
    args: user ? [user] : undefined,
    query: { enabled: !!user },
  })
}

export function useRequestRedeem() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const requestRedeem = (sWattAmount: bigint) =>
    writeContract({
      address: CONTRACT_ADDRESSES.wevQueue,
      abi: wevQueueAbi,
      functionName: 'requestRedeem',
      args: [sWattAmount],
    })

  return { requestRedeem, hash, isPending: isPending || isConfirming, isSuccess, error }
}

export function useCancelRequest() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const cancel = (requestId: `0x${string}`) =>
    writeContract({
      address: CONTRACT_ADDRESSES.wevQueue,
      abi: wevQueueAbi,
      functionName: 'cancelRequest',
      args: [requestId],
    })

  return { cancel, hash, isPending: isPending || isConfirming, isSuccess, error }
}
