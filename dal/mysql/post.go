package mysql

import "github.com/yifeng-coding/VUniversity/model/sql"

// GetPostList 分页加载帖子信息
func GetPostList(pageNum, pageSize int) ([]*sql.Post, error) {
	var posts []*sql.Post
	err := GetDB().Where("state != ?", 2).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("create_time desc").
		Find(&posts).Error
	return posts, err
}
