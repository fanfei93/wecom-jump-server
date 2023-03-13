package redirect

import "github.com/gin-gonic/gin"

func HandleRedirect(c *gin.Context) {
	redirectUrl := c.Query("url")
	if redirectUrl == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "url can not be empty",
		})
		return
	}

	c.Redirect(301, redirectUrl)
}
