package model

// UserData 用户数据
type UserData struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Avatar          string `json:"avatar"`
	Email           string `json:"email"`
	CreateTimestamp string `json:"create_timestamp"` // 创建时间戳（秒级）
}

// TokenPairData Jwt token
type TokenPairData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// CreateUserRequest 创建用户请求参数
type CreateUserRequest struct {
	Username   string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Email      string `form:"email" json:"email" binding:"required,email"`
	VerifyCode string `form:"verify_code" json:"verify_code" binding:"required,len=6"`
}

// GetRegisterVerifyCodeRequest 获取注册验证码请求参数
type GetRegisterVerifyCodeRequest struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// UserLoginRequest 用户登录请求参数
type UserLoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserLoginData 用户登录响应数据
type UserLoginData struct {
	UserData      UserData      `json:"user_data"`
	TokenPairData TokenPairData `json:"token_pair_data"`
}

// GetNewTokenByRefreshTokenRequest 获取新的token请求参数
type GetNewTokenByRefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" binding:"required"`
}

// UpdatePasswordRequest 修改密码请求参数
type UpdatePasswordRequest struct {
	OldPassword string `form:"old_password" json:"old_password" binding:"required"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required"`
}

// UpdateUserRequest 修改用户信息请求参数
type UpdateUserRequest struct {
	Username string `form:"username" json:"username"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// UploadAvatarData 上传头像响应数据
type UploadAvatarData struct {
	AvatarPath string `json:"avatar_path"`
}
