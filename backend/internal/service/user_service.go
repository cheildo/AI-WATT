package service

import (
	"context"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
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
	// TODO: inject UserRepo, JWT config, KYC client
}

// NewUserService constructs a UserService.
func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(ctx context.Context, req dto.CreateUserRequest) (dto.AuthResponse, error) {
	// TODO: hash password, persist user, issue JWT
	return dto.AuthResponse{}, nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	// TODO: look up user, verify password, issue JWT
	return dto.AuthResponse{}, nil
}

func (s *UserService) WalletLogin(ctx context.Context, req dto.WalletLoginRequest) (dto.AuthResponse, error) {
	// TODO: verify EIP-4361 signature, upsert user, issue JWT
	return dto.AuthResponse{}, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (dto.UserResponse, error) {
	// TODO: fetch from UserRepo, map model → dto
	return dto.UserResponse{}, nil
}

func (s *UserService) Update(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	// TODO: fetch, update, persist, return dto
	return dto.UserResponse{}, nil
}

func (s *UserService) List(ctx context.Context, q dto.ListUsersQuery) (dto.ListUsersResponse, error) {
	// TODO: paginated fetch from UserRepo
	return dto.ListUsersResponse{}, nil
}
