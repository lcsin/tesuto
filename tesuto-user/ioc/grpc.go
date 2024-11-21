package ioc

import (
	"github.com/lcsin/tesuto/pkg/grpcx"
	"github.com/lcsin/tesuto/tesuto-user/internel/rpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPCXServer(userServer *rpc.UserServer) *grpcx.Server {
	type Config struct {
		port int `yaml:"port"`
	}

	port := viper.Get("app.server.port").(int)
	if port == 0 {
		panic("invalid port")
	}

	server := grpc.NewServer()

	return &grpcx.Server{
		Server: server,
		Port:   port,
	}
}
