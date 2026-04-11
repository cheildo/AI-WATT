package dto

// TelemetryPayload is the signed payload posted by the Veriflow agent every 5 minutes.
type TelemetryPayload struct {
	AssetID         string  `json:"asset_id"          binding:"required,uuid"`
	Timestamp       int64   `json:"timestamp"         binding:"required"`
	HMACSignature   string  `json:"hmac_signature"    binding:"required"`
	GPUUtilization  float64 `json:"gpu_utilization"`
	GPUTemperature  float64 `json:"gpu_temperature"`
	GPUMemoryUsedMB int64   `json:"gpu_memory_used_mb"`
	GPUErrorRate    float64 `json:"gpu_error_rate"`
	PowerDrawWatts  float64 `json:"power_draw_watts"`
	FanSpeedRPM     int     `json:"fan_speed_rpm"`
}

// HealthScoreResponse returns the current computed health score for an asset.
type HealthScoreResponse struct {
	AssetID     string  `json:"asset_id"`
	HealthScore float64 `json:"health_score"`
	Status      string  `json:"status"`
	ComputedAt  string  `json:"computed_at"`
}

// AttestationResponse returns the latest on-chain attestation for an asset.
type AttestationResponse struct {
	AssetID     string  `json:"asset_id"`
	HealthScore float64 `json:"health_score"`
	HealthHash  string  `json:"health_hash"`
	XDCTxHash   string  `json:"xdc_tx_hash"`
	Timestamp   string  `json:"timestamp"`
}
