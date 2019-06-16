package prediction

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type fixturesResult struct {
	result []common.Fixture
	err    error
}
type predictionsResult struct {
	result []common.Prediction
	err    error
}
type fixturePredictionsResult struct {
	result []model.FixturePrediction
	err    error
}
type formResult struct {
	result map[string]model.TeamForm
	err    error
}
