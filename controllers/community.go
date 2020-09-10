package controllers

import (
	"bluebell/logic"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//-- 社关

func CommunityHandler(c *gin.Context) {
	// 查询到所有的ID, 及名称 一列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println(data)
	ResponseSuccess(c, data)
}

//
func CommunityDetailHandler(c *gin.Context) {
	// get community ID
	idStr := c.Param("id")
	// python int()
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("ParseInt Failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询到所有的ID, 及名称 一列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println(data)
	ResponseSuccess(c, data)
}
