package main

import (
	"context"
	"testing"

	userv1 "github.com/lcsin/tesuto/proto/v1/rpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCClient(t *testing.T) {
	cc, err := grpc.Dial("localhost:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	client := userv1.NewUserClient(cc)
	resp, err := client.AddUser(context.Background(), &userv1.AddUserReq{
		Email:         "1847@qq.com",
		Username:      "root",
		Passwd:        "root",
		ConfirmPasswd: "root",
	})
	if err != nil {
		panic(err)
	}
	t.Logf("resp: %v", resp)
}
