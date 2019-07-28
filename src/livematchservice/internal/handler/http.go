package handler

import (
	"encoding/json"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
)

type ServerError struct {
	Message string `json:"message"`
}

const (
	invalidRequestBody = "Invalid Request Body"
	invalidPin         = "Invalid Pin"
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

	router.HandleFunc("/health", health.Health).
		Methods(http.MethodGet)

	router.HandleFunc("/upcoming", h.getUpcomingMatches).
		Methods(http.MethodGet)
	router.HandleFunc("/match/{matchId}/user/{userId}", h.getMatchSummary).
		Methods(http.MethodGet)

	return router
}

func (h *httpHandler) getUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	upcomingMatches, err := h.service.GetUpcomingMatches()

	h.sendResponse(upcomingMatches, err, w)
}

func (h *httpHandler) getMatchSummary(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	ctx := metadata.NewIncomingContext(r.Context(), metadata.MD{"token": []string{token}})

	matchId := mux.Vars(r)["matchId"]
	userId := mux.Vars(r)["userId"]

	req := model.PredictionRequest{
		UserId:  userId,
		MatchId: matchId,
	}

	league, err := h.service.GetMatchSummary(ctx, req)

	h.sendResponse(league, err, w)
}

func (h *httpHandler) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		w.WriteHeader(http.StatusOK)
		if data != nil {
			_ = json.NewEncoder(w).Encode(data)
			return
		}
		return

	case err == model.ErrMatchNotFound:
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: err.Error(),
		})

	case errors.Cause(err) == m.ErrInvalidRequestData:
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: util.GetErrorMessage(err),
		})

	default:
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: err.Error(),
		})
	}

	log.Println(err.Error())
}
