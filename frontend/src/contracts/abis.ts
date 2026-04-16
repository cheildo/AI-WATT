// Minimal ABI slices — only the functions the frontend calls.
// Re-run abigen / copy from backend/internal/blockchain/abis/ if contracts change.

export const erc20Abi = [
  { name: 'balanceOf',  type: 'function', stateMutability: 'view',       inputs: [{ name: 'account', type: 'address' }],                                            outputs: [{ type: 'uint256' }] },
  { name: 'allowance',  type: 'function', stateMutability: 'view',       inputs: [{ name: 'owner', type: 'address' }, { name: 'spender', type: 'address' }],        outputs: [{ type: 'uint256' }] },
  { name: 'approve',    type: 'function', stateMutability: 'nonpayable', inputs: [{ name: 'spender', type: 'address' }, { name: 'amount', type: 'uint256' }],        outputs: [{ type: 'bool' }] },
  { name: 'decimals',   type: 'function', stateMutability: 'view',       inputs: [],                                                                                 outputs: [{ type: 'uint8' }] },
  { name: 'totalSupply',type: 'function', stateMutability: 'view',       inputs: [],                                                                                 outputs: [{ type: 'uint256' }] },
] as const

export const mintEngineAbi = [
  { name: 'deposit', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'stablecoin', type: 'address' }, { name: 'amount', type: 'uint256' }],
    outputs: [{ name: 'wattMinted', type: 'uint256' }] },
  { name: 'redeem', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'wattAmount', type: 'uint256' }],
    outputs: [{ name: 'stableReturned', type: 'uint256' }] },
  { name: 'mintFee', type: 'function', stateMutability: 'view', inputs: [], outputs: [{ type: 'uint256' }] },
] as const

export const sWattUSDabi = [
  { name: 'balanceOf',       type: 'function', stateMutability: 'view',       inputs: [{ name: 'account', type: 'address' }],                                                                       outputs: [{ type: 'uint256' }] },
  { name: 'deposit',         type: 'function', stateMutability: 'nonpayable', inputs: [{ name: 'assets', type: 'uint256' }, { name: 'receiver', type: 'address' }],                                 outputs: [{ name: 'shares', type: 'uint256' }] },
  { name: 'redeem',          type: 'function', stateMutability: 'nonpayable', inputs: [{ name: 'shares', type: 'uint256' }, { name: 'receiver', type: 'address' }, { name: 'owner', type: 'address' }], outputs: [{ name: 'assets', type: 'uint256' }] },
  { name: 'convertToAssets', type: 'function', stateMutability: 'view',       inputs: [{ name: 'shares', type: 'uint256' }],                                                                        outputs: [{ type: 'uint256' }] },
  { name: 'convertToShares', type: 'function', stateMutability: 'view',       inputs: [{ name: 'assets', type: 'uint256' }],                                                                        outputs: [{ type: 'uint256' }] },
  { name: 'totalAssets',     type: 'function', stateMutability: 'view',       inputs: [],                                                                                                            outputs: [{ type: 'uint256' }] },
  { name: 'totalSupply',     type: 'function', stateMutability: 'view',       inputs: [],                                                                                                            outputs: [{ type: 'uint256' }] },
  { name: 'maxWithdraw',     type: 'function', stateMutability: 'view',       inputs: [{ name: 'owner', type: 'address' }],                                                                         outputs: [{ type: 'uint256' }] },
] as const

export const lendingPoolAbi = [
  { name: 'originateLoan', type: 'function', stateMutability: 'nonpayable',
    inputs: [
      { name: 'assetId',      type: 'bytes32' },
      { name: 'borrower',     type: 'address' },
      { name: 'principal',    type: 'uint256' },
      { name: 'interestRate', type: 'uint256' },
      { name: 'termDays',     type: 'uint256' },
      { name: 'engineType',   type: 'uint8'   },
    ],
    outputs: [{ name: 'loanId', type: 'bytes32' }] },
  { name: 'repay', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'loanId', type: 'bytes32' }, { name: 'amount', type: 'uint256' }],
    outputs: [] },
  { name: 'fullRepay', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'loanId', type: 'bytes32' }],
    outputs: [] },
  { name: 'getLoan', type: 'function', stateMutability: 'view',
    inputs: [{ name: 'loanId', type: 'bytes32' }],
    outputs: [{
      type: 'tuple', name: 'loan',
      components: [
        { name: 'loanId',       type: 'bytes32' },
        { name: 'assetId',      type: 'bytes32' },
        { name: 'borrower',     type: 'address' },
        { name: 'curator',      type: 'address' },
        { name: 'principal',    type: 'uint256' },
        { name: 'outstanding',  type: 'uint256' },
        { name: 'interestRate', type: 'uint256' },
        { name: 'status',       type: 'uint8'   },
        { name: 'engineType',   type: 'uint8'   },
        { name: 'originatedAt', type: 'uint256' },
        { name: 'maturityAt',   type: 'uint256' },
      ],
    }] },
  { name: 'getBorrowerLoans', type: 'function', stateMutability: 'view',
    inputs: [{ name: 'borrower', type: 'address' }],
    outputs: [{ type: 'bytes32[]' }] },
] as const

export const wevQueueAbi = [
  { name: 'requestRedeem', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'sWattAmount', type: 'uint256' }], outputs: [{ name: 'requestId', type: 'bytes32' }] },
  { name: 'requestPriorityRedeem', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'sWattAmount', type: 'uint256' }, { name: 'priorityFee', type: 'uint256' }],
    outputs: [{ name: 'requestId', type: 'bytes32' }] },
  { name: 'cancelRequest', type: 'function', stateMutability: 'nonpayable',
    inputs: [{ name: 'requestId', type: 'bytes32' }], outputs: [] },
  { name: 'getQueueDepth', type: 'function', stateMutability: 'view', inputs: [], outputs: [{ type: 'uint256' }] },
  { name: 'getUserRequests', type: 'function', stateMutability: 'view',
    inputs: [{ name: 'user', type: 'address' }], outputs: [{ type: 'bytes32[]' }] },
  { name: 'getRequest', type: 'function', stateMutability: 'view',
    inputs: [{ name: 'requestId', type: 'bytes32' }],
    outputs: [{
      type: 'tuple', name: 'req',
      components: [
        { name: 'requestId',   type: 'bytes32' },
        { name: 'user',        type: 'address' },
        { name: 'sWattAmount', type: 'uint256' },
        { name: 'priorityFee', type: 'uint256' },
        { name: 'requestedAt', type: 'uint256' },
        { name: 'status',      type: 'uint8'   },
      ],
    }] },
] as const
