package logic

import (
	"go.uber.org/zap"
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

func GetPostById(pid int64) (*models.PostDetail, error) {
	// 获取帖子信息
	post, err := mysql.GetPostById(pid)

	if err != nil {
		zap.L().Error("mysql.GetPostById(id) failed.", zap.Int64("post_id", pid))
		return nil, err
	}
	// 获取用户信息
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.getUserById(post.AuthorId) failed.",
			zap.Int64("userid", post.AuthorId),
			zap.Error(err))
		return nil, err
	}

	// 获取分类详情
	community, err := mysql.GetCommunityById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityById(post.CommunityId) failed",
			zap.Int64("post_id", post.CommunityId),
			zap.Error(err))
		return nil, err
	}
	postDetail := &models.PostDetail{
		User:      user,
		Post:      post,
		Community: community,
	}

	return postDetail, err
}
