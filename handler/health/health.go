package health

import "github.com/gin-gonic/gin"

func HandleCheckHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
	return
}
