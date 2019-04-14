package chat

import (
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"log"
)

type service struct {
	repository interfaces.Repository
	notifier   common.Notifier
}

const newMessage = "New Message"

func NewService(repository interfaces.Repository, notifier common.Notifier) (interfaces.Service, error) {
	return &service{
		repository: repository,
		notifier: notifier,
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
	id, err := s.repository.SaveMessage(message)
	if err != nil {
		return "", err
	}

	go func() {
		readReceipt := model.ReadReceipt{
			SenderId: message.SenderId,
			MessageId: id,
			ChatId: message.ChatId,
			DateTime: message.DateTime,
		}

		err := s.repository.SaveReadReceipt(readReceipt)
		if err != nil {
			log.Println(err)
		}

		err = s.sendNotifications(message)
		if err != nil {
			log.Println(err)
		}
	}()

	return id, nil
}

func (s *service) sendNotifications(message model.Message) error {
	chat, _ := s.repository.GetChatById(message.ChatId)

	var userIds []string
	for _, u := range chat.Users {
		if u.Id != message.SenderId {
			userIds = append(userIds, u.Id)
		}
	}

	notification := m.Notification{
		Title:   newMessage,
		Message: message.Text,
	}

	return s.notifier.Send(notification, userIds...)
}
