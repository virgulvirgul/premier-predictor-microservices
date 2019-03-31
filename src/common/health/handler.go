package health

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type healthServiceServer struct{}

func (s *healthServiceServer) createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods("GET")

	return r
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Service is healthy")
	w.WriteHeader(http.StatusOK)
}
