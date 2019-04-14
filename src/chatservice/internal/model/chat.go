package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	_ "github.com/golang/protobuf/proto"
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
	ReadTime        time.Time
}

type Message struct {
	Id       string
	SenderId string
	ChatId   string
	Text     string
	DateTime time.Time
}

func MessageFromGrpc(req *gen.SendRequest) Message {
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

func MessageToGrpcMessage(msg Message) (*gen.Message, error) {
	t, err := ptypes.TimestampProto(msg.DateTime)
	if err != nil {
		return nil, err
	}

	return &gen.Message{
		MessageId: msg.Id,
		SenderId:  msg.SenderId,
		Type:      gen.Message_MESSAGE,
		Text:      msg.Text,
		DateTime:  t,
	}, nil
}
