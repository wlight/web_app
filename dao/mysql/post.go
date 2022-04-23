package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func CreatePost(post *models.ParamCreatePost) (err error) {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return
}

func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)

	sqlStr := `select post_id,title, content, author_id, community_id, create_time
	from post 
	where post_id = ?
	`
	err = db.Get(post, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvaildId
		}
	}

	return post, err
}

// GetPostList 获取帖子列表
func GetPostList(pageIndex, pageSize int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post 
	limit ?, ?`

	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, (pageIndex-1)*pageSize, pageSize)
	if err != nil {
		zap.L().Error("db.Select failed", zap.Error(err))
		if err == sql.ErrNoRows {
			err = ErrorInvaildId
		}
	}
	return
}
