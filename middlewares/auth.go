package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web_app/controllers"
	"web_app/pkg/jwt"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	/*
		GET /ping HTTP/1.1
		Host: 127.0.0.1:8081
		Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Ijk5OSIsInVzZXJfaWQiOjQ0OTgxNzM2ODM2MTczODI0LCJleHAiOjE2NDI5MjQ0NTIsImlzcyI6IndlYl9hcHAifQ.bfYureBxGS5PnvB1mlRyiL7-5grfPi3HTEyZN-0OD7U
		cache-control: no-cache
		Postman-Token: 34d15b4e-ec69-4230-b40a-339b479993b3
	*/
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2003,
			//	"msg":  "请求头中的auth为空",
			//})
			c.Abort()
			return
		}
		// 按照空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2004,
			//	"msg":  "请求头中的auth格式有误",
			//})
			c.Abort()
			return
		}
		// parts[1] 是获取的tokenString， 我们使用之前定义好的解析JWT的函数来解析它
		myClaims, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			//c.JSON(http.StatusOK, gin.H{
			//	"code": 2004,
			//	"msg":  "无效的token",
			//})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.CtxUsernameKey, myClaims.Username)
		c.Set(controllers.CtxUserIdKey, myClaims.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
