package factory

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
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

func (p *predictionClientFactory) NewPredictionClient() (gen.PredictionServiceClient, error) {
	err := p.connect()
	if err != nil {
		return nil, err
	}
	p.connectionOnState()

	return gen.NewPredictionServiceClient(p.conn), nil
}

func (p *predictionClientFactory) connect() error {
	conn, err := grpc.Dial(p.addr, grpc.WithInsecure(), options.ClientKeepAlive)
	if err != nil {
		return err
	}
	p.conn = conn

	return nil
}

func (p *predictionClientFactory) connectionOnState() {
	go func() {
		for {
			p.conn.WaitForStateChange(context.Background(), p.conn.GetState())

			currentState := p.conn.GetState()
			log.Printf("predictionservice - connection state change - currentState: %s", currentState)

			if currentState == connectivity.Connecting {
				continue
			}

			if currentState != connectivity.Ready {
				log.Println("reconnecting predictionservice connection")

				err := p.connect()
				if err != nil {
					log.Println("failed to reconnect predictionservice connection")
					continue
				}

				log.Println("reconnected predictionservice connection!")
			}
		}
	}()
}

func (p *predictionClientFactory) CloseConnection() error {
	return p.conn.Close()
}
