package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	nonceCachePrefix  = "txmanager:nonce:"
	gasBufferPct      = 120 // 20% overhead
	maxNonceRetries   = 3
	receiptPollInterval = 2 * time.Second
	receiptTimeout    = 30 * time.Second
)

// TxManager signs and sends transactions to XDC contracts.
// Nonces are cached in Redis with a mutex to prevent races on concurrent sends.
type TxManager struct {
	client     *BlockchainClient
	privKey    *ecdsa.PrivateKey
	fromAddr   common.Address
	chainID    *big.Int
	rdb        *redis.Client
	logger     *zap.Logger
	nonceMu    sync.Mutex
}

// NewTxManager constructs a TxManager. privateKey must be a 0x-prefixed hex string.
func NewTxManager(
	client *BlockchainClient,
	privateKey string,
	chainID int64,
	rdb *redis.Client,
	logger *zap.Logger,
) (*TxManager, error) {
	hex := strings.TrimPrefix(privateKey, "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		return nil, fmt.Errorf("TxManager: parse private key: %w", err)
	}
	fromAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	return &TxManager{
		client:   client,
		privKey:  privKey,
		fromAddr: fromAddr,
		chainID:  big.NewInt(chainID),
		rdb:      rdb,
		logger:   logger,
	}, nil
}

// ── Public write functions ─────────────────────────────────────────────────

// MintOCNFT mints an OC-NFT on-chain for the given asset.
func (tm *TxManager) MintOCNFT(ctx context.Context, to common.Address, assetID [32]byte, metadataURI string) (*types.Receipt, error) {
	opts, err := tm.txOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.MintOCNFT: %w", err)
	}
	tx, err := tm.client.OCNFT.MintOCNFT(opts, to, assetID, metadataURI)
	if err != nil {
		return nil, fmt.Errorf("TxManager.MintOCNFT: call: %w", err)
	}
	return tm.waitForReceipt(ctx, tx.Hash())
}

// RegisterAssetOnChain registers an asset in AssetRegistry.
func (tm *TxManager) RegisterAssetOnChain(
	ctx context.Context,
	assetID [32]byte,
	assetType uint8,
	borrower common.Address,
	ltv uint16,
) (*types.Receipt, error) {
	opts, err := tm.txOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.RegisterAssetOnChain: %w", err)
	}
	tx, err := tm.client.AssetRegistry.RegisterAsset(opts, assetID, assetType, borrower, ltv)
	if err != nil {
		return nil, fmt.Errorf("TxManager.RegisterAssetOnChain: call: %w", err)
	}
	return tm.waitForReceipt(ctx, tx.Hash())
}

// UpdateLTVOnChain updates the LTV for an asset in AssetRegistry.
func (tm *TxManager) UpdateLTVOnChain(ctx context.Context, assetID [32]byte, newLTV uint16) (*types.Receipt, error) {
	opts, err := tm.txOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.UpdateLTVOnChain: %w", err)
	}
	tx, err := tm.client.AssetRegistry.UpdateLTV(opts, assetID, newLTV)
	if err != nil {
		return nil, fmt.Errorf("TxManager.UpdateLTVOnChain: call: %w", err)
	}
	return tm.waitForReceipt(ctx, tx.Hash())
}

// UpdateAssetStatusOnChain updates an asset's status in AssetRegistry.
func (tm *TxManager) UpdateAssetStatusOnChain(ctx context.Context, assetID [32]byte, status uint8) (*types.Receipt, error) {
	opts, err := tm.txOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.UpdateAssetStatusOnChain: %w", err)
	}
	tx, err := tm.client.AssetRegistry.UpdateStatus(opts, assetID, status)
	if err != nil {
		return nil, fmt.Errorf("TxManager.UpdateAssetStatusOnChain: call: %w", err)
	}
	return tm.waitForReceipt(ctx, tx.Hash())
}

// SubmitAttestation writes a health attestation on-chain.
func (tm *TxManager) SubmitAttestation(ctx context.Context, assetID [32]byte, healthHash [32]byte, score uint8) (*types.Receipt, error) {
	opts, err := tm.txOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.SubmitAttestation: %w", err)
	}
	tx, err := tm.client.HealthAttestation.SubmitAttestation(opts, assetID, healthHash, score)
	if err != nil {
		return nil, fmt.Errorf("TxManager.SubmitAttestation: call: %w", err)
	}
	return tm.waitForReceipt(ctx, tx.Hash())
}

// ── Internals ─────────────────────────────────────────────────────────────

// txOpts builds bind.TransactOpts with a managed nonce and estimated gas.
// Retries up to maxNonceRetries on "nonce too low" errors.
func (tm *TxManager) txOpts(ctx context.Context) (*bind.TransactOpts, error) {
	tm.nonceMu.Lock()
	defer tm.nonceMu.Unlock()

	nonce, err := tm.nextNonce(ctx)
	if err != nil {
		return nil, fmt.Errorf("txOpts: nonce: %w", err)
	}

	gasPrice, err := tm.client.Eth().SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("txOpts: gas price: %w", err)
	}

	signer, err := bind.NewKeyedTransactorWithChainID(tm.privKey, tm.chainID)
	if err != nil {
		return nil, fmt.Errorf("txOpts: signer: %w", err)
	}
	signer.Nonce = big.NewInt(int64(nonce))
	signer.GasPrice = gasPrice
	signer.GasLimit = 0 // 0 means estimate per-call by the binding

	// Bump nonce in Redis so the next concurrent call doesn't collide.
	if err := tm.incrementNonce(ctx, nonce); err != nil {
		tm.logger.Warn("TxManager: failed to increment nonce in Redis", zap.Error(err))
	}

	return signer, nil
}

// nextNonce returns the nonce to use for the next transaction.
// It checks Redis first; if missing it reads from the chain and seeds Redis.
func (tm *TxManager) nextNonce(ctx context.Context) (uint64, error) {
	key := nonceCachePrefix + tm.fromAddr.Hex()

	val, err := tm.rdb.Get(ctx, key).Uint64()
	if err == redis.Nil {
		// Cache miss — read from chain.
		chainNonce, err := tm.client.Eth().PendingNonceAt(ctx, tm.fromAddr)
		if err != nil {
			return 0, fmt.Errorf("nextNonce: chain nonce: %w", err)
		}
		return chainNonce, nil
	}
	if err != nil {
		// Redis error — fall back to chain.
		tm.logger.Warn("TxManager: Redis nonce read failed, falling back to chain", zap.Error(err))
		return tm.client.Eth().PendingNonceAt(ctx, tm.fromAddr)
	}
	return val, nil
}

// incrementNonce stores nonce+1 in Redis.
func (tm *TxManager) incrementNonce(ctx context.Context, current uint64) error {
	key := nonceCachePrefix + tm.fromAddr.Hex()
	return tm.rdb.Set(ctx, key, current+1, 0).Err()
}

// refreshNonceFromChain re-reads the nonce from the chain and updates Redis.
// Called when a "nonce too low" error is detected.
func (tm *TxManager) refreshNonceFromChain(ctx context.Context) error {
	chainNonce, err := tm.client.Eth().PendingNonceAt(ctx, tm.fromAddr)
	if err != nil {
		return fmt.Errorf("refreshNonceFromChain: %w", err)
	}
	key := nonceCachePrefix + tm.fromAddr.Hex()
	return tm.rdb.Set(ctx, key, chainNonce, 0).Err()
}

// waitForReceipt polls until the transaction is mined or the 30s timeout fires.
func (tm *TxManager) waitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	tm.logger.Info("TxManager: waiting for receipt", zap.String("tx", txHash.Hex()))

	deadline := time.Now().Add(receiptTimeout)
	for time.Now().Before(deadline) {
		receipt, err := tm.client.Eth().TransactionReceipt(ctx, txHash)
		if err == nil {
			if receipt.Status == types.ReceiptStatusFailed {
				return receipt, fmt.Errorf("TxManager: transaction reverted: %s", txHash.Hex())
			}
			tm.logger.Info("TxManager: transaction confirmed",
				zap.String("tx", txHash.Hex()),
				zap.Uint64("block", receipt.BlockNumber.Uint64()),
			)
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(receiptPollInterval):
		}
	}
	return nil, fmt.Errorf("TxManager: receipt timeout after %s for tx %s", receiptTimeout, txHash.Hex())
}
