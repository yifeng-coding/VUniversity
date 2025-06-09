package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/service"
)

// @Summary		获取帖子列表
// @Description	获取帖子列表，支持分页
// @Tags			帖子
// @Accept			json
// @Produce		json
// @Param			page_num	query		int														true	"页码"	minimum(1)	default(1)
// @Param			page_size	query		int														true	"每页数量"	minimum(1)	maximum(50)	default(10)
// @Success		200			{object}	model.ResponseWrapper{code=int,data=[]model.PostData}	"code为0表示成功，否则为失败"
// @Router			/post/list [get]
func GetPostList(c *gin.Context) {
	req := &model.GetPostListRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := service.GetPostList(c.Request.Context(), req.PageNum, req.PageSize)
	model.Response(c, err, data)
	return
}
