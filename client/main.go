package main

import (
	"context"
	"fmt"
	"github.com/gireesh-devtron/totality/client/UserTests"
	server "github.com/gireesh-devtron/totality/server/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := getConnection()
	if err != nil {
		panic("error in connecting with grpc service")
	}
	userServer := server.NewUserServiceClient(conn)
	config := zap.NewProductionConfig()
	log, err := config.Build()
	if err != nil {
		panic("failed to create the logger: " + err.Error())
	}
	logger := log.Sugar()
	userServerTester := UserTests.NewUserServerTester(logger, userServer)
	userServerTester.RunTests()
}

func getConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts = append(opts,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(20*1024*1024),
		),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	endpoint := fmt.Sprintf("dns:///%s", "localhost:3000")
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return conn, err
}
