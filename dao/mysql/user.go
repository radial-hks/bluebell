package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
)

const secret = "radial.cool"

// 把每一步数据操作封装成函数
// 待logic 层根据业务需求调用

// CheckUserExist 查看用户是否存在（检查制定用户名的用户是否存在）
func CheckUserExist(name string) (b bool, err error) {
	sqlStr := `select count(user_id) from user where username = ?;`
	var count int
	if err = DB.Get(&count, sqlStr, name); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(u *models.User) (err error) {
	//  不能直接存储明文密码 需要加密
	u.Password = encryptPassword(u.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = DB.Exec(sqlStr, u.UserID, u.UserName, u.Password)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 户登录时上传的密码
	sqlStr := `select user_id,username,password from user where username = ?`
	err = DB.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		// 数据库查询失败
		fmt.Println(err)
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}

func GetUserIDByID(uid int64) (user *models.User, err error) {
	//
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = DB.Get(user, sqlStr, uid)
	return
}
