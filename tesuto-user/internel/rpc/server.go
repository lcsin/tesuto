package rpc

import (
	"context"

	userv1 "github.com/lcsin/tesuto/proto/v1/rpc"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
	"google.golang.org/grpc"
)

type UserServer struct {
	userv1.UnimplementedUserServer

	svc service.IUserService
}

func NewUserServer(svc service.IUserService) *UserServer {
	return &UserServer{svc: svc}
}

func (u *UserServer) GetUserByEmail(ctx context.Context, req *userv1.GetUserByEmailReq) (*userv1.GetUserByEmailRep, error) {
	resp, err := u.svc.GetUserByEmail(ctx, req.Email)
	return &userv1.GetUserByEmailRep{
		Id:       resp.ID,
		Username: resp.Username,
		Email:    resp.Email,
	}, err
}

func (u *UserServer) AddUser(ctx context.Context, req *userv1.AddUserReq) (*userv1.AddUserRep, error) {
	if err := u.svc.AddUser(ctx, req.Email, req.Username, req.Passwd, req.ConfirmPasswd); err != nil {
		return nil, err
	}
	return &userv1.AddUserRep{}, nil
}

func (u *UserServer) Register(server *grpc.Server) {
	userv1.RegisterUserServer(server, u)
}
