package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
	"go.uber.org/zap"
)

// VeriflowHandler handles telemetry ingestion and health score queries.
type VeriflowHandler struct {
	veriflowService service.VeriflowServicer
	logger          *zap.Logger
}

// NewVeriflowHandler constructs a VeriflowHandler.
func NewVeriflowHandler(svc service.VeriflowServicer, logger *zap.Logger) *VeriflowHandler {
	return &VeriflowHandler{veriflowService: svc, logger: logger}
}

// IngestTelemetry godoc
// @Summary Ingest a signed telemetry payload from a Veriflow agent
// @Tags veriflow
// @Accept json
// @Produce json
// @Param body body dto.TelemetryPayload true "Signed telemetry payload"
// @Success 200 {object} response.Envelope
// @Failure 400 {object} response.Envelope
// @Failure 401 {object} response.Envelope
// @Router /api/v1/veriflow/telemetry [post]
func (h *VeriflowHandler) IngestTelemetry(c *gin.Context) {
	var payload dto.TelemetryPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	if err := h.veriflowService.IngestTelemetry(c.Request.Context(), payload); err != nil {
		h.logger.Error("VeriflowHandler.IngestTelemetry", zap.Error(err))
		response.Unauthorized(c, "telemetry rejected")
		return
	}
	response.OK(c, gin.H{"message": "telemetry accepted"})
}

// GetHealthScore godoc
// @Summary Get the current health score for an asset
// @Tags veriflow
// @Produce json
// @Param id path string true "Asset ID"
// @Success 200 {object} response.Envelope{data=dto.HealthScoreResponse}
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/veriflow/assets/{id}/score [get]
func (h *VeriflowHandler) GetHealthScore(c *gin.Context) {
	scoreResp, err := h.veriflowService.GetHealthScore(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.NotFound(c, "health score not found")
		return
	}
	response.OK(c, scoreResp)
}

// GetAttestation godoc
// @Summary Get the latest on-chain attestation for an asset
// @Tags veriflow
// @Produce json
// @Param id path string true "Asset ID"
// @Success 200 {object} response.Envelope{data=dto.AttestationResponse}
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/veriflow/assets/{id}/attestation [get]
func (h *VeriflowHandler) GetAttestation(c *gin.Context) {
	attResp, err := h.veriflowService.GetAttestation(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.NotFound(c, "attestation not found")
		return
	}
	response.OK(c, attResp)
}
