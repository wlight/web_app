package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (list []*models.Community, err error) {
	sqlStr := "select community_id, community_name, introduction from community"

	if err = db.Select(&list, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is a no rows error", zap.Error(err))
		}
		return nil, err
	}
	return list, nil
}

func GetCommunityById(id int64) (*models.Community, error) {
	detail := new(models.Community)
	sqlStr := "select community_id, community_name, introduction from community where community_id = ?"

	err := db.Get(detail, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvaildId
		}
	}
	return detail, err
}
