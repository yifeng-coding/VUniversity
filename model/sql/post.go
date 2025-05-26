package sql

import "time"

type Post struct {
	ID           int       `json:"id" gorm:"primaryKey" `
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CommentCount int       `json:"comment_count"`
	Score        int       `json:"score"`
	Type         int       `json:"type"`
	State        int       `json:"state"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}
