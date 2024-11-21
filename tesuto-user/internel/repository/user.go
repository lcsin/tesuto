package repository

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-user/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)

	AddUser(ctx context.Context, user *model.User) error
}

type UserRepository struct {
	dao dao.IUserDAO
}

func NewUserRepository(dao dao.IUserDAO) IUserRepository {
	return &UserRepository{dao: dao}
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.dao.SelectUserByEmail(ctx, email)
}

func (u *UserRepository) AddUser(ctx context.Context, user *model.User) error {
	return u.dao.InsertUser(ctx, user)
}
