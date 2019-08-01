package legacy

import (
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository interfaces.Repository
}

func NewService(repository interfaces.Repository) (interfaces.Service, error) {
	return &service{
		repository: repository,
	}, nil
}

func (s *service) GetUserById(id int) (*model.User, error) {
	return s.repository.GetUserById(id)
}

func (s *service) LegacyLogin(email, password string) (*model.User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err == model.ErrLegacyUserNotFound {
		return nil, model.ErrLegacyLoginFailed
	}

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, model.ErrLegacyLoginFailed
	}

	return user, nil
}
