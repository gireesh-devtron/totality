package main

import (
	"github.com/gireesh-devtron/totality/server/pkg/service"
)

func main() {
	app := Init()
	app.Start()
}

func Init() *App {

	userService := service.NewUserServiceImpl()
	userServiceWrapper := service.NewUserServiceServerImpl(userService)
	app := NewApp(userServiceWrapper)
	return app
}
