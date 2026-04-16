import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { lendingPoolAbi } from '@/contracts/abis'

export function useLoan(loanId?: `0x${string}`) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.lendingPool,
    abi: lendingPoolAbi,
    functionName: 'getLoan',
    args: loanId ? [loanId] : undefined,
    query: { enabled: !!loanId },
  })
}

export function useBorrowerLoans(borrower?: Address) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.lendingPool,
    abi: lendingPoolAbi,
    functionName: 'getBorrowerLoans',
    args: borrower ? [borrower] : undefined,
    query: { enabled: !!borrower },
  })
}

export function useOriginateLoan() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const originate = (params: {
    assetId: `0x${string}`
    borrower: Address
    principal: bigint
    interestRate: bigint
    termDays: bigint
    engineType: number
  }) =>
    writeContract({
      address: CONTRACT_ADDRESSES.lendingPool,
      abi: lendingPoolAbi,
      functionName: 'originateLoan',
      args: [params.assetId, params.borrower, params.principal, params.interestRate, params.termDays, params.engineType],
    })

  return { originate, hash, isPending: isPending || isConfirming, isSuccess, error }
}

export function useRepayLoan() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const repay = (loanId: `0x${string}`, amount: bigint) =>
    writeContract({
      address: CONTRACT_ADDRESSES.lendingPool,
      abi: lendingPoolAbi,
      functionName: 'repay',
      args: [loanId, amount],
    })

  return { repay, hash, isPending: isPending || isConfirming, isSuccess, error }
}
