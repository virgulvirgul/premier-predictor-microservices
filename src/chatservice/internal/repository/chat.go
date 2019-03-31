package chat

import (
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"time"
)

type chatEntity struct {
	Id       string     `bson:"_id"`
	Users    []chatUser `bson:"users"`
	Messages []message  `bson:"messages"`
}

type chatUser struct {
	Id              string `bson:"id"`
	LastReadMessage int64  `bson:"lastReadMessage"`
}

type message struct {
	Id       string    `bson:"id"`
	SenderId string    `bson:"senderId"`
	ChatId   string    `bson:"chatId"`
	Text     string    `bson:"text"`
	DateTime time.Time `bson:"dateTime"`
}

func toChat(chat *chatEntity) *model.Chat {
	c := model.Chat{
		Id:       chat.Id,
		Users:    []model.ChatUser{},
		Messages: []model.Message{},
	}

	for i := range chat.Users {
		u := model.ChatUser{
			Id:              chat.Users[i].Id,
			LastReadMessage: chat.Users[i].LastReadMessage,
		}

		c.Users = append(c.Users, u)
	}

	for i := range chat.Messages {
		u := model.Message{
			Id:       chat.Messages[i].Id,
			SenderId: chat.Messages[i].SenderId,
			ChatId:   chat.Messages[i].ChatId,
			Text:     chat.Messages[i].Text,
			DateTime: chat.Messages[i].DateTime,
		}

		c.Messages = append(c.Messages, u)
	}

	return &c
}

func fromChat(chat *model.Chat) *chatEntity {
	c := chatEntity{
		Id:       chat.Id,
		Users:    []chatUser{},
		Messages: []message{},
	}

	for i := range chat.Users {
		u := chatUser{
			Id:              chat.Users[i].Id,
			LastReadMessage: chat.Users[i].LastReadMessage,
		}

		c.Users = append(c.Users, u)
	}

	for i := range chat.Messages {
		u := message{
			Id:       chat.Messages[i].Id,
			SenderId: chat.Messages[i].SenderId,
			ChatId:   chat.Messages[i].ChatId,
			Text:     chat.Messages[i].Text,
			DateTime: chat.Messages[i].DateTime,
		}

		c.Messages = append(c.Messages, u)
	}

	return &c
}
