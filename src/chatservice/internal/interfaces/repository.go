//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=chatmocks github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"

type Repository interface {
	GetChatById(chatId string) (*model.Chat, error)
	CreateChat(chatId, userId string) error
	JoinChat(chatId, userId string) error
	LeaveChat(chatId, userId string) error
	GetLatestMessages(chatId string) ([]model.Message, error)
	GetPreviousMessages(chatId, messageId string) ([]model.Message, error)
	SaveReadReceipt(readReceipt model.ReadReceipt) error
	SaveMessage(message model.Message) (string, error)
}
