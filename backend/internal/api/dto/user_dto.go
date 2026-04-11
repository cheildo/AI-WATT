package dto

// CreateUserRequest is sent when a user registers.
type CreateUserRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Email         string `json:"email"          binding:"required,email"`
	Password      string `json:"password"       binding:"required,min=8"`
}

// LoginRequest is sent for email/password login.
type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// WalletLoginRequest is sent for Sign-In with XDC (EIP-4361).
type WalletLoginRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Message       string `json:"message"        binding:"required"`
	Signature     string `json:"signature"      binding:"required"`
}

// UpdateUserRequest allows updating user profile fields.
type UpdateUserRequest struct {
	Email string `json:"email" binding:"omitempty,email"`
}

// UserResponse is the public representation of a user — no password hash.
type UserResponse struct {
	ID            string `json:"id"`
	WalletAddress string `json:"wallet_address"`
	Email         string `json:"email,omitempty"`
	Role          string `json:"role"`
	KYCStatus     string `json:"kyc_status"`
	CreatedAt     string `json:"created_at"`
}

// AuthResponse wraps a JWT token returned after login.
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// ListUsersQuery holds pagination and filter params for the user list.
type ListUsersQuery struct {
	Page     int    `form:"page,default=1"      binding:"min=1"`
	PageSize int    `form:"page_size,default=20" binding:"min=1,max=100"`
	Role     string `form:"role"                binding:"omitempty,oneof=depositor borrower curator admin"`
}

// ListUsersResponse wraps a paginated user list.
type ListUsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
}
