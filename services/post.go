package services

import (
	"github.com/spf13/cast"
	"github.com/yifeng-coding/VUniversity/dal/mysql"
	"github.com/yifeng-coding/VUniversity/model/errs"
	"github.com/yifeng-coding/VUniversity/model/response"
)

// GetPostList 分页加载帖子信息
func GetPostList(pageNum, pageSize int) ([]*response.PostData, *errs.BizError) {
	dbPostList, err := mysql.GetPostList(pageNum, pageSize)
	if err != nil {
		return nil, errs.ServerError.WithMessage(err.Error())
	}
	result := make([]*response.PostData, len(dbPostList))
	for i, dbPost := range dbPostList {
		result[i] = &response.PostData{
			ID:              dbPost.ID,
			Title:           dbPost.Title,
			Content:         dbPost.Content,
			CreateTimestamp: cast.ToString(dbPost.CreateTime.Unix()),
			UpdateTimestamp: cast.ToString(dbPost.UpdateTime.Unix()),
			UserID:          dbPost.UserID,
			Username:        "张三",
			Avatar:          "https://avatars.githubusercontent.com/u/117701379?s=48&v=4",
		}
	}
	return result, nil
}
