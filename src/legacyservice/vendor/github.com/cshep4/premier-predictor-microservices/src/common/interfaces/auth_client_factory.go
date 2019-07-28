//go:generate mockgen -destination=../factory/mocks/mock_auth_client_factory.go -package=factorymocks github.com/cshep4/premier-predictor-microservices/src/common/interfaces AuthClientFactory

package interfaces

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type AuthClientFactory interface {
	NewAuthClient() (model.AuthServiceClient, error)
	CloseConnection() error
}
