package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

var secret = "light"

func Login(user *models.User) (err error) {
	var oPassword = user.Password
	sqlStr := "select userid, username, password from users where username=?"
	err = db.Get(user, sqlStr, user.Password)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		// 数据库执行失败
		return err
	}

	if encryptPassword(oPassword) != user.Password {
		return errors.New("密码错误")
	}
	return
}

// FindUserByUsername 根据用户名查找用户
func FindUserByUsername(username string) (user models.User, err error) {
	sqlStr := "select userid, username, password from users where username=?"
	err = db.Get(&user, sqlStr, username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// CheckUserPassword 验证用户密码是否正确
func CheckUserPassword(userPassword string, originPassword string) bool {
	return encryptPassword(userPassword) == originPassword
}

// CheckUserExist 检查用户是否在用户表中存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(userid) from users where username = ?"
	//query, err := db.Query(sqlStr, username)
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertUser 向数据库插入一条新的用户数据
func InsertUser(user *models.User) (err error) {
	// 用户密码加密
	user.Password = encryptPassword(user.Password)
	// 执行插入
	sqlStr := "insert into users(userid, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, user.UserId, user.Username, user.Password)
	if err != nil {
		return err
	}

	return
}

func encryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(password)))
}
