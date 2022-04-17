package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

//CommunityHandler 获取分类列表
func CommunityHandler(c *gin.Context) {
	// 获取类型列表
	list, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, list)
	return
}

// CommunityDetailHandler 获取分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取类型id
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// logic 获取分类详情
	community, err := logic.GetCommunityDetail(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, community)

	return

}
