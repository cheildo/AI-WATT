package dto

// RedemptionRequest is sent when a user wants to redeem sWATT.
type RedemptionRequest struct {
	SWattAmount uint64 `json:"swatt_amount" binding:"required,gt=0"`
	Priority    bool   `json:"priority"`
	PriorityFee uint64 `json:"priority_fee"` // required when priority=true
}

// RedemptionResponse is returned after a successful queue entry.
type RedemptionResponse struct {
	RequestID     string `json:"request_id"`
	EstimatedDays int    `json:"estimated_days"` // 3 for priority, 30 for standard
	Status        string `json:"status"`
}

// QueueStatusResponse describes the current state of the WEV redemption queue.
type QueueStatusResponse struct {
	DepthSWatt      uint64 `json:"depth_swatt"`       // total sWATT in queue
	QueueDepth      int64  `json:"queue_depth"`       // number of pending requests
	NextProcessing  string `json:"next_processing"`   // ISO8601 timestamp
	StandardDays    int    `json:"standard_days"`
	PriorityDays    int    `json:"priority_days"`
	PriorityFeeBPS  int    `json:"priority_fee_bps"`
}
