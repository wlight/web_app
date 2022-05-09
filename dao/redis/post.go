package redis

import "web_app/models"

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	//order := p.Order
	// 1、根据用户请求中的order参数来确定要查询的redis key
	var key string
	switch p.Order {
	case models.OrderTime:
		key = KeyPostTimeZSet
	case models.OrderScore:
		key = KeyPostScoreZSet
	}
	// 2、确定查询的索引起始点
	start := (p.PageIndex - 1) * p.PageSize
	stop := p.PageIndex*p.PageSize - 1

	// 3、查询
	return rdb.ZRevRange(key, start, stop).Result()
}
