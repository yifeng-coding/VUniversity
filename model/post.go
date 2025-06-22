package model

// PostStatus 帖子状态
type PostStatus int

const (
	PostStatusNormal  PostStatus = 0 // 正常
	PostStatusTop     PostStatus = 1 // 置顶
	PostStatusDeleted PostStatus = 2 // 删除
)

// PostData 帖子数据
type PostData struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	CreateTimestamp string `json:"create_timestamp"` // 创建时间戳（秒级）
	UpdateTimestamp string `json:"update_timestamp"` // 更新时间戳（秒级）
	UserID          int    `json:"user_id"`
	Username        string `json:"username"`
	Avatar          string `json:"avatar"`
}

// GetPostListRequest 分页加载帖子请求
type GetPostListRequest struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"required,gt=0"`          // 页码
	PageSize int `form:"page_size" json:"page_size" binding:"required,gt=0,lte=50"` // 每页数量
	UserID   int `form:"user_id" json:"user_id"`                                    // 用户ID
}

// GetPostListData 分页加载帖子响应数据
type GetPostListData struct {
	PostData   []*PostData `json:"post_list"`
	TotalCount int         `json:"total_count"`
}

// CreatePostRequest 创建帖子请求
type CreatePostRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

// UpdatePostRequest 更新帖子请求
type UpdatePostRequest struct {
	PostID  int    `form:"post_id" json:"post_id" binding:"required"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

// DeletePostRequest 删除帖子请求
type DeletePostRequest struct {
	PostID int `form:"post_id" json:"post_id" binding:"required,gt=1"`
}
