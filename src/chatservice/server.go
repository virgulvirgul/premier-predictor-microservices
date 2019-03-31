package main

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/handler"
	chat2 "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/repository"
	chat3 "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/service"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
)

func main() {
	var exitCode codes.Code
	var grpcServer *grpc.Server
	var healthServer *http.Server

	sigs := make(chan os.Signal)

	go func() {
		grpcServer = startGrpcServer()
		sigs <- syscall.SIGQUIT
	}()

	go func() {
		healthServer = health.StartHealthServer()
		sigs <- syscall.SIGQUIT
	}()

	switch sig := <-sigs; sig {
	case os.Interrupt, syscall.SIGINT, syscall.SIGQUIT:
		log.Print("Shutting down")

		grpcServer.GracefulStop()
		err := healthServer.Shutdown(context.Background())
		if err != nil {
			log.Printf("Error shutting down health server: %v\n", err)
		}

		exitCode = codes.Aborted
	case syscall.SIGTERM:
		exitCode = codes.OK
	}

	os.Exit(int(exitCode))
}

func startGrpcServer() *grpc.Server {
	//cer, err := tls.LoadX509KeyPair("certs/tls.crt", "certs/tls.key")
	//if err != nil {
	//	log.Fatalf("Failed to load key pair: %v\n", err)
	//}
	//
	//path := ":" + os.Getenv("PORT")
	//config := &tls.Config{Certificates: []tls.Certificate{cer}}
	//
	//lis, err := tls.Listen("tcp", path, config)
	//if err != nil {
	//	log.Fatalf("Failed to listen: %v\n", err)
	//}

	path := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", path)

	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(auth.Interceptor))
	grpcServer := grpc.NewServer()

	repository, err := chat2.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	service, err := chat3.NewService(repository)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	server, err := handler.NewChatServiceServer(service)
	if err != nil {
		log.Fatalf("failed to grpc handler: %v", err)
	}

	chat.RegisterChatServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}
