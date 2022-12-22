package post

import (
	"database/sql"
	"go_code/gintest/app/models"
	validatorPost "go_code/gintest/app/validators/post"
	"go_code/gintest/bootstrap/gdb"
	"strconv"
	"strings"
	"time"
)

// StorePost 新增帖子
func StorePost(post *models.Post) error {
	sqlStr := "INSERT INTO post" +
		" (post_id, user_id, community_id, title, content, created_at)" +
		" VALUES(?, ?, ?, ?, ?, ?)"
	_, err := gdb.DB.Exec(
		sqlStr,
		post.PostID,
		post.UserID,
		post.CommunityID,
		post.Title,
		post.Content,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return err
}

// GetPostByID 查询指定的帖子
func GetPostByID(postID int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := "SELECT post_id, user_id, community_id, title, content, created_at" +
		" FROM post WHERE post_id = ?"

	err = gdb.DB.Get(post, sqlStr, postID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

// GetOrderedPosts 帖子排序列表
func GetOrderedPosts(
	request *validatorPost.OrderedPostsRequest,
) (
	[]*models.Post,
	error,
) {
	page := request.Page
	pageSize := request.PageSize
	offset := (page - 1) * pageSize
	var posts []*models.Post
	var err error
	var sqlStr string

	sqlStr = "SELECT post_id, user_id, community_id, title, content, created_at" +
		" FROM post :s ORDER BY ?, ? LIMIT ?, ?"
	if request.CommunityID != 0 {
		s := "WHERE community_id = " + strconv.FormatInt(request.CommunityID, 10)
		sqlStr = strings.Replace(sqlStr, ":s", s, 1)
	} else {
		sqlStr = strings.Replace(sqlStr, ":s", "", 1)
	}
	err = gdb.DB.Select(
		&posts,
		sqlStr,
		request.OrderKey,
		request.OrderDirection,
		offset,
		pageSize,
	)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// CountPosts 查询帖子总条数
func CountPosts(request *validatorPost.OrderedPostsRequest) int {
	var (
		total int
		sqlStr string
		err error
	)
	sqlStr = "SELECT COUNT(*)"+" FROM post :s"
	if request.CommunityID != 0 {
		s := "WHERE community_id = " + strconv.FormatInt(request.CommunityID, 10)
		sqlStr = strings.Replace(sqlStr, ":s", s, 1)
	} else {
		sqlStr = strings.Replace(sqlStr, ":s", "", 1)
	}

	err = gdb.DB.QueryRow(sqlStr).Scan(&total)
	if err != nil {
		return 0
	}
	return total
}

// GetOrderedPostsForCommunity 社区下的帖子排序列表
func GetOrderedPostsForCommunity(
	request *validatorPost.OrderedPostsForCommunityRequest,
) (
	[]*models.Post,
	error,
) {
	var posts []*models.Post
	sqlStr := "SELECT post_id, user_id, community_id, title, content, created_at" +
		" FROM post WHERE community_id = ? ORDER BY ?, ? LIMIT ?, ?"

	page := request.Page
	pageSize := request.PageSize
	offset := (page - 1) * pageSize

	err := gdb.DB.Select(
		&posts,
		sqlStr,
		request.CommunityID,
		request.OrderKey,
		request.OrderDirection,
		offset,
		pageSize,
	)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
