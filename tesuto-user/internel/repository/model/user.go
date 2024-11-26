package model

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email       string    `json:"email" gorm:"unique"`
	Username    string    `json:"username"`
	Passwd      string    `json:"passwd"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

func (u *User) TableName() string {
	return "user_tbl"
}
