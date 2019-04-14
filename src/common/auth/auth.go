package auth

import (
	"context"
	. "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	request, err := createRequest(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot create validate request")
	}

	svcAddr := os.Getenv("AUTH_ADDR")

	if svcAddr == "" {
		log.Println("could not connect to auth service")
		return nil, status.Error(codes.Unauthenticated, "could not connect to auth service")
	}

	//creds, err := credentials.NewClientTLSFromFile("certs/tls.crt", "auth.gyme.uk")
	//if err != nil {
	//	log.Fatalf("Failed to load key pair: %v\n", err)
	//}

	//conn, err := grpc.DialContext(ctx, svcAddr, grpc.WithTransportCredentials(creds))
	conn, err := grpc.DialContext(ctx, svcAddr, grpc.WithInsecure())

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

func createRequest(ctx context.Context) (*ValidateRequest, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(meta["token"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token")
	}

	return &ValidateRequest{Token: meta["token"][0]}, nil
}