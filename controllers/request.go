package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	CtxUsernameKey = "username"
	CtxUserIdKey   = "userId"
)

var ErrUserNotLogin = errors.New("用户未登录")

func getUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIdKey)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	return
}
