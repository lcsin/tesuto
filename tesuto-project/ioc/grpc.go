package ioc

import (
	"github.com/lcsin/tesuto/pkg/grpcx"
	"github.com/lcsin/tesuto/tesuto-project/internel/rpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPCXServer(projectServer *rpc.ProjectServer) *grpcx.Server {
	type Config struct {
		port int `yaml:"port"`
	}

	port := viper.Get("app.server.port").(int)
	if port == 0 {
		panic("invalid port")
	}

	server := grpc.NewServer()
	projectServer.RegisterServer(server)

	return &grpcx.Server{
		Server: server,
		Port:   port,
	}
}
