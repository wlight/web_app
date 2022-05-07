package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1000， // 错误码
	"msg": "xxx", // 提示信息
	"data": {}    // 返回的数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 返回错误
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.getMsg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 返回错误携带自定义信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccessWithMsg 返回成功携带自定义信息
func ResponseSuccessWithMsg(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.getMsg(),
		Data: data,
	})
}
