package chat

import (
	"errors"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/repository/mocks"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/notification/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestService_UpdateReadMessage(t *testing.T) {

}

func TestService_CreateChat(t *testing.T) {

}

func TestService_JoinChat(t *testing.T) {

}

func TestService_LeaveChat(t *testing.T) {

}

func TestService_GetLatestMessages(t *testing.T) {

}

func TestService_GetPreviousMessages(t *testing.T) {

}

func TestService_SendMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	senderId := "1"
	chatId := "2"
	messageId := "3"
	userId := "4"
	dateTime := time.Now()
	text := "message text"

	msg := model.Message{
		SenderId: senderId,
		ChatId:   chatId,
		Text:     text,
		DateTime: dateTime,
	}

	t.Run("Saves message to db, sends notification and returns id", func(t *testing.T) {
		repo.EXPECT().SaveMessage(msg).Return(messageId, nil)

		read := model.ReadReceipt{
			SenderId: senderId,
			ChatId: chatId,
			MessageId: messageId,
			DateTime: dateTime,
		}
		repo.EXPECT().SaveReadReceipt(read).Return(nil)

		chat := &model.Chat{
			Users: []model.ChatUser{
				{
					Id: senderId,
				},
				{
					Id: userId,
				},
			},
		}

		repo.EXPECT().GetChatById(chatId).Return(chat, nil)

		notification := m.Notification{
			Title:   newMessage,
			Message: text,
		}

		notifier.EXPECT().Send(notification, userId).Return(nil)

		time.Sleep(200 * time.Millisecond)

		id, err := service.SendMessage(msg)
		require.NoError(t, err)

		assert.Equal(t, messageId, id)
	})

	t.Run("Returns an error if there is a problem with saving message to db", func(t *testing.T) {
		e := errors.New("error saving to db")

		repo.EXPECT().SaveMessage(msg).Return("", e)
		repo.EXPECT().SaveReadReceipt(gomock.Any()).Times(0)
		notifier.EXPECT().Send(gomock.Any()).Times(0)

		time.Sleep(200 * time.Millisecond)

		id, err := service.SendMessage(msg)
		require.Error(t, err)
		assert.Equal(t, "", id)
		assert.Equal(t, e, err)
	})
}
