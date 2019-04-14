package chat

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strconv"
	"testing"
	"time"
)

const (
	chatId = "chatId"
	userId = "userId"
)

func Test_Repository(t *testing.T) {
	err := os.Setenv("MONGO_PORT", "27017")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	assert.NoError(t, err)

	repository, err := NewRepository()
	assert.NoError(t, err)

	createChat := func() {
		err := repository.CreateChat(chatId, userId)
		assert.NoError(t, err)
	}

	cleanupDb := func() {
		_, _ = repository.
			client.
			Database(db).
			Collection(collection).
			DeleteOne(
				context.Background(),
				bson.M{
					"_id": chatId,
				},
			)
	}

	t.Run("CreateChat & GetChatById", func(t *testing.T) {
		createChat()
		defer cleanupDb()

		c, err := repository.GetChatById(chatId)

		assert.NoError(t, err)
		assert.Equal(t, chatId, c.Id)
		assert.Equal(t, userId, c.Users[0].Id)
		assert.Equal(t, "", c.Users[0].LastReadMessage)
	})

	t.Run("LeaveChat", func(t *testing.T) {
		createChat()
		defer cleanupDb()

		err = repository.LeaveChat(chatId, userId)
		assert.NoError(t, err)

		c, err := repository.GetChatById(chatId)

		assert.NoError(t, err)
		assert.Equal(t, chatId, c.Id)
		assert.Equal(t, 0, len(c.Users))
	})

	t.Run("JoinChat", func(t *testing.T) {
		createChat()
		defer cleanupDb()

		const userId2 = "user 2"
		err = repository.JoinChat(chatId, userId2)
		assert.NoError(t, err)

		c, err := repository.GetChatById(chatId)

		assert.NoError(t, err)
		assert.Equal(t, chatId, c.Id)
		assert.Equal(t, 2, len(c.Users))
		assert.Equal(t, userId, c.Users[0].Id)
		assert.Equal(t, "", c.Users[0].LastReadMessage)
		assert.Equal(t, userId2, c.Users[1].Id)
		assert.Equal(t, "", c.Users[1].LastReadMessage)
	})

	t.Run("SaveReadReceipt", func(t *testing.T) {
		t.Run("Updates users last read message", func(t *testing.T) {
			createChat()
			defer cleanupDb()

			messageId := primitive.NewObjectID()

			now := time.Now().Round(time.Second)

			readReceipt := model.ReadReceipt{
				ChatId: chatId,
				SenderId: userId,
				MessageId: messageId.Hex(),
				DateTime: now,
			}

			err := repository.SaveReadReceipt(readReceipt)
			assert.NoError(t, err)

			c, err := repository.GetChatById(chatId)

			assert.NoError(t, err)
			assert.Equal(t, chatId, c.Id)
			assert.Equal(t, userId, c.Users[0].Id)
			assert.Equal(t, messageId.Hex(), c.Users[0].LastReadMessage)
			assert.Equal(t, now, c.Users[0].ReadTime)
		})

		t.Run("Doesn't update other user's", func(t *testing.T) {
			createChat()
			defer cleanupDb()

			const userId2 = "user 2"
			err = repository.JoinChat(chatId, userId2)
			assert.NoError(t, err)

			messageId := primitive.NewObjectID()

			now := time.Now().Round(time.Second)

			readReceipt := model.ReadReceipt{
				ChatId: chatId,
				SenderId: userId,
				MessageId: messageId.Hex(),
				DateTime: now,
			}

			err := repository.SaveReadReceipt(readReceipt)
			assert.NoError(t, err)

			c, err := repository.GetChatById(chatId)

			assert.NoError(t, err)
			assert.Equal(t, chatId, c.Id)
			assert.Equal(t, userId, c.Users[0].Id)
			assert.Equal(t, messageId.Hex(), c.Users[0].LastReadMessage)
			assert.Equal(t, now, c.Users[0].ReadTime)
			assert.Equal(t, userId2, c.Users[1].Id)
			assert.Equal(t, "", c.Users[1].LastReadMessage)
		})
	})

	cleanupChat := func() {
		_, _ = repository.
			client.
			Database(db).
			Collection(repository.getChatCollection(chatId)).
			DeleteMany(
				context.Background(),
				bson.M{},
			)

		_ = repository.
			client.
			Database(db).
			Collection(repository.getChatCollection(chatId)).
			Drop(context.Background())
	}

	t.Run("SendMessage", func(t *testing.T) {
		t.Run("returns error if message id is not valid", func(t *testing.T) {
			msg := model.Message{
				Id: "Invalid objectId",
			}
			id, err := repository.SaveMessage(msg)
			defer cleanupChat()

			assert.Error(t, err)
			assert.Equal(t, "", id)
		})

		t.Run("returns error if message already has id", func(t *testing.T) {
			msg := model.Message{
				Id: primitive.NewObjectID().Hex(),
			}
			id, err := repository.SaveMessage(msg)
			defer cleanupChat()

			assert.Error(t, err)
			assert.Equal(t, ErrMessageAlreadyExists, err)
			assert.Equal(t, "", id)
		})

		t.Run("inserts message and returns id as a string", func(t *testing.T) {
			const messageText = "Test message"

			msg := model.Message{
				SenderId: userId,
				ChatId:   chatId,
				Text:     messageText,
				DateTime: time.Now().Round(time.Second),
			}

			id, err := repository.SaveMessage(msg)
			assert.NoError(t, err)
			defer cleanupChat()

			var m message

			err = repository.client.
				Database(db).
				Collection(repository.getChatCollection(chatId)).
				FindOne(
					context.Background(),
					bson.M{
						"senderId": userId,
					},
				).
				Decode(&m)
			assert.NoError(t, err)

			result := toMessage(m)

			assert.Equal(t, id, result.Id)
			assert.Equal(t, userId, result.SenderId)
			assert.Equal(t, chatId, result.ChatId)
			assert.Equal(t, messageText, result.Text)
			assert.Equal(t, msg.DateTime, result.DateTime)
		})
	})

	const numMessages = 50
	t.Run("GetLatestMessages", func(t *testing.T) {
		t.Run("Gets last 50 messages", func(t *testing.T) {
			var ids []string
			defer cleanupChat()

			for i := 0; i < 60; i++ {
				id, err := repository.SaveMessage(model.Message{ChatId: chatId, Text: strconv.Itoa(i)})
				assert.NoError(t, err)

				ids = append(ids, id)
			}

			messages, err := repository.GetLatestMessages(chatId)
			assert.NoError(t, err)
			assert.NotNil(t, messages)
			assert.Equal(t, numMessages, len(messages))

			for i := range messages {
				assert.Equal(t, ids[i+10], messages[i].Id)
				assert.Equal(t, strconv.Itoa(i+10), messages[i].Text)
			}
		})

		t.Run("Gets all messages if there is not 50", func(t *testing.T) {
			var ids []string
			defer cleanupChat()

			for i := 0; i < 10; i++ {
				id, err := repository.SaveMessage(model.Message{ChatId: chatId, Text: strconv.Itoa(i)})
				assert.NoError(t, err)

				ids = append(ids, id)
			}

			messages, err := repository.GetLatestMessages(chatId)
			assert.NoError(t, err)
			assert.NotNil(t, messages)
			assert.Equal(t, 10, len(messages))

			for i := range messages {
				assert.Equal(t, ids[i], messages[i].Id)
				assert.Equal(t, strconv.Itoa(i), messages[i].Text)
			};
		})
	})

	t.Run("GetPreviousMessages", func(t *testing.T) {
		t.Run("Returns error if messageId is invalid", func(t *testing.T) {
			messages, err := repository.GetPreviousMessages(chatId, "")
			assert.Error(t, err)
			assert.Nil(t, messages)
		})

		t.Run("Gets previous 50 messages from specified id", func(t *testing.T) {
			const messageIndex = 100

			var ids []string
			defer cleanupChat()

			var messageId string

			for i := 0; i < 200; i++ {
				id, err := repository.SaveMessage(model.Message{ChatId: chatId, Text: strconv.Itoa(i)})
				assert.NoError(t, err)

				if i == messageIndex {
					messageId = id
				}

				ids = append(ids, id)
			}

			messages, err := repository.GetPreviousMessages(chatId, messageId)
			assert.NoError(t, err)
			assert.NotNil(t, messages)
			assert.Equal(t, numMessages, len(messages))

			for i := range messages {
				assert.Equal(t, ids[i+messageIndex-numMessages], messages[i].Id)
				assert.Equal(t, strconv.Itoa(i+messageIndex-numMessages), messages[i].Text)
			}
		})

		t.Run("Gets remaining messages if there is not 50 previous messages", func(t *testing.T) {
			const messageIndex = 20

			var ids []string
			defer cleanupChat()

			var messageId string

			for i := 0; i < 120; i++ {
				id, err := repository.SaveMessage(model.Message{ChatId: chatId, Text: strconv.Itoa(i)})
				assert.NoError(t, err)

				if i == messageIndex {
					messageId = id
				}

				ids = append(ids, id)
			}

			messages, err := repository.GetPreviousMessages(chatId, messageId)
			assert.NoError(t, err)
			assert.NotNil(t, messages)
			assert.Equal(t, messageIndex, len(messages))

			for i := range messages {
				assert.Equal(t, ids[i], messages[i].Id)
				assert.Equal(t, strconv.Itoa(i), messages[i].Text)
			}
		})
	})

	t.Run("GetPreviousMessages", func(t *testing.T) {
		t.Run("Returns error if messageId is invalid", func(t *testing.T) {
			messages, err := repository.GetRecentMessages(chatId, "")
			assert.Error(t, err)
			assert.Nil(t, messages)
		})

		t.Run("Gets all latest messages after specified id", func(t *testing.T) {
			const messageIndex = 100

			var ids []string
			defer cleanupChat()

			var messageId string

			for i := 0; i < 200; i++ {
				id, err := repository.SaveMessage(model.Message{ChatId: chatId, Text: strconv.Itoa(i)})
				assert.NoError(t, err)

				if i == messageIndex {
					messageId = id
				}

				ids = append(ids, id)
			}

			messages, err := repository.GetRecentMessages(chatId, messageId)
			assert.NoError(t, err)
			assert.NotNil(t, messages)

			for i := range messages {
				assert.Equal(t, ids[i+messageIndex+1], messages[i].Id)
				assert.Equal(t, strconv.Itoa(i+messageIndex+1), messages[i].Text)
			}
		})
	})
}
