package controllers

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/dao"
	serviceCommunity "go_code/gintest/app/services"
	"go_code/gintest/bootstrap/glog"
	"strconv"
)

// Communities 社区分类列表
func Communities(c *gin.Context) {
	communities, err := serviceCommunity.GetCommunities()
	if err == dao.DataNotExists {
		SuccessRes(c, &Data{Data: nil})
		return
	}
	if err != nil {
		glog.SL.Error("查询社区列表错误", err)
		ErrorRes(c, CodeInternalError, nil)
		return
	}

	SuccessRes(c, &Data{Data: communities})
	return
}

// CommunityDetail 社区分类详情
func CommunityDetail(c *gin.Context) {
	idStr := c.Query("community_id")
	if idStr == "" {
		ErrorRes(c, CodeUnexpectedParam, nil)
		return
	}
	ID, _ := strconv.ParseInt(idStr, 10, 64)
	community, err := serviceCommunity.GetCommunityByID(ID)
	if err == dao.DataNotExists {
		ErrorRes(c, CodeDataNotExists, nil)
		return
	}
	if err != nil {
		ErrorRes(c, CodeInternalError, nil)
		return
	}
	SuccessRes(c, &Data{Data: community})
	return
}
