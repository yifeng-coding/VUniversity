package service

import (
	"context"
	"github.com/spf13/cast"
	"github.com/yifeng-coding/VUniversity/dal/mysql"
	"github.com/yifeng-coding/VUniversity/model"
)

// GetPostList 分页加载帖子信息
func GetPostList(ctx context.Context, pageNum, pageSize int) (*model.GetPostListData, *model.BizError) {
	// 获取帖子列表
	dbPostList, err := mysql.GetPostList(ctx, pageNum, pageSize)
	if err != nil {
		return nil, model.ServerError.WithMessage(err.Error())
	}
	// 获取总数量
	totalCount, err := mysql.GetPostTotalCount(ctx)
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
