package auth

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
)

type authenticator struct {
	client gen.AuthServiceClient
}

func NewAuthenticator(client gen.AuthServiceClient) (interfaces.Authenticator, error) {
	return &authenticator{
		client: client,
	}, nil
}

func (a *authenticator) doAuth(token string) error {
	request := &gen.ValidateRequest{Token: token}

	_, err := a.client.Validate(context.Background(), request)

	if err != nil {
		return err
	}

	return nil
}