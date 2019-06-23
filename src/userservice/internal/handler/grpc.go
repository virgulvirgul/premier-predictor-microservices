package handler

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type userServiceServer struct {
	service interfaces.Service
}

func NewUserServiceServer(service interfaces.Service) (*userServiceServer, error) {
	log.Print("Registered userServiceServer handler")

	return &userServiceServer{
		service: service,
	}, nil
}

func (u *userServiceServer) GetAllUsers(ctx context.Context, req *empty.Empty) (*gen.UserResponse, error) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return u.buildResponse(users), nil
}

func (u *userServiceServer) GetLeagueUsers(ctx context.Context, req *gen.LeagueRequest) (*gen.UserResponse, error) {
	users, err := u.service.GetLeagueUsers(req.Ids)
	if err != nil {
		return nil, err
	}

	return u.buildResponse(users), nil
}

func (u *userServiceServer) buildResponse(users []*model.User) *gen.UserResponse {
	var usrs []*gen.User
	for _, u := range users {
		usrs = append(usrs, model.UserToGrpc(u))
	}

	return &gen.UserResponse{
		Users: usrs,
	}
}
