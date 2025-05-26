package request

type GetPostListRequest struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"required,gt=0"`          // 页码
	PageSize int `form:"page_size" json:"page_size" binding:"required,gt=0,lte=50"` // 每页数量
}
