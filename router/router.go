package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置为发布模式，启动时不再提示warning
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})

	// 注册路由
	r.POST("/signup", controllers.SignUpHandler)

	// 登录
	r.POST("/login", controllers.LoginHandler)

	return r
}
