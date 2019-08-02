//go:generate mockgen -destination=./mocks/mock_user_service_client.go -package=usermocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen UserServiceClient

package user

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	usermocks "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	email = "📧"
)

func TestUserService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userClient := usermocks.NewMockUserServiceClient(ctrl)

	userService, err := NewUserService(userClient)
	require.NoError(t, err)

	user := &gen.User{
		Id:        "1",
		FirstName: "name",
	}

	req := gen.EmailRequest{
		Email: email,
	}

	expectedUser := &model.User{
		FirstName: "name",
	}

	t.Run("returns error if there is a problem retrieving user", func(t *testing.T) {
		userClient.EXPECT().GetUserByEmail(context.Background(), req).Return(nil, errors.New("error"))

		result, err := userService.GetUserByEmail(email)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("gets user from UserService for email specified", func(t *testing.T) {
		userClient.EXPECT().GetUserByEmail(context.Background(), req).Return(user, nil)

		result, err := userService.GetUserByEmail(email)
		require.NoError(t, err)

		assert.Equal(t, expectedUser, result)
	})
}
