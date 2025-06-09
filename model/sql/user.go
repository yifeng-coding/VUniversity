package sql

import "time"

type User struct {
	ID         int       `json:"id" gorm:"primaryKey" `
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar"`
	State      int       `json:"state"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
