package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/yifeng-coding/VUniversity/dal/mysql"
	"github.com/yifeng-coding/VUniversity/dal/redis"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/util"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// GetPostList 分页加载帖子信息
func GetPostList(ctx context.Context, pageNum, pageSize, userID int) (*model.GetPostListData, *model.BizError) {
	// 获取帖子列表
	totalCount, dbPostList, err := mysql.GetPostList(ctx, pageNum, pageSize, userID)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 获取帖子对应的用户信息
	userIDMap := make(map[int]struct{})
	for _, dbPost := range dbPostList {
		userIDMap[dbPost.UserID] = struct{}{}
	}
	userIDs := make([]int, len(userIDMap))
	for userID := range userIDMap {
		userIDs = append(userIDs, userID)
	}
	userMap, err := mysql.MGetUserByIDs(ctx, userIDs)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}

	// 封装结果
	postDataList := make([]*model.PostData, len(dbPostList))
	for i, dbPost := range dbPostList {
		postDataList[i] = &model.PostData{
			ID:              dbPost.ID,
			Title:           dbPost.Title,
			Content:         dbPost.Content,
			CreateTimestamp: cast.ToString(dbPost.CreateTime.Unix()),
			UpdateTimestamp: cast.ToString(dbPost.UpdateTime.Unix()),
			UserID:          dbPost.UserID,
		}
		user := userMap[dbPost.UserID]
		if user != nil {
			postDataList[i].Username = user.Username
			postDataList[i].Avatar = user.Avatar
		}
	}
	return &model.GetPostListData{
		PostData:   postDataList,
		TotalCount: totalCount,
	}, nil
}

// CreatePost 创建帖子
func CreatePost(ctx context.Context, userID int, title, content string) (*model.PostData, *model.BizError) {
	// 限制一天最多发表100条帖子
	key := fmt.Sprintf("[create_post]%s_%d", time.Now().Format(time.DateOnly), userID)
	count, err := redis.GetRedis().Incr(ctx, key).Result()
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 设置key的过期时间，失败了忽略
	go func() {
		if err := redis.GetRedis().Expire(ctx, key, 24*time.Hour).Err(); err != nil {
			zap.L().Warn("set redis key expire failed", zap.String("key", key), zap.Error(err))
		}
	}()
	if count > 100 {
		return nil, model.ServerError.WithMessage("一天最多发表100条贴子")
	}

	// 创建帖子
	title = util.GetSensitiveFilter().Filter(title)     // 敏感词过滤
	content = util.GetSensitiveFilter().Filter(content) // 敏感词过滤
	dbPost, err := mysql.CreatePost(ctx, userID, title, content)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	result := &model.PostData{
		ID:              dbPost.ID,
		Title:           dbPost.Title,
		Content:         dbPost.Content,
		CreateTimestamp: cast.ToString(dbPost.CreateTime.Unix()),
	}
	return result, nil
}

// UpdatePost 更改帖子
func UpdatePost(ctx context.Context, id int, userID int, title, content string) *model.BizError {
	dbPost, err := mysql.GetPostByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ParamInvalid.WithMessage("帖子不存在")
		}
		return model.ServerError.WithMessage(err.Error())
	}
	if dbPost.UserID != userID {
		return model.ParamInvalid.WithMessage("仅能修改个人帖子")
	}
	updates := make(map[string]any)
	if title != "" {
		title = util.GetSensitiveFilter().Filter(title) // 敏感词过滤
		updates["title"] = title
	}
	if content != "" {
		content = util.GetSensitiveFilter().Filter(content) // 敏感词过滤
		updates["content"] = content
	}
	if err := mysql.UpdatePost(ctx, id, updates); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	return nil
}

// DeletePost 删除帖子
func DeletePost(ctx context.Context, id int, userID int) *model.BizError {
	dbPost, err := mysql.GetPostByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ParamInvalid.WithMessage("帖子不存在")
		}
		return model.ServerError.WithMessage(err.Error())
	}
	if dbPost.UserID != userID {
		return model.ParamInvalid.WithMessage("仅能删除个人帖子")
	}
	// 修改帖子状态实现删除（即软删除）
	updates := map[string]any{
		"status": 2,
	}
	if err := mysql.UpdatePost(ctx, id, updates); err != nil {
		return model.ServerError.WithMessage(err.Error())
	}
	return nil
}

// GetPostByID 获取帖子
func GetPostByID(ctx context.Context, id int) (*model.PostData, *model.BizError) {
	dbPost, err := mysql.GetPostByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ParamInvalid.WithMessage("帖子不存在")
		}
		return nil, model.ServerError.WithMessage(err.Error())
	}
	dbUser, err := mysql.GetUserByID(ctx, dbPost.UserID)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	return &model.PostData{
		ID:              dbPost.ID,
		Title:           dbPost.Title,
		Content:         dbPost.Content,
		CreateTimestamp: cast.ToString(dbPost.CreateTime.Unix()),
		UpdateTimestamp: cast.ToString(dbPost.UpdateTime.Unix()),
		UserID:          dbPost.UserID,
		Username:        dbUser.Username,
		Avatar:          dbUser.Avatar,
	}, nil
}
