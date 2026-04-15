package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
	"go.uber.org/zap"
)

// MintHandler handles HTTP requests for WATT minting, redemption, and vault stats.
type MintHandler struct {
	mintService  service.MintServicer
	yieldService service.YieldServicer
	logger       *zap.Logger
}

// NewMintHandler constructs a MintHandler.
func NewMintHandler(svc service.MintServicer, yieldSvc service.YieldServicer, logger *zap.Logger) *MintHandler {
	return &MintHandler{mintService: svc, yieldService: yieldSvc, logger: logger}
}

// Mint godoc
// @Summary Mint WATT by depositing USDC or USDT
// @Tags mint
// @Accept json
// @Produce json
// @Param body body dto.MintRequest true "Mint payload"
// @Success 201 {object} response.Envelope{data=dto.MintResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/mint [post]
func (h *MintHandler) Mint(c *gin.Context) {
	var req dto.MintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	mintResp, err := h.mintService.Mint(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("MintHandler.Mint", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.Created(c, mintResp)
}

// Redeem godoc
// @Summary Redeem WATT back to stablecoin
// @Tags mint
// @Accept json
// @Produce json
// @Param body body dto.RedeemRequest true "Redeem payload"
// @Success 200 {object} response.Envelope{data=dto.MintResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/mint/redeem [post]
func (h *MintHandler) Redeem(c *gin.Context) {
	var req dto.RedeemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	mintResp, err := h.mintService.Redeem(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("MintHandler.Redeem", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, mintResp)
}

// GetNAV godoc
// @Summary Get the current NAV per sWATT share
// @Tags mint
// @Produce json
// @Success 200 {object} response.Envelope{data=dto.NAVResponse}
// @Security BearerAuth
// @Router /api/v1/mint/nav [get]
func (h *MintHandler) GetNAV(c *gin.Context) {
	navResp, err := h.mintService.GetNAV(c.Request.Context())
	if err != nil {
		h.logger.Error("MintHandler.GetNAV", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, navResp)
}

// GetVaultStats godoc
// @Summary Get sWattUSD vault statistics (NAV, APR, deployed capital)
// @Tags vault
// @Produce json
// @Success 200 {object} response.Envelope{data=dto.VaultStatsResponse}
// @Security BearerAuth
// @Router /api/v1/vault/stats [get]
func (h *MintHandler) GetVaultStats(c *gin.Context) {
	stats, err := h.yieldService.GetVaultStats(c.Request.Context())
	if err != nil {
		h.logger.Error("MintHandler.GetVaultStats", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, stats)
}
