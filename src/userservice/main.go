package main

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	"github.com/cshep4/premier-predictor-microservices/src/common/factory"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces"
	repo "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/repository"
	svc "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service"
	"github.com/gorilla/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	authAddress, ok := os.LookupEnv("AUTH_ADDR")
	if !ok {
		log.Fatalf("failed to get authservice address")
	}

	authFactory := factory.NewAuthClientFactory(authAddress)
	authClient, err := authFactory.NewAuthClient()
	clientConnCloseFunc = append(clientConnCloseFunc, authFactory.CloseConnection)

	authenticator, err := auth.NewAuthenticator(authClient)
	if err != nil {
		log.Fatalf("failed to create authenticator: %v", err)
	}

	repository, err := repo.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	service, err := svc.NewService(repository)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	var exitCode codes.Code
	var grpcServer *grpc.Server
	var httpServer *http.Server

	sigs := make(chan os.Signal)

	go func() {
		grpcServer = startGrpcServer(service, authenticator)
		sigs <- syscall.SIGQUIT
	}()

	go func() {
		httpServer = startHttpServer(service, authenticator)
		sigs <- syscall.SIGQUIT
	}()

	switch sig := <-sigs; sig {
	case os.Interrupt, syscall.SIGINT, syscall.SIGQUIT:
		log.Print("Shutting down")

		for i := range clientConnCloseFunc {
			err := clientConnCloseFunc[i]()
			if err != nil {
				log.Printf("Error closing client connection: %v\n", err)
			}
		}

		grpcServer.GracefulStop()
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Printf("Error shutting down http server: %v\n", err)
		}

		exitCode = codes.Aborted
	case syscall.SIGTERM:
		exitCode = codes.OK
	}

	os.Exit(int(exitCode))
}

var clientConnCloseFunc []func() error

func startHttpServer(service interfaces.Service, authenticator common.Authenticator) *http.Server {
	h, err := handler.NewHttpHandler(service, authenticator)
	if err != nil {
		log.Fatalf("failed to create http handler: %v", err)
	}

	path := ":" + os.Getenv("HTTP_PORT")

	http := &http.Server{
		Addr:         path,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handlers.CORS()(h.Route()),
	}

	log.Printf("Http server listening on %s", path)

	err = http.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start http server: %v\n", err)
	}

	return http
}

func startGrpcServer(service interfaces.Service, authenticator common.Authenticator) *grpc.Server {
	path := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Grpc server listening on %s", path)

	server, err := handler.NewUserServiceServer(service)
	if err != nil {
		log.Fatalf("failed to create grpc handler: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authenticator.GrpcUnaryInterceptor),
	)

	gen.RegisterUserServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}
