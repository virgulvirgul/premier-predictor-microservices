package handler

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type legacyUserServiceServer struct {
	service interfaces.Service
}

func NewLegacyUserServiceServer(service interfaces.Service) (*legacyUserServiceServer, error) {
	log.Print("Registered legacyUserServiceServer handler")

	return &legacyUserServiceServer{
		service: service,
	}, nil
}

func (l *legacyUserServiceServer) GetLegacyUserById(ctx context.Context, req *gen.LegacyIdRequest) (*gen.LegacyUserResponse, error) {
	user, err := l.service.GetUserById(int(req.Id))

	if err != nil {
		return nil, err
	}

	return model.LegacyUserToGrpc(user), nil
}

func (l *legacyUserServiceServer) LegacyLogin(ctx context.Context, req *gen.LoginRequest) (*gen.LegacyUserResponse, error) {
	user, err := l.service.LegacyLogin(req.Email, req.Password)

	if err == model.ErrLegacyLoginFailed {
		return nil, status.Error(codes.Unauthenticated, model.ErrLegacyLoginFailed.Error())
	}

	if err != nil {
		return nil, err
	}

	return model.LegacyUserToGrpc(user), nil
}
