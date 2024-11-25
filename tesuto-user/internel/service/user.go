package service

import (
	"context"
	"errors"
	"time"

	"github.com/lcsin/tesuto/pkg/errcode"
	"github.com/lcsin/tesuto/pkg/response"
	"github.com/lcsin/tesuto/tesuto-user/internel/domain"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	Register(ctx context.Context, email, passwd, confirmPasswd string) error

	Login(ctx context.Context, email, passwd string) (*domain.User, error)

	UpdateUserInfo(ctx context.Context, user domain.User) error

	UpdateUserPasswd(ctx context.Context, email, oldPasswd, newPasswd string) error
}

var (
	ErrUserLoginFailed        = errors.New("用户不存在或密码错误")
	ErrUserNotExists          = errors.New("用户不存在")
	ErrEmailAlreadyRegistered = errors.New("邮箱已被注册")
	ErrPasswdIncorrect        = errors.New("密码错误")
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (u *UserService) Register(ctx context.Context, email, passwd, confirmPasswd string) error {
	// 根据邮箱找到用户，说明该邮箱已被注册
	_, err := u.repo.GetUserByEmail(ctx, email)
	if err == nil {
		return ErrEmailAlreadyRegistered
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return response.Failed(errcode.UNKnownError)
	}

	return u.repo.AddUser(ctx, model.User{
		Email:     email,
		Passwd:    string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (u *UserService) Login(ctx context.Context, email, passwd string) (*domain.User, error) {
	user, err := u.repo.GetUserByEmailPasswd(ctx, email, passwd)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserLoginFailed
		}
		return nil, errcode.ErrInternalServer
	}

	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserService) UpdateUserInfo(ctx context.Context, user domain.User) error {
	_, err := u.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotExists
		}
		return errcode.ErrInternalServer
	}

	return u.repo.UpdateUserInfo(ctx, model.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		UpdatedAt: time.Now(),
		UpdatedBy: user.Username,
	})
}

func (u *UserService) UpdateUserPasswd(ctx context.Context, email, oldPasswd, newPasswd string) error {
	// 先判断密码是否正确
	hash, err := bcrypt.GenerateFromPassword([]byte(oldPasswd), bcrypt.DefaultCost)
	if err != nil {
		return response.Failed(errcode.UNKnownError)
	}
	user, err := u.repo.GetUserByEmailPasswd(ctx, email, string(hash))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPasswdIncorrect
		}
		return errcode.ErrInternalServer
	}

	// 修改为新密码
	hash, err = bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost)
	if err != nil {
		return response.Failed(errcode.UNKnownError)
	}

	return u.repo.UpdateUserPasswd(ctx, user.ID, string(hash))
}
