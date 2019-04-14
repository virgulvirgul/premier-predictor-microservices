//go:generate mockgen -destination=../service/mocks/mock_service.go -package=chatmocks github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/interfaces Service

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"

type Service interface {
	UpdateReadMessage(readReceipt model.ReadReceipt) error
	CreateChat(chatId, userId string) error
	JoinChat(chatId, userId string) error
	LeaveChat(chatId, userId string) error
	GetLatestMessages(chatId string) ([]model.Message, error)
	GetPreviousMessages(chatId, messageId string) ([]model.Message, error)
	SendMessage(message model.Message) (string, error)
}

