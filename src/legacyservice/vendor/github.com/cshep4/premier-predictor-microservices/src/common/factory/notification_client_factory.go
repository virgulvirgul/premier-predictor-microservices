package factory

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"google.golang.org/grpc"
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
	conn, err := grpc.Dial(n.addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return gen.NewNotificationServiceClient(conn), nil
}

func (n *notificationClientFactory) CloseConnection() error {
	return n.conn.Close()
}
