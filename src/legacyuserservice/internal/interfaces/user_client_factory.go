//go:generate mockgen -destination=../factory/mocks/mock_user_client_factory.go -package=factorymocks github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces UserClientFactory

package interfaces

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type UserClientFactory interface {
	NewUserClient() (model.UserServiceClient, error)
	CloseConnection() error
}
