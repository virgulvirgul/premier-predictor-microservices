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

func (u *userServiceServer) GetAllUsersByIds(ctx context.Context, req *gen.GroupIdRequest) (*gen.UserResponse, error) {
	users, err := u.service.GetAllUsersByIds(req.Ids)
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

func (u *userServiceServer) GetOverallRank(ctx context.Context, req *gen.IdRequest) (*gen.RankResponse, error) {
	rank, err := u.service.GetOverallRank(req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.RankResponse{
		Rank: rank,
	}, nil
}

func (u *userServiceServer) GetRankForGroup(ctx context.Context, req *gen.GroupRankRequest) (*gen.RankResponse, error) {
	rank, err := u.service.GetRankForGroup(req.Id, req.Ids)
	if err != nil {
		return nil, err
	}

	return &gen.RankResponse{
		Rank: rank,
	}, nil
}

func (u *userServiceServer) GetUserCount(context.Context, *empty.Empty) (*gen.CountResponse, error) {
	count, err := u.service.GetUserCount()
	if err != nil {
		return nil, err
	}

	return &gen.CountResponse{
		Count: count,
	}, nil
}

func (u *userServiceServer) GetUserByEmail(ctx context.Context, req *gen.EmailRequest) (*gen.User, error) {
	user, err := u.service.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return model.UserToGrpc(user), nil
}
