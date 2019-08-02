package legacy

import (
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type service struct {
	repository  interfaces.Repository
	userService interfaces.UserService
}

func NewService(repository interfaces.Repository, userService interfaces.UserService) (interfaces.Service, error) {
	return &service{
		repository:  repository,
		userService: userService,
	}, nil
}

type legacyLoginResult struct {
	user *model.User
	err  error
}

func (s *service) GetUserById(id int) (*model.User, error) {
	return s.repository.GetUserById(id)
}

func (s *service) LegacyLogin(email, password string) (*model.User, error) {
	email = strings.ToLower(email)

	legacyLoginChan := s.doLegacyLogin(email, password)
	isExistingUserChan := s.isUserAlreadyRegistered(email)

	loginResult := <-legacyLoginChan
	if loginResult.err != nil {
		return nil, loginResult.err
	}

	isExistingUser := <-isExistingUserChan
	if isExistingUser {
		return nil, model.ErrLegacyUserNotFound
	}

	return loginResult.user, nil
}

func (s *service) doLegacyLogin(email string, password string) chan legacyLoginResult {
	legacyLoginChan := make(chan legacyLoginResult)
	go func() {
		user, err := s.repository.GetUserByEmail(email)
		if err == model.ErrLegacyUserNotFound {
			legacyLoginChan <- legacyLoginResult{
				user: nil,
				err:  err,
			}
			return
		}

		if err != nil {
			legacyLoginChan <- legacyLoginResult{
				user: nil,
				err:  err,
			}
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			legacyLoginChan <- legacyLoginResult{
				user: nil,
				err:  err,
			}
			return
		}

		legacyLoginChan <- legacyLoginResult{
			user: user,
			err:  nil,
		}
	}()

	return legacyLoginChan
}

func (s *service) isUserAlreadyRegistered(email string) chan bool {
	isExistingUserChan := make(chan bool)

	go func() {
		user, err := s.userService.GetUserByEmail(email)
		if err != nil || user == nil {
			isExistingUserChan <- false
		}

		isExistingUserChan <- true
	}()

	return isExistingUserChan
}
