package service

import (
	"context"
	"errors"
	"time"

	"github.com/lcsin/tesuto/pkg/errcode"
	"github.com/lcsin/tesuto/tesuto-project/internel/domain"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository/model"
	"gorm.io/gorm"
)

type IProjectService interface {
	GetProjectByID(ctx context.Context, id int64) (*domain.Project, error)
	GetProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*domain.Project, int64, error)

	CreateProject(ctx context.Context, project domain.Project) error
	UpdateProjectByID(ctx context.Context, id int64, name, desc string) error
	DeleteProjectByID(ctx context.Context, id int64) error
}

var (
	ErrProjectNotFound = errors.New("项目不存在")
)

type ProjectService struct {
	repo repository.IProjectRepository
}

func NewProjectService(repo repository.IProjectRepository) IProjectService {
	return &ProjectService{repo: repo}
}

func (p *ProjectService) GetProjectByID(ctx context.Context, id int64) (*domain.Project, error) {
	project, err := p.repo.GetProjectByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProjectNotFound
		}
		return nil, errcode.ErrInternalServer
	}

	return &domain.Project{
		ID:          project.ID,
		UID:         project.UID,
		Name:        project.Name,
		Desc:        project.Desc,
		CreatedTime: project.CreatedTime,
	}, nil
}

func (p *ProjectService) GetProjectsByUID(ctx context.Context, uid int64, pageNo, pageSize int64) ([]*domain.Project, int64, error) {
	projects, total, err := p.repo.GetProjectsByUID(ctx, uid, pageNo, pageSize)
	if err != nil {
		return nil, 0, errcode.ErrInternalServer
	}

	projectList := make([]*domain.Project, 0, len(projects))
	for _, v := range projects {
		projectList = append(projectList, &domain.Project{
			ID:          v.ID,
			UID:         v.UID,
			Name:        v.Name,
			Desc:        v.Desc,
			CreatedTime: v.CreatedTime,
		})
	}
	return projectList, total, nil
}

func (p *ProjectService) CreateProject(ctx context.Context, project domain.Project) error {
	return p.repo.AddProject(ctx, model.Project{
		UID:         project.UID,
		Name:        project.Name,
		Desc:        project.Desc,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
}

func (p *ProjectService) UpdateProjectByID(ctx context.Context, id int64, name, desc string) error {
	return p.repo.UpdateProjectByID(ctx, model.Project{
		ID:          id,
		Name:        name,
		Desc:        desc,
		UpdatedTime: time.Now(),
	})
}

func (p *ProjectService) DeleteProjectByID(ctx context.Context, id int64) error {
	return p.repo.DeleteProjectByID(ctx, id)
}
