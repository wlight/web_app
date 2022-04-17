package mysql

import "web_app/models"

func CreatePost(post *models.ParamCreatePost) (err error) {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return
}
