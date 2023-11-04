package UserTests

import (
	"context"
	"fmt"
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
	impl.logger.Infow("test1 : get user by id 1")
	impl.testGetByUserId(1)
	impl.logger.Infow("test1 : get user by id 2")
	//
	//impl.logger.Infow("test1 : get user by id 3")
	//
	//impl.logger.Infow("test1 : get users by ids 1,2,3")
	//
	//impl.logger.Infow("test1 : get user by id 4,5")
	//
	//impl.logger.Infow("test1 : get user by id 1000")
}

func (impl *UserServerTester) testGetByUserId(id int32) {
	req := &client.UserRequest{
		Id: id,
	}
	user, err := impl.userServer.GetUser(context.Background(), req)
	if err != nil {
		impl.logger.Errorw("error in finding the user by Id", "id", id, "err", err)
		return
	}
	fmt.Println(fmt.Sprintf("user found for id : %v , user : %v", id, user))
	return
}

func (impl *UserServerTester) testGetByUserIds(ids []int) {

}
