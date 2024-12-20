// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository"
	"github.com/lcsin/tesuto/tesuto-project/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-project/internel/rpc"
	"github.com/lcsin/tesuto/tesuto-project/internel/service"
	"github.com/lcsin/tesuto/tesuto-project/ioc"
)

// Injectors from wire.go:

func InitApp() *App {
	db := ioc.InitMySQL()
	iProjectDAO := dao.NewProjectDAO(db)
	iProjectRepository := repository.NewProjectRepository(iProjectDAO)
	iProjectService := service.NewProjectService(iProjectRepository)
	projectServer := rpc.NewProjectServer(iProjectService)
	iModuleDAO := dao.NewModuleDAO(db)
	iModuleRepository := repository.NewModuleRepository(iModuleDAO)
	iModuleService := service.NewModuleService(iModuleRepository)
	moduleServer := rpc.NewModuleServer(iModuleService)
	server := ioc.InitGRPCXServer(projectServer, moduleServer)
	app := &App{
		server: server,
	}
	return app
}

// wire.go:

var iocSet = wire.NewSet(ioc.InitMySQL, ioc.InitGRPCXServer)

var projectSvcProvider = wire.NewSet(dao.NewProjectDAO, repository.NewProjectRepository, service.NewProjectService, rpc.NewProjectServer)

var moduleSvcProvider = wire.NewSet(dao.NewModuleDAO, repository.NewModuleRepository, service.NewModuleService, rpc.NewModuleServer)
