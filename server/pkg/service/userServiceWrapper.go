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
	users, err := impl.userService.GetUsers(request)
	if err != nil {
		return nil, err
	}
	userResponse := &client.UsersResponse{
		UserResponses: make([]*client.UserResponse, 0),
	}

	for _, user := range users {
		userResponse.UserResponses = append(userResponse.UserResponses, user.ConvertToUserResponse())
	}
	return userResponse, nil
}

func (impl *UserServiceServerImpl) GetUser(ctx context.Context, request *client.UserRequest) (*client.UserResponse, error) {
	user, err := impl.userService.GetUser(request)
	if err != nil {
		return nil, err
	}
	return user.ConvertToUserResponse(), nil
}
