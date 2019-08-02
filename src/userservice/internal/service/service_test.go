package user

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const (
	userId           = "1"
	emailAddress     = "example@test.com"
	oldPassword      = "old password"
	newPassword      = "new password"
	newValidPassword = "Qwerty123"
)

var (
	e = errors.New("error")
)

func TestService_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().GetUserById(userId).Return(nil, e)

		result, err := service.GetUserById(userId)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("Gets user from db", func(t *testing.T) {
		user := &model.User{
			Id: userId,
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		result, err := service.GetUserById(userId)

		require.NoError(t, err)
		assert.Equal(t, user, result)
	})
}

func TestService_UpdateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("returns error if the email address is not valid", func(t *testing.T) {
		userInfo := model.UserInfo{
			Email: "invalid email address",
		}

		err := service.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, invalidEmail, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if the email address is already taken by a different user", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:    userId,
			Email: emailAddress,
		}

		repository.EXPECT().IsEmailTakenByADifferentUser(userId, emailAddress).Return(true)

		err := service.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, emailAlreadyTaken, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if the first name is blank", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:    userId,
			Email: emailAddress,
		}

		repository.EXPECT().IsEmailTakenByADifferentUser(userId, emailAddress).Return(false)

		err := service.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, firstNameIsBlank, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if the surname is blank", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
		}

		repository.EXPECT().IsEmailTakenByADifferentUser(userId, emailAddress).Return(false)

		err := service.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, surnameIsBlank, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if details cannot be updated", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
			Surname:   "surname",
		}

		repository.EXPECT().IsEmailTakenByADifferentUser(userId, emailAddress).Return(false)
		repository.EXPECT().UpdateUserInfo(userInfo).Return(e)

		err := service.UpdateUserInfo(userInfo)
		require.Error(t, err)

		assert.Equal(t, e, err)
	})

	t.Run("returns nil if details are updated successfully", func(t *testing.T) {
		userInfo := model.UserInfo{
			Id:        userId,
			Email:     emailAddress,
			FirstName: "first name",
			Surname:   "surname",
		}

		repository.EXPECT().IsEmailTakenByADifferentUser(userId, emailAddress).Return(false)
		repository.EXPECT().UpdateUserInfo(userInfo).Return(nil)

		err := service.UpdateUserInfo(userInfo)
		require.NoError(t, err)

		assert.Nil(t, err)
	})
}

func TestService_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	//newHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newValidPassword), 10)

	t.Run("returns error if the user cannot be retrieved", func(t *testing.T) {
		repository.EXPECT().GetUserById(userId).Return(nil, e)

		err := service.UpdatePassword(model.UpdatePassword{Id: userId})
		require.Error(t, err)

		assert.Equal(t, e, err)
	})

	t.Run("returns error if old password does not match what is currently stored", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:          userId,
			OldPassword: "different old password",
		}

		err := service.UpdatePassword(updatePassword)
		require.Error(t, err)

		assert.Equal(t, oldPasswordDoesNotMatch, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if new password does not match the confirmation", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newPassword,
			ConfirmPassword: "different confirmation password",
		}

		err := service.UpdatePassword(updatePassword)
		require.Error(t, err)

		assert.Equal(t, confirmationDoesNotMatch, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if new password is not valid", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newPassword,
			ConfirmPassword: newPassword,
		}

		err := service.UpdatePassword(updatePassword)
		require.Error(t, err)

		assert.Equal(t, invalidPassword, util.GetErrorMessage(err))
		assert.Equal(t, common.ErrInvalidRequestData, errors.Cause(err))
	})

	t.Run("returns error if password cannot be updated", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newValidPassword,
			ConfirmPassword: newValidPassword,
		}

		repository.EXPECT().UpdatePassword(userId, gomock.Any()).Return(e)

		err := service.UpdatePassword(updatePassword)
		require.Error(t, err)

		assert.Equal(t, e, err)
	})

	t.Run("returns nil if password is updated successfully", func(t *testing.T) {
		user := &model.User{
			Id:       userId,
			Password: hashPassword(oldPassword),
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		updatePassword := model.UpdatePassword{
			Id:              userId,
			OldPassword:     oldPassword,
			NewPassword:     newValidPassword,
			ConfirmPassword: newValidPassword,
		}

		repository.EXPECT().UpdatePassword(userId, gomock.Any()).Return(nil)

		err := service.UpdatePassword(updatePassword)
		require.NoError(t, err)

		assert.Nil(t, err)
	})
}

func TestService_GetUserScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("returns error if user cannot be retrieved", func(t *testing.T) {
		repository.EXPECT().GetUserById(userId).Return(nil, e)

		result, err := service.GetUserScore(userId)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Equal(t, 0, result)
	})

	t.Run("returns specified user's score", func(t *testing.T) {
		const score = 1234

		user := &model.User{
			Id:    userId,
			Score: score,
		}

		repository.EXPECT().GetUserById(userId).Return(user, nil)

		result, err := service.GetUserScore(userId)

		require.NoError(t, err)
		assert.Equal(t, score, result)
	})
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func TestService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		repository.EXPECT().GetAllUsers().Return(nil, e)

		result, err := service.GetAllUsers()

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: userId,
			},
		}

		repository.EXPECT().GetAllUsers().Return(users, nil)

		result, err := service.GetAllUsers()

		require.NoError(t, err)

		assert.Equal(t, users, result)
	})
}

func TestService_GetAllUsersByIds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	ids := []string{userId}

	t.Run("returns error if users cannot be retrieved", func(t *testing.T) {
		repository.EXPECT().GetAllUsersByIds(ids).Return(nil, e)

		result, err := service.GetAllUsersByIds(ids)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("returns all users", func(t *testing.T) {
		users := []*model.User{
			{
				Id: userId,
			},
		}

		repository.EXPECT().GetAllUsersByIds(ids).Return(users, nil)

		result, err := service.GetAllUsersByIds(ids)

		require.NoError(t, err)

		assert.Equal(t, users, result)
	})
}

func TestService_GetRankForGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	ids := []string{userId}

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().GetRankForGroup(userId, ids).Return(int64(0), e)

		result, err := service.GetRankForGroup(userId, ids)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets rank for group", func(t *testing.T) {
		repository.EXPECT().GetRankForGroup(userId, ids).Return(int64(1), nil)

		result, err := service.GetRankForGroup(userId, ids)

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetOverallRank(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().GetOverallRank(userId).Return(int64(0), e)

		result, err := service.GetOverallRank(userId)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets overall rank", func(t *testing.T) {
		repository.EXPECT().GetOverallRank(userId).Return(int64(1), nil)

		result, err := service.GetOverallRank(userId)

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetUserCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().GetUserCount().Return(int64(0), e)

		result, err := service.GetUserCount()

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets user count", func(t *testing.T) {
		repository.EXPECT().GetUserCount().Return(int64(1), nil)

		result, err := service.GetUserCount()

		require.NoError(t, err)
		assert.Equal(t, int64(1), result)
	})
}

func TestService_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := usermocks.NewMockRepository(ctrl)

	service, err := NewService(repository)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().GetUserByEmail(emailAddress).Return(nil, e)

		result, err := service.GetUserByEmail(emailAddress)

		require.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, result)
	})

	t.Run("Gets user from db", func(t *testing.T) {
		user := &model.User{
			Id:    userId,
			Email: emailAddress,
		}

		repository.EXPECT().GetUserByEmail(emailAddress).Return(user, nil)

		result, err := service.GetUserByEmail(emailAddress)

		require.NoError(t, err)
		assert.Equal(t, user, result)
	})
}
