//go:generate mockgen -destination=./mocks/mock_service.go -package=chatmocks github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/service Service

package chat

import (
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/repository"
)

type Service interface {
	UpdateReadMessage(readReceipt model.ReadReceipt) error
	CreateChat(chatId, userId string) error
	JoinChat(chatId, userId string) error
	LeaveChat(chatId, userId string) error
	GetLatestMessages(chatId string) ([]model.Message, error)
	GetPreviousMessages(chatId, messageId string) ([]model.Message, error)
	SendMessage(message model.Message) (string, error)
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

func (s *service) CreateChat(chatId, userId string) error {
	panic("implement me")
}

func (s *service) JoinChat(chatId, userId string) error {
	panic("implement me")
}

func (s *service) LeaveChat(chatId, userId string) error {
	panic("implement me")
}

func (s *service) GetLatestMessages(chatId string) ([]model.Message, error) {
	panic("implement me")
}

func (s *service) GetPreviousMessages(chatId, messageId string) ([]model.Message, error) {
	panic("implement me")
}

func (s *service) SendMessage(message model.Message) (string, error) {
	//Add to DB & update read receipt
	//Send notifications
	panic("implement me")
}
