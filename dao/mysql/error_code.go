package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvaildPassword = errors.New("密码错误")
	ErrorInvaildId       = errors.New("id 不正确")
)
