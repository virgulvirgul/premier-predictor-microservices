package model

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type Chat struct {
	Id    string
	Users []ChatUser
}

type ChatUser struct {
	Id              string
	LastReadMessage string
}

type Message struct {
	Id       string
	SenderId string
	ChatId   string
	Text     string
	DateTime time.Time
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

func MessageToGrpcMessage(msg Message) (*chat.Message, error) {
	t, err := ptypes.TimestampProto(msg.DateTime)
	if err != nil {
		return nil, err
	}

	return &chat.Message{
		MessageId: msg.Id,
		SenderId: msg.SenderId,
		Type: chat.Message_MESSAGE,
		Text:     msg.Text,
		DateTime: t,
	}, nil
}
