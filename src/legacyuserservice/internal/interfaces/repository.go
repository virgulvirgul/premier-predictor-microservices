//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=legacyusermocks github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"

type Repository interface {
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
