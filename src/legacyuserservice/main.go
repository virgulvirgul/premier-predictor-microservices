package main

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	uFactory "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/factory"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	repo "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/repository"
	svc "github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/service"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/user"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

var clientConnCloseFunc []func() error

func main() {
	userAddress, ok := os.LookupEnv("USER_ADDR")
	if !ok {
		log.Fatalf("failed to get userservice address")
	}

	userFactory := uFactory.NewUserClientFactory(userAddress)
	userClient, err := userFactory.NewUserClient()
	clientConnCloseFunc = append(clientConnCloseFunc, userFactory.CloseConnection)

	userService, err := user.NewUserService(userClient)
	if err != nil {
		log.Fatalf("failed to create userService: %v", err)
	}

	repository, err := repo.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	service, err := svc.NewService(repository, userService)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	var exitCode codes.Code
	var grpcServer *grpc.Server
	var httpServer *http.Server

	sigs := make(chan os.Signal)

	go func() {
		grpcServer = startGrpcServer(service)
		sigs <- syscall.SIGQUIT
	}()

	go func() {
		httpServer = startHttpServer(service)
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

func startHttpServer(service interfaces.Service) *http.Server {
	h, err := handler.NewHttpHandler(service)
	if err != nil {
		log.Fatalf("failed to create http handler: %v", err)
	}

	path := ":" + os.Getenv("HTTP_PORT")

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{"*"},
	})

	http := &http.Server{
		Addr:         path,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      corsOpts.Handler(h.Route()),
	}

	log.Printf("Http server listening on %s", path)

	err = http.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start http server: %v\n", err)
	}

	return http
}

func startGrpcServer(service interfaces.Service) *grpc.Server {
	path := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Grpc server listening on %s", path)

	server, err := handler.NewLegacyUserServiceServer(service)
	if err != nil {
		log.Fatalf("failed to create grpc handler: %v", err)
	}

	grpcServer := grpc.NewServer(options.ServerKeepAlive)

	gen.RegisterLegacyUserServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}
