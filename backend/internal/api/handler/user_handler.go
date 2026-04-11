package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
	"go.uber.org/zap"
)

// UserHandler handles HTTP requests for the users domain.
type UserHandler struct {
	userService service.UserServicer
	logger      *zap.Logger
}

// NewUserHandler constructs a UserHandler.
func NewUserHandler(svc service.UserServicer, logger *zap.Logger) *UserHandler {
	return &UserHandler{userService: svc, logger: logger}
}

// Register godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.CreateUserRequest true "Registration payload"
// @Success 201 {object} response.Envelope{data=dto.AuthResponse}
// @Failure 400 {object} response.Envelope
// @Router /api/v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	authResp, err := h.userService.Register(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("UserHandler.Register", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.Created(c, authResp)
}

// Login godoc
// @Summary Log in with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login payload"
// @Success 200 {object} response.Envelope{data=dto.AuthResponse}
// @Failure 400 {object} response.Envelope
// @Failure 401 {object} response.Envelope
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	authResp, err := h.userService.Login(c.Request.Context(), req)
	if err != nil {
		response.Unauthorized(c, "invalid credentials")
		return
	}
	response.OK(c, authResp)
}

// WalletLogin godoc
// @Summary Log in with a signed XDC wallet message (EIP-4361)
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.WalletLoginRequest true "Wallet login payload"
// @Success 200 {object} response.Envelope{data=dto.AuthResponse}
// @Failure 400 {object} response.Envelope
// @Failure 401 {object} response.Envelope
// @Router /api/v1/auth/wallet-login [post]
func (h *UserHandler) WalletLogin(c *gin.Context) {
	var req dto.WalletLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	authResp, err := h.userService.WalletLogin(c.Request.Context(), req)
	if err != nil {
		response.Unauthorized(c, "wallet signature verification failed")
		return
	}
	response.OK(c, authResp)
}

// GetByID godoc
// @Summary Get a user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Envelope{data=dto.UserResponse}
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	userResp, err := h.userService.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.NotFound(c, "user not found")
		return
	}
	response.OK(c, userResp)
}

// Update godoc
// @Summary Update user profile
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param body body dto.UpdateUserRequest true "Update payload"
// @Success 200 {object} response.Envelope{data=dto.UserResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/users/{id} [patch]
func (h *UserHandler) Update(c *gin.Context) {
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	userResp, err := h.userService.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		h.logger.Error("UserHandler.Update", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, userResp)
}
