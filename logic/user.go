package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 查找是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//  生成uid
	userId := snowflake.GenerateId()
	// 构造一个User 实例

	u := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 存入数据库
	err = mysql.InsertUser(u)

	return
}
