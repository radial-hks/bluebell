package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/bwmarrin/snowflake"
)

// 存放业务逻辑代码
func SignUp(key *models.ParamSignUp) (err error) {
	// 1、判断用户存不存在
	exist, err := mysql.CheckUserExist(key.Username)
	if err != nil {
		// 数据库查询出错
		return err
	}
	if exist {
		// 用户已存在
		return errors.New("Name had Exist")
	}
	// 2、生成 user_id
	userID := snowflake.GenID()
	// 构造 user 实例
	u := models.User{
		UserID:   userID,
		UserName: key.Username,
		Password: key.PassWord,
	}
	// 3、保存进数据库
	return mysql.InsertUser(&u)
}

//func Login(info *models.ParamLogin) (err error) {
//	user := &models.User{
//		UserName: info.Username,
//		Password: info.PassWord,
//	}
//	//if err := mysql.Login(user); err != nil {
//	//	return errors.New("Login Failed")
//	//}
//
//	return mysql.Login(user)
//
//}

func Login(info *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: info.Username,
		Password: info.PassWord,
	}
	if err := mysql.Login(user); err != nil {
		return "", errors.New("Login Failed")
	}
	return jwt.GenToken(user.UserID, user.UserName)

}
