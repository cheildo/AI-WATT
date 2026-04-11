package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
	"go.uber.org/zap"
)

// AssetHandler handles HTTP requests for the assets domain.
type AssetHandler struct {
	assetService service.AssetServicer
	logger       *zap.Logger
}

// NewAssetHandler constructs an AssetHandler.
func NewAssetHandler(svc service.AssetServicer, logger *zap.Logger) *AssetHandler {
	return &AssetHandler{assetService: svc, logger: logger}
}

// List godoc
// @Summary List hardware assets
// @Tags assets
// @Produce json
// @Param query query dto.ListAssetsQuery false "Filter params"
// @Success 200 {object} response.Envelope{data=dto.ListAssetsResponse}
// @Security BearerAuth
// @Router /api/v1/assets [get]
func (h *AssetHandler) List(c *gin.Context) {
	var q dto.ListAssetsQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	result, err := h.assetService.List(c.Request.Context(), q)
	if err != nil {
		h.logger.Error("AssetHandler.List", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, result)
}

// Register godoc
// @Summary Register a new hardware asset
// @Tags assets
// @Accept json
// @Produce json
// @Param body body dto.RegisterAssetRequest true "Asset registration payload"
// @Success 201 {object} response.Envelope{data=dto.AssetResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/assets [post]
func (h *AssetHandler) Register(c *gin.Context) {
	var req dto.RegisterAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	assetResp, err := h.assetService.Register(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("AssetHandler.Register", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.Created(c, assetResp)
}

// GetByID godoc
// @Summary Get an asset by ID
// @Tags assets
// @Produce json
// @Param id path string true "Asset ID"
// @Success 200 {object} response.Envelope{data=dto.AssetResponse}
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/assets/{id} [get]
func (h *AssetHandler) GetByID(c *gin.Context) {
	assetResp, err := h.assetService.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.NotFound(c, "asset not found")
		return
	}
	response.OK(c, assetResp)
}

// UpdateLTV godoc
// @Summary Update asset LTV (admin only — triggered by Veriflow scoring)
// @Tags assets
// @Accept json
// @Produce json
// @Param id path string true "Asset ID"
// @Param body body dto.UpdateAssetLTVRequest true "LTV update payload"
// @Success 200 {object} response.Envelope{data=dto.AssetResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/assets/{id}/ltv [patch]
func (h *AssetHandler) UpdateLTV(c *gin.Context) {
	var req dto.UpdateAssetLTVRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	assetResp, err := h.assetService.UpdateLTV(c.Request.Context(), c.Param("id"), req.NewLTV)
	if err != nil {
		h.logger.Error("AssetHandler.UpdateLTV", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, assetResp)
}
