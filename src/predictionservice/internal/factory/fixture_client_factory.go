package factory

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	interfaces2 "github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"google.golang.org/grpc"
)

type fixtureClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewFixtureClientFactory(addr string) interfaces2.FixtureClientFactory {
	return &fixtureClientFactory{
		addr: addr,
	}
}

func (n *fixtureClientFactory) NewFixtureClient() (gen.FixtureServiceClient, error) {
	conn, err := grpc.Dial(n.addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return gen.NewFixtureServiceClient(conn), nil
}

func (n *fixtureClientFactory) CloseConnection() error {
	return n.conn.Close()
}
