// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-user/internel/rpc"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
	"github.com/lcsin/tesuto/tesuto-user/ioc"
)

// Injectors from wire.go:

func InitApp() *App {
	db := ioc.InitMySQL()
	iUserDAO := dao.NewUserDAO(db)
	iUserRepository := repository.NewUserRepository(iUserDAO)
	iUserService := service.NewUserService(iUserRepository)
	userServer := rpc.NewUserServer(iUserService)
	server := ioc.InitGRPCXServer(userServer)
	app := &App{
		server: server,
	}
	return app
}

// wire.go:

var iocSet = wire.NewSet(ioc.InitMySQL, ioc.InitGRPCXServer)

var userSvcProvider = wire.NewSet(dao.NewUserDAO, repository.NewUserRepository, service.NewUserService, rpc.NewUserServer)
