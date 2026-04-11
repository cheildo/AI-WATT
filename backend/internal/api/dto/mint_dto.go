package dto

// MintRequest is sent when a depositor wants to mint WATT by depositing stablecoin.
type MintRequest struct {
	Amount        float64 `json:"amount"         binding:"required,gt=0"`
	StablecoinType string `json:"stablecoin_type" binding:"required,oneof=USDC USDT"`
	DepositorID   string  `json:"depositor_id"   binding:"required,uuid"`
}

// RedeemRequest is sent when a holder wants to redeem WATT back to stablecoin.
type RedeemRequest struct {
	Amount      float64 `json:"amount"       binding:"required,gt=0"`
	HolderID    string  `json:"holder_id"    binding:"required,uuid"`
}

// MintResponse is returned after a successful mint or redeem operation.
type MintResponse struct {
	TxHash        string  `json:"tx_hash"`
	AmountMinted  float64 `json:"amount_minted,omitempty"`
	AmountRedeemed float64 `json:"amount_redeemed,omitempty"`
	Fee           float64 `json:"fee"`
}

// NAVResponse returns the current NAV per sWATT share.
type NAVResponse struct {
	NAVPerShare    float64 `json:"nav_per_share"`
	TotalAssets    float64 `json:"total_assets"`
	TotalSupply    float64 `json:"total_supply"`
	APR            float64 `json:"apr"`
	LastUpdatedAt  string  `json:"last_updated_at"`
}
