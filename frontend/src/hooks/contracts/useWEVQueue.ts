import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { wevQueueAbi } from '@/contracts/abis'

export function useQueueStatus() {
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

  const requestRedeem = (sWattAmount: bigint, priority = false, priorityFee = 0n) => {
    if (priority) {
      writeContract({
        address: CONTRACT_ADDRESSES.wevQueue,
        abi: wevQueueAbi,
        functionName: 'requestPriorityRedeem',
        args: [sWattAmount, priorityFee],
      })
    } else {
      writeContract({
        address: CONTRACT_ADDRESSES.wevQueue,
        abi: wevQueueAbi,
        functionName: 'requestRedeem',
        args: [sWattAmount],
      })
    }
  }

  return { requestRedeem, hash, isPending: isPending || isConfirming, isSuccess, error }
}

export function useCancelRedeem() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const cancelRedeem = (requestId: `0x${string}`) =>
    writeContract({
      address: CONTRACT_ADDRESSES.wevQueue,
      abi: wevQueueAbi,
      functionName: 'cancelRequest',
      args: [requestId],
    })

  return { cancelRedeem, hash, isPending: isPending || isConfirming, isSuccess, error }
}
