package main

import (
	"fmt"
	"github.com/caarlos0/env"
	client "github.com/gireesh-devtron/totality/server/grpc"
	"github.com/gireesh-devtron/totality/server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

type ServerConfig struct {
	Port int `env:"port" envDefault:"3000"`
}

type App struct {
	ServerImpl *service.UserServiceServerImpl
}

func NewApp(ServerImpl *service.UserServiceServerImpl) *App {
	return &App{
		ServerImpl: ServerImpl,
	}
}

func (app *App) Start() {

	config := &ServerConfig{}
	env.Parse(config)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 10 * time.Second,
		}),
	}

	grpcServer := grpc.NewServer(opts...)
	client.RegisterUserServiceServer(grpcServer, app.ServerImpl)
	fmt.Println("grpc server running in port : ", config.Port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic(fmt.Sprintf("failed to listen: err %v", err.Error()))
	}
	fmt.Println("server stoped!")
}
