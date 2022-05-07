package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)

// PostVote 为贴子投票
func PostVote(userId int64, voteData *models.ParamVoteData) error {
	err := redis.PostVote(strconv.FormatInt(userId, 10), strconv.FormatInt(voteData.PostId, 10), float64(voteData.Direction))
	return err
}
