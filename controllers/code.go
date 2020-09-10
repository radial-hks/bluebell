package controllers

type ResCode int64

const (
	CodeSuccess = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidAuth
	CodeNeedAuth
	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeServerBusy:      "服务忙",
	CodeNeedAuth:        "待认证",
	CodeNeedLogin:       "需要登陆",
	CodeInvalidAuth:     "无效认证",
	CodeInvalidToken:    "无效的TOKEN",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[c]
	}
	return msg
}

func (c ResCode) Success() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[c]
	}
	return msg
}
