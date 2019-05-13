package interfaces

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
)

type Authenticator interface {
	GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
	HttpMiddleware(next http.Handler) http.Handler
}
