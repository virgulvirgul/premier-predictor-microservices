package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type ReadReceipt struct {
	SenderId string
	ChatId string
	MessageId string
	DateTime time.Time
}

func ReceiptFromGrpc(receipt *gen.ReadReceipt) ReadReceipt {
	t, err := ptypes.Timestamp(receipt.DateTime)
	if err != nil {
		t = time.Now()
	}

	return ReadReceipt{
		SenderId: receipt.UserId,
		ChatId: receipt.ChatId,
		MessageId: receipt.MessageId,
		DateTime: t,
	}
}