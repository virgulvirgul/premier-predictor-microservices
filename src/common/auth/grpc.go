package auth

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func (a *authenticator) GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	token, err := a.getTokenFromGrpcMetadata(ctx)
	if err != nil {
		log.Printf("auth error: %s", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = a.doAuth(token)
	if err != nil {
		log.Printf("auth error: %s", err)
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return handler(ctx, req)
}

func (a *authenticator) GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	token, err := a.getTokenFromGrpcMetadata(stream.Context())
	if err != nil {
		log.Printf("auth error: %s", err)
		return status.Error(codes.Unauthenticated, err.Error())
	}

	err = a.doAuth(token)
	if err != nil {
		log.Printf("auth error: %s", err)
		return status.Error(codes.Unauthenticated, err.Error())
	}

	return handler(srv, stream)
}

func (a *authenticator) getTokenFromGrpcMetadata(ctx context.Context) (string, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing context metadata")
	}

	if len(meta["token"]) != 1 {
		return "", errors.New("invalid access token")
	}

	return meta["token"][0], nil
}
