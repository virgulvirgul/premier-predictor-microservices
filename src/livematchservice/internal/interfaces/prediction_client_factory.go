//go:generate mockgen -destination=../factory/mocks/mock_prediction_client_factory.go -package=factorymocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces PredictionClientFactory

package interfaces

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type PredictionClientFactory interface {
	NewPredictionClient() (model.PredictionServiceClient, error)
	CloseConnection() error
}
