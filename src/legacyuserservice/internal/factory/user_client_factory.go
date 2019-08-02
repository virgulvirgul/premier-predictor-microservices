package factory

import (
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"google.golang.org/grpc"
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

func (n *userClientFactory) NewUserClient() (model.UserServiceClient, error) {
	conn, err := grpc.Dial(n.addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return model.NewUserServiceClient(conn), nil
}

func (n *userClientFactory) CloseConnection() error {
	return n.conn.Close()
}
