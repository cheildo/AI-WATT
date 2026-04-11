package dto

// CreateLoanRequest is sent when a borrower submits a financing request.
type CreateLoanRequest struct {
	AssetID    string  `json:"asset_id"    binding:"required,uuid"`
	Amount     float64 `json:"amount"      binding:"required,gt=0"`
	EngineType int     `json:"engine_type" binding:"required,oneof=1 2 3"`
	BorrowerID string  `json:"borrower_id" binding:"required,uuid"`
}

// UpdateLoanRequest allows a curator to update loan status.
type UpdateLoanRequest struct {
	Status    string `json:"status"     binding:"required,oneof=active repaying settled defaulted"`
	CuratorID string `json:"curator_id" binding:"omitempty,uuid"`
}

// LoanResponse is the public shape of a loan record.
type LoanResponse struct {
	ID            string  `json:"id"`
	AssetID       string  `json:"asset_id"`
	BorrowerID    string  `json:"borrower_id"`
	CuratorID     string  `json:"curator_id,omitempty"`
	Amount        float64 `json:"amount"`
	EngineType    int     `json:"engine_type"`
	Status        string  `json:"status"`
	OnChainTxHash string  `json:"on_chain_tx_hash,omitempty"`
	CreatedAt     string  `json:"created_at"`
}

// ListLoansQuery holds pagination and filter params.
type ListLoansQuery struct {
	Page       int    `form:"page,default=1"       binding:"min=1"`
	PageSize   int    `form:"page_size,default=20" binding:"min=1,max=100"`
	Status     string `form:"status"               binding:"omitempty,oneof=pending active repaying settled defaulted"`
	BorrowerID string `form:"borrower_id"          binding:"omitempty,uuid"`
}

// ListLoansResponse wraps a paginated loan list.
type ListLoansResponse struct {
	Loans []LoanResponse `json:"loans"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
}
