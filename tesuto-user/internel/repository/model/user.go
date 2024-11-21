package model

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Passwd    string    `json:"passwd"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

func (u *User) TableName() string {
	return "user_tbl"
}
