package service

import (
	"context"
	client "github.com/gireesh-devtron/totality/server/grpc"
	"go.uber.org/zap"
)

type UserServiceServerImpl struct {
	client.UnimplementedUserServiceServer
	logger      *zap.SugaredLogger
	userService UserService
}

func NewUserServiceServerImpl(logger *zap.SugaredLogger, userService UserService) *UserServiceServerImpl {
	return &UserServiceServerImpl{
		logger:      logger,
		userService: userService,
	}
}
func (impl *UserServiceServerImpl) MustEmbedUnimplementedApplicationServiceServer() {
	panic("implement me")
}

func (impl *UserServiceServerImpl) GetUsers(ctx context.Context, request *client.UsersRequest) (*client.UsersResponse, error) {
	impl.logger.Infow("Hello from GetUsers Method")
	return &client.UsersResponse{}, nil
}
