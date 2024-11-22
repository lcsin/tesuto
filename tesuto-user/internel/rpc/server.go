package rpc

import (
	"context"

	"github.com/lcsin/tesuto/proto/v1/rpc/common"
	"github.com/lcsin/tesuto/proto/v1/rpc/user"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
)

type UserServer struct {
	svc service.IUserService
}

func NewUserServer(svc service.IUserService) *UserServer {
	return &UserServer{svc: svc}
}

func (u *UserServer) Register(ctx context.Context, req *user.RegisterReq) (*common.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) Login(ctx context.Context, req *user.LoginReq) (*user.LoginRep, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (*common.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) UpdateUserPasswd(ctx context.Context, req *user.UpdateUserPasswdReq) (*common.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServer) mustEmbedUnimplementedUserServer() {
	//TODO implement me
	panic("implement me")
}
