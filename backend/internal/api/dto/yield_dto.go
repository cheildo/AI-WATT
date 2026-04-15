package dto

// VaultStatsResponse describes the current state of the sWattUSD vault.
type VaultStatsResponse struct {
	NAVPerShare    float64 `json:"nav_per_share"`
	TotalAssets    float64 `json:"total_assets"`
	TotalSupply    float64 `json:"total_supply"`
	DeployedPct    float64 `json:"deployed_pct"`    // fraction in active loans
	TBillReserve   float64 `json:"t_bill_reserve"`  // idle capital in T-bills (Phase 12)
	APR7D          float64 `json:"apr_7d"`
	APR30D         float64 `json:"apr_30d"`
	LastUpdatedAt  string  `json:"last_updated_at"`
}
