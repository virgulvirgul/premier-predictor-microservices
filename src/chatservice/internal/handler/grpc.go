package handler

import (
	"context"
	. "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/service"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"
	"strconv"
)

type chatServiceServer struct {
	service chat.Service
	msg     map[string]map[string]chan Message
}

func (c *chatServiceServer) CreateChat(ctx context.Context, req *AddRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (c *chatServiceServer) JoinChat(ctx context.Context, req *AddRequest) (*empty.Empty, error) {
	panic("implement me")
}

func (c *chatServiceServer) LeaveChat(ctx context.Context, req *AddRequest) (*empty.Empty, error) {
	panic("implement me")
}

func NewChatServiceServer(service chat.Service) (*chatServiceServer, error) {
	log.Print("Registered chatServiceServer handler")

	return &chatServiceServer{
		service: service,
		msg:     make(map[string]map[string]chan Message),
	}, nil
}

func (c *chatServiceServer) GetLatestMessages(ctx context.Context, req *LatestMessagesRequest) (*MessageList, error) {
	panic("implement me")
}

func (c *chatServiceServer) GetPreviousMessages(ctx context.Context, req *PreviousMessagesRequest) (*MessageList, error) {
	panic("implement me")
}

func (c *chatServiceServer) Send(ctx context.Context, req *SendRequest) (*empty.Empty, error) {
	//Add to DB & update read receipt
	//Send notifications
	//_ := model.MessageFromGrpc(req)

	//Send to other clients
	m := Message{
		SenderId: req.UserId,
		Type:     Message_MESSAGE,
		Text:     req.Message,
		DateTime: req.DateTime,
	}
	for u := range c.msg[req.ChatId] {
		if u != req.UserId {
			c.msg[req.ChatId][u] <- m
		}
	}

	return &empty.Empty{}, nil
}

func (c *chatServiceServer) Subscribe(stream ChatService_SubscribeServer) error {
	initialised := false
	chatId := make(chan string)
	userId := make(chan string)
	sendErr := make(chan error)

	//Sending messages to client
	go func() {
		cId := <-chatId
		uId := <-userId
		for {
			m := <-c.msg[cId][uId]
			if err := stream.Send(&m); err != nil {
				// Put message back to the channel
				c.msg[cId][uId] <- m
				log.Printf("Sending message stream connection failed: %v", err)
				sendErr <- err
			}
			log.Printf("Message sent to user %s for chat %s: %+v", uId, cId, m)
		}
	}()

	var cId, uId string

	defer func() {
		if initialised {
			c.closeSession(cId, uId)
		}
	}()

	//Receiving ReadReceipt
	for {
		select {
		case err := <-sendErr:
			return err
		default:
		}

		readReceipt, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if !initialised {
			c.initialiseChat(readReceipt.ChatId, readReceipt.UserId)
			chatId <- readReceipt.ChatId
			userId <- readReceipt.UserId
			cId = readReceipt.ChatId
			uId = readReceipt.UserId
			initialised = true
		}

		//If initialisation message then skip to next message
		if readReceipt.MessageId == 0 {
			continue
		}

		//Send to DB
		go func() {
			receipt := model.ReceiptFromGrpc(readReceipt)
			if err := c.service.UpdateReadMessage(receipt); err != nil {
				log.Println(err)
			}
		}()

		//Add send read receipt message for other users
		m := Message{
			SenderId: uId,
			Type:     Message_READ,
			Text:     strconv.FormatInt(readReceipt.MessageId, 10),
			DateTime: readReceipt.DateTime,
		}
		for u := range c.msg[cId] {
			if u != uId {
				c.msg[cId][u] <- m
			}
		}
	}
}

func (c *chatServiceServer) initialiseChat(chatId, userId string) {
	if _, exists := c.msg[chatId]; !exists {
		c.msg[chatId] = make(map[string]chan Message)
	}
	if _, exists := c.msg[chatId][userId]; !exists {
		c.msg[chatId][userId] = make(chan Message, 1000)
	}
}

func (c *chatServiceServer) closeSession(chatId, userId string) {
	if _, exists := c.msg[chatId]; !exists {
		return
	}
	if _, exists := c.msg[chatId][userId]; exists {
		delete(c.msg[chatId], userId)
	}
}
