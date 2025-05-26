package response

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
