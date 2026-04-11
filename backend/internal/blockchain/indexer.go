package blockchain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// EventIndexer subscribes to XDC contract events and writes them to the DB.
type EventIndexer struct {
	client    *Client
	addresses []common.Address
	logger    *zap.Logger
	// TODO: inject EventRepo
}

// NewEventIndexer constructs an EventIndexer watching the given contract addresses.
func NewEventIndexer(client *Client, addresses []common.Address, logger *zap.Logger) *EventIndexer {
	return &EventIndexer{
		client:    client,
		addresses: addresses,
		logger:    logger,
	}
}

// Start subscribes to logs and processes them in a goroutine.
// Call cancel() on the returned context to stop.
func (idx *EventIndexer) Start(ctx context.Context) error {
	query := ethereum.FilterQuery{Addresses: idx.addresses}
	logsCh := make(chan types.Log, 100)
	sub, err := idx.client.Eth().SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		return fmt.Errorf("EventIndexer.Start: subscribe: %w", err)
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				idx.logger.Error("EventIndexer subscription error", zap.Error(err))
				return
			case log := <-logsCh:
				idx.processLog(ctx, log)
			case <-ctx.Done():
				sub.Unsubscribe()
				return
			}
		}
	}()
	return nil
}

func (idx *EventIndexer) processLog(ctx context.Context, log types.Log) {
	// TODO: parse event type, persist to chain_events table via EventRepo
	idx.logger.Debug("EventIndexer: received log",
		zap.String("tx_hash", log.TxHash.Hex()),
		zap.Uint64("block", log.BlockNumber),
	)
}
