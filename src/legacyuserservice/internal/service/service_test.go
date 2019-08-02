package legacy

import (
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	legacyusermocks "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/repository/mocks"
	usermocks "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"testing"
)

const (
	userId   = 1
	email    = "email"
	password = "password"
)

func TestService_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := legacyusermocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)

	service, err := NewService(repository, userService)
	require.NoError(t, err)

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		repository.EXPECT().GetUserById(userId).Return(nil, e)

		user, err := service.GetUserById(userId)
		require.Error(t, err)

		assert.Empty(t, user)
		assert.Equal(t, e, err)
	})

	t.Run("gets the user by id", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		result, err := service.GetUserById(userId)
		require.NoError(t, err)

		assert.Equal(t, user, result)
	})
}

func TestService_LegacyLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := legacyusermocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)

	service, err := NewService(repository, userService)
	require.NoError(t, err)

	hashedPassword := hashPassword(password)

	t.Run("returns login failed error if user cannot be found", func(t *testing.T) {
		repository.EXPECT().GetUserByEmailAndPassword(email, hashedPassword).Return(nil, model.ErrLegacyUserNotFound)

		user, err := service.LegacyLogin(email, password)
		require.Error(t, err)

		assert.Empty(t, user)
		assert.Equal(t, model.ErrLegacyLoginFailed, err)
	})

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		e := errors.New("error")

		repository.EXPECT().GetUserByEmailAndPassword(email, hashedPassword).Return(nil, e)

		user, err := service.LegacyLogin(email, password)
		require.Error(t, err)

		assert.Empty(t, user)
		assert.Equal(t, e, err)
	})

	t.Run("returns error if password is wrong", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
			Password:  "incorrect password",
		}

		repository.EXPECT().GetUserByEmailAndPassword(email, hashedPassword).Return(user, nil)

		result, err := service.LegacyLogin(email, password)
		require.Error(t, err)

		assert.Empty(t, result)
		assert.Equal(t, model.ErrLegacyLoginFailed, err)
	})

	t.Run("returns error if user already exists in userService", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
			Password:  hashedPassword,
		}

		repository.EXPECT().GetUserByEmailAndPassword(email, hashedPassword).Return(user, nil)
		userService.EXPECT().GetUserByEmail(email).Return(user, nil)

		result, err := service.LegacyLogin(email, password)
		require.Error(t, err)

		assert.Empty(t, result)
		assert.Equal(t, model.ErrLegacyLoginFailed, err)
	})

	t.Run("gets the user and checks the password", func(t *testing.T) {
		user := &model.User{
			FirstName: "first",
			Password:  hashedPassword,
		}

		repository.EXPECT().GetUserByEmailAndPassword(email, hashedPassword).Return(user, nil)
		userService.EXPECT().GetUserByEmail(email).Return(nil, errors.New("user exists"))

		result, err := service.LegacyLogin(email, password)
		require.NoError(t, err)

		assert.Equal(t, user, result)
	})
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}
