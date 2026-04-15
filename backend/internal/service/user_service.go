package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/models"
	"github.com/neurowatt/aiwatt-backend/internal/repository"
	pkgjwt "github.com/neurowatt/aiwatt-backend/pkg/jwt"
)

// UserServicer defines the user business logic interface.
type UserServicer interface {
	Register(ctx context.Context, req dto.CreateUserRequest) (dto.AuthResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error)
	WalletLogin(ctx context.Context, req dto.WalletLoginRequest) (dto.AuthResponse, error)
	GetByID(ctx context.Context, id string) (dto.UserResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error)
	List(ctx context.Context, q dto.ListUsersQuery) (dto.ListUsersResponse, error)
}

// UserService implements UserServicer.
type UserService struct {
	userRepo  repository.UserRepo
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepo, jwtSecret string) *UserService {
	return &UserService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s *UserService) Register(ctx context.Context, req dto.CreateUserRequest) (dto.AuthResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.Register: hash: %w", err)
	}

	user := &models.User{
		ID:            uuid.NewString(),
		WalletAddress: strings.ToLower(req.WalletAddress),
		Email:         req.Email,
		PasswordHash:  string(hash),
		Role:          models.RoleDepositor,
		KYCStatus:     models.KYCPending,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.Register: create: %w", err)
	}

	token, err := pkgjwt.Generate(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.Register: jwt: %w", err)
	}

	return dto.AuthResponse{Token: token, User: toUserDTO(user)}, nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.AuthResponse{}, errors.New("invalid credentials")
		}
		return dto.AuthResponse{}, fmt.Errorf("user_service.Login: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return dto.AuthResponse{}, errors.New("invalid credentials")
	}
	token, err := pkgjwt.Generate(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.Login: jwt: %w", err)
	}
	return dto.AuthResponse{Token: token, User: toUserDTO(user)}, nil
}

// WalletLogin verifies an EIP-191 personal_sign signature and issues a JWT.
// The client must sign the exact message text provided. The recovered address is
// compared against the claimed wallet address (case-insensitive).
func (s *UserService) WalletLogin(ctx context.Context, req dto.WalletLoginRequest) (dto.AuthResponse, error) {
	wallet := strings.ToLower(req.WalletAddress)

	if err := verifyPersonalSign(req.Message, req.Signature, wallet); err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.WalletLogin: signature invalid: %w", err)
	}

	// Upsert — find or create the user for this wallet.
	user, err := s.userRepo.GetByWallet(ctx, wallet)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &models.User{
			ID:            uuid.NewString(),
			WalletAddress: wallet,
			Role:          models.RoleDepositor,
			KYCStatus:     models.KYCPending,
		}
		if createErr := s.userRepo.Create(ctx, user); createErr != nil {
			return dto.AuthResponse{}, fmt.Errorf("user_service.WalletLogin: create: %w", createErr)
		}
	} else if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.WalletLogin: lookup: %w", err)
	}

	token, err := pkgjwt.Generate(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("user_service.WalletLogin: jwt: %w", err)
	}
	return dto.AuthResponse{Token: token, User: toUserDTO(user)}, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("user_service.GetByID: %w", err)
	}
	return toUserDTO(user), nil
}

func (s *UserService) Update(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("user_service.Update: %w", err)
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if err := s.userRepo.Update(ctx, user); err != nil {
		return dto.UserResponse{}, fmt.Errorf("user_service.Update: save: %w", err)
	}
	return toUserDTO(user), nil
}

func (s *UserService) List(ctx context.Context, q dto.ListUsersQuery) (dto.ListUsersResponse, error) {
	offset := (q.Page - 1) * q.PageSize
	users, total, err := s.userRepo.List(ctx, q.Role, offset, q.PageSize)
	if err != nil {
		return dto.ListUsersResponse{}, fmt.Errorf("user_service.List: %w", err)
	}
	resp := make([]dto.UserResponse, len(users))
	for i, u := range users {
		resp[i] = toUserDTO(u)
	}
	return dto.ListUsersResponse{Users: resp, Total: total, Page: q.Page}, nil
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func toUserDTO(u *models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:            u.ID,
		WalletAddress: u.WalletAddress,
		Email:         u.Email,
		Role:          u.Role,
		KYCStatus:     u.KYCStatus,
		CreatedAt:     u.CreatedAt.Format(time.RFC3339),
	}
}

// verifyPersonalSign recovers the signer from an EIP-191 personal_sign signature
// and compares it to the expected wallet address.
func verifyPersonalSign(message, sigHex, expectedWallet string) error {
	// Strip 0x prefix from signature
	sigHex = strings.TrimPrefix(sigHex, "0x")
	if len(sigHex) != 130 {
		return errors.New("signature must be 65 bytes (130 hex chars)")
	}

	sigBytes := common.Hex2Bytes(sigHex)

	// EIP-191: prefix the message with "\x19Ethereum Signed Message:\n<length>"
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256([]byte(msg))

	// Adjust v byte: Ethereum uses 27/28, go-ethereum crypto uses 0/1
	if sigBytes[64] >= 27 {
		sigBytes[64] -= 27
	}

	pubKey, err := crypto.SigToPub(hash, sigBytes)
	if err != nil {
		return fmt.Errorf("recover pubkey: %w", err)
	}
	recovered := strings.ToLower(crypto.PubkeyToAddress(*pubKey).Hex())
	if recovered != strings.ToLower(expectedWallet) {
		return fmt.Errorf("signature mismatch: got %s, expected %s", recovered, expectedWallet)
	}
	return nil
}
