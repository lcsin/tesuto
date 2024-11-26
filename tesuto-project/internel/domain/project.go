package domain

import "time"

type Project struct {
	ID          int64     `json:"id"`
	UID         int64     `json:"uid"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
