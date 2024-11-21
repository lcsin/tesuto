//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-user/internel/rpc"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
	"github.com/lcsin/tesuto/tesuto-user/ioc"
)

var iocSet = wire.NewSet(
	ioc.InitMySQL,
	ioc.InitGRPCXServer,
)

var userSvcProvider = wire.NewSet(
	dao.NewUserDAO,
	repository.NewUserRepository,
	service.NewUserService,
	rpc.NewUserServer,
)

func InitApp() *App {
	wire.Build(
		userSvcProvider,
		iocSet,
		wire.Struct(new(App), "*"),
	)

	return new(App)
}
