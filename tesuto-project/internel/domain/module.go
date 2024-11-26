package domain

import "time"

type Module struct {
	ID          int64     `json:"id"`
	UID         int64     `json:"uid"`
	ProjectID   int64     `json:"project_id"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
