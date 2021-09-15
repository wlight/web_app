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
	Data interface{} `json:"data"`
}

// ResponseError 返回错误
func ResponseError(c *gin.Context, code ResCode) {
	responseData := &ResponseData{
		Code: code,
		Msg:  code.getMsg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, &responseData)
}

// ResponseError 返回错误
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	responseData := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, &responseData)
}

// ResponseSuccessWithMsg 返回成功携带自定义信息
func ResponseSuccessWithMsg(c *gin.Context, msg interface{}, data interface{}) {
	responseData := &ResponseData{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, &responseData)
}

// ResponseSuccess 返回成功
func ResponseSuccess(c *gin.Context, data interface{}) {
	responseData := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.getMsg(),
		Data: data,
	}
	c.JSON(http.StatusOK, &responseData)
}
