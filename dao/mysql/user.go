package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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
		return ErrorUserExist
	}

	return
}

// InsertUser 插入数据库
func InsertUser(user *models.User) (err error) {
	// 密码加密
	password := encryptPassword(user.Password)
	// 执行插入语句
	sqlStr := "insert into user (user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user, user.Username, password)
	return
}

func Login(user *models.User) (err error) {
	password := encryptPassword(user.Password)

	sqlStr := "select user_id, password from user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if password != user.Password {
		return ErrorInvaildPassword
	}
	return
}

func encryptPassword(oPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(oPassword)))
}

// GetUserById 获取用户详情
func GetUserById(uid int64) (*models.User, error) {
	user := new(models.User)
	sqlStr := "select user_id, username from user where user_id = ?"
	err := db.Get(user, sqlStr, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorUserNotExist
		}
	}
	return user, err
}
