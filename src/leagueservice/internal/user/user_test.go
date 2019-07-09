//go:generate mockgen -destination=./mocks/mock_user_service_client.go -package=usermocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen UserServiceClient

package user

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	usermocks "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

const (
	id1 = "🆔"
	id2 = "⚽️"
)

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	var users []*gen.User
	var expectedUsers []*model.LeagueUser

	for i:=0; i<200000; i++ {
		users = append(users, &gen.User{
			Id: strconv.Itoa(i),
		})
		expectedUsers = append(expectedUsers, &model.LeagueUser{
			Id: strconv.Itoa(i),
		})
	}

	resp := &gen.UserResponse{
		Users: users,
	}

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		userClient.EXPECT().GetAllUsers(context.Background(), &empty.Empty{}).Return(nil, errors.New("error"))

		result, err := userService.GetAllUsers()
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		userClient.EXPECT().GetAllUsers(context.Background(), &empty.Empty{}).Return(resp, nil)

		result, err := userService.GetAllUsers()
		require.NoError(t, err)

		assert.Equal(t, expectedUsers, result)
	})
}

func TestUserService_GetLeagueUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	var users []*gen.User
	var expectedUsers []*model.LeagueUser

	for i:=0; i<5; i++ {
		users = append(users, &gen.User{
			Id: strconv.Itoa(i),
		})
		expectedUsers = append(expectedUsers, &model.LeagueUser{
			Id: strconv.Itoa(i),
		})
	}

	ids := []string{id1, id2}

	req := &gen.GroupIdRequest{
		Ids: ids,
	}

	resp := &gen.UserResponse{
		Users: users,
	}


	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		userClient.EXPECT().GetAllUsersByIds(context.Background(), req).Return(nil, errors.New("error"))

		result, err := userService.GetLeagueUsers(ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		userClient.EXPECT().GetAllUsersByIds(context.Background(), req).Return(resp, nil)

		result, err := userService.GetLeagueUsers(ids)
		require.NoError(t, err)

		assert.Equal(t, expectedUsers, result)
	})
}

func TestUserService_GetOverallRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	rank := int64(12345)

	req := &gen.IdRequest{
		Id: id1,
	}

	resp := &gen.RankResponse{
		Rank: rank,
	}


	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		userClient.EXPECT().GetOverallRank(context.Background(), req).Return(nil, errors.New("error"))

		result, err := userService.GetOverallRank(id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		userClient.EXPECT().GetOverallRank(context.Background(), req).Return(resp, nil)

		result, err := userService.GetOverallRank(id1)
		require.NoError(t, err)

		assert.Equal(t, rank, result)
	})
}

func TestUserService_GetLeagueRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	rank := int64(12345)

	ids := []string{id1, id2}

	req := &gen.GroupRankRequest{
		Id: id1,
		Ids: ids,
	}

	resp := &gen.RankResponse{
		Rank: rank,
	}


	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		userClient.EXPECT().GetRankForGroup(context.Background(), req).Return(nil, errors.New("error"))

		result, err := userService.GetLeagueRank(id1, ids)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		userClient.EXPECT().GetRankForGroup(context.Background(), req).Return(resp, nil)

		result, err := userService.GetLeagueRank(id1, ids)
		require.NoError(t, err)

		assert.Equal(t, rank, result)
	})
}

func TestUserService_GetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	count := int64(1234)

	resp := &gen.CountResponse{
		Count: count,
	}

	t.Run("returns error if there is a problem retrieving users", func(t *testing.T) {
		userClient.EXPECT().GetUserCount(context.Background(), &empty.Empty{}).Return(nil, errors.New("error"))

		result, err := userService.GetUserCount()
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets all users from UserService", func(t *testing.T) {
		userClient.EXPECT().GetUserCount(context.Background(), &empty.Empty{}).Return(resp, nil)

		result, err := userService.GetUserCount()
		require.NoError(t, err)

		assert.Equal(t, count, result)
	})
}