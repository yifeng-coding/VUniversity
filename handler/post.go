package handler

import (
	"fmt"
	"github.com/SniperCoding/VUniversity/model/errs"
	"github.com/SniperCoding/VUniversity/model/request"
	"github.com/SniperCoding/VUniversity/model/response"
	"github.com/SniperCoding/VUniversity/services"
	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")
	fmt.Printf("Raw page_num: %s, page_size: %s\n", pageNumStr, pageSizeStr)
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
