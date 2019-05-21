package prediction

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type service struct {
	repository    interfaces.Repository
	fixtureClient gen.FixtureServiceClient
}

func NewService(repository interfaces.Repository, fixtureClient gen.FixtureServiceClient) (interfaces.Service, error) {
	return &service{
		repository:    repository,
		fixtureClient: fixtureClient,
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
