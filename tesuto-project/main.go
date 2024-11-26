package main

import (
	"log"

	"github.com/lcsin/tesuto/tesuto-project/ioc"
)

func main() {
	ioc.InitConfig()
	app := InitApp()

	log.Printf("server will started on: %v\n", app.server.Port)
	if err := app.server.Serve(); err != nil {
		panic(err)
	}
}
