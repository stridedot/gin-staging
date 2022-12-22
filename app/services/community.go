package services

import (
	daoCommunity "go_code/gintest/app/dao"
	"go_code/gintest/app/models"
)

// GetCommunities 查询社区列表
func GetCommunities() ([]*models.Community, error) {
	return daoCommunity.GetCommunities()
}

// GetCommunityByID 查询社区分类详情
func GetCommunityByID(ID int64) (*models.Community, error) {
	return daoCommunity.GetCommunityByID(ID)
}