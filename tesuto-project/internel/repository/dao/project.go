package dao

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-project/internel/repository/model"
	"gorm.io/gorm"
)

type IProjectDAO interface {
	SelectProjectByID(ctx context.Context, id int64) (*model.Project, error)
	SelectProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*model.Project, int64, error)

	InsertProject(ctx context.Context, project model.Project) error
	UpdateProjectByID(ctx context.Context, project model.Project) error
	DeleteProjectByID(ctx context.Context, id int64) error
}

type ProjectDAO struct {
	db *gorm.DB
}

func NewProjectDAO(db *gorm.DB) IProjectDAO {
	return &ProjectDAO{db: db}
}

func (p *ProjectDAO) SelectProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	var project model.Project
	if err := p.db.Where("id = ?", id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (p *ProjectDAO) SelectProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*model.Project, int64, error) {
	var projects []*model.Project
	if err := p.db.Where("uid = ?", uid).Limit(int(pageSize)).Offset(int((pageNo - 1) * pageSize)).Find(&projects).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := p.db.Where("uid = ?", uid).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

func (p *ProjectDAO) InsertProject(ctx context.Context, project model.Project) error {
	return p.db.Create(&project).Error
}

func (p *ProjectDAO) UpdateProjectByID(ctx context.Context, project model.Project) error {
	return p.db.Where("id = ?", project.ID).UpdateColumns(map[string]any{
		"name":         project.Name,
		"desc":         project.Desc,
		"updated_time": project.UpdatedTime,
	}).Error
}

func (p *ProjectDAO) DeleteProjectByID(ctx context.Context, id int64) error {
	return p.db.Where("id = ?", id).Delete(&model.Project{}).Error
}
