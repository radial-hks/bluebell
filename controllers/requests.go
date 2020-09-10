package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("User not Login")

const ContextUserIDKey = "userID"

func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	fmt.Println(uid, ok)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	fmt.Println(userID, ok)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
