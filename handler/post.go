package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
// @Param			user_id	query		int															false	"用户ID"	minimum(1)
// @Success		200			{object}	model.ResponseWrapper{code=int,data=[]model.PostData}	"code为0表示成功，否则为失败"
// @Router			/post/list [get]
func GetPostList(c *gin.Context) {
	req := &model.GetPostListRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := service.GetPostList(c.Request.Context(), req.PageNum, req.PageSize, req.UserID)
	model.Response(c, err, data)
	return
}

// @Summary		创建帖子
// @Description	创建帖子
// @Tags			帖子
// @Accept			json
// @Produce		json
// @Param			title	query		string												true	"标题"
// @Param			content	query		string												true	"正文"
// @Success		200		{object}	model.ResponseWrapper{code=int,data=model.PostData}	"code为0表示成功，否则为失败"
// @Router			/post/create [post]
// @Security		ApiKeyAuth
func CreatePost(c *gin.Context) {
	req := &model.CreatePostRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	// 获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		model.Fail(c, model.UserNotLogin)
		return
	}
	// 创建帖子
	postData, err := service.CreatePost(c.Request.Context(), cast.ToInt(userID), req.Title, req.Content)
	model.Response(c, err, postData)
	return
}

// @Summary		更改帖子
// @Description	更改帖子
// @Tags			帖子
// @Accept			json
// @Produce		json
// @Param			title	query		string							true	"标题"
// @Param			content	query		string							true	"正文"
// @Param			post_id	query		int								true	"帖子ID"
// @Success		200		{object}	model.ResponseWrapper{code=int}	"code为0表示成功，否则为失败"
// @Router			/post/update [post]
// @Security		ApiKeyAuth
func UpdatePost(c *gin.Context) {
	req := &model.UpdatePostRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	// 获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		model.Fail(c, model.UserNotLogin)
		return
	}
	// 更新帖子
	err := service.UpdatePost(c.Request.Context(), req.PostID, cast.ToInt(userID), req.Title, req.Content)
	model.Response(c, err, nil)
	return
}

// @Summary		删除帖子
// @Description	根据帖子ID删除帖子
// @Tags			帖子
// @Accept			json
// @Produce		json
// @Param			post_id	query		int								true	"帖子ID"
// @Success		200		{object}	model.ResponseWrapper{code=int}	"code为0表示成功，否则为失败"
// @Router			/post/delete [post]
// @Security		ApiKeyAuth
func DeletePost(c *gin.Context) {
	req := &model.DeletePostRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	// 获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		model.Fail(c, model.UserNotLogin)
		return
	}
	// 删除帖子
	err_ := service.DeletePost(c.Request.Context(), req.PostID, cast.ToInt(userID))
	model.Response(c, err_, nil)
	return
}

// @Summary		获取帖子
// @Description	获取帖子
// @Tags			帖子
// @Accept			json
// @Produce		json
// @Param			post_id	path		int													true	"帖子ID"
// @Success		200		{object}	model.ResponseWrapper{code=int,data=model.PostData}	"code为0表示成功，否则为失败"
// @Router			/post/get/{post_id} [get]
func GetPostByID(c *gin.Context) {
	// 从路径参数中获取user_id
	postIDStr := c.Param("post_id")
	postID, err := cast.ToIntE(postIDStr)
	if err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	// 获取帖子信息
	data, err_ := service.GetPostByID(c.Request.Context(), postID)
	model.Response(c, err_, data)
	return
}
