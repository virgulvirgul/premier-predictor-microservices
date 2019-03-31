package health

import (
	"log"
	"net/http"
	"os"
	"time"
)

func StartHealthServer() *http.Server {
	svc := &healthServiceServer{}
	r := svc.createRouter()

	log.Print("Registered healthServiceServer handler")

	path := ":" + os.Getenv("HEALTH_PORT")

	healthServer := &http.Server{
		Addr:         path,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	log.Printf("Health server listening on %s", path)

	err := healthServer.ListenAndServe()
	if err != nil {
		log.Printf("Failed to start health server: %v\n", err)
	}

	return healthServer
}
