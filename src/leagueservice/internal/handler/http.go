package handler

import (
	"encoding/json"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
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
	invalidPin = "Invalid Pin"
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

	router.HandleFunc("/{id}", h.getUsersLeagueList).
		Methods(http.MethodGet)
	router.HandleFunc("/", h.addLeague).
		Methods(http.MethodPost)
	router.HandleFunc("/join", h.joinLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/leave", h.leaveLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/rename", h.renameLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/standings/{id}", h.getLeagueTable).
		Methods(http.MethodGet)
	router.HandleFunc("/standings", h.getOverallTable).
		Methods(http.MethodGet)

	return router
}

func (h *httpHandler) getUsersLeagueList(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	league, err := h.service.GetUsersLeagueList(id)

	h.sendResponse(league, err, w)
}

func (h *httpHandler) addLeague(w http.ResponseWriter, r *http.Request) {
	var req addLeagueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	league, err := h.service.AddUserLeague(req.Id, req.Name)

	h.sendResponse(league, err, w)
}

func (h *httpHandler) joinLeague(w http.ResponseWriter, r *http.Request) {
	var req leagueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	leagueOverview, err := h.service.JoinUserLeague(req.Id, req.Pin)

	h.sendResponse(leagueOverview, err, w)
}

func (h *httpHandler) leaveLeague(w http.ResponseWriter, r *http.Request) {
	var req leagueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	err = h.service.LeaveUserLeague(req.Id, req.Pin)

	h.sendResponse(nil, err, w)
}

func (h *httpHandler) renameLeague(w http.ResponseWriter, r *http.Request) {
	var req renameRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidRequestBody), w)
		return
	}

	err = h.service.RenameUserLeague(req.Pin, req.Name)

	h.sendResponse(nil, err, w)
}

func (h *httpHandler) getLeagueTable(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	pin, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.sendResponse(nil, errors.Wrap(m.ErrInvalidRequestData, invalidPin), w)
		return
	}

	league, err := h.service.GetLeagueTable(pin)
	h.sendResponse(league, err, w)
}

func (h *httpHandler) getOverallTable(w http.ResponseWriter, r *http.Request) {
	league, err := h.service.GetOverallLeagueTable()
	h.sendResponse(league, err, w)
}

func (h *httpHandler) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	_, leagueCreated := data.(*model.League)

	switch {
	case err == nil && leagueCreated:
		w.WriteHeader(http.StatusCreated)
		if data != nil {
			_ = json.NewEncoder(w).Encode(data)
		}
		return

	case err == nil:
		w.WriteHeader(http.StatusOK)
		if data != nil {
			_ = json.NewEncoder(w).Encode(data)
		}
		return

	case err == model.ErrLeagueNotFound:
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
