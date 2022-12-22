package models

// Post model
type Post struct {
	PostID      int64  `db:"post_id" json:"post_id,string"`
	UserID      int64  `db:"user_id" json:"user_id,string"`
	CommunityID int64  `db:"community_id" json:"community_id,string"`
	Title       string `db:"title" json:"title"`
	Content     string `db:"content" json:"content"`
	CreatedAt   string `db:"created_at" json:"created_at"`
}

type PostWithOptions struct {
	AgreeVoteNum string `json:"agree_vote_num"`
	*Post
	*User      `json:"user"`
	*Community `json:"community"`
}
