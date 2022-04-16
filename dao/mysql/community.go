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
