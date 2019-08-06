package factory

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"

	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

type notificationClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewNotificationClientFactory(addr string) interfaces.NotificationClientFactory {
	return &notificationClientFactory{
		addr: addr,
	}
}

func (n *notificationClientFactory) NewNotificationClient() (gen.NotificationServiceClient, error) {
	err := n.connect()
	if err != nil {
		return nil, err
	}
	n.connectionOnState()

	return gen.NewNotificationServiceClient(n.conn), nil
}

func (n *notificationClientFactory) connect() error {

	conn, err := grpc.Dial(n.addr, grpc.WithInsecure(), options.ClientKeepAlive)
	if err != nil {
		return err
	}
	n.conn = conn

	return nil
}

func (n *notificationClientFactory) connectionOnState() {
	go func() {
		for {
			n.conn.WaitForStateChange(context.Background(), n.conn.GetState())

			currentState := n.conn.GetState()
			log.Printf("notificationservice - connection state change - currentState: %s", currentState)

			if currentState == connectivity.Connecting {
				continue
			}

			if currentState != connectivity.Ready {
				log.Println("reconnecting notificationservice connection")

				err := n.connect()
				if err != nil {
					log.Println("failed to reconnect notificationservice connection")
					continue
				}

				log.Println("reconnected notificationservice connection!")
			}
		}
	}()
}

func (n *notificationClientFactory) CloseConnection() error {
	return n.conn.Close()
}
