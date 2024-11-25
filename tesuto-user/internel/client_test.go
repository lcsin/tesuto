package internel

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCClient(t *testing.T) {
	cc, err := grpc.Dial("localhost:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	client := user_pb.NewUserClient(cc)
	resp, err := client.Register(context.Background(), &user_pb.RegisterReq{
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
