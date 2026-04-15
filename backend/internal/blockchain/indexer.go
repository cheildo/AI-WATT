package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/blockchain/contracts"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
)

const redisLastBlockKey = "indexer:last_block"

// EventIndexer subscribes to XDC contract events and writes them to the DB.
// On startup it backfills from the last indexed block stored in Redis, then
// switches to a live subscription.
type EventIndexer struct {
	client    *BlockchainClient
	eventRepo repository.EventRepo
	rdb       *redis.Client
	logger    *zap.Logger

	// pre-built filterers — one per contract
	assetRegistryF    *contracts.AssetRegistryFilterer
	lendingPoolF      *contracts.LendingPoolFilterer
	wattUSDf          *contracts.WattUSDFilterer
	sWattUSDf         *contracts.SWattUSDFilterer
	mintEngineF       *contracts.MintEngineFilterer
	ocnftF            *contracts.OCNFTFilterer
	healthAttestationF *contracts.HealthAttestationFilterer
	wevQueueF         *contracts.WEVQueueFilterer
}

// NewEventIndexer constructs an EventIndexer.
func NewEventIndexer(
	client *BlockchainClient,
	eventRepo repository.EventRepo,
	rdb *redis.Client,
	logger *zap.Logger,
) (*EventIndexer, error) {
	eth := client.Eth()

	assetRegF, err := contracts.NewAssetRegistryFilterer(client.Addrs.AssetRegistry, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: AssetRegistryFilterer: %w", err)
	}
	lpF, err := contracts.NewLendingPoolFilterer(client.Addrs.LendingPool, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: LendingPoolFilterer: %w", err)
	}
	wusdF, err := contracts.NewWattUSDFilterer(client.Addrs.WattUSD, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: WattUSDFilterer: %w", err)
	}
	swusdF, err := contracts.NewSWattUSDFilterer(client.Addrs.SWattUSD, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: SWattUSDFilterer: %w", err)
	}
	meF, err := contracts.NewMintEngineFilterer(client.Addrs.MintEngine, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: MintEngineFilterer: %w", err)
	}
	ocnftF, err := contracts.NewOCNFTFilterer(client.Addrs.OCNFT, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: OCNFTFilterer: %w", err)
	}
	haF, err := contracts.NewHealthAttestationFilterer(client.Addrs.HealthAttestation, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: HealthAttestationFilterer: %w", err)
	}
	wevF, err := contracts.NewWEVQueueFilterer(client.Addrs.WEVQueue, eth)
	if err != nil {
		return nil, fmt.Errorf("EventIndexer: WEVQueueFilterer: %w", err)
	}

	return &EventIndexer{
		client:             client,
		eventRepo:          eventRepo,
		rdb:                rdb,
		logger:             logger,
		assetRegistryF:     assetRegF,
		lendingPoolF:       lpF,
		wattUSDf:           wusdF,
		sWattUSDf:          swusdF,
		mintEngineF:        meF,
		ocnftF:             ocnftF,
		healthAttestationF: haF,
		wevQueueF:          wevF,
	}, nil
}

// Start backfills historical events then starts a live log subscription.
// Blocks until ctx is cancelled.
func (idx *EventIndexer) Start(ctx context.Context) error {
	if err := idx.backfill(ctx); err != nil {
		idx.logger.Error("EventIndexer: backfill failed", zap.Error(err))
		// Non-fatal — continue to live subscription.
	}

	query := ethereum.FilterQuery{Addresses: idx.client.AllAddresses()}
	logsCh := make(chan types.Log, 256)
	sub, err := idx.client.Eth().SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		return fmt.Errorf("EventIndexer.Start: subscribe: %w", err)
	}
	idx.logger.Info("EventIndexer: live subscription active")

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("EventIndexer: subscription error: %w", err)
		case log := <-logsCh:
			idx.processLog(ctx, log)
		case <-ctx.Done():
			sub.Unsubscribe()
			return nil
		}
	}
}

// backfill fetches historical logs from last indexed block to the chain head.
func (idx *EventIndexer) backfill(ctx context.Context) error {
	fromBlock, err := idx.lastIndexedBlock(ctx)
	if err != nil {
		return fmt.Errorf("EventIndexer.backfill: last block: %w", err)
	}

	head, err := idx.client.GetLatestBlock(ctx)
	if err != nil {
		return fmt.Errorf("EventIndexer.backfill: get head: %w", err)
	}
	if fromBlock >= head {
		return nil
	}

	from := big.NewInt(int64(fromBlock + 1))
	to := big.NewInt(int64(head))
	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: idx.client.AllAddresses(),
	}
	logs, err := idx.client.Eth().FilterLogs(ctx, query)
	if err != nil {
		return fmt.Errorf("EventIndexer.backfill: filter logs: %w", err)
	}
	idx.logger.Info("EventIndexer: backfilling",
		zap.Uint64("from", fromBlock+1),
		zap.Uint64("to", head),
		zap.Int("logs", len(logs)),
	)
	for _, l := range logs {
		idx.processLog(ctx, l)
	}
	return idx.setLastIndexedBlock(ctx, head)
}

// processLog dispatches a raw log to the appropriate parser based on the
// emitting contract address, then persists the structured event.
func (idx *EventIndexer) processLog(ctx context.Context, log types.Log) {
	if len(log.Topics) == 0 {
		return
	}

	addr := log.Address
	var eventType string
	var argsMap map[string]any

	switch addr {
	case idx.client.Addrs.AssetRegistry:
		eventType, argsMap = idx.parseAssetRegistryLog(log)
	case idx.client.Addrs.LendingPool:
		eventType, argsMap = idx.parseLendingPoolLog(log)
	case idx.client.Addrs.WattUSD:
		eventType, argsMap = idx.parseWattUSDLog(log)
	case idx.client.Addrs.SWattUSD:
		eventType, argsMap = idx.parseSWattUSDLog(log)
	case idx.client.Addrs.MintEngine:
		eventType, argsMap = idx.parseMintEngineLog(log)
	case idx.client.Addrs.OCNFT:
		eventType, argsMap = idx.parseOCNFTLog(log)
	case idx.client.Addrs.HealthAttestation:
		eventType, argsMap = idx.parseHealthAttestationLog(log)
	case idx.client.Addrs.WEVQueue:
		eventType, argsMap = idx.parseWEVQueueLog(log)
	default:
		return
	}

	if eventType == "" {
		return // unrecognised / non-indexed event
	}

	argsJSON, _ := json.Marshal(argsMap)
	row := &models.ChainEvent{
		EventType:       eventType,
		ContractAddress: strings.ToLower(addr.Hex()),
		TxHash:          log.TxHash.Hex(),
		BlockNumber:     log.BlockNumber,
		LogIndex:        uint(log.Index),
		ParsedArgs:      string(argsJSON),
		CreatedAt:       time.Now().UTC(),
	}

	if err := idx.eventRepo.Create(ctx, row); err != nil {
		idx.logger.Error("EventIndexer: insert failed",
			zap.String("event", eventType),
			zap.String("tx", log.TxHash.Hex()),
			zap.Error(err),
		)
		return
	}

	if err := idx.setLastIndexedBlock(ctx, log.BlockNumber); err != nil {
		idx.logger.Warn("EventIndexer: failed to update last_block in Redis", zap.Error(err))
	}

	idx.logger.Debug("EventIndexer: indexed event",
		zap.String("type", eventType),
		zap.String("tx", log.TxHash.Hex()),
	)
}

// ── Per-contract parsers ───────────────────────────────────────────────────

func (idx *EventIndexer) parseAssetRegistryLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.assetRegistryF.ParseAssetRegistered(log); err == nil {
		return "AssetRegistered", map[string]any{
			"asset_id": fmt.Sprintf("0x%x", ev.AssetId),
			"borrower": ev.Borrower.Hex(),
			"ltv":      ev.Ltv,
		}
	}
	if ev, err := idx.assetRegistryF.ParseLTVUpdated(log); err == nil {
		return "LTVUpdated", map[string]any{
			"asset_id": fmt.Sprintf("0x%x", ev.AssetId),
			"old_ltv":  ev.OldLTV,
			"new_ltv":  ev.NewLTV,
		}
	}
	if ev, err := idx.assetRegistryF.ParseStatusChanged(log); err == nil {
		return "StatusChanged", map[string]any{
			"asset_id":   fmt.Sprintf("0x%x", ev.AssetId),
			"old_status": ev.OldStatus,
			"new_status": ev.NewStatus,
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseLendingPoolLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.lendingPoolF.ParseLoanOriginated(log); err == nil {
		return "LoanOriginated", map[string]any{
			"loan_id":   fmt.Sprintf("0x%x", ev.LoanId),
			"asset_id":  fmt.Sprintf("0x%x", ev.AssetId),
			"borrower":  ev.Borrower.Hex(),
			"principal": ev.Principal.String(),
		}
	}
	if ev, err := idx.lendingPoolF.ParseRepaymentReceived(log); err == nil {
		return "RepaymentReceived", map[string]any{
			"loan_id": fmt.Sprintf("0x%x", ev.LoanId),
			"amount":  ev.Amount.String(),
		}
	}
	if ev, err := idx.lendingPoolF.ParseLoanSettled(log); err == nil {
		return "LoanSettled", map[string]any{
			"loan_id": fmt.Sprintf("0x%x", ev.LoanId),
		}
	}
	if ev, err := idx.lendingPoolF.ParseLoanLiquidated(log); err == nil {
		return "LoanLiquidated", map[string]any{
			"loan_id": fmt.Sprintf("0x%x", ev.LoanId),
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseWattUSDLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.wattUSDf.ParseTransfer(log); err == nil {
		return "WattUSD.Transfer", map[string]any{
			"from":  ev.From.Hex(),
			"to":    ev.To.Hex(),
			"value": ev.Value.String(),
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseSWattUSDLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.sWattUSDf.ParseDeposit(log); err == nil {
		return "sWattUSD.Deposit", map[string]any{
			"sender": ev.Sender.Hex(),
			"owner":  ev.Owner.Hex(),
			"assets": ev.Assets.String(),
			"shares": ev.Shares.String(),
		}
	}
	if ev, err := idx.sWattUSDf.ParseWithdraw(log); err == nil {
		return "sWattUSD.Withdraw", map[string]any{
			"sender":   ev.Sender.Hex(),
			"receiver": ev.Receiver.Hex(),
			"owner":    ev.Owner.Hex(),
			"assets":   ev.Assets.String(),
			"shares":   ev.Shares.String(),
		}
	}
	if ev, err := idx.sWattUSDf.ParseYieldReceived(log); err == nil {
		return "sWattUSD.YieldReceived", map[string]any{
			"amount":          ev.Amount.String(),
			"new_total_assets": ev.NewTotalAssets.String(),
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseMintEngineLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.mintEngineF.ParseMinted(log); err == nil {
		return "MintEngine.Minted", map[string]any{
			"depositor":      ev.Depositor.Hex(),
			"watt_minted":    ev.WattMinted.String(),
			"deposit_amount": ev.DepositAmount.String(),
			"fee":            ev.Fee.String(),
		}
	}
	if ev, err := idx.mintEngineF.ParseRedeemed(log); err == nil {
		return "MintEngine.Redeemed", map[string]any{
			"redeemer":            ev.Redeemer.Hex(),
			"watt_burned":         ev.WattBurned.String(),
			"stablecoin_returned": ev.StablecoinReturned.String(),
			"fee":                 ev.Fee.String(),
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseOCNFTLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.ocnftF.ParseOCNFTMinted(log); err == nil {
		return "OCNFT.Minted", map[string]any{
			"to":       ev.To.Hex(),
			"token_id": ev.TokenId.String(),
			"asset_id": fmt.Sprintf("0x%x", ev.AssetId),
		}
	}
	if ev, err := idx.ocnftF.ParseOCNFTBurned(log); err == nil {
		return "OCNFT.Burned", map[string]any{
			"token_id": ev.TokenId.String(),
			"asset_id": fmt.Sprintf("0x%x", ev.AssetId),
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseHealthAttestationLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.healthAttestationF.ParseAttestationSubmitted(log); err == nil {
		return "HealthAttestation.Submitted", map[string]any{
			"asset_id":    fmt.Sprintf("0x%x", ev.AssetId),
			"health_hash": fmt.Sprintf("0x%x", ev.HealthHash),
			"score":       ev.Score,
		}
	}
	return "", nil
}

func (idx *EventIndexer) parseWEVQueueLog(log types.Log) (string, map[string]any) {
	if ev, err := idx.wevQueueF.ParseRedemptionRequested(log); err == nil {
		return "WEVQueue.RedemptionRequested", map[string]any{
			"request_id":   fmt.Sprintf("0x%x", ev.RequestId),
			"user":         ev.User.Hex(),
			"swatt_amount": ev.SWattAmount.String(),
			"is_priority":  ev.IsPriority,
		}
	}
	if ev, err := idx.wevQueueF.ParseBatchProcessed(log); err == nil {
		return "WEVQueue.BatchProcessed", map[string]any{
			"count": ev.Count.String(),
		}
	}
	return "", nil
}

// ── Redis helpers ──────────────────────────────────────────────────────────

func (idx *EventIndexer) lastIndexedBlock(ctx context.Context) (uint64, error) {
	val, err := idx.rdb.Get(ctx, redisLastBlockKey).Uint64()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("EventIndexer.lastIndexedBlock: redis get: %w", err)
	}
	return val, nil
}

func (idx *EventIndexer) setLastIndexedBlock(ctx context.Context, block uint64) error {
	return idx.rdb.Set(ctx, redisLastBlockKey, block, 0).Err()
}
