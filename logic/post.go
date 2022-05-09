package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// CreatePost 创建贴子
func CreatePost(post *models.ParamCreatePost) (err error) {
	// 生成帖子id
	post.PostId = snowflake.GenerateId()

	// mysql 处理帖子存储
	err = mysql.CreatePost(post)
	if err != nil {
		return err
	}
	err = redis.CreatePost(post.PostId)
	if err != nil {
		return err
	}
	return
}

// GetPostById 根据id获取贴子详情
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

// GetPostList 获取帖子列表
func GetPostList(pageIndex, pageSize int64) (data []*models.PostDetail, err error) {
	postList, err := mysql.GetPostList(pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	var posts = make([]*models.PostDetail, 0, len(postList))
	for _, post := range postList {
		// 获取用户信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.getUserById(post.AuthorId) failed.",
				zap.Int64("userid", post.AuthorId),
				zap.Error(err))
			continue
		}

		// 获取分类详情
		community, err := mysql.GetCommunityById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityId) failed",
				zap.Int64("post_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			User:      user,
			Post:      post,
			Community: community,
		}
		posts = append(posts, postDetail)
	}
	return posts, err
}

// GetPostList2 获取帖子列表
func GetPostList2(p *models.ParamPostList) (data []*models.PostDetail, err error) {
	// 1、redis 中获取贴子的id列表
	ids, err := redis.GetPostIdsInOrder(p)
	if len(ids) == 0 || err != nil {
		return
	}
	// 2、mysql 中查询贴子的详情
	postList, err := mysql.GetPostListByIds(ids)

	var posts = make([]*models.PostDetail, 0, len(ids))
	for _, post := range postList {
		// 获取用户信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.getUserById(post.AuthorId) failed.",
				zap.Int64("userid", post.AuthorId),
				zap.Error(err))
			continue
		}

		// 获取分类详情
		community, err := mysql.GetCommunityById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityId) failed",
				zap.Int64("post_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			User:      user,
			Post:      post,
			Community: community,
		}
		posts = append(posts, postDetail)
	}

	return posts, err
}
