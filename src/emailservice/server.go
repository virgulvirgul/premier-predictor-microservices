package main

import (
	"context"
	. "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/grpc/options"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	"github.com/cshep4/premier-predictor-microservices/src/emailservice/internal/server"
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
	path := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", path)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", path)

	grpcServer := grpc.NewServer(options.ServerKeepAlive)

	svc := server.NewEmailServiceServer()
	RegisterEmailServiceServer(grpcServer, svc)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}

	return grpcServer
}
