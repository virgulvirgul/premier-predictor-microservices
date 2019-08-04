package factory

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

type authClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewAuthClientFactory(addr string) interfaces.AuthClientFactory {
	return &authClientFactory{
		addr: addr,
	}
}

func (a *authClientFactory) NewAuthClient() (gen.AuthServiceClient, error) {
	err := a.connect()
	if err != nil {
		return nil, err
	}
	a.connectionOnState()

	return gen.NewAuthServiceClient(a.conn), nil
}

func (a *authClientFactory) connect() error {
	conn, err := grpc.Dial(a.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	a.conn = conn

	return nil
}

func (a *authClientFactory) connectionOnState() {
	go func() {
		for {
			a.conn.WaitForStateChange(context.Background(), a.conn.GetState())

			currentState := a.conn.GetState()
			log.Printf("connection state change, currentState: %s", currentState)
			if currentState != connectivity.Ready {
				log.Println("reconnecting authservice connection")

				err := a.connect()
				if err != nil {
					log.Println("failed to reconnect authservice connection")
					continue
				}

				log.Println("reconnected authservice connection!")
			}
		}
	}()
}

func (a *authClientFactory) CloseConnection() error {
	return a.conn.Close()
}
