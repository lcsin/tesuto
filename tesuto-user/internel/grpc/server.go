package grpc

import (
	"context"

	userv1 "github.com/lcsin/tesuto/proto/v1/rpc"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
)

type UserServer struct {
	userv1.UnimplementedUserServer

	svc service.UserService
}

func (u *UserServer) GetUserByEmail(ctx context.Context, req *userv1.GetUserByEmailReq) (*userv1.GetUserByEmailRep, error) {
	_, err := u.svc.GetUserByEmail(ctx, req.Email)
	return &userv1.GetUserByEmailRep{}, err
}
