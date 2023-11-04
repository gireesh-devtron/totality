package service

import (
	"fmt"
	"github.com/gireesh-devtron/totality/server/bean"
	client "github.com/gireesh-devtron/totality/server/grpc"
)

type UserService interface {
	GetUsers(request *client.UsersRequest) ([]*bean.UserModel, error)
	GetUser(request *client.UserRequest) (*bean.UserModel, error)
}

type UserServiceImpl struct {
	dataBase map[int32]*bean.UserModel
}

func NewUserServiceImpl() *UserServiceImpl {
	dataBase := map[int32]*bean.UserModel{
		1: {
			Id:        1,
			FirstName: "Steve",
			City:      "LA",
			Phone:     "+37123456",
			Married:   true,
			Height:    5.4,
		},
		2: {
			Id:        1,
			FirstName: "Harvey",
			City:      "LA",
			Phone:     "+37456661",
			Married:   true,
			Height:    5.10,
		},
		3: {
			Id:        1,
			FirstName: "Jordan",
			City:      "N.Dacota",
			Phone:     "+37123362",
			Married:   true,
			Height:    6.7,
		},
		4: {
			Id:        1,
			FirstName: "Kobey",
			City:      "Chicago",
			Phone:     "+37145632",
			Married:   true,
			Height:    6.6,
		},
		5: {
			Id:        1,
			FirstName: "Satnam",
			City:      "Punjab",
			Phone:     "+910014314312",
			Married:   true,
			Height:    6.8,
		},
		6: {
			Id:        1,
			FirstName: "Siva",
			City:      "Vizag",
			Phone:     "+910014313333",
			Married:   false,
			Height:    5.5,
		},
	}
	return &UserServiceImpl{
		dataBase: dataBase,
	}
}

func (impl *UserServiceImpl) GetUsers(request *client.UsersRequest) ([]*bean.UserModel, error) {
	fmt.Println("serving get users by ids request...")
	defer fmt.Println("request served!")
	usersFound := make([]*bean.UserModel, 0, len(request.UserRequests))
	for _, userRequest := range request.UserRequests {
		user, found := impl.dataBase[userRequest.Id]
		if !found {
			continue
		}
		usersFound = append(usersFound, user)
	}
	return []*bean.UserModel{}, nil
}

func (impl *UserServiceImpl) GetUser(request *client.UserRequest) (*bean.UserModel, error) {
	fmt.Println("serving get user by id request...")
	defer fmt.Println("request served!")
	user, found := impl.dataBase[request.Id]
	if !found {
		return nil, fmt.Errorf(fmt.Sprintf("user not found for given id : %v", request.Id))
	}
	return user, nil
}
