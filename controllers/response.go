package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*

{
	"code":"10001", // 程序中的错误码
	"message": xx, // 提示信息
	"data":{},     // 数据
}


*/

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"mag"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	rd := &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess,
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, Msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  Msg,
		Data: nil,
	})
}
