package main

import (
	"github.com/google/wire"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository"
	"github.com/lcsin/tesuto/tesuto-user/internel/repository/dao"
	"github.com/lcsin/tesuto/tesuto-user/internel/service"
	"github.com/lcsin/tesuto/tesuto-user/ioc"
)

var thirdPartySet = wire.NewSet(
	ioc.InitMySQL,
)

var userSvcProvider = wire.NewSet(
	dao.NewUserDAO,
	repository.NewUserRepository,
	service.NewUserService,
)

func InitApp() *App {
	wire.Build(
		thirdPartySet,
		userSvcProvider,
		wire.Struct(new(App), "*"),
	)

	return new(App)
}
