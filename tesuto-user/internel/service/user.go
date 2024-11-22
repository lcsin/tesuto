package service

import (
	"context"
	"errors"

	"github.com/lcsin/tesuto/tesuto-user/internel/domain"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
)

type IUserService interface {
	Register(ctx context.Context, user domain.User) error

	Login(ctx context.Context, email, passwd string) (*domain.User, error)

	UpdateUserInfo(ctx context.Context, user domain.User) error

	UpdateUserPasswd(ctx context.Context, email, oldPasswd, newPasswd string) error
}

var (
	ErrUserNotFoundOrPasswdIncorrect = errors.New("用户不存在或密码错误")
	ErrUserNotExists                 = errors.New("用户不存在")
	ErrEmailAlreadyRegistered        = errors.New("邮箱已被注册")
	ErrPasswdInconsistency           = errors.New("密码不一致")
	ErrPasswdIncorrect               = errors.New("密码错误")
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

func (u *UserService) Register(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) Login(ctx context.Context, email, passwd string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUserInfo(ctx context.Context, user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUserPasswd(ctx context.Context, email, oldPasswd, newPasswd string) error {
	//TODO implement me
	panic("implement me")
}
