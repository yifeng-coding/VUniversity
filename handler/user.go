package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/service"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// @Summary		创建用户
// @Description	创建用户
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			username	query		string												true	"用户名"
// @Param			password	query		string												true	"密码"
// @Param			email		query		string												true	"邮箱"
// @Param			verify_code	query		string												true	"六位验证码"
// @Success		200			{object}	model.ResponseWrapper{code=int,data=model.UserData}	"code为0表示成功，否则为失败"
// @Router			/user/register [post]
func CreateUser(c *gin.Context) {
	req := &model.CreateUserRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := service.CreateUser(c.Request.Context(), req.Username, req.Password, req.Email, req.VerifyCode)
	model.Response(c, err, data)
	return
}

// @Summary		获取注册邮箱验证码
// @Description	获取注册邮箱验证码
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			email	query		string							true	"邮箱"
// @Success		200		{object}	model.ResponseWrapper{code=int}	"code为0表示成功，否则为失败"
// @Router			/user/register_code [get]
func GetRegisterVerifyCode(c *gin.Context) {
	req := &model.GetRegisterVerifyCodeRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	err := service.GetRegisterVerifyCode(c.Request.Context(), req.Email)
	model.Response(c, err, nil)
	return
}

// @Summary		用户登录
// @Description	用户登录，返回token
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			email		query		string														true	"邮箱"
// @Param			password	query		string														true	"密码"
// @Success		200			{object}	model.ResponseWrapper{code=int,data=model.UserLoginData}	"code为0表示成功，否则为失败"
// @Router			/user/login [post]
func UserLogin(c *gin.Context) {
	req := &model.UserLoginRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := service.UserLogin(c.Request.Context(), req.Email, req.Password)
	model.Response(c, err, data)
	return
}

// @Summary		获取用户信息
// @Description	获取用户信息
// @Tags			用户
// @Accept			json
// @Produce		json
// @Success		200	{object}	model.ResponseWrapper{code=int,data=model.UserData}	"code为0表示成功，否则为失败"
// @Router			/user/info [get]
// @Security		ApiKeyAuth
func GetUserInfo(c *gin.Context) {
	// 判断用户是否登录
	userID, exists := c.Get("user_id")
	if !exists {
		model.Fail(c, model.UserNotLogin)
		return
	}
	data, err := service.GetUserInfo(c.Request.Context(), cast.ToInt(userID))
	model.Response(c, err, data)
	return
}

// @Summary		刷新token
// @Description	根据refresh_token获取新token
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			refresh_token	query		string														true	"刷新token"
// @Success		200				{object}	model.ResponseWrapper{code=int,data=model.TokenPairData}	"code为0表示成功，否则为失败"
// @Router			/user/refresh_token [post]
func GetNewTokenByRefreshToken(c *gin.Context) {
	req := &model.GetNewTokenByRefreshTokenRequest{}
	// 绑定并校验参数
	if err := c.ShouldBind(req); err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	data, err := service.GetNewTokenByRefreshToken(c.Request.Context(), req.RefreshToken)
	model.Response(c, err, data)
	return
}

// @Summary		更新用户信息
// @Description	更新用户信息
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			username	query		string							false	"用户名"
// @Param			avatar		query		string							false	"头像"
// @Success		200			{object}	model.ResponseWrapper{code=int}	"code为0表示成功，否则为失败"
// @Router			/user/update/info [post]
// @Security		ApiKeyAuth
func UpdateUser(c *gin.Context) {
	req := &model.UpdateUserRequest{}
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
	// 更新用户信息
	err := service.UpdateUser(c.Request.Context(), cast.ToInt(userID), req.Username, req.Avatar)
	model.Response(c, err, nil)
	return
}

// @Summary		更新密码
// @Description	更新密码
// @Tags			用户
// @Accept			json
// @Produce		json
// @Param			old_password	query		string							true	"旧密码"
// @Param			new_password	query		string							true	"新密码"
// @Success		200				{object}	model.ResponseWrapper{code=int}	"code为0表示成功，否则为失败"
// @Router			/user/update/password [post]
// @Security		ApiKeyAuth
func UpdatePassword(c *gin.Context) {
	req := &model.UpdatePasswordRequest{}
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
	// 更新密码
	err := service.UpdatePassword(c.Request.Context(), cast.ToInt(userID), req.OldPassword, req.NewPassword)
	model.Response(c, err, nil)
	return
}

// @Summary		上传用户头像
// @Description	上传用户头像并返回存储路径
// @Tags			用户
// @Accept			multipart/form-data
// @Produce		json
// @Param			avatar	formData	file														true	"头像文件"
// @Success		200		{object}	model.ResponseWrapper{code=int,data=model.UploadAvatarData}	"code为0表示成功，否则为失败"
// @Router			/user/upload/avatar [post]
// @Security		ApiKeyAuth
func UploadAvatar(c *gin.Context) {
	// 判断登录状态
	userID, exists := c.Get("user_id")
	if !exists {
		model.Fail(c, model.UserNotLogin)
		return
	}
	// 解析上传的文件
	file, err := c.FormFile("avatar")
	if err != nil {
		model.Fail(c, model.ParamInvalid.WithMessage(err.Error()))
		return
	}
	// 文件大小限制10M
	maxSize := int64(10 * 1024 * 1024)
	if file.Size > maxSize {
		model.Fail(c, model.ParamInvalid.WithMessage("文件大小不能超过10M"))
		return
	}
	// 文件扩展名限制为jpg、jpeg、png
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !slices.Contains([]string{".jpg", ".jpeg", ".png"}, ext) {
		model.Fail(c, model.ParamInvalid.WithMessage("文件扩展名只能为jpg、jpeg、png"))
		return
	}
	// 生成随机文件名
	fileName := fmt.Sprintf("%v-%v%v", userID, uuid.New().String(), ext)
	// 构建文件保存路径
	rootDir, _ := os.Getwd()
	filePath := filepath.Join(rootDir, "static/avatars", fileName)
	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		model.Fail(c, model.ServerError.WithMessage(err.Error()))
		return
	}
	model.Success(c, model.UploadAvatarData{
		AvatarPath: "/static/avatars/" + fileName,
	})
	return
}
