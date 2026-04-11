package blockchain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client wraps go-ethereum ethclient for XDC RPC interaction.
type Client struct {
	eth    *ethclient.Client
	rpcURL string
}

// NewClient dials the XDC RPC endpoint and returns a connected Client.
func NewClient(rpcURL string) (*Client, error) {
	eth, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		return nil, fmt.Errorf("blockchain.NewClient: dial %s: %w", rpcURL, err)
	}
	return &Client{eth: eth, rpcURL: rpcURL}, nil
}

// Eth exposes the raw ethclient for use by abigen bindings.
func (c *Client) Eth() *ethclient.Client {
	return c.eth
}

// Close shuts down the underlying RPC connection.
func (c *Client) Close() {
	c.eth.Close()
}
