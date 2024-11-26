package dao

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-project/internel/repository/model"
	"gorm.io/gorm"
)

type IModuleDAO interface {
	SelectModuleByID(ctx context.Context, id int64, uid int64) (*model.Module, error)
	SelectModuleListByProjectID(ctx context.Context, projectID, uid int64) ([]*model.Module, error)

	InsertModule(ctx context.Context, module model.Module) error
	UpdateModuleByID(ctx context.Context, module model.Module) error
	DeleteModuleByID(ctx context.Context, id int64, uid int64) error
}

type ModuleDAO struct {
	db *gorm.DB
}

func NewModuleDAO(db *gorm.DB) IModuleDAO {
	return &ModuleDAO{db: db}
}

func (m *ModuleDAO) SelectModuleByID(ctx context.Context, id int64, uid int64) (*model.Module, error) {
	var module model.Module
	if err := m.db.Where("id = ? and uid = ?", id, uid).First(&module).Error; err != nil {
		return nil, err
	}
	return &module, nil
}

func (m *ModuleDAO) SelectModuleListByProjectID(ctx context.Context, projectID, uid int64) ([]*model.Module, error) {
	var modules []*model.Module
	if err := m.db.Where("project_id = ? and uid = ?", projectID, uid).Find(&modules).Error; err != nil {
		return nil, err
	}
	return modules, nil
}

func (m *ModuleDAO) InsertModule(ctx context.Context, module model.Module) error {
	return m.db.Create(&module).Error
}

func (m *ModuleDAO) UpdateModuleByID(ctx context.Context, module model.Module) error {
	return m.db.Where("id = ? and uid = ?", module.ID, module.UID).UpdateColumns(map[string]any{
		"name":         module.Name,
		"desc":         module.Desc,
		"updated_time": module.UpdatedTime,
	}).Error
}

func (m *ModuleDAO) DeleteModuleByID(ctx context.Context, id int64, uid int64) error {
	return m.db.Where("id = ? and uid = ?", id, uid).Delete(&model.Module{}).Error
}
