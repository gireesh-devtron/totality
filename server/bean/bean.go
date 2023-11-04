package bean

import client "github.com/gireesh-devtron/totality/server/grpc"

type UserModel struct {
	Id        int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

func (model *UserModel) ConvertToUserResponse() *client.UserResponse {
	userResponse := &client.UserResponse{
		Id:        model.Id,
		FirstName: model.FirstName,
		City:      model.City,
		Phone:     model.Phone,
		Height:    model.Height,
		Married:   model.Married,
	}
	return userResponse
}
