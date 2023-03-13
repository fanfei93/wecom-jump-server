package permission

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"wecom-jump-server/config"
	"wecom-jump-server/service/permission"
)

func HandleGetPermissionConfig(c *gin.Context) {
	urlParam := c.Query("url")
	u, err := url.Parse(urlParam)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}

	var redirectUrl string
	for k, v := range query {
		if k == "redirectUrl" && len(v) > 0 {
			redirectUrl = fmt.Sprintf("%s/redirect?url=%s", config.Domain, v[0])
			break
		}
	}

	if redirectUrl == "" {
		c.JSON(200, gin.H{
			"code":    500,
			"message": "redirectUrl can not be empty",
		})
		return
	}

	reply, err := permission.GetPermissionConfig(urlParam)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data": gin.H{
			"corpid":      config.AppID,
			"agentid":     config.AgentID,
			"timestamp":   reply.NowTime,
			"nonceStr":    reply.NonceStr,
			"signature":   reply.Signature,
			"redirectUrl": redirectUrl,
		},
	})
}
