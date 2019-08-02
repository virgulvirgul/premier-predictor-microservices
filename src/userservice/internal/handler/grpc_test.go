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

func TestUserServiceServer_GetAllUsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	ids := []string{id1, id2}

	req := &gen.GroupIdRequest{
		Ids: ids,
	}

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetAllUsersByIds(ids).Return(nil, e)

		resp, err := userService.GetAllUsersByIds(context.Background(), req)
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
		service.EXPECT().GetAllUsersByIds(ids).Return(users, nil)

		resp, err := userService.GetAllUsersByIds(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Users[0].Id)
		assert.Equal(t, id2, resp.Users[1].Id)
	})
}

func TestUserServiceServer_GetOverallRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	req := &gen.IdRequest{
		Id: id1,
	}

	t.Run("returns error if rank cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetOverallRank(id1).Return(int64(0), e)

		resp, err := userService.GetOverallRank(context.Background(), req)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns all league users", func(t *testing.T) {
		rank := int64(1231)
		service.EXPECT().GetOverallRank(id1).Return(rank, nil)

		resp, err := userService.GetOverallRank(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, rank, resp.Rank)
	})
}

func TestUserServiceServer_GetRankForGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	ids := []string{id1, id2}

	req := &gen.GroupRankRequest{
		Id:  id1,
		Ids: ids,
	}

	t.Run("returns error if rank cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetRankForGroup(id1, ids).Return(int64(0), e)

		resp, err := userService.GetRankForGroup(context.Background(), req)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns all league users", func(t *testing.T) {
		rank := int64(2)
		service.EXPECT().GetRankForGroup(id1, ids).Return(rank, nil)

		resp, err := userService.GetRankForGroup(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, rank, resp.Rank)
	})
}

func TestUserServiceServer_GetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	t.Run("returns error if user count cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetUserCount().Return(int64(0), e)

		resp, err := userService.GetUserCount(context.Background(), &empty.Empty{})
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns total user count", func(t *testing.T) {
		count := int64(2)
		service.EXPECT().GetUserCount().Return(count, nil)

		resp, err := userService.GetUserCount(context.Background(), &empty.Empty{})
		require.NoError(t, err)

		assert.Equal(t, count, resp.Count)
	})
}

func TestUserServiceServer_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)

	userService, err := NewUserServiceServer(service)
	require.NoError(t, err)

	email := "📧"

	req := &gen.EmailRequest{
		Email: email,
	}

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		service.EXPECT().GetUserByEmail(email).Return(nil, e)

		resp, err := userService.GetUserByEmail(context.Background(), req)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Nil(t, resp)
	})

	t.Run("returns user", func(t *testing.T) {
		users := []*model.User{
			{
				Id: id1,
			},
		}
		service.EXPECT().GetUserByEmail(email).Return(users, nil)

		resp, err := userService.GetUserByEmail(context.Background(), req)
		require.NoError(t, err)

		assert.Equal(t, id1, resp.Id)
	})
}
