package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const secret = "encrypt"

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	//query, err := db.Query(sqlStr, username)
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户已存在")
	}

	return
}

// InsertUser 插入数据库
func InsertUser(user *models.User) (err error) {
	// 密码加密
	password := encryptPassword(user.Password)
	// 执行插入语句
	sqlStr := "insert into user (user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, password)
	return
}

func CheckUserPassword(username string, oPassword string) (err error) {
	sqlStr := "select user_id, password from user where username = ?"
	var user models.User
	err = db.Get(&user, sqlStr, username)
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("用户密码错误")
	}
	return
}

func encryptPassword(oPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(oPassword)))
}
