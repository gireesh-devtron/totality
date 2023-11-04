package UserTests

import (
	client "github.com/gireesh-devtron/totality/server/grpc"
	"go.uber.org/zap"
)

type UserServerTester struct {
	logger     *zap.SugaredLogger
	userServer client.UserServiceClient
}

func NewUserServerTester(logger *zap.SugaredLogger,
	userServer client.UserServiceClient) *UserServerTester {

	return &UserServerTester{
		logger:     logger,
		userServer: userServer,
	}
}

func (impl *UserServerTester) RunTests() {

}
