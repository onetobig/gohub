package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetUpRoute(router *gin.Engine) {
	// 全局中间件
	registerGlobalMiddleware(router)
	// API 路由
	routes.RegisterAPIRoutes(router)
	// 404 路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusOK, "页面返回 4040")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"error_code": 404,
				"error_msg":  "路由未注册，请检查 url 和 请求方法是否正确。",
			})
		}
	})
}
