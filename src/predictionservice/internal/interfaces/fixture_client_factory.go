//go:generate mockgen -destination=../factory/mocks/mock_fixture_client_factory.go -package=factorymocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces FixtureClientFactory

package interfaces

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type FixtureClientFactory interface {
	NewFixtureClient() (model.FixtureServiceClient, error)
	CloseConnection() error
}
