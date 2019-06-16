//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=predictionmocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/common/model"

type Repository interface {
	GetPrediction(userId, matchId string) (*model.Prediction, error)
	GetPredictionsByUserId(id string) ([]model.Prediction, error)
	UpdatePredictions(predictions []model.Prediction) error
	GetMatchPredictionSummary(id string) (int, int, int, error)
}
