package response

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func Error(c *gin.Context, status int, message, errors string) {
	c.JSON(status, &errorResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
	})
}
