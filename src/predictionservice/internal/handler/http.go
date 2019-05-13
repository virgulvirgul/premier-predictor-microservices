package handler

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type httpHandler struct {
	service       interfaces.Service
	authenticator common.Authenticator
}

func NewHttpHandler(service interfaces.Service, authenticator common.Authenticator) (*httpHandler, error) {
	log.Print("Registered httpServer handler")

	return &httpHandler{
		service:       service,
		authenticator: authenticator,
	}, nil
}

func (h *httpHandler) Route() http.Handler {
	router := mux.NewRouter()
	router.Use(h.authenticator.HttpMiddleware)

	return router
}
