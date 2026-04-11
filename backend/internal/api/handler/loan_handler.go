package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/neurowatt/aiwatt-backend/internal/api/dto"
	"github.com/neurowatt/aiwatt-backend/internal/service"
	"github.com/neurowatt/aiwatt-backend/pkg/response"
	"go.uber.org/zap"
)

// LoanHandler handles HTTP requests for the loans domain.
type LoanHandler struct {
	loanService service.LoanServicer
	logger      *zap.Logger
}

// NewLoanHandler constructs a LoanHandler.
func NewLoanHandler(svc service.LoanServicer, logger *zap.Logger) *LoanHandler {
	return &LoanHandler{loanService: svc, logger: logger}
}

// List godoc
// @Summary List loans with optional filters
// @Tags loans
// @Produce json
// @Param query query dto.ListLoansQuery false "Filter params"
// @Success 200 {object} response.Envelope{data=dto.ListLoansResponse}
// @Security BearerAuth
// @Router /api/v1/loans [get]
func (h *LoanHandler) List(c *gin.Context) {
	var q dto.ListLoansQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	result, err := h.loanService.List(c.Request.Context(), q)
	if err != nil {
		h.logger.Error("LoanHandler.List", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, result)
}

// Create godoc
// @Summary Submit a new loan / financing request
// @Tags loans
// @Accept json
// @Produce json
// @Param body body dto.CreateLoanRequest true "Loan request"
// @Success 201 {object} response.Envelope{data=dto.LoanResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/loans [post]
func (h *LoanHandler) Create(c *gin.Context) {
	var req dto.CreateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	loanResp, err := h.loanService.Create(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("LoanHandler.Create", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.Created(c, loanResp)
}

// GetByID godoc
// @Summary Get a loan by ID
// @Tags loans
// @Produce json
// @Param id path string true "Loan ID"
// @Success 200 {object} response.Envelope{data=dto.LoanResponse}
// @Failure 404 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/loans/{id} [get]
func (h *LoanHandler) GetByID(c *gin.Context) {
	loanResp, err := h.loanService.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		response.NotFound(c, "loan not found")
		return
	}
	response.OK(c, loanResp)
}

// Update godoc
// @Summary Update loan status (curator/admin only)
// @Tags loans
// @Accept json
// @Produce json
// @Param id path string true "Loan ID"
// @Param body body dto.UpdateLoanRequest true "Update payload"
// @Success 200 {object} response.Envelope{data=dto.LoanResponse}
// @Failure 400 {object} response.Envelope
// @Security BearerAuth
// @Router /api/v1/loans/{id} [patch]
func (h *LoanHandler) Update(c *gin.Context) {
	var req dto.UpdateLoanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	loanResp, err := h.loanService.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		h.logger.Error("LoanHandler.Update", zap.Error(err))
		response.InternalError(c)
		return
	}
	response.OK(c, loanResp)
}
