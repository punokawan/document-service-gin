package helpers

import "github.com/gin-gonic/gin"

func ResponseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"error": true,
	})
}

func ResponseSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"error": true,
		"data": data,
	})
}