package service

import (
	"context"
	"time"

	"github.com/lcsin/tesuto/pkg/errcode"
	"github.com/lcsin/tesuto/pkg/response"
	"github.com/lcsin/tesuto/tesuto-user/internel/domain"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
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
		return nil, response.Failed(errcode.EmailIsEmpty)
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, response.Failed(errcode.UNKnownError)
	}

	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserService) AddUser(ctx context.Context, email, username, passwd, confirmPasswd string) error {
	if passwd != confirmPasswd {
		return response.Failed(errcode.PasswordInconsistency)
	}
	if len(email) == 0 {
		return response.Failed(errcode.EmailIsEmpty)
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return response.Failed(errcode.UNKnownError)
	}
	if user.ID > 0 {
		return response.Failed(errcode.EmailAlreadyRegistered)
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return response.Failed(errcode.UNKnownError)
	}

	return u.repo.AddUser(ctx, &model.User{
		Email:     email,
		Username:  username,
		Passwd:    string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
