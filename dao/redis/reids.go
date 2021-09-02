package redis

import (
	"fmt"
	"web_app/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(conf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	_ = rdb.Close()
}
