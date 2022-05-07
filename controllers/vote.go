package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		// 处理错误
		zap.L().Error("PostVoteHandler with invalid param", zap.Error(err))
		// 获取validator.ValidationErrors类型的errors，判断是否为validator.ValidationErrors类型
		errors, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errors.Translate(trans)))
		return
	}
	// 获取当前登录用户的id
	userId, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 投票逻辑处理
	err = logic.PostVote(userId, p)
	if err != nil {
		zap.L().Error("logic.PostVote failed", zap.Int64("userId", userId), zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerErr, err.Error())
		return
	}
	// 返回响应
	ResponseSuccess(c, nil)
	return
}
