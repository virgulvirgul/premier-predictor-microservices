package handler

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	legacyusermocks "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

const (
	userId   = 1
	email    = "email"
	password = "password"
)

func TestLegacyUserServiceServer_GetLegacyUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := legacyusermocks.NewMockService(ctrl)

	grpc, err := NewLegacyUserServiceServer(service)
	require.Error(t, err)

	req := &gen.LegacyIdRequest{
		Id: userId,
	}

	t.Run("returns error if a user could not be retrieved", func(t *testing.T) {
		service.EXPECT().GetUserById(userId).Return(nil, model.ErrLegacyUserNotFound)

		result, err := grpc.GetLegacyUserById(context.Background(), req)
		require.Error(t, err)

		assert.Empty(t, result)
		assert.Equal(t, model.ErrLegacyUserNotFound, err)
	})

	t.Run("returns legacy user", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
		}

		service.EXPECT().GetUserById(userId).Return(user, nil)

		result, err := grpc.GetLegacyUserById(context.Background(), req)
		require.NoError(t, err)

		expectedResult := model.LegacyUserToGrpc(user)

		assert.Equal(t, expectedResult, result)
	})
}

func TestLegacyUserServiceServer_LegacyLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := legacyusermocks.NewMockService(ctrl)

	grpc, err := NewLegacyUserServiceServer(service)
	require.Error(t, err)

	req := &gen.LoginRequest{
		Email:    email,
		Password: password,
	}

	t.Run("returns unauthorised if a user could not be found", func(t *testing.T) {
		service.EXPECT().LegacyLogin(email, password).Return(nil, model.ErrLegacyLoginFailed)

		result, err := grpc.LegacyLogin(context.Background(), req)
		require.Error(t, err)

		assert.Empty(t, result)
		assert.Equal(t, status.Error(codes.Unauthenticated, model.ErrLegacyLoginFailed.Error()), err)
	})

	t.Run("returns error if a user could not be retrieved", func(t *testing.T) {
		e := errors.New("error")
		service.EXPECT().LegacyLogin(email, password).Return(nil, e)

		result, err := grpc.LegacyLogin(context.Background(), req)
		require.Error(t, err)

		assert.Empty(t, result)
		assert.Equal(t, e, err)
	})

	t.Run("returns legacy user", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
		}

		service.EXPECT().LegacyLogin(email, password).Return(user, nil)

		result, err := grpc.LegacyLogin(context.Background(), req)
		require.NoError(t, err)

		expectedResult := model.LegacyUserToGrpc(user)

		assert.Equal(t, expectedResult, result)
	})
}
