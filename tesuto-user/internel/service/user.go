package service

import (
	"context"
	"errors"
	"time"

	"github.com/lcsin/tesuto/pkg/errcode"
	"github.com/lcsin/tesuto/tesuto-user/internel/domain"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(ctx context.Context, uid int64) (*domain.User, error)
	// GetUserByEmail 根据邮件获取用户信息
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)

	// AddUser 新增用户
	AddUser(ctx context.Context, email, username, passwd, confirmPasswd string) error
	// UpdateUserInfo 更新用户信息
	UpdateUserInfo(ctx context.Context, user *domain.User) error
	// UpdateUserPasswd 修改用户密码
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

func (u *UserService) GetUserByID(ctx context.Context, uid int64) (*domain.User, error) {
	if uid == 0 {
		return nil, errcode.ErrInvalidParams
	}

	user, err := u.repo.GetUserByID(ctx, uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errcode.ErrInternalServer
	}
	if user == nil || user.ID == 0 {
		return nil, ErrUserNotExists
	}

	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserService) UpdateUserInfo(ctx context.Context, user *domain.User) error {
	if user == nil || user.ID == 0 {
		return errcode.ErrInvalidParams
	}

	if err := u.repo.UpdateUserInfo(ctx, &model.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		UpdatedAt: time.Now(),
	}); err != nil {
		return errcode.ErrInternalServer
	}

	return nil
}

func (u *UserService) UpdateUserPasswd(ctx context.Context, email, oldPasswd, newPasswd string) error {
	if len(email) == 0 || len(oldPasswd) == 0 || len(newPasswd) == 0 {
		return errcode.ErrInvalidParams
	}
	// 检查旧密码是否正确
	oldPasswdHash, err := bcrypt.GenerateFromPassword([]byte(oldPasswd), bcrypt.DefaultCost)
	if err != nil {
		return errcode.ErrInternalServer
	}

	user, err := u.repo.GetUserByEmailPasswd(ctx, email, string(oldPasswdHash))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errcode.ErrInternalServer
	}
	if user == nil {
		return ErrPasswdIncorrect
	}

	// 加密新密码
	newPasswdHash, err := bcrypt.GenerateFromPassword([]byte(newPasswd), bcrypt.DefaultCost)
	if err != nil {
		return errcode.ErrInternalServer
	}
	if err = u.repo.UpdateUserPasswd(ctx, user.ID, string(newPasswdHash)); err != nil {
		return errcode.ErrInternalServer
	}

	return nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if len(email) == 0 {
		return nil, errcode.ErrInvalidParams
	}

	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errcode.ErrInternalServer
	}
	if user == nil || user.ID == 0 {
		return nil, ErrUserNotExists
	}

	return &domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserService) AddUser(ctx context.Context, email, username, passwd, confirmPasswd string) error {
	if passwd != confirmPasswd {
		return ErrPasswdInconsistency
	}
	if len(email) == 0 {
		return errcode.ErrInvalidParams
	}

	// 邮箱已被注册
	if _, err := u.GetUserByEmail(ctx, email); err == nil {
		return ErrEmailAlreadyRegistered
	}

	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return errcode.ErrInternalServer
	}

	return u.repo.AddUser(ctx, &model.User{
		Email:     email,
		Username:  username,
		Passwd:    string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
