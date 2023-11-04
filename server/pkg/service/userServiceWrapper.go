package service

import (
	"context"
	client "github.com/gireesh-devtron/totality/server/grpc"
)

type UserServiceServerImpl struct {
	client.UnimplementedUserServiceServer
	userService UserService
}

func NewUserServiceServerImpl(userService UserService) *UserServiceServerImpl {
	return &UserServiceServerImpl{
		userService: userService,
	}
}
func (impl *UserServiceServerImpl) MustEmbedUnimplementedApplicationServiceServer() {
	panic("implement me")
}

func (impl *UserServiceServerImpl) GetUsers(ctx context.Context, request *client.UsersRequest) (*client.UsersResponse, error) {
	impl.userService.GetUsers(request)
	return &client.UsersResponse{}, nil
}

func (impl *UserServiceServerImpl) GetUser(ctx context.Context, request *client.UserRequest) (*client.UserResponse, error) {
	impl.userService.GetUser(request)
	return &client.UserResponse{}, nil
}
