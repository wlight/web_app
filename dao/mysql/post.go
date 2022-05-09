package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
	"web_app/models"
)

// CreatePost 创建贴子
func CreatePost(post *models.ParamCreatePost) (err error) {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return
}

// GetPostById 根据id获取贴子详情
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

// GetPostListByIds 根据ids获取贴子列表
func GetPostListByIds(ids []string) (posts []*models.Post, err error) {
	// 动态填充id
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time
				FROM post 
				WHERE post.post_id IN (?) 
				ORDER BY FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&posts, query, args...)
	return
}
