package response

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

type successResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Error(c *gin.Context, status int, message, errors string) {
	c.JSON(status, &errorResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
	})
}

func Success(c *gin.Context, status int, message string, data any) {
	c.JSON(status, &successResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
