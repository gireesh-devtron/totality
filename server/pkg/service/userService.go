package service

import (
	"github.com/gireesh-devtron/totality/server/bean"
	client "github.com/gireesh-devtron/totality/server/grpc"
	"go.uber.org/zap"
	"sync"
)

type UserService interface {
	GetUsers(request *client.UsersRequest) (*client.UsersResponse, error)
}

type UserServiceImpl struct {
	locker   *sync.Mutex
	logger   *zap.SugaredLogger
	dataBase map[int]*bean.UserModel
}

func NewUserServiceImpl(logger *zap.SugaredLogger) *UserServiceImpl {
	dataBase := map[int]*bean.UserModel{
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
	}
	return &UserServiceImpl{
		logger:   logger,
		locker:   &sync.Mutex{},
		dataBase: dataBase,
	}
}

func (impl *UserServiceImpl) GetUsers(request *client.UsersRequest) (*client.UsersResponse, error) {
	impl.logger.Errorw("hello from UserService GetUsers method")
	return &client.UsersResponse{}, nil
}
