package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/neurowatt/aiwatt-backend/internal/blockchain/contracts"
)

// ContractAddresses holds the deployed proxy address for every contract.
type ContractAddresses struct {
	WattUSD          common.Address
	SWattUSD         common.Address
	MintEngine       common.Address
	AssetRegistry    common.Address
	OCNFT            common.Address
	HealthAttestation common.Address
	LendingPool      common.Address
	WEVQueue         common.Address
}

// BlockchainClient wraps go-ethereum ethclient and exposes typed contract bindings.
type BlockchainClient struct {
	eth               *ethclient.Client
	rpcURL            string
	Addrs             ContractAddresses
	WattUSD           *contracts.WattUSD
	SWattUSD          *contracts.SWattUSD
	MintEngine        *contracts.MintEngine
	AssetRegistry     *contracts.AssetRegistry
	OCNFT             *contracts.OCNFT
	HealthAttestation *contracts.HealthAttestation
	LendingPool       *contracts.LendingPool
	WEVQueue          *contracts.WEVQueue
}

// NewBlockchainClient dials the XDC RPC endpoint, initialises all contract bindings,
// and returns a ready-to-use BlockchainClient.
func NewBlockchainClient(rpcURL string, addrs ContractAddresses) (*BlockchainClient, error) {
	eth, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: dial %s: %w", rpcURL, err)
	}

	wattUSD, err := contracts.NewWattUSD(addrs.WattUSD, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: WattUSD binding: %w", err)
	}
	sWattUSD, err := contracts.NewSWattUSD(addrs.SWattUSD, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: sWattUSD binding: %w", err)
	}
	mintEngine, err := contracts.NewMintEngine(addrs.MintEngine, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: MintEngine binding: %w", err)
	}
	assetRegistry, err := contracts.NewAssetRegistry(addrs.AssetRegistry, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: AssetRegistry binding: %w", err)
	}
	ocnft, err := contracts.NewOCNFT(addrs.OCNFT, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: OCNFT binding: %w", err)
	}
	healthAtt, err := contracts.NewHealthAttestation(addrs.HealthAttestation, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: HealthAttestation binding: %w", err)
	}
	lendingPool, err := contracts.NewLendingPool(addrs.LendingPool, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: LendingPool binding: %w", err)
	}
	wevQueue, err := contracts.NewWEVQueue(addrs.WEVQueue, eth)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewBlockchainClient: WEVQueue binding: %w", err)
	}

	return &BlockchainClient{
		eth:               eth,
		rpcURL:            rpcURL,
		Addrs:             addrs,
		WattUSD:           wattUSD,
		SWattUSD:          sWattUSD,
		MintEngine:        mintEngine,
		AssetRegistry:     assetRegistry,
		OCNFT:             ocnft,
		HealthAttestation: healthAtt,
		LendingPool:       lendingPool,
		WEVQueue:          wevQueue,
	}, nil
}

// Eth exposes the raw ethclient for use by abigen bindings and TxManager.
func (c *BlockchainClient) Eth() *ethclient.Client {
	return c.eth
}

// Close shuts down the underlying RPC connection.
func (c *BlockchainClient) Close() {
	c.eth.Close()
}

// GetLatestBlock returns the current head block number.
func (c *BlockchainClient) GetLatestBlock(ctx context.Context) (uint64, error) {
	n, err := c.eth.BlockNumber(ctx)
	if err != nil {
		return 0, fmt.Errorf("BlockchainClient.GetLatestBlock: %w", err)
	}
	return n, nil
}

// GetTransactionReceipt returns the receipt for the given tx hash, or nil if not yet mined.
func (c *BlockchainClient) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	receipt, err := c.eth.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("BlockchainClient.GetTransactionReceipt: %w", err)
	}
	return receipt, nil
}

// AllAddresses returns the slice of all watched contract addresses (for the EventIndexer).
func (c *BlockchainClient) AllAddresses() []common.Address {
	return []common.Address{
		c.Addrs.WattUSD,
		c.Addrs.SWattUSD,
		c.Addrs.MintEngine,
		c.Addrs.AssetRegistry,
		c.Addrs.OCNFT,
		c.Addrs.HealthAttestation,
		c.Addrs.LendingPool,
		c.Addrs.WEVQueue,
	}
}

// NAVPerShare reads the current NAV per sWATT share from the sWattUSD vault.
func (c *BlockchainClient) NAVPerShare(ctx context.Context) (*big.Int, error) {
	one := big.NewInt(1e18) // 1 share in 18-decimal units
	assets, err := c.SWattUSD.ConvertToAssets(nil, one)
	if err != nil {
		return nil, fmt.Errorf("BlockchainClient.NAVPerShare: %w", err)
	}
	return assets, nil
}

// IsAssetActive returns true when the given assetId (bytes32) is ACTIVE in AssetRegistry.
func (c *BlockchainClient) IsAssetActive(ctx context.Context, assetID [32]byte) (bool, error) {
	active, err := c.AssetRegistry.IsActive(nil, assetID)
	if err != nil {
		return false, fmt.Errorf("BlockchainClient.IsAssetActive: %w", err)
	}
	return active, nil
}
