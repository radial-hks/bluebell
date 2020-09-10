package models

// 定义请求参数结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" form:"username" binding:"required"`
	PassWord   string `json:"password" form:"password" binding:"required"`
	RePassWord string `json:"re_password" form:"re_password" binding:"required,eqfield=PassWord"`
}

// ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	PassWord string `json:"password" form:"password" binding:"required"`
}
