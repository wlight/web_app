package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 查找是否存在
	exist, err := mysql.CheckUserExist(p.Username)

	if err != nil {
		return err
	}

	if exist {
		return errors.New("用户已存在")
	}
	//  生成uid
	userId := snowflake.GenerateId()
	// 构造一个User 实例

	// 存入数据库
	//

	mysql.InsertUser()
}
