package chat

import (
	"context"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	readReceipt := model.ReadReceipt{}

	t.Run("Updates read receipt in db", func(t *testing.T) {
		repo.EXPECT().SaveReadReceipt(readReceipt).Return(nil)

		err := service.UpdateReadMessage(readReceipt)
		assert.NoError(t, err)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().SaveReadReceipt(readReceipt).Return(e)

		err := service.UpdateReadMessage(readReceipt)
		assert.Error(t, err)
		assert.Equal(t, e, err)
	})
}

func TestService_CreateChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"
	userId := "2"

	t.Run("Adds user to chat", func(t *testing.T) {
		repo.EXPECT().JoinChat(chatId, userId).Return(nil)

		err := service.JoinChat(chatId, userId)
		assert.NoError(t, err)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().JoinChat(chatId, userId).Return(e)

		err := service.JoinChat(chatId, userId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
	})
}

func TestService_JoinChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"
	userId := "2"

	t.Run("Adds user to chat", func(t *testing.T) {
		repo.EXPECT().JoinChat(chatId, userId).Return(nil)

		err := service.JoinChat(chatId, userId)
		assert.NoError(t, err)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().JoinChat(chatId, userId).Return(e)

		err := service.JoinChat(chatId, userId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
	})
}

func TestService_LeaveChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"
	userId := "2"

	t.Run("Removes user from chat", func(t *testing.T) {
		repo.EXPECT().LeaveChat(chatId, userId).Return(nil)

		err := service.LeaveChat(chatId, userId)
		assert.NoError(t, err)
	})

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().LeaveChat(chatId, userId).Return(e)

		err := service.LeaveChat(chatId, userId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
	})
}

func TestService_GetLatestMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"

	t.Run("Retrieves latest messages from db", func(t *testing.T) {
		var m []model.Message
		repo.EXPECT().GetLatestMessages(chatId).Return(m, nil)

		messages, err := service.GetLatestMessages(chatId)
		assert.NoError(t, err)
		assert.Equal(t, m, messages)
	})

	t.Run("Returns error if there is a problem retrieving messages", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().GetLatestMessages(chatId).Return(nil, e)

		messages, err := service.GetLatestMessages(chatId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, messages)
	})
}

func TestService_GetPreviousMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"
	messageId := "2"

	t.Run("Retrieves previous messages from db", func(t *testing.T) {
		var m []model.Message
		repo.EXPECT().GetPreviousMessages(chatId, messageId).Return(m, nil)

		messages, err := service.GetPreviousMessages(chatId, messageId)
		assert.NoError(t, err)
		assert.Equal(t, m, messages)
	})

	t.Run("Returns error if there is a problem retrieving messages", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().GetPreviousMessages(chatId, messageId).Return(nil, e)

		messages, err := service.GetPreviousMessages(chatId, messageId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, messages)
	})
}

func TestService_GetRecentMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := chatmocks.NewMockRepository(ctrl)
	notifier := notificationmocks.NewMockNotifier(ctrl)

	service, err := NewService(repo, notifier)
	require.NoError(t, err)

	chatId := "1"
	messageId := "2"

	t.Run("Retrieves recent messages from db", func(t *testing.T) {
		var m []model.Message
		repo.EXPECT().GetRecentMessages(chatId, messageId).Return(m, nil)

		messages, err := service.GetRecentMessages(chatId, messageId)
		assert.NoError(t, err)
		assert.Equal(t, m, messages)
	})

	t.Run("Returns error if there is a problem retrieving messages", func(t *testing.T) {
		e := errors.New("db error")
		repo.EXPECT().GetRecentMessages(chatId, messageId).Return(nil, e)

		messages, err := service.GetRecentMessages(chatId, messageId)
		assert.Error(t, err)
		assert.Equal(t, e, err)
		assert.Nil(t, messages)
	})
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

	ctx := context.Background()

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

		notifier.EXPECT().Send(ctx, notification, userId).Return(nil)

		time.Sleep(200 * time.Millisecond)

		id, err := service.SendMessage(ctx, msg)
		require.NoError(t, err)

		assert.Equal(t, messageId, id)
	})

	t.Run("Returns an error if there is a problem with saving message to db", func(t *testing.T) {
		e := errors.New("error saving to db")

		repo.EXPECT().SaveMessage(msg).Return("", e)
		repo.EXPECT().SaveReadReceipt(gomock.Any()).Times(0)
		notifier.EXPECT().Send(ctx, gomock.Any()).Times(0)

		time.Sleep(200 * time.Millisecond)

		id, err := service.SendMessage(ctx, msg)
		require.Error(t, err)
		assert.Equal(t, "", id)
		assert.Equal(t, e, err)
	})
}
