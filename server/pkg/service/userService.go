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
	dataBase := make(map[int]*bean.UserModel)
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
