package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "world",
			})
		})
	}
}
