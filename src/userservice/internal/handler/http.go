package handler

import (
	"encoding/json"
	"github.com/cshep4/premier-predictor-microservices/src/common/health"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type ServerError struct {
	Message string `json:"message"`
}

const invalidRequestBody = "Invalid Request Body"

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

	router.HandleFunc("/users/{id}", h.getUser).
		Methods(http.MethodGet)
	router.HandleFunc("/users", h.updateUserInfo).
		Methods(http.MethodPut)
	router.HandleFunc("/users/password", h.updatePassword).
		Methods(http.MethodPut)
	router.HandleFunc("/users/score/{id}", h.getUserScore).
		Methods(http.MethodGet)
	//TODO - add tests, fix middleware
	router.HandleFunc("/legacy", h.storeLegacyUser).
		Methods(http.MethodPost)

	router.HandleFunc("/health", health.Health).
		Methods(http.MethodGet)

	return router
}

func (h *httpHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.service.GetUser(id)

	h.sendResponse(user, err, w)
}

func (h *httpHandler) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	var userInfo model.UserInfo
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	err = h.service.UpdateUserInfo(userInfo)

	h.sendResponse(nil, err, w)
}

func (h *httpHandler) updatePassword(w http.ResponseWriter, r *http.Request) {
	var updatePassword model.UpdatePassword
	err := json.NewDecoder(r.Body).Decode(&updatePassword)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	err = h.service.UpdatePassword(updatePassword)

	h.sendResponse(nil, err, w)
}

func (h *httpHandler) getUserScore(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	score, err := h.service.GetUserScore(id)

	h.sendResponse(model.UserScore{
		Score: score,
	}, err, w)
}

func (h *httpHandler) storeLegacyUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	err = h.service.StoreLegacyUser(user)

	h.sendResponse(nil, err, w)
}

func (h *httpHandler) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		if data != nil {
			_ = json.NewEncoder(w).Encode(data)
			return
		}
		w.WriteHeader(http.StatusOK)
		return

	case err == model.ErrUserNotFound:
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
