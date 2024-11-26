package dao

import (
	"context"
	"time"

	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
	"gorm.io/gorm"
)

type IUserDAO interface {
	// SelectUserByID 根据用户ID获取用户信息
	SelectUserByID(ctx context.Context, uid int64) (*model.User, error)
	// SelectUserByEmail 根据邮件获取用户信息
	SelectUserByEmail(ctx context.Context, email string) (*model.User, error)
	// SelectUserByEmailPasswd 根据用户邮箱和密码获取用户信息
	SelectUserByEmailPasswd(ctx context.Context, email, passwd string) (*model.User, error)

	// InsertUser 新增用户
	InsertUser(ctx context.Context, user model.User) error
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(ctx context.Context, user model.User) error
	// UpdateUserPasswd 修改用户密码
	UpdateUserPasswd(ctx context.Context, uid int64, passwd string) error
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) IUserDAO {
	return &UserDAO{db: db}
}

func (u *UserDAO) SelectUserByEmailPasswd(ctx context.Context, email, passwd string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ? and passwd = ?", email, passwd).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDAO) SelectUserByID(ctx context.Context, uid int64) (*model.User, error) {
	var user model.User
	if err := u.db.Where("id = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDAO) UpdateUserInfo(ctx context.Context, user model.User) error {
	return u.db.Where("id = ?", user.ID).Updates(map[string]interface{}{
		"email":      user.Email,
		"username":   user.Username,
		"updated_at": user.UpdatedTime,
	}).Error
}

func (u *UserDAO) UpdateUserPasswd(ctx context.Context, uid int64, passwd string) error {
	return u.db.Where("id = ?", uid).UpdateColumns(map[string]interface{}{
		"passwd":     passwd,
		"updated_at": time.Now(),
	}).Error
}

func (u *UserDAO) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDAO) InsertUser(ctx context.Context, user model.User) error {
	return u.db.Create(&user).Error
}
