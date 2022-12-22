package post

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/dao"
	"go_code/gintest/app/dao/gredis"
	daoPost "go_code/gintest/app/dao/post"
	"go_code/gintest/app/models"
	"go_code/gintest/app/services"
	serviceUser "go_code/gintest/app/services/user"
	validatorPost "go_code/gintest/app/validators/post"
	"go_code/gintest/pkg/snowflake"
	"math"
	"strconv"
)

const (
	DefaultPageSize       = 20     // DefaultPageSize 默认每页条数
	DefaultOrderKey       = "id"   // DefaultOrderByKey 默认排序 key
	DefaultOrderDirection = "DESC" // DefaultOrderByDirection 默认排序方向
)

// StorePost 新增帖子
func StorePost(c *gin.Context, post *validatorPost.StorePostRequest) error {
	postModel := &models.Post{
		PostID:      snowflake.GenID(),
		UserID:      post.UserID,
		CommunityID: post.CommunityID,
		Title:       post.Title,
		Content:     post.Content,
	}

	err := daoPost.StorePost(postModel)
	if err != nil {
		return err
	}
	gredis.ZStorePostToRedis(c, postModel.PostID, postModel.CommunityID)

	return nil
}

// GetPostByID 查询指定帖子
func GetPostByID(postID int64) (*models.Post, error) {
	row, err := daoPost.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, dao.DataNotExists
	}
	return row, nil
}

// GetOrderedPosts 帖子排序列表
func GetOrderedPosts(
	c *gin.Context,
	request *validatorPost.OrderedPostsRequest,
) (
	*services.Pagination,
	error,
) {
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = DefaultPageSize
	}
	if request.OrderKey == "" {
		request.OrderKey = DefaultOrderKey
	}
	if request.OrderDirection == "" {
		request.OrderDirection = DefaultOrderDirection
	}

	total := daoPost.CountPosts(request)
	posts, err := daoPost.GetOrderedPosts(request)
	pagination := &services.Pagination{
		Meta: &services.Meta{
			Total: total,
			PerPage: request.PageSize,
			CurrentPage: request.Page,
			LastPage: int(math.Ceil(float64(total) / float64(request.PageSize))),
		},
		Data: nil,
	}

	if err != nil {
		return  pagination, err
	}
	if len(posts) <= 0 {
		return  pagination, dao.DataNotExists
	}

	var data []*models.PostWithOptions
	for _, post := range posts {
		item := &models.PostWithOptions{
			Post: post,
		}
		// 获取用户信息
		item.User, _ = serviceUser.GetUserByID(post.UserID)
		item.Community, _ = services.GetCommunityByID(post.CommunityID)
		agreeVoteNum, _ := gredis.GetPostVotes(c, strconv.FormatInt(post.PostID, 10))
		item.AgreeVoteNum = strconv.FormatInt(agreeVoteNum, 10)
		data = append(data, item)
	}
	pagination.Data = data

	return pagination, nil
}

// GetOrderedPostsForCommunity 帖子排序列表
func GetOrderedPostsForCommunity(
	c *gin.Context,
	request *validatorPost.OrderedPostsForCommunityRequest,
) (
	*services.Pagination,
	error,
) {
	if request.Page == 0 {
		request.Page = 1
	}
	if request.PageSize == 0 {
		request.PageSize = DefaultPageSize
	}
	if request.OrderKey == "" {
		request.OrderKey = DefaultOrderKey
	}
	if request.OrderDirection == "" {
		request.OrderDirection = DefaultOrderDirection
	}

	// total := daoPost.CountPosts(request)
	total := 0
	posts, err := daoPost.GetOrderedPostsForCommunity(request)
	pagination := &services.Pagination{
		Meta: &services.Meta{
			Total: total,
			PerPage: request.PageSize,
			CurrentPage: request.Page,
			LastPage: int(math.Ceil(float64(total) / float64(request.PageSize))),
		},
		Data: nil,
	}

	if err != nil {
		return  pagination, err
	}
	if len(posts) <= 0 {
		return  pagination, dao.DataNotExists
	}

	var data []*models.PostWithOptions
	for _, post := range posts {
		item := &models.PostWithOptions{
			Post: post,
		}
		// 获取用户信息
		item.User, _ = serviceUser.GetUserByID(post.UserID)
		item.Community, _ = services.GetCommunityByID(post.CommunityID)
		agreeVoteNum, _ := gredis.GetPostVotes(c, strconv.FormatInt(post.PostID, 10))
		item.AgreeVoteNum = strconv.FormatInt(agreeVoteNum, 10)
		data = append(data, item)
	}
	pagination.Data = data

	return pagination, nil
}