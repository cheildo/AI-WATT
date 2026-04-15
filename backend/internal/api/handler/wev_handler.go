package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
)

// WEVHandler handles HTTP requests for the sWATT redemption queue.
type WEVHandler struct {
	wevService service.WEVServicer
	logger     *zap.Logger
}

// NewWEVHandler constructs a WEVHandler.
func NewWEVHandler(svc service.WEVServicer, logger *zap.Logger) *WEVHandler {
	return &WEVHandler{wevService: svc, logger: logger}
}

// RequestRedeem godoc
// @Summary Enter the sWATT redemption queue
// @Tags wev
// @Accept json
// @Produce json
// @Param body body dto.RedemptionRequest true "Redemption request"
// @Success 201 {object} response.Envelope{data=dto.RedemptionResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/wev/redeem [post]
func (h *WEVHandler) RequestRedeem(c *gin.Context) {
	var req dto.RedemptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	userID, _ := c.Get("user_id")
	resp, err := h.wevService.Enqueue(c.Request.Context(), userID.(string), req)
	if err != nil {
		h.logger.Error("WEVHandler.RequestRedeem", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.Created(c, resp)
}

// CancelRedeem godoc
// @Summary Cancel a QUEUED sWATT redemption request
// @Tags wev
// @Produce json
// @Param requestId path string true "Request ID (bytes32 hex)"
// @Success 200 {object} response.Envelope
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/wev/redeem/{requestId} [delete]
func (h *WEVHandler) CancelRedeem(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if err := h.wevService.Cancel(c.Request.Context(), userID.(string), c.Param("requestId")); err != nil {
		h.logger.Error("WEVHandler.CancelRedeem", zap.Error(err))
		response.NotFound(c, "request not found or cannot be cancelled")
		return
	}
	response.OK(c, gin.H{"cancelled": true})
}

// GetQueue godoc
// @Summary Get WEV queue status
// @Tags wev
// @Produce json
// @Success 200 {object} response.Envelope{data=dto.QueueStatusResponse}
// @Security BearerAuth
// @Router /api/v1/wev/queue [get]
func (h *WEVHandler) GetQueue(c *gin.Context) {
	status, err := h.wevService.GetQueueStatus(c.Request.Context())
	if err != nil {
		h.logger.Error("WEVHandler.GetQueue", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, status)
}

// GetMyQueue godoc
// @Summary Get the authenticated user's redemption requests
// @Tags wev
// @Produce json
// @Success 200 {object} response.Envelope{data=[]dto.RedemptionResponse}
// @Security BearerAuth
// @Router /api/v1/wev/queue/me [get]
func (h *WEVHandler) GetMyQueue(c *gin.Context) {
	userID, _ := c.Get("user_id")
	requests, err := h.wevService.GetUserQueue(c.Request.Context(), userID.(string))
	if err != nil {
		h.logger.Error("WEVHandler.GetMyQueue", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, requests)
}
