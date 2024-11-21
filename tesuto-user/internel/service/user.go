package service

import (
	"context"
	"fmt"
	"time"

	"github.com/lcsin/tesuto/user/internel/domain"
	"github.com/lcsin/tesuto/user/internel/repository"
	"github.com/lcsin/tesuto/user/internel/repository/model"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)

	AddUser(ctx context.Context, email, username, passwd, confirmPasswd string) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("参数错误")
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserService) AddUser(ctx context.Context, email, username, passwd, confirmPasswd string) error {
	if passwd != confirmPasswd {
		return fmt.Errorf("两次输入的密码不一致")
	}
	if len(email) == 0 {
		return fmt.Errorf("参数无效")
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user.ID > 0 {
		return fmt.Errorf("邮箱已被注册")
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return u.repo.AddUser(ctx, &model.User{
		Email:     email,
		Username:  username,
		Passwd:    string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
