package redis

// 程序运行中不变的key设置为常量
// redis key 使用命名空间，方便查询和拆分

const (
	KeyPostTimeZSet  = "webapp:post:time"   // 贴子按照发布时间排序的队列
	KeyPostScoreZSet = "webapp:post:score"  // 贴子按照投票分数排序的队列
	KeyPostVotedPre  = "webapp:post:voted:" // 贴子投票记录前缀，参数是post_id
)
