package mysql

import (
	"context"
	"github.com/yifeng-coding/VUniversity/model/sql"
)

// GetPostList 分页加载帖子信息
func GetPostList(ctx context.Context, pageNum, pageSize int) ([]*sql.Post, error) {
	var posts []*sql.Post
	err := GetDB().WithContext(ctx).Where("state != ?", 2).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("create_time desc").
		Find(&posts).Error
	return posts, err
}

// GetPostTotalCount 获取帖子总数量
func GetPostTotalCount(ctx context.Context) (int, error) {
	var count int64
	err := GetDB().WithContext(ctx).Model(&sql.Post{}).Where("state != ?", 2).Count(&count).Error
	return int(count), err
}
