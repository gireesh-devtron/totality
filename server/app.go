package main

import (
	"fmt"
	client "github.com/gireesh-devtron/totality/server/grpc"
	"github.com/gireesh-devtron/totality/server/pkg/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

type App struct {
	Logger     *zap.SugaredLogger
	ServerImpl *service.UserServiceServerImpl
}

func NewApp(Logger *zap.SugaredLogger, ServerImpl *service.UserServiceServerImpl) *App {
	return &App{
		Logger:     Logger,
		ServerImpl: ServerImpl,
	}
}

func (app *App) Start() {

	port := 3000 //TODO: extract from environment variable
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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
	err = grpcServer.Serve(listener)
	if err != nil {
		app.Logger.Fatalw("failed to listen: %v", "err", err)
	}

}
