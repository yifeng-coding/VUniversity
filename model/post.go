package model

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
}

// GetPostListData 分页加载帖子响应数据
type GetPostListData struct {
	PostData   []*PostData `json:"post_list"`
	TotalCount int         `json:"total_count"`
}
