package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controllers"
	"web_app/logger"
	"web_app/middlewares"
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

	v1 := r.Group("/api/v1")

	// 注册路由
	v1.POST("/signup", controllers.SignUpHandler)

	// 登录
	v1.POST("/login", controllers.LoginHandler)
	// 刷新access token
	//v1.POST("/refreshAccessToken", controllers.RefreshAccessTokenHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用jwt认证中间件

	{
		// 获取帖子的分类
		v1.GET("/community", controllers.CommunityHandler)
		// 获取帖子分类详情
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		// 创建帖子
		v1.POST("/post", controllers.CreatePostHandler)

		// 帖子详情
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		// 帖子列表
		v1.GET("/post/list", controllers.GetPostListHandler)
		// 帖子列表(按照时间或者分数查询)
		v1.GET("/post/list2", controllers.GetPostListHandler2)
		// 贴子投票
		v1.POST("/post/vote", controllers.PostVoteHandler)

	}

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
