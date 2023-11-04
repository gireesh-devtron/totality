package main

import (
	"github.com/gireesh-devtron/totality/server/pkg/service"
	"go.uber.org/zap"
)

func main() {
	app := Init()
	app.Start()
}

func Init() *App {
	config := zap.NewProductionConfig()
	log, err := config.Build()
	if err != nil {
		panic("failed to create the logger: " + err.Error())
	}
	logger := log.Sugar()
	userService := service.NewUserServiceImpl(logger)
	userServiceWrapper := service.NewUserServiceServerImpl(logger, userService)
	app := NewApp(logger, userServiceWrapper)
	return app
}
