package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yifeng-coding/VUniversity/model/errs"
	"github.com/yifeng-coding/VUniversity/model/request"
	"github.com/yifeng-coding/VUniversity/model/response"
	"github.com/yifeng-coding/VUniversity/services"
)

// @Summary		获取帖子列表
// @Description	获取帖子列表，支持分页
// @Tags		帖子
// @Accept		json
// @Produce		json
// @Param		page_num	query		int															true	"页码"	minimum(1)	default(1)
// @Param		page_size	query		int															true	"每页数量"	minimum(1)	maximum(50)	default(10)
// @Success		200			{object}	response.ResponseWrapper{code=int,data=[]response.PostData}	"code为0表示成功，否则为失败"
// @Router		/post/list [get]
func GetPostList(c *gin.Context) {
	req := &request.GetPostListRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		response.Fail(c, errs.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := services.GetPostList(req.PageNum, req.PageSize)
	response.Response(c, err, data)
	return
}
