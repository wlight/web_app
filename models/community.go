package models

type Community struct {
	CommunityId   int    `json:"id" db:"community_id"`
	CommunityName string `json:"name" db:"community_name"`
	Introduction  string `json:"introduction" db:"introduction"`
}
