package post

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/controllers"
	"go_code/gintest/app/dao"
	servicePost "go_code/gintest/app/services/post"
	"go_code/gintest/app/validators"
	validatorPost "go_code/gintest/app/validators/post"
)

// VotePost 投票
func VotePost(c *gin.Context)  {
	request := new(validatorPost.VotePostRequest)
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
	err := servicePost.VotePost(c, request)
	if err == dao.VoteTimeExpired {
		controllers.ErrorRes(c, controllers.CodeVoteTimeExpired, nil)
		return
	}
	if err == dao.VoteRepeated {
		controllers.ErrorRes(c, controllers.CodeVoteRepeated, nil)
		return
	}
	if err != nil {
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	controllers.SuccessRes(c, nil)
	return
}
