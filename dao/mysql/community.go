package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

func GetCommunityList() (CommunityList []*models.Community, err error) {
	sqlStr := "SELECT community_id,community_name FROM coumunity;"
	if err = DB.Select(&CommunityList, sqlStr); err != nil {
		zap.L().Warn("NO CommunityList in DB")
		err = nil
	}
	fmt.Println(CommunityList)
	return
}

// get community detail
func GetCommunityByID(id int64) (CommunityDetail *models.CommunityDetail, err error) {

	CommunityDetail = new(models.CommunityDetail)
	sqlStr := "SELECT community_id,community_name,introduction,create_time,udpate_time FROM coumunity WHERE  community_id = ?"
	err = DB.Get(CommunityDetail, sqlStr, id)
	if err == sql.ErrNoRows {
		//err = ErrorCommunityInvalidID
		err = nil
	}
	if err != nil {
		zap.L().Warn("GetCommunityByID Failed")
		err = nil
	}
	return
}
