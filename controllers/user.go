package controllers

import (
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	// 1、获取参数和校验
	var p = new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//// 非validator.ValidationErrors类型错误直接返回
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}

		// validator.ValidationErrors类型错误则进行翻译
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)),
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))

		return
	}
	// 2、登录逻辑
	err := logic.Login(p)
	if err != nil {
		zap.L().Error("登录失败", zap.String("username", p.Username), zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或密码错误",
		//})
		ResponseSuccessWithMsg(c, "用户名或密码错误", nil)
		return
	}
	// 3、返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登录成功",
	//})
	ResponseSuccessWithMsg(c, "登录成功", nil)
}

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	// 1、获取参数和校验
	var p = new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 手动校验参数不能为空
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SignUpParam with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg":"请求参数有误",
	// 	})
	// 	return
	// }

	// 2、注册逻辑
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3、返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}
