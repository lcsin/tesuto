//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-project/internel/rpc"
	"github.com/lcsin/tesuto/tesuto-project/internel/service"
	"github.com/lcsin/tesuto/tesuto-project/ioc"
)

var iocSet = wire.NewSet(
	ioc.InitMySQL,
	ioc.InitGRPCXServer,
)

var projectSvcProvider = wire.NewSet(
	dao.NewProjectDAO,
	repository.NewProjectRepository,
	service.NewProjectService,
	rpc.NewProjectServer,
)

func InitApp() *App {
	wire.Build(
		projectSvcProvider,
		iocSet,
		wire.Struct(new(App), "*"),
	)

	return new(App)
}
