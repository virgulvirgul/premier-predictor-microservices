//go:generate mockgen -destination=../service/mocks/mock_service.go -package=predictionmocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces Service

package interfaces

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type Service interface {
	GetFixturesWithPredictions(id string) ([]model.FixturePrediction, error)
	GetPredictorData(id string) (*model.PredictorData, error)
	GetUsersPastPredictions(id string) (*model.PredictionSummary, error)
	UpdatePredictions(predictions []common.Prediction) error
	GetPrediction(userId, matchId string) (*common.Prediction, error)
	GetMatchPredictionSummary(id string) (*common.MatchPredictionSummary, error)
}
