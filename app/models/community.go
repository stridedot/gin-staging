package models

type Community struct {
	CommunityId   int64  `db:"community_id" json:"community_id,string"`
	CommunityName string `db:"community_name" json:"community_name"`
	Introduction  string `db:"introduction" json:"introduction"`
}
