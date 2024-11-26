package repository

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-project/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository/model"
)

type IModuleRepository interface {
	GetModuleByID(ctx context.Context, id int64, uid int64) (*model.Module, error)
	GetModuleByProjectID(ctx context.Context, projectID int64, uid int64) ([]*model.Module, error)

	CreateModule(ctx context.Context, module model.Module) error
	UpdateModuleByID(ctx context.Context, module model.Module) error
	DeleteModuleByID(ctx context.Context, id int64, uid int64) error
}

type ModuleRepository struct {
	dao dao.IModuleDAO
}

func NewModuleRepository(dao dao.IModuleDAO) IModuleRepository {
	return &ModuleRepository{dao: dao}
}

func (m *ModuleRepository) GetModuleByID(ctx context.Context, id int64, uid int64) (*model.Module, error) {
	return m.dao.SelectModuleByID(ctx, id, uid)
}

func (m *ModuleRepository) GetModuleByProjectID(ctx context.Context, projectID int64, uid int64) ([]*model.Module, error) {
	return m.dao.SelectModuleListByProjectID(ctx, projectID, uid)
}

func (m *ModuleRepository) CreateModule(ctx context.Context, module model.Module) error {
	return m.dao.InsertModule(ctx, module)
}

func (m *ModuleRepository) UpdateModuleByID(ctx context.Context, module model.Module) error {
	return m.dao.UpdateModuleByID(ctx, module)
}

func (m *ModuleRepository) DeleteModuleByID(ctx context.Context, id int64, uid int64) error {
	return m.dao.DeleteModuleByID(ctx, id, uid)
}
