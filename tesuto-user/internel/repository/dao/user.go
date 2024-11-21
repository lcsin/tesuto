package dao

import (
	"context"

	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
	"gorm.io/gorm"
)

type IUserDAO interface {
	SelectUserByEmail(ctx context.Context, email string) (*model.User, error)

	InsertUser(ctx context.Context, user *model.User) error
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) IUserDAO {
	return &UserDAO{db: db}
}

func (u *UserDAO) SelectUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDAO) InsertUser(ctx context.Context, user *model.User) error {
	return u.db.Create(&user).Error
}
