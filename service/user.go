package service

import (
	"context"
	"errors"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yifeng-coding/VUniversity/dal/email"
	"github.com/yifeng-coding/VUniversity/dal/mysql"
	"github.com/yifeng-coding/VUniversity/dal/redis"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CreateUser 创建用户
func CreateUser(ctx context.Context, username, password, email, verifyCode string) (*model.UserData, *model.BizError) {
	// 从redis中获取验证码
	registerCodeKey := fmt.Sprintf("[register_key]%s", email)
	val, err := redis.GetRedis().Get(ctx, registerCodeKey).Result()
	if err != nil {
		if err == go_redis.Nil {
			return nil, model.ParamInvalid.WithMessage("验证码不存在或已过期，请重新获取")
		}
		return nil, model.ServerError.WithMessage(err.Error())
	}
	if val != verifyCode {
		return nil, model.ParamInvalid.WithMessage("验证码错误")
	}
	// 密码加密处理
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 创建用户
	createUser, err := mysql.CreateUser(ctx, username, string(hashedPasswordBytes), email)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, model.UserAlreadyExist
		}
		return nil, model.ServerError.WithMessage(err.Error())
	}

	result := &model.UserData{
		ID:              createUser.ID,
		Username:        createUser.Username,
		Email:           createUser.Email,
		CreateTimestamp: cast.ToString(createUser.CreateTime.Unix()),
	}
	return result, nil
}

func GetRegisterVerifyCode(ctx context.Context, registerEmail string) *model.BizError {
	registerCodeKey := fmt.Sprintf("[register_key]%s", registerEmail)
	// 判断是否已经发过
	if err := redis.GetRedis().Get(ctx, registerCodeKey).Err(); err == nil {
		return model.ParamInvalid.WithMessage("验证码已发送，请稍后再试")
	} else if err != go_redis.Nil {
		return model.ServerError.WithMessage(err.Error())
	}
	//  生成六位数字验证码
	verifyCode := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 将验证码存储到redis中（过期时间五分钟）
	if err := redis.GetRedis().Set(ctx, registerCodeKey, verifyCode, 300*time.Second).Err(); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	// 判断用户是否已经注册
	_, err := mysql.GetUserByEmail(ctx, registerEmail)
	if err == nil {
		return model.UserAlreadyExist
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ServerError.WithMessage(err.Error())
	}
	// 给用户发送邮件
	msg := fmt.Sprintf("您的验证码为【%s】，有效期五分钟，请尽快验证！", verifyCode)
	if err := email.Send(ctx, registerEmail, "微学堂注册验证码", msg); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	return nil
}

func UserLogin(ctx context.Context, email string, password string) (*model.UserLoginData, *model.BizError) {
	// 查询用户信息
	user, err := mysql.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.UserNotExist
		}
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, model.ParamInvalid.WithMessage("密码错误")
	}
	// 生成 access token
	accessToken, err := util.GenerateToken(ctx, user.ID, util.AccessTokenExpire)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 生成 refresh token
	refreshToken, err := util.GenerateToken(ctx, user.ID, util.RefreshTokenExpire)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	return &model.UserLoginData{
		UserData: model.UserData{
			ID:              user.ID,
			Username:        user.Username,
			Avatar:          user.Avatar,
			Email:           user.Email,
			CreateTimestamp: cast.ToString(user.CreateTime.Unix()),
		},
		TokenPairData: model.TokenPairData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func GetUserInfo(ctx context.Context, userID int) (*model.UserData, *model.BizError) {
	user, err := mysql.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.UserNotExist
		}
		return nil, model.ServerError.WithMessage(err.Error())
	}
	return &model.UserData{
		ID:              user.ID,
		Username:        user.Username,
		Avatar:          user.Avatar,
		Email:           user.Email,
		CreateTimestamp: cast.ToString(user.CreateTime.Unix()),
	}, nil
}

func GetNewTokenByRefreshToken(ctx context.Context, refreshToken string) (*model.TokenPairData, *model.BizError) {
	// 验证刷新令牌
	claims, err := util.ParseToken(refreshToken)
	if err != nil {
		return nil, model.TokenInvalid
	}
	// 获取用户ID
	userID := claims.UserID
	if userID == 0 {
		return nil, model.TokenInvalid
	}
	// 生成新的访问&刷新令牌
	accessToken, err := util.GenerateToken(ctx, userID, util.AccessTokenExpire)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	newRefreshToken, err := util.GenerateToken(ctx, userID, util.RefreshTokenExpire)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	return &model.TokenPairData{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func UpdateUser(ctx context.Context, userID int, username, avatar string) *model.BizError {
	// 更新用户信息
	updates := map[string]any{}
	if username != "" {
		updates["username"] = username
	}
	if avatar != "" {
		rootDir, _ := os.Getwd()
		avatarPath := filepath.Join(rootDir, avatar)
		if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
			return model.ParamInvalid.WithMessage("头像不存在")
		}
		updates["avatar"] = avatar
	}
	if err := mysql.UpdateUser(ctx, userID, updates); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	return nil
}

func UpdatePassword(ctx context.Context, userID int, oldPassword, newPassword string) *model.BizError {
	// 查询用户信息
	user, err := mysql.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserNotExist
		}
		return model.ServerError.WithMessage(err.Error())
	}
	// 校验旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return model.ParamInvalid.WithMessage("旧密码错误")
	}
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	// 更新密码
	if err := mysql.UpdateUser(ctx, userID, map[string]any{"password": string(hashedPasswordBytes)}); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	return nil
}
