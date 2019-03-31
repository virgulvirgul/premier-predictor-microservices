package chat

import (
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/repository"
)

type Service interface {
	UpdateReadMessage(readReceipt model.ReadReceipt) error
}

type service struct {
	repository chat.Repository
}

func NewService(repository chat.Repository) (Service, error) {
	return &service{
		repository: repository,
	}, nil
}

func (s *service) UpdateReadMessage(readReceipt model.ReadReceipt) error {
	return nil
}
