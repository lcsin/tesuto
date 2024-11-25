package rpc

import (
	"context"

	"github.com/lcsin/tesuto/proto/v1/rpc/commonpb"
	"github.com/lcsin/tesuto/proto/v1/rpc/userpb"
	"github.com/lcsin/tesuto/tesuto-user/internel/domain"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
	"google.golang.org/grpc"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer

	svc service.IUserService
}

func NewUserServer(svc service.IUserService) *UserServer {
	return &UserServer{svc: svc}
}

func (u *UserServer) Register(ctx context.Context, req *userpb.RegisterReq) (*commonpb.Empty, error) {
	if err := u.svc.Register(ctx, req.Email, req.Passwd, req.ConfirmPasswd); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (u *UserServer) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRep, error) {
	user, err := u.svc.Login(ctx, req.Email, req.Passwd)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginRep{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (u *UserServer) UpdateUserInfo(ctx context.Context, req *userpb.UpdateUserInfoReq) (*commonpb.Empty, error) {
	if err := u.svc.UpdateUserInfo(ctx, domain.User{
		ID:       req.Id,
		Email:    req.Email,
		Username: req.Username,
	}); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (u *UserServer) UpdateUserPasswd(ctx context.Context, req *userpb.UpdateUserPasswdReq) (*commonpb.Empty, error) {
	if err := u.svc.UpdateUserPasswd(ctx, req.Email, req.OldPasswd, req.NewPasswd); err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func (u *UserServer) RegisterServer(server *grpc.Server) {
	userpb.RegisterUserServiceServer(server, u)
}
