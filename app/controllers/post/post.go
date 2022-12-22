package post

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/controllers"
	"go_code/gintest/app/dao"
	servicePost "go_code/gintest/app/services/post"
	"go_code/gintest/app/validators"
	validatorPost "go_code/gintest/app/validators/post"
	"go_code/gintest/bootstrap/glog"
	"strconv"
)

// StorePost 新增帖子
func StorePost(c *gin.Context) {
	request := new(validatorPost.StorePostRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		s := validators.Error(err, request)
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, s)
		return
	}

	userID, res := controllers.GetLoggedUserID(c)
	if res == false {
		controllers.ErrorRes(c, controllers.CodeUnauthorized, nil)
		return
	}
	request.UserID = userID.(int64)
	err := servicePost.StorePost(c, request)
	if err != nil {
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	controllers.SuccessRes(c, nil)
	return
}

// ShowPost 获取一个帖子
func ShowPost(c *gin.Context) {
	idStr := c.Query("post_id")
	if idStr == "" {
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, nil)
		return
	}

	postID, _ := strconv.ParseInt(idStr, 10, 64)
	post, err := servicePost.GetPostByID(postID)
	if err == dao.DataNotExists {
		controllers.ErrorRes(c, controllers.CodeDataNotExists, nil)
		return
	}
	if err != nil {
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}
	controllers.SuccessRes(c, &controllers.Data{Data: post})
	return
}

// OrderedPosts 排序帖子列表
func OrderedPosts(c *gin.Context) {
	request := new(validatorPost.OrderedPostsRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, err.Error())
		return
	}

	pagination, err := servicePost.GetOrderedPosts(c, request)
	if err == dao.DataNotExists {
		controllers.SuccessRes(c, &controllers.Data{Data: nil})
		return
	}
	if err != nil {
		glog.SL.Error("查询排序列表 error:", err)
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	controllers.SuccessRes(c, &controllers.Data{
		Data: pagination.Data,
		Meta: pagination.Meta,
	})
	return
}

// OrderedPostsForCommunity 社区下的帖子列表
func OrderedPostsForCommunity(c *gin.Context) {
	request := new(validatorPost.OrderedPostsForCommunityRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, err.Error())
		return
	}

	pagination ,err := servicePost.GetOrderedPostsForCommunity(c, request)
	if err == dao.DataNotExists {
		controllers.SuccessRes(c, &controllers.Data{Data: nil})
		return
	}

	if err != nil {
		glog.SL.Error("查询社区下的帖子排序列表 error:", err)
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	controllers.SuccessRes(c, &controllers.Data{
		Data: pagination.Data,
		Meta: pagination.Meta,
	})
	return
}
