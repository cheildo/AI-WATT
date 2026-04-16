import { useReadContract, useWriteContract, useWaitForTransactionReceipt } from 'wagmi'
import type { Address } from 'viem'
import { CONTRACT_ADDRESSES } from '@/contracts/addresses'
import { erc20Abi, mintEngineAbi } from '@/contracts/abis'

export function useWattBalance(address?: Address) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.wattUSD,
    abi: erc20Abi,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
    query: { enabled: !!address },
  })
}

export function useWattAllowance(owner?: Address, spender?: Address) {
  return useReadContract({
    address: CONTRACT_ADDRESSES.wattUSD,
    abi: erc20Abi,
    functionName: 'allowance',
    args: owner && spender ? [owner, spender] : undefined,
    query: { enabled: !!owner && !!spender },
  })
}

export function useMintWatt() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const mint = (stablecoin: Address, amount: bigint) =>
    writeContract({
      address: CONTRACT_ADDRESSES.mintEngine,
      abi: mintEngineAbi,
      functionName: 'deposit',
      args: [stablecoin, amount],
    })

  return { mint, hash, isPending: isPending || isConfirming, isSuccess, error }
}

export function useRedeemWatt() {
  const { writeContract, data: hash, isPending, error } = useWriteContract()
  const { isLoading: isConfirming, isSuccess } = useWaitForTransactionReceipt({ hash })

  const redeem = (wattAmount: bigint) =>
    writeContract({
      address: CONTRACT_ADDRESSES.mintEngine,
      abi: mintEngineAbi,
      functionName: 'redeem',
      args: [wattAmount],
    })

  return { redeem, hash, isPending: isPending || isConfirming, isSuccess, error }
}
