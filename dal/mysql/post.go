package mysql

import (
	"context"
	"github.com/yifeng-coding/VUniversity/model"
	"github.com/yifeng-coding/VUniversity/model/sql"
	"time"
)

// GetPostList 分页加载帖子信息
func GetPostList(ctx context.Context, pageNum, pageSize int, userID int) (int, []*sql.Post, error) {
	db := GetDB().WithContext(ctx).Model(&sql.Post{}).Where("state != ?", model.PostStatusDeleted)
	if userID != 0 { // 如果用户ID不为0，则只查询该用户发布的帖子
		db = db.Where("user_id = ?", userID)
	}
	// 1.查询总数
	var totalCount int64
	if err := db.Count(&totalCount).Error; err != nil {
		return 0, nil, err
	}
	// 2.查询列表
	var posts []*sql.Post
	if err := db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&posts).Error; err != nil {
		return 0, nil, err
	}
	return int(totalCount), posts, nil
}

// CreatePost 创建帖子
func CreatePost(ctx context.Context, userID int, title, content string) (*sql.Post, error) {
	post := &sql.Post{
		UserID:     userID,
		Title:      title,
		Content:    content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := GetDB().WithContext(ctx).Create(post).Error
	return post, err
}

// UpdatePost 更新帖子信息
func UpdatePost(ctx context.Context, id int, updates map[string]any) error {
	return GetDB().WithContext(ctx).Model(&sql.Post{}).Where("id = ?", id).Updates(updates).Error
}

// GetPostByID 根据ID获取帖子信息
func GetPostByID(ctx context.Context, id int) (*sql.Post, error) {
	var post sql.Post
	err := GetDB().WithContext(ctx).Where("id = ?", id).Where("state != ?", model.PostStatusDeleted).First(&post).Error
	return &post, err
}
