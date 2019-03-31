package model

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type Chat struct {
	Id       string     `bson:"_id"`
	Users    []ChatUser `bson:"users"`
	Messages []Message  `bson:"messages"`
}

type ChatUser struct {
	Id              string `bson:"id"`
	LastReadMessage int64  `bson:"lastReadMessage"`
}

type Message struct {
	Id       string    `bson:"id"`
	SenderId string    `bson:"senderId"`
	ChatId   string    `bson:"chatId"`
	Text     string    `bson:"text"`
	DateTime time.Time `bson:"dateTime"`
}

func MessageFromGrpc(req *chat.SendRequest) Message {
	t, err := ptypes.Timestamp(req.DateTime)
	if err != nil {
		t = time.Now()
	}

	return Message{
		SenderId: req.UserId,
		ChatId:   req.ChatId,
		Text:     req.Message,
		DateTime: t,
	}
}
