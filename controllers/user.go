package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
)

// SignUpHandler 处理用户注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1、获取参数和校验
	var p = new(models.ParamSignUp) // 获取一个指针
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors，判断是否为validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}

		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2、注册逻辑
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3、返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理用户登录请求的函数
func LoginHandler(c *gin.Context) {
	// 1. 参数校验
	var p = new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errors.Translate(trans)))
		return
	}
	// 2. 登录逻辑
	aToken, rToken, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorInvaildPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	token := map[string]string{
		"aToken": aToken,
		"rToken": rToken,
	}

	// 3. 返回响应
	ResponseSuccess(c, token)
}
