package repository

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-project/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository/model"
)

type IProjectRepository interface {
	GetProjectByID(ctx context.Context, id int64) (*model.Project, error)
	GetProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*model.Project, int64, error)

	AddProject(ctx context.Context, project model.Project) error
	UpdateProjectByID(ctx context.Context, project model.Project) error
	DeleteProjectByID(ctx context.Context, id int64) error
}

type ProjectRepository struct {
	dao dao.IProjectDAO
}

func NewProjectRepository(dao dao.IProjectDAO) IProjectRepository {
	return &ProjectRepository{dao: dao}
}

func (p *ProjectRepository) GetProjectByID(ctx context.Context, id int64) (*model.Project, error) {
	return p.dao.SelectProjectByID(ctx, id)
}

func (p *ProjectRepository) GetProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*model.Project, int64, error) {
	return p.dao.SelectProjectsByUID(ctx, uid, pageNo, pageSize)
}

func (p *ProjectRepository) AddProject(ctx context.Context, project model.Project) error {
	return p.dao.InsertProject(ctx, project)
}

func (p *ProjectRepository) UpdateProjectByID(ctx context.Context, project model.Project) error {
	return p.dao.UpdateProjectByID(ctx, project)
}

func (p *ProjectRepository) DeleteProjectByID(ctx context.Context, id int64) error {
	return p.dao.DeleteProjectByID(ctx, id)
}
