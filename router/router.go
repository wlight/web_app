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
	r.GET("/ping", func(c *gin.Context) {
		// 如果是登录的用户，判断请求头中是否有效的JWT

		c.JSON(http.StatusOK, gin.H{
			"msg": "pang",
		})
	})

	// 注册路由
	r.POST("/signup", controllers.SignUpHandler)

	// 登录
	r.POST("/login", controllers.LoginHandler)
	// 刷新access token
	//r.POST("/refreshAccessToken", controllers.RefreshAccessTokenHandler)

	return r
}
