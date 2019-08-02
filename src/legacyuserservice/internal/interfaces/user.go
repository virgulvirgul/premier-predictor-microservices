//go:generate mockgen -destination=../user/mocks/mock_user.go -package=usermocks github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces UserService

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"

type UserService interface {
	GetUserByEmail(email string) (*model.User, error)
}
