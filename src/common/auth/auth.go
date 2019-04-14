package auth

import (
	"context"
	. "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	request, err := createRequest(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot create validate request")
	}

	svcAddr := os.Getenv("AUTH_ADDR")

	if svcAddr == "" {
		log.Println("could not connect to auth service")
		return nil, status.Error(codes.Unauthenticated, "could not connect to auth service")
	}

	creds, err := credentials.NewClientTLSFromFile("certs/tls.crt", "auth.gyme.uk")
	if err != nil {
		log.Fatalf("Failed to load key pair: %v\n", err)
	}

	conn, err := grpc.DialContext(ctx, svcAddr, grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Unauthenticated, "could not connect to auth service")
	}

	defer conn.Close()

	c := NewAuthServiceClient(conn)
	_, err = c.Validate(ctx, request)

	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func createRequest(ctx context.Context, req interface{}) (*ValidateRequest, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["token"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token")
	}

	id, err := getUserId(req)
	if err != nil {
		return nil, err
	}

	return &ValidateRequest{Id: id, Token: meta["token"][0]}, nil
}

func getUserId(req interface{}) (string, error) {
	request, ok := req.(*IdRequest)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "cannot get user id")
	}

	return request.Id, nil
}