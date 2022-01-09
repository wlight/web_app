package router

import (
	"net/http"
	"web_app/controllers"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})

	// 注册路由
	r.POST("/signup", controllers.SignUpHandler)

	return r
}
