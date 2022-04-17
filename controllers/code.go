package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 10000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeServerErr

	CodeNeedLogin
	CodeInvalidToken
)

var codeMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "系统繁忙",
	CodeServerErr:       "系统错误，请重试",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
}

func (r ResCode) getMsg() string {
	s, ok := codeMap[r]
	if !ok {
		return codeMap[CodeServerBusy]
	}
	return s
}
