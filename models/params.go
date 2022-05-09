package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamCreatePost 创建贴子参数
type ParamCreatePost struct {
	PostId      int64  `json:"post_id"`
	AuthorId    int64  `json:"author_id"`
	CommunityId int64  `json:"community_id" binding:"required"`
	Status      int8   `json:"status"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`
}

// ParamVoteData 贴子投票参数
type ParamVoteData struct {
	PostId    int64 `json:"post_id,string" binding:"required"`       // 贴子id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成(1)反对(-1)取消投票(0)
}

// ParamPostList 获取贴子列表参数
type ParamPostList struct {
	PageIndex int64  `json:"pageIndex" form:"pageIndex"`
	PageSize  int64  `json:"pageSize" form:"pageSize"`
	Order     string `form:"order"`
}
