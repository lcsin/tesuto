package grayscale

import (
	"context"
	"math/rand"

	"github.com/lcsin/tesuto/proto/v1/rpc/commonpb"
	"github.com/lcsin/tesuto/proto/v1/rpc/userpb"
	"go.uber.org/atomic"
)

type Client struct {
	userpb.UnimplementedUserServiceServer

	// 调用远程的连接
	remote userpb.UserServiceClient
	// 调用本地的连接
	local userpb.UserServiceClient

	// 流量阈值
	threshold *atomic.Int64
}

func (c *Client) Register(ctx context.Context, req *userpb.RegisterReq) (*commonpb.Empty, error) {
	return c.client().Register(ctx, req)
}

func (c *Client) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRep, error) {
	return c.client().Login(ctx, req)
}

func (c *Client) UpdateUserInfo(ctx context.Context, req *userpb.UpdateUserInfoReq) (*commonpb.Empty, error) {
	return c.client().UpdateUserInfo(ctx, req)
}

func (c *Client) UpdateUserPasswd(ctx context.Context, req *userpb.UpdateUserPasswdReq) (*commonpb.Empty, error) {
	return c.client().UpdateUserPasswd(ctx, req)
}

func (c *Client) UpdateThreshold(threshold int64) {
	c.threshold.Store(threshold)
}

func (c *Client) client() userpb.UserServiceClient {
	threshold := c.threshold.Load()

	num := rand.Int63n(100)
	if num < threshold {
		return c.remote
	}
	return c.local
}
