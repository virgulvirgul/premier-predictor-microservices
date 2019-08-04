package factory

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

type userClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewUserClientFactory(addr string) interfaces.UserClientFactory {
	return &userClientFactory{
		addr: addr,
	}
}

func (u *userClientFactory) NewUserClient() (model.UserServiceClient, error) {
	err := u.connect()
	if err != nil {
		return nil, err
	}
	u.connectionOnState()

	return model.NewUserServiceClient(u.conn), nil
}

func (u *userClientFactory) connect() error {
	conn, err := grpc.Dial(u.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	u.conn = conn

	return nil
}

func (u *userClientFactory) connectionOnState() {
	go func() {
		for {
			u.conn.WaitForStateChange(context.Background(), u.conn.GetState())

			currentState := u.conn.GetState()
			log.Printf("connection state change, currentState: %s", currentState)
			if currentState != connectivity.Ready {
				log.Println("reconnecting userservice connection")

				err := u.connect()
				if err != nil {
					log.Println("failed to reconnect userservice connection")
					continue
				}

				log.Println("reconnected userservice connection!")
			}
		}
	}()
}

func (u *userClientFactory) CloseConnection() error {
	return u.conn.Close()
}
