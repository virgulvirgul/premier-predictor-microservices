package prediction

import (
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
)

type service struct {
	repository interfaces.Repository
}

func NewService(repository interfaces.Repository) (interfaces.Service, error) {
	return &service{
		repository: repository,
	}, nil
}
