package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

func CommunityHandler(c *gin.Context) {
	// 获取类型列表
	list, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, list)
}
