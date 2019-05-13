package factory

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"google.golang.org/grpc"
)

type predictionClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewPredictionClientFactory(addr string) interfaces.PredictionClientFactory {
	return &predictionClientFactory{
		addr: addr,
	}
}

func (n *predictionClientFactory) NewPredictionClient() (gen.PredictionServiceClient, error) {
	conn, err := grpc.Dial(n.addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	n.conn = conn

	return gen.NewPredictionServiceClient(conn), nil
}

func (n *predictionClientFactory) CloseConnection() error {
	return n.conn.Close()
}
