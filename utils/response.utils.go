package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess ...
func ResponseSuccess(c *gin.Context, success bool, message string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

// ResponseUnauthorized ...
func ResponseUnauthorized(c *gin.Context, success bool) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"success": success,
		"message": "Unauthorized",
	})
}

// ResponseBadRequest ...
func ResponseBadRequest(c *gin.Context, success bool, message string, data gin.H) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

// ResponseServerError ...
func ResponseServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
	})
}
