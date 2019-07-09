package user

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

const (
	invalidEmail      = "Invalid Email Address"
	emailAlreadyTaken = "Email Address is Already Taken"
	firstNameIsBlank  = "First Name Cannot be Blank"
	surnameIsBlank    = "Surname Cannot be Blank"
	oldPasswordDoesNotMatch = "Old Password does not Match"
	confirmationDoesNotMatch = "New Password does not Match Confirmation"
	invalidPassword = "Invalid Password"

	emailRegex = "^([_a-zA-Z0-9-]+(\\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*(\\.[a-zA-Z]{1,6}))?$"
)

type service struct {
	repository interfaces.Repository
}

func NewService(repository interfaces.Repository) (interfaces.Service, error) {
	return &service{
		repository: repository,
	}, nil
}

func (s *service) GetUser(id string) (*model.User, error) {
	return s.repository.GetUserById(id)
}

func (s *service) UpdateUserInfo(userInfo model.UserInfo) error {
	userInfo.Email = strings.ToLower(userInfo.Email)

	switch {
	case !regexp.MustCompile(emailRegex).MatchString(userInfo.Email):
		return errors.Wrap(common.ErrInvalidRequestData, invalidEmail)
	case s.repository.IsEmailTakenByADifferentUser(userInfo.Id, userInfo.Email):
		return errors.Wrap(common.ErrInvalidRequestData, emailAlreadyTaken)
	case userInfo.FirstName == "":
		return errors.Wrap(common.ErrInvalidRequestData, firstNameIsBlank)
	case userInfo.Surname == "":
		return errors.Wrap(common.ErrInvalidRequestData, surnameIsBlank)
	}

	return s.repository.UpdateUserInfo(userInfo)
}

func (s *service) UpdatePassword(updatePassword model.UpdatePassword) error {
	user, err := s.repository.GetUserById(updatePassword.Id)
	if err != nil {
		return err
	}

	switch {
	case bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePassword.OldPassword)) != nil:
		return errors.Wrap(common.ErrInvalidRequestData, oldPasswordDoesNotMatch)
	case updatePassword.NewPassword != updatePassword.ConfirmPassword:
		return errors.Wrap(common.ErrInvalidRequestData, confirmationDoesNotMatch)
	case !util.VerifyPassword(updatePassword.NewPassword):
		return errors.Wrap(common.ErrInvalidRequestData, invalidPassword)
	}

	return s.repository.UpdatePassword(updatePassword.Id, updatePassword.NewPassword)
}

func (s *service) GetUserScore(id string) (int, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return 0, err
	}

	return user.Score, nil
}

func (s *service) GetAllUsers() ([]*model.User, error) {
	return s.repository.GetAllUsers()
}

func (s *service) GetAllUsersByIds(ids []string) ([]*model.User, error) {
	return s.repository.GetAllUsersByIds(ids)
}

func (s *service) GetRankForGroup(id string, ids []string) (int64, error) {
	return s.repository.GetRankForGroup(id, ids)
}

func (s *service) GetOverallRank(id string) (int64, error) {
	return s.repository.GetOverallRank(id)
}

func (s *service) GetUserCount() (int64, error) {
	return s.repository.GetUserCount()
}
