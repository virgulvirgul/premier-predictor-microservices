//go:generate mockgen -destination=../service/mocks/mock_service.go -package=legacyusermocks github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces Service

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"

type Service interface {
	GetUserById(id int) (*model.User, error)
	LegacyLogin(email, password string) (*model.User, error)
}
