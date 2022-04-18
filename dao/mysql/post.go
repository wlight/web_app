package mysql

import (
	"database/sql"
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
