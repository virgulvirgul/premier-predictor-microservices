package main

import (
	"context"
	"crypto/tls"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/handler"
	repo "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/repository"
	svc "github.com/cshep4/premier-predictor-microservices/src/chatservice/internal/service"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	"github.com/cshep4/premier-predictor-microservices/src/common/factory"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	"github.com/cshep4/premier-predictor-microservices/src/common/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
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

		for i := range  clientConnCloseFunc{
			err := clientConnCloseFunc[i]()
			if err != nil {
				log.Printf("Error closing client connection: %v\n", err)
			}
		}

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

var clientConnCloseFunc []func() error

func startGrpcServer() *grpc.Server {
	cer, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/privkey.pem")
	if err != nil {
		log.Fatalf("Failed to load key pair: %v\n", err)
	}

	path := ":" + os.Getenv("PORT")
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	lis, err := tls.Listen("tcp", path, config)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	//path := ":" + os.Getenv("PORT")
	//
	//lis, err := net.Listen("tcp", path)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}

	log.Printf("Listening on %s", path)

	authAddress, ok := os.LookupEnv("AUTH_ADDR")
	if !ok {
		log.Fatalf("failed to get authservice address")
	}

	notificationAddress, ok := os.LookupEnv("NOTIFICATION_ADDR")
	if !ok {
		log.Fatalf("failed to get notificationservice address")
	}

	authFactory := factory.NewAuthClientFactory(authAddress)
	authClient, err := authFactory.NewAuthClient()
	clientConnCloseFunc = append(clientConnCloseFunc, authFactory.CloseConnection)

	notificationFactory := factory.NewNotificationClientFactory(notificationAddress)
	notificationClient, err := notificationFactory.NewNotificationClient()
	clientConnCloseFunc = append(clientConnCloseFunc, notificationFactory.CloseConnection)

	authenticator, err := auth.NewAuthenticator(authClient)
	if err != nil {
		log.Fatalf("failed to create authenticator: %v", err)
	}

	notifier, err := notification.NewNotifier(notificationClient)
	if err != nil {
		log.Fatalf("failed to create notifier: %v", err)
	}

	repository, err := repo.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	service, err := svc.NewService(repository, notifier)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	server, err := handler.NewChatServiceServer(service)
	if err != nil {
		log.Fatalf("failed to grpc handler: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authenticator.GrpcUnaryInterceptor),
		grpc.StreamInterceptor(authenticator.GrpcStreamInterceptor),
	)

	gen.RegisterChatServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}
