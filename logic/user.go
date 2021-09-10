package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)


func SignUp(p *models.ParamSignUp) {
	// 查找是否存在
	mysql.FindUserByUsername()
	//  生成uid
	snowflake.GenerateId()
	// 存入数据库
	mysql.InsertUser()
}
