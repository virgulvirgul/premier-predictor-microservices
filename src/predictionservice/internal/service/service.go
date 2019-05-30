package prediction

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type service struct {
	repository     interfaces.Repository
	fixtureService interfaces.FixtureService
}

func NewService(repository interfaces.Repository, fixtureService interfaces.FixtureService) (interfaces.Service, error) {
	return &service{
		repository:     repository,
		fixtureService: fixtureService,
	}, nil
}

func (s *service) GetFixturesWithPredictions(id string) ([]*model.FixturePrediction, error) {
	panic("implement me")
}

func (s *service) GetPredictorData(id string) (*model.PredictorData, error) {
	panic("implement me")
}

func (s *service) GetUsersPastPredictions(id string) ([]*model.FixturePrediction, error) {
	panic("implement me")
}

func (s *service) UpdatePredictions(predictions []common.Prediction) error {
	panic("implement me")
}

func (s *service) GetPrediction(userId, matchId string) (*common.Prediction, error) {
	panic("implement me")
}
