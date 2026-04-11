package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Envelope is the standard API response wrapper for all endpoints.
type Envelope struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *string     `json:"error"`
}

// OK sends a 200 response with data.
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Envelope{Success: true, Data: data, Error: nil})
}

// Created sends a 201 response with data.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Envelope{Success: true, Data: data, Error: nil})
}

// BadRequest sends a 400 response with an error message.
func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Envelope{Success: false, Data: nil, Error: &msg})
}

// Unauthorized sends a 401 response.
func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, Envelope{Success: false, Data: nil, Error: &msg})
}

// Forbidden sends a 403 response.
func Forbidden(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, Envelope{Success: false, Data: nil, Error: &msg})
}

// NotFound sends a 404 response.
func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Envelope{Success: false, Data: nil, Error: &msg})
}

// UnprocessableEntity sends a 422 response.
func UnprocessableEntity(c *gin.Context, msg string) {
	c.JSON(http.StatusUnprocessableEntity, Envelope{Success: false, Data: nil, Error: &msg})
}

// InternalError sends a 500 response with a generic message (never expose internal errors).
func InternalError(c *gin.Context) {
	msg := "internal server error"
	c.JSON(http.StatusInternalServerError, Envelope{Success: false, Data: nil, Error: &msg})
}
