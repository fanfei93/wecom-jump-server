package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wecom-jump-server/handler/group"
	"wecom-jump-server/handler/health"
	"wecom-jump-server/handler/permission"
	"wecom-jump-server/handler/redirect"
)

const (
	VerifyURL   = "/xxxxxxxxxxxxxxx.txt"
	VerifyValue = "xxxxxxx"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.Static("static", "static")
	r.LoadHTMLFiles("static/html/index.html")

	r.GET("/health", health.HandleCheckHealth)
	r.POST("/group", group.HandleCreateGroup)
	// 企业微信可信域名验证
	//r.GET(VerifyURL, func(c *gin.Context) {
	//	c.String(200, VerifyValue)
	//})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/redirect", redirect.HandleRedirect)
	r.GET("/getWeiXinPermissionsValidationConfig", permission.HandleGetPermissionConfig)

	return r
}
