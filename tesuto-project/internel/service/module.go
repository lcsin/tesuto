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

type IModuleService interface {
	GetModuleByID(ctx context.Context, id int64, uid int64) (*domain.Module, error)
	GetModuleByProjectID(ctx context.Context, projectID int64, uid int64) ([]*domain.Module, error)

	CreateModule(ctx context.Context, module domain.Module) error
	UpdateModuleByID(ctx context.Context, id int64, uid int64, name, desc string) error
	DeleteModuleByID(ctx context.Context, id int64, uid int64) error
}

var (
	ErrModuleNotFound = errors.New("模块不存在")
)

type ModuleService struct {
	repo repository.IModuleRepository
}

func NewModuleService(repo repository.IModuleRepository) IModuleService {
	return &ModuleService{repo: repo}
}

func (m *ModuleService) GetModuleByID(ctx context.Context, id int64, uid int64) (*domain.Module, error) {
	module, err := m.repo.GetModuleByID(ctx, id, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrModuleNotFound
		}
		return nil, errcode.ErrInternalServer
	}

	return &domain.Module{
		ID:          module.ID,
		UID:         module.UID,
		ProjectID:   module.ProjectID,
		Name:        module.Name,
		Desc:        module.Desc,
		CreatedTime: module.CreatedTime,
		UpdatedTime: module.UpdatedTime,
	}, err
}

func (m *ModuleService) GetModuleByProjectID(ctx context.Context, projectID int64, uid int64) ([]*domain.Module, error) {
	modules, err := m.repo.GetModuleByProjectID(ctx, projectID, uid)
	if err != nil {
		return nil, errcode.ErrInternalServer
	}

	list := make([]*domain.Module, 0, len(modules))
	for _, v := range modules {
		list = append(list, &domain.Module{
			ID:          v.ID,
			UID:         v.UID,
			ProjectID:   v.ProjectID,
			Name:        v.Name,
			Desc:        v.Desc,
			CreatedTime: v.CreatedTime,
			UpdatedTime: v.UpdatedTime,
		})
	}
	return list, nil
}

func (m *ModuleService) CreateModule(ctx context.Context, module domain.Module) error {
	if err := m.repo.CreateModule(ctx, model.Module{
		UID:         module.UID,
		ProjectID:   module.ProjectID,
		Name:        module.Name,
		Desc:        module.Desc,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}); err != nil {
		return errcode.ErrInternalServer
	}
	return nil
}

func (m *ModuleService) UpdateModuleByID(ctx context.Context, id int64, uid int64, name, desc string) error {
	if err := m.repo.UpdateModuleByID(ctx, model.Module{
		ID:          id,
		UID:         uid,
		Name:        name,
		Desc:        desc,
		UpdatedTime: time.Now(),
	}); err != nil {
		return errcode.ErrInternalServer
	}
	return nil
}

func (m *ModuleService) DeleteModuleByID(ctx context.Context, id int64, uid int64) error {
	if err := m.repo.DeleteModuleByID(ctx, id, uid); err != nil {
		return errcode.ErrInternalServer
	}
	return nil
}
