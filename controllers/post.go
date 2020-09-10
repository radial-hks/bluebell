package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//get the flag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		// zap.L().Error("ShouldBindJSON Failed", zap.Error(err))
		fmt.Println(p)
		ResponseError(c, CodeInvalidParam)
		return
	}
	// ResponseSuccess(c, "hahh")
	// get user id
	userid, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userid
	zap.L().Debug("GetCurrentUser", zap.Any("userid", userid))
	// createh acticle
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("CreatePost Failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "ok")
}

func GetPostHandler(c *gin.Context) {
	// get the flag : post_id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt Failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//
	data, err := logic.GetPostHandler(pid)
	if err != nil {
		zap.L().Error("logic.GetPostHandler Failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	// get data
	data, err := logic.GetPostList()
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// return
}
