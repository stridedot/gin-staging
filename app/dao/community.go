package dao

import (
	"database/sql"
	"go_code/gintest/app/models"
	"go_code/gintest/bootstrap/gdb"
	"strconv"
)

// GetCommunities 社区列表
func GetCommunities() ([]*models.Community, error) {
	communities := new([]*models.Community)
	sqlStr := "SELECT community_id, community_name"+" FROM community"
	err := gdb.DB.Select(communities, sqlStr)

	if err != nil {
		return nil, err
	}
	if len(*communities) <= 0 {
		return nil, DataNotExists
	}

	return *communities, nil
}

// GetCommunityByID 查询社区分类详情
func GetCommunityByID(ID int64) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := "SELECT community_id, community_name, introduction"+
		" FROM community WHERE community_id = ?"
	err = gdb.DB.Get(community, sqlStr, strconv.FormatInt(ID, 10))

	if err == sql.ErrNoRows {
		return nil, DataNotExists
	}
	return community, err
}
