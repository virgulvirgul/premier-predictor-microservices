package factory

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"log"
)

type fixtureClientFactory struct {
	conn *grpc.ClientConn
	addr string
}

func NewFixtureClientFactory(addr string) interfaces.FixtureClientFactory {
	return &fixtureClientFactory{
		addr: addr,
	}
}

func (f *fixtureClientFactory) NewFixtureClient() (gen.FixtureServiceClient, error) {
	err := f.connect()
	if err != nil {
		return nil, err
	}
	f.connectionOnState()

	return gen.NewFixtureServiceClient(f.conn), nil
}

func (f *fixtureClientFactory) connect() error {
	conn, err := grpc.Dial(f.addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	f.conn = conn

	return nil
}

func (f *fixtureClientFactory) connectionOnState() {
	go func() {
		for {
			f.conn.WaitForStateChange(context.Background(), f.conn.GetState())

			currentState := f.conn.GetState()
			log.Printf("connection state change, currentState: %s", currentState)
			if currentState != connectivity.Ready {
				log.Println("reconnecting fixtureservice connection")

				err := f.connect()
				if err != nil {
					log.Println("failed to reconnect fixtureservice connection")
					continue
				}

				log.Println("reconnected fixtureservice connection!")
			}
		}
	}()
}

func (f *fixtureClientFactory) CloseConnection() error {
	return f.conn.Close()
}
