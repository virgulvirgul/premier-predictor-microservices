package factory

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"google.golang.org/grpc"
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

func (n *authClientFactory) NewAuthClient() (gen.AuthServiceClient, error) {
	conn, err := grpc.Dial(n.addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return gen.NewAuthServiceClient(conn), nil
}

func (n *authClientFactory) CloseConnection() error {
	return n.conn.Close()
}
