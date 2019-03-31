package chat

import (
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type chatEntity struct {
	Id    string     `bson:"_id,omitempty"`
	Users []chatUser `bson:"users,omitempty"`
}

type chatUser struct {
	Id              string              `bson:"id,omitempty"`
	LastReadMessage *primitive.ObjectID `bson:"lastReadMessage,omitempty"`
}

type message struct {
	Id       *primitive.ObjectID `bson:"_id,omitempty"`
	SenderId string              `bson:"senderId,omitempty"`
	ChatId   string              `bson:"chatId,omitempty"`
	Text     string              `bson:"text,omitempty"`
	DateTime time.Time           `bson:"dateTime,omitempty"`
}

func toChat(chat chatEntity) *model.Chat {
	c := model.Chat{
		Id:    chat.Id,
		Users: []model.ChatUser{},
	}

	for i := range chat.Users {
		u := model.ChatUser{
			Id:              chat.Users[i].Id,
			LastReadMessage: objectIdToString(chat.Users[i].LastReadMessage),
		}

		c.Users = append(c.Users, u)
	}

	return &c
}

func objectIdToString(objectId *primitive.ObjectID) string {
	if objectId == nil || objectId.IsZero() {
		return ""
	}

	return objectId.Hex()
}

func fromChat(chat model.Chat) *chatEntity {
	c := chatEntity{
		Id:    chat.Id,
		Users: []chatUser{},
	}

	for i := range chat.Users {
		messageId, err := primitive.ObjectIDFromHex(chat.Users[i].LastReadMessage)
		if err != nil {
			return nil
		}

		u := chatUser{
			Id:              chat.Users[i].Id,
			LastReadMessage: &messageId,
		}

		c.Users = append(c.Users, u)
	}

	return &c
}

func toMessage(msg message) *model.Message {
	return &model.Message{
		Id:       msg.Id.Hex(),
		SenderId: msg.SenderId,
		ChatId:   msg.ChatId,
		Text:     msg.Text,
		DateTime: msg.DateTime,
	}
}

func fromMessage(msg model.Message) (*message, error) {
	m := message{
		SenderId: msg.SenderId,
		ChatId:   msg.ChatId,
		Text:     msg.Text,
		DateTime: msg.DateTime,
	}

	if msg.Id == "" {
		id := primitive.NewObjectID()
		m.Id = &id

		return &m, nil
	}

	messageId, err := primitive.ObjectIDFromHex(msg.Id)
	if err != nil {
		return nil, err
	}

	m.Id = &messageId

	return &m, nil
}
