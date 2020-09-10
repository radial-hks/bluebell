package controllers

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//  服务入口 负责处理路由、参数校验、 请求转发

// 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1、获取参数与 参数校验
	//var key models.ParamSignUp
	// 传入的值比较大 使用值传递
	key := new(models.ParamSignUp)
	// fmt.Println(key)
	if err := c.ShouldBindJSON(key); err != nil {
		zap.L().Error("SignUp with invalid ...", zap.Error(err))
		// 判断是不是 validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg_zh": "请求参数有误",
		//	"msg":    removeTopStruct(errs.Translate(trans)),
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return

		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "请求参数有误",
		//	"err": err.Error(),
		//})
		//return

	}

	//  手动进行请求参数进行详细的业务规则校验
	//if len(key.Username) == 0 || len(key.PassWord) == 0 || len(key.RePassWord) == 0 || key.PassWord != key.RePassWord {
	//	zap.L().Error("请求参数不完整...")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "注册信息不完整",
	//	})
	//	return
	//}

	// fmt.Println(key)
	// 2、业务处理
	if err := logic.SignUp(key); err != nil {
		zap.L().Error("SignUp with invalid ...", zap.Error(err))
		// 判断是不是 validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg_zh": "failed ",
		//	"msg":    removeTopStruct(errs.Translate(trans)),
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 3、返回响应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "ok",
	//})
	ResponseSuccess(c, CodeSuccess)
}

//  LoginHandler  请求登陆
func LoginHandler(c *gin.Context) {
	// 1、获取参数及参数校验
	var info models.ParamLogin
	// 解析请求的参数
	if err := c.ShouldBindJSON(&info); err != nil {
		zap.L().Error("Login with invalid ...", zap.Error(err))
		// 判断是不是 validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			ResponseError(c, CodeInvalidParam)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"msg_zh": "登陆信息出现错误",
		//	"msg":    removeTopStruct(errs.Translate(trans)),
		//})
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2、与数据库中的数据进行比对
	token, err := logic.Login(&info)
	if err != nil {
		zap.L().Error("Login with invalid ...", zap.String("username", info.Username), zap.Error(err))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": "用户名或密码错误",
		//})
		ResponseErrorWithMsg(c, CodeUserNotExist, "用户名或密码错误")
		return
	}
	// 3、返回登陆凭证 （验证通过）

	// 4、将数据写入 redis 数据库中 （维持登陆状态）

	// 2、进行业务逻辑判断
	// 3、返回相应
	//c.JSON(http.StatusOK, gin.H{
	//	"msg_zh": "Login Success",
	//})
	ResponseSuccess(c, token)
	//ResponseError(c,CodeSuccess,token)

}
