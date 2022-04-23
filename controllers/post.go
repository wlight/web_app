package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {
	// 获取参数并校验
	p := new(models.ParamCreatePost)
	if err := c.ShouldBindJSON(p); err != nil {
		// 处理错误
		zap.L().Error("CreatePostHandler with invalid param", zap.Error(err))
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
	// 获取当前登录用户id
	userId, err := getUserId(c)
	if err != nil {
		ResponseError(c, CodeServerErr)
		return
	}
	// 创建帖子 logic
	p.AuthorId = userId
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic createPost err", zap.Error(err))
		ResponseError(c, CodeServerErr)
		return
	}
	// 返回响应
	ResponseSuccess(c, CodeSuccess)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取帖子详情
	post, err := logic.GetPostById(id)
	if err != nil {
		zap.L().Error("logic getPostById failed", zap.Error(err))
		ResponseError(c, CodeServerErr)
		return
	}
	// 返回响应
	ResponseSuccess(c, post)
	return
}

// GetPsotListHandler 获取帖子列表处理函数
func GetPsotListHandler(c *gin.Context) {
	// 获取分页参数
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")

	var (
		pageIndex int64
		pageSize  int64
	)
	pageIndex, err := strconv.ParseInt(pageIndexStr, 10, 64)
	if err != nil {
		pageIndex = 1
	}
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 10
	}
	list, err := logic.GetPostList(pageIndex, pageSize)
	if err != nil {
		zap.L().Error("logic GetPostList failed", zap.Int64("pageIndex", pageIndex), zap.Int64("pageSize", pageSize), zap.Error(err))
		ResponseError(c, CodeServerErr)
		return
	}
	// 返回响应
	ResponseSuccess(c, list)

	return
}
