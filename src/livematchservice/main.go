package main

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/auth"
	"github.com/cshep4/premier-predictor-microservices/src/common/factory"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	pFactory "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/factory"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/prediction"
	repo "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/repository"
	svc "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/service"
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

func main() {
	authAddress, ok := os.LookupEnv("AUTH_ADDR")
	if !ok {
		log.Fatalf("failed to get authservice address")
	}

	predictionAddress, ok := os.LookupEnv("PREDICTION_ADDR")
	if !ok {
		log.Fatalf("failed to get predictionservice address")
	}

	authFactory := factory.NewAuthClientFactory(authAddress)
	authClient, err := authFactory.NewAuthClient()
	clientConnCloseFunc = append(clientConnCloseFunc, authFactory.CloseConnection)

	predictionFactory := pFactory.NewPredictionClientFactory(predictionAddress)
	predictionClient, err := predictionFactory.NewPredictionClient()
	clientConnCloseFunc = append(clientConnCloseFunc, predictionFactory.CloseConnection)

	authenticator, err := auth.NewAuthenticator(authClient)
	if err != nil {
		log.Fatalf("failed to create authenticator: %v", err)
	}

	predictor, err := prediction.NewPredictor(predictionClient)
	if err != nil {
		log.Fatalf("failed to create predictor: %v", err)
	}

	repository, err := repo.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	service, err := svc.NewService(repository, predictor)
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
			log.Printf("Error shutting down health server: %v\n", err)
		}

		exitCode = codes.Aborted
	case syscall.SIGTERM:
		exitCode = codes.OK
	}

	os.Exit(int(exitCode))
}

var clientConnCloseFunc []func() error

func startGrpcServer(service interfaces.Service, authenticator common.Authenticator) *grpc.Server {
	path := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", path)

	server, err := handler.NewLiveMatchServiceServer(service, time.Minute)
	if err != nil {
		log.Fatalf("failed to grpc handler: %v", err)
	}

	grpcServer := grpc.NewServer(
		options.ServerKeepAlive,
		grpc.UnaryInterceptor(authenticator.GrpcUnaryInterceptor),
		grpc.StreamInterceptor(authenticator.GrpcStreamInterceptor),
	)

	gen.RegisterLiveMatchServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}

func startHttpServer(service interfaces.Service, authenticator common.Authenticator) *http.Server {
	h, err := handler.NewHttpHandler(service, authenticator)
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
