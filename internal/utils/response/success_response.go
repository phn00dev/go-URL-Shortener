package response

import "github.com/gin-gonic/gin"

// successResponse struct definition
type successResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// SuccessResponse function to return a success response
func Success(c *gin.Context, status int, message string, data any) {
	c.JSON(status, &successResponse{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
