package user

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
)

type userService struct {
	userClient gen.UserServiceClient
}

func NewUserService(userClient gen.UserServiceClient) (interfaces.UserService, error) {
	return &userService{
		userClient: userClient,
	}, nil
}

func (u *userService) GetUserByEmail(email string) (*model.User, error) {
	req := &gen.EmailRequest{
		Email: email,
	}
	user, err := u.userClient.GetUserByEmail(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return model.UserFromGrpc(user), nil
}
