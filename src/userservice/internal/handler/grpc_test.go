package handler

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	id1 = "1"
	id2 = "2"
)

func TestUserServiceServer_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetAllUsers().Return(nil, e)

		resp, err := userService.GetAllUsers(context.Background(), &empty.Empty{})
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: id1,
			},
			{
				Id: id2,
			},
		}
		service.EXPECT().GetAllUsers().Return(users, nil)

		resp, err := userService.GetAllUsers(context.Background(), &empty.Empty{})
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Users[0].Id)
		assert.Equal(t, id2, resp.Users[1].Id)
	})
}

func TestUserServiceServer_GetLeagueUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	ids := []string{id1, id2}

	req := &gen.LeagueRequest{
		Ids: ids,
	}

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetLeagueUsers(ids).Return(nil, e)

		resp, err := userService.GetLeagueUsers(context.Background(), req)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns all league users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: id1,
			},
			{
				Id: id2,
			},
		}
		service.EXPECT().GetLeagueUsers(ids).Return(users, nil)

		resp, err := userService.GetLeagueUsers(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Users[0].Id)
		assert.Equal(t, id2, resp.Users[1].Id)
	})
}
