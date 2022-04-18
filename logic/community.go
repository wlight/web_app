package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 获取类型列表
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.Community, error) {
	return mysql.GetCommunityById(id)
}
