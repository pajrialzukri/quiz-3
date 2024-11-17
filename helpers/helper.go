package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, statusCode int, message string, status string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"status":  status,
		"data":    data,
	})
}
