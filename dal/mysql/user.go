package mysql

import (
	"context"
	"fmt"
	"github.com/yifeng-coding/VUniversity/model/sql"
	"time"
)

// CreateUser 创建用户
func CreateUser(ctx context.Context, username, password, email string) (*sql.User, error) {
	user := &sql.User{
		Username:   username,
		Password:   password,
		Email:      email,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := GetDB().WithContext(ctx).Create(user).Error
	return user, err
}

// GetUserByEmail 根据邮箱获取用户
func GetUserByEmail(ctx context.Context, email string) (*sql.User, error) {
	var user sql.User
	err := GetDB().WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

// GetUserByID 根据ID获取用户
func GetUserByID(ctx context.Context, id int) (*sql.User, error) {
	var user sql.User
	err := GetDB().WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

// UpdateUser 更新用户
func UpdateUser(ctx context.Context, userID int, updates map[string]any) error {
	if userID <= 0 {
		return fmt.Errorf("userID <= 0")
	}
	err := GetDB().WithContext(ctx).
		Model(&sql.User{}).
		Where("id = ?", userID).
		Updates(updates).Error
	return err
}

// MGetUserByIDs 批量获取用户信息
func MGetUserByIDs(ctx context.Context, userIDs []int) (map[int]*sql.User, error) {
	var users []*sql.User
	err := GetDB().WithContext(ctx).Where("id in ?", userIDs).Find(&users).Error
	if err != nil {
		return nil, err
	}
	userMap := make(map[int]*sql.User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, err
}
