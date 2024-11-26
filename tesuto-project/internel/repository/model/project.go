package model

import "time"

type Project struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UID         int64     `json:"uid"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func (p *Project) TableName() string {
	return "project_tbl"
}
