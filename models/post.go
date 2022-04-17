package models

import "time"

// 内存对齐，尽量把相同类型的字段放到一起
type Post struct {
	Id          int64     `json:"id" db:"id"`
	PostId      int64     `json:"post_id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id"`
	Status      int8      `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
