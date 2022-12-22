package post

// StorePostRequest 新增帖子 validator
type StorePostRequest struct {
	UserID      int64  `json:"user_id"`
	CommunityID int64  `json:"community_id" binding:"required" label:"社区ID"`
	Title       string `json:"title" binding:"required" label:"帖子标题"`
	Content     string `json:"content" binding:"required" label:"帖子内容"`
}

// OrderedPostsRequest 帖子排序 validator
type OrderedPostsRequest struct {
	Page           int    `json:"page"`
	PageSize       int    `json:"page_size"`
	CommunityID    int64  `json:"community_id,string"`
	OrderKey       string `json:"order_key"`
	OrderDirection string `json:"order_direction"`
}

// VotePostRequest 给帖子排序 validator
type VotePostRequest struct {
	UserID    int64  `json:"user_id"`
	Direction int    `json:"direction,string" binding:"oneof=-1 0 1"`
	PostID    string `json:"post_id" binding:"required"`
}

// OrderedPostsForCommunityRequest 社区下的帖子排序 validator
type OrderedPostsForCommunityRequest struct {
	Page           int    `json:"page,string"`
	PageSize       int    `json:"page_size,string"`
	CommunityID    int64  `json:"community_id,string"`
	OrderKey       string `json:"order_key"`
	OrderDirection string `json:"order_direction"`
}
