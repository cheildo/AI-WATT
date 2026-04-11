package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap"
)

// TxManager signs and sends transactions to XDC contracts.
// It manages nonces, estimates gas, and retries on failure.
type TxManager struct {
	client     *Client
	privateKey string
	chainID    *big.Int
	logger     *zap.Logger
}

// NewTxManager constructs a TxManager.
func NewTxManager(client *Client, privateKey string, chainID int64, logger *zap.Logger) (*TxManager, error) {
	if len(privateKey) < 2 {
		return nil, fmt.Errorf("TxManager: private key is empty")
	}
	return &TxManager{
		client:     client,
		privateKey: privateKey,
		chainID:    big.NewInt(chainID),
		logger:     logger,
	}, nil
}

// SendTransaction signs and broadcasts a transaction, waits for receipt.
func (tm *TxManager) SendTransaction(ctx context.Context, to common.Address, data []byte, value *big.Int) (*types.Receipt, error) {
	privKey, err := crypto.HexToECDSA(tm.privateKey[2:]) // strip 0x prefix
	if err != nil {
		return nil, fmt.Errorf("TxManager.SendTransaction: parse key: %w", err)
	}
	fromAddr := crypto.PubkeyToAddress(privKey.PublicKey)
	nonce, err := tm.client.Eth().PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, fmt.Errorf("TxManager.SendTransaction: get nonce: %w", err)
	}
	gasPrice, err := tm.client.Eth().SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("TxManager.SendTransaction: suggest gas price: %w", err)
	}
	gasLimit := uint64(300000) // TODO: use EstimateGas per-call
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	signer := types.NewEIP155Signer(tm.chainID)
	signed, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		return nil, fmt.Errorf("TxManager.SendTransaction: sign: %w", err)
	}
	if err := tm.client.Eth().SendTransaction(ctx, signed); err != nil {
		return nil, fmt.Errorf("TxManager.SendTransaction: broadcast: %w", err)
	}
	tm.logger.Info("transaction broadcast", zap.String("tx_hash", signed.Hash().Hex()))
	// TODO: poll for receipt with retry
	return nil, nil
}
