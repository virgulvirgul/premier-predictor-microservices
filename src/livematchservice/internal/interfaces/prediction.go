//go:generate mockgen -destination=../prediction/mocks/mock_prediction_client.go -package=predictionmocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces Predictor

package interfaces

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	. "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
)

type Predictor interface {
	GetPrediction(ctx context.Context, req PredictionRequest) (*model.Prediction, error)
	GetPredictionSummary(ctx context.Context, matchId string) (*model.MatchPredictionSummary, error)
}
