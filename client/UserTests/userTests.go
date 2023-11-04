package UserTests

import (
	"context"
	"fmt"
	client "github.com/gireesh-devtron/totality/server/grpc"
	"log"
)

type UserServerTester struct {
	userServer client.UserServiceClient
}

func NewUserServerTester(userServer client.UserServiceClient) *UserServerTester {

	return &UserServerTester{
		userServer: userServer,
	}
}

func (impl *UserServerTester) RunTests() {
	fmt.Println("running tests on users server......")
	fmt.Println("test1 : get user by id 1")
	impl.testGetByUserId(1)
	fmt.Println("test1 : get user by id 2")
	impl.testGetByUserId(2)
	fmt.Println("test1 : get user by id 3")
	impl.testGetByUserId(3)
	fmt.Println("test1 : get users by ids 1,2,3")
	impl.testGetByUserIds([]int32{1, 2, 3})
	fmt.Println("test1 : get user by id 4,5")
	impl.testGetByUserIds([]int32{4, 5})
	fmt.Println("test1 : get user by id 1000")
	impl.testGetByUserId(1000)
	fmt.Println("running tests on users server ended!")

}

func (impl *UserServerTester) testGetByUserId(id int32) {
	req := &client.UserRequest{
		Id: id,
	}
	user, err := impl.userServer.GetUser(context.Background(), req)
	if err != nil {
		log.Println("error in finding the user by Id", id, "err", err)
		return
	}
	fmt.Println(fmt.Sprintf("user found for id : %v , user : %v", id, user))
	return
}

func (impl *UserServerTester) testGetByUserIds(ids []int32) {
	req := &client.UsersRequest{}
	userRequests := make([]*client.UserRequest, len(ids))
	for i, id := range ids {
		userRequests[i] = &client.UserRequest{
			Id: id,
		}
	}
	req.UserRequests = userRequests

	users, err := impl.userServer.GetUsers(context.Background(), req)
	if err != nil {
		log.Println(fmt.Sprintf("error in finding the user by ids : %v , err : %v", ids, err))
		return
	}
	fmt.Println(fmt.Sprintf("user found for ids : %v , user : %v", ids, users))
	return
}
