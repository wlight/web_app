package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(post *models.ParamCreatePost) (err error) {
	// 生成帖子id
	post.PostId = snowflake.GenerateId()

	// mysql 处理帖子存储
	return mysql.CreatePost(post)
}
