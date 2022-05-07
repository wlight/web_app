package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

/*
投一票为贴子加 432 分 86400/200 -> 200张赞同票可以让贴子在首页一天
投票的几种情况
direction=1时，有两种情况
	1、之前没有投过票，现在投赞成票  +432
	2、之前投反对票，现在改投赞成票  +432*2
direction=0时，有两种情况
	1、之前投赞成票，现在取消投票	-432
	2、之前投反对票，现在取消投票	+432
direction=-1时，有两种情况
	1、之前没有投过票，现在投反对票	-432
	2、之前投赞成票，现在改投反对票	-432*2
投票的限制：
每个贴子投票时间只有一周
	1、到期之后将redis中保存的赞成和反对都保存到mysql中
	2、到期之后删除redis中的数据 KeyPostVotedZSetPre 加 post_id
*/

const (
	oneWeekInSecond = 24 * 3600 * 7
	scorePreVote    = 432 // 每一票的分值
)

// PostVote 贴子投票
func PostVote(userId, postId string, value float64) error {
	// 1、判断投票限制
	postTime := rdb.ZScore(KeyPostTimeZSet, postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return errors.New("投票时间已过")
	}
	// 2和3需要放到一个事务中
	pipeline := rdb.Pipeline()
	// 2、更新贴子的分数
	// 查询贴子的投票记录
	ov := rdb.ZScore(KeyPostVotedPre+postId, userId).Val()
	var op float64
	if value > ov { // 说明现在投的是赞成票
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline.ZIncrBy(KeyPostScoreZSet, op*diff*scorePreVote, postId)

	// 3、记录该用户的投票数据
	if value == 0 { // 如果是取消投票，则移除投票
		pipeline.ZRem(KeyPostVotedPre+postId, userId)
	} else {
		pipeline.ZAdd(KeyPostVotedPre+postId, redis.Z{
			Score:  value,
			Member: userId,
		})
	}
	_, err := pipeline.Exec()
	return err
}

// CreatePost 创建贴子redis
func CreatePost(postId int64) (err error) {
	// 事务操作
	pipeline := rdb.TxPipeline()

	// 贴子发布时间
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 贴子分数
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  0,
		Member: postId,
	})
	_, err = pipeline.Exec()
	if err != nil {
		return err
	}
	return
}
