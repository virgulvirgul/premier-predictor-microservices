package handler

import (
	"encoding/json"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/legacyuserservice/internal/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)

type ServerError struct {
	Message string `json:"message"`
}

const (
	invalidRequestBody = "Invalid Request Body"
	invalidUserId      = "Invalid user id"
)

type httpHandler struct {
	service interfaces.Service
}

func NewHttpHandler(service interfaces.Service) (*httpHandler, error) {
	log.Print("Registered httpServer handler")

	return &httpHandler{
		service: service,
	}, nil
}

func (h *httpHandler) Route() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/health", health.Health).
		Methods(http.MethodGet)

	router.HandleFunc("/{id}", h.getLegacyUserById).
		Methods(http.MethodGet)
	router.HandleFunc("/", h.legacyLogin).
		Methods(http.MethodPost)

	return router
}

func (h *httpHandler) getLegacyUserById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]

	id, err := strconv.Atoi(param)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidUserId), w)
	}

	legacyUser, err := h.service.GetUserById(id)

	h.sendResponse(legacyUser, err, w)
}

func (h *httpHandler) legacyLogin(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	legacyUser, err := h.service.LegacyLogin(req.Email, req.Password)

	h.sendResponse(legacyUser, err, w)
}

func (h *httpHandler) sendResponse(legacyUser *model.User, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		w.WriteHeader(http.StatusOK)
		if legacyUser != nil {
			_ = json.NewEncoder(w).Encode(legacyUser)
			return
		}
		return

	case err == model.ErrLegacyUserNotFound:
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: err.Error(),
		})

	case err == model.ErrLegacyLoginFailed:
		w.WriteHeader(http.StatusUnauthorized)
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
