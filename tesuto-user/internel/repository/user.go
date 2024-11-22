package repository

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-user/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
)

type IUserRepository interface {
	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(ctx context.Context, uid int64) (*model.User, error)
	// GetUserByEmail 根据邮件获取用户信息
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	// GetUserByEmailPasswd 根据用户邮箱和密码获取用户信息
	GetUserByEmailPasswd(ctx context.Context, email, passwd string) (*model.User, error)

	// AddUser 新增用户
	AddUser(ctx context.Context, user *model.User) error
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(ctx context.Context, user *model.User) error
	// UpdateUserPasswd 修改用户密码
	UpdateUserPasswd(ctx context.Context, uid int64, passwd string) error
}

type UserRepository struct {
	dao dao.IUserDAO
}

func NewUserRepository(dao dao.IUserDAO) IUserRepository {
	return &UserRepository{dao: dao}
}

func (u *UserRepository) GetUserByEmailPasswd(ctx context.Context, email, passwd string) (*model.User, error) {
	return u.dao.SelectUserByEmailPasswd(ctx, email, passwd)
}

func (u *UserRepository) GetUserByID(ctx context.Context, uid int64) (*model.User, error) {
	return u.dao.SelectUserByID(ctx, uid)
}

func (u *UserRepository) UpdateUserInfo(ctx context.Context, user *model.User) error {
	return u.dao.UpdateUserInfo(ctx, user)
}

func (u *UserRepository) UpdateUserPasswd(ctx context.Context, uid int64, passwd string) error {
	return u.dao.UpdateUserPasswd(ctx, uid, passwd)
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.dao.SelectUserByEmail(ctx, email)
}

func (u *UserRepository) AddUser(ctx context.Context, user *model.User) error {
	return u.dao.InsertUser(ctx, user)
}
