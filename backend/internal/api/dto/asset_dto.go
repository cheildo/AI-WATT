package dto

// RegisterAssetRequest is sent when a new hardware asset is onboarded.
type RegisterAssetRequest struct {
	AssetType string  `json:"asset_type" binding:"required,oneof=gpu_cluster robotics energy"`
	OwnerID   string  `json:"owner_id"   binding:"required,uuid"`
	InitialLTV float64 `json:"initial_ltv" binding:"required,gt=0,lte=1"`
}

// UpdateAssetLTVRequest allows backend to update LTV after a Veriflow score change.
type UpdateAssetLTVRequest struct {
	NewLTV float64 `json:"new_ltv" binding:"required,gt=0,lte=1"`
}

// AssetResponse is the public shape of a hardware asset record.
type AssetResponse struct {
	ID           string  `json:"id"`
	AssetType    string  `json:"asset_type"`
	LoanID       string  `json:"loan_id,omitempty"`
	OwnerID      string  `json:"owner_id"`
	HealthScore  float64 `json:"health_score"`
	CurrentLTV   float64 `json:"current_ltv"`
	Status       string  `json:"status"`
	OCNFTTokenID string  `json:"ocnft_token_id,omitempty"`
	MetadataURI  string  `json:"metadata_uri,omitempty"`
	CreatedAt    string  `json:"created_at"`
}

// ListAssetsQuery holds pagination and filter params.
type ListAssetsQuery struct {
	Page     int    `form:"page,default=1"       binding:"min=1"`
	PageSize int    `form:"page_size,default=20" binding:"min=1,max=100"`
	OwnerID  string `form:"owner_id"             binding:"omitempty,uuid"`
	Status   string `form:"status"               binding:"omitempty,oneof=pending active flagged liquidated"`
}

// ListAssetsResponse wraps a paginated asset list.
type ListAssetsResponse struct {
	Assets []AssetResponse `json:"assets"`
	Total  int64           `json:"total"`
	Page   int             `json:"page"`
}
