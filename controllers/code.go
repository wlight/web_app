package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 10000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
)

var codeMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "系统繁忙",
}

func (r ResCode) getMsg() string {
	s, ok := codeMap[r]
	if !ok {
		return codeMap[CodeServerBusy]
	}
	return s
}
