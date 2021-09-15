package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// Login 用户登录逻辑
func Login(p *models.LoginParam) (err error) {
	//// 查找用户
	//user, err := mysql.FindUserByUsername(p.Username)
	//if err != nil {
	//	// 查询出错
	//	return errors.New("用户不存在")
	//}
	//
	//// 验证密码是否正确
	//if !mysql.CheckUserPassword(p.Password, user.Password) {
	//	return errors.New("用户密码不正确")
	//}
	var user = &models.User{
		Username: p.Username,
		Password: p.Username,
	}

	err = mysql.Login(user)
	if err != nil {
		return err
	}
	return
}

// SignUp 用户注册逻辑
func SignUp(p *models.SignUpParam) (err error) {
	// 查找是否存在
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		// 查询出错
		return err
	}

	if exist {
		// 用户已存在
		return errors.New("用户已存在")
	}
	//  生成uid
	userId := snowflake.GenerateId()

	// 构造一个User 实例
	u := models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 存入数据库

	err = mysql.InsertUser(&u)
	if err != nil {
		return err
	}

	return
}
