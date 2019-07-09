package user

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
)

type userService struct {
	userClient gen.UserServiceClient
}

func NewUserService(userClient gen.UserServiceClient) (interfaces.UserService, error) {
	return &userService{
		userClient: userClient,
	}, nil
}

func (u *userService) GetAllUsers() ([]*model.LeagueUser, error) {
	resp, err := u.userClient.GetAllUsers(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, errors.New("error getting users")
	}

	return u.toLeagueUsers(resp), nil
}

func (u *userService) GetLeagueUsers(ids []string) ([]*model.LeagueUser, error) {
	req := &gen.GroupIdRequest{
		Ids: ids,
	}

	resp, err := u.userClient.GetAllUsersByIds(context.Background(), req)
	if err != nil {
		return nil, errors.New("error getting users")
	}

	return u.toLeagueUsers(resp), nil
}

func (u *userService) toLeagueUsers(resp *gen.UserResponse) []*model.LeagueUser {
	var users []*model.LeagueUser
	for _, u := range resp.Users {
		users = append(users, model.LeagueUserFromGrpc(u))
	}
	return users
}

func (u *userService) GetOverallRank(id string) (int64, error) {
	req := &gen.IdRequest{
		Id: id,
	}

	resp, err := u.userClient.GetOverallRank(context.Background(), req)
	if err != nil {
		return 0, errors.New("error getting rank")
	}

	return resp.Rank, nil
}

func (u *userService) GetLeagueRank(id string, ids []string) (int64, error) {
	req := &gen.GroupRankRequest{
		Id: id,
		Ids: ids,
	}

	resp, err := u.userClient.GetRankForGroup(context.Background(), req)
	if err != nil {
		return 0, errors.New("error getting rank")
	}

	return resp.Rank, nil
}

func (u *userService) GetUserCount() (int64, error) {
	resp, err := u.userClient.GetUserCount(context.Background(), &empty.Empty{})
	if err != nil {
		return 0, errors.New("error getting count")
	}

	return resp.Count, nil
}
