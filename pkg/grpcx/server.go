package grpcx

import (
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	cancel func()
	Port   int
}

func (s *Server) Serve() error {
	_, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	port := strconv.Itoa(s.Port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// 要先确保启动成功，再注册服务
	//err = s.register(ctx, port)
	//if err != nil {
	//	return err
	//}
	return s.Server.Serve(l)
}
