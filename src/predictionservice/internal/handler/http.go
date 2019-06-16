package handler

import (
	"encoding/json"
	common "github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/repository"
	"github.com/gorilla/mux"
	"io/ioutil"
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

	router.HandleFunc("/fixtures/predicted/{id}", h.getFixturesWithPredictions).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions/{id}", h.getPredictorData).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions", h.getFixturesWithPredictions).
		Methods(http.MethodPost)
	router.HandleFunc("/predictions/summary/{id}", h.getUsersPastPredictions).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions/{userId}/{matchId}", h.getPrediction).
		Methods(http.MethodGet)

	return router
}

func (h *httpHandler) getFixturesWithPredictions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fixturePredictions, err := h.service.GetFixturesWithPredictions(id)

	h.sendResponse(fixturePredictions, err, w)
}

func (h *httpHandler) getPredictorData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictorData, err := h.service.GetPredictorData(id)

	h.sendResponse(predictorData, err, w)
}

func (h *httpHandler) getUsersPastPredictions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictions, err := h.service.GetUsersPastPredictions(id)

	h.sendResponse(predictions, err, w)
}

func (h *httpHandler) updatePredictions(w http.ResponseWriter, r *http.Request) {
	var predictions []model.Prediction

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("cannot read request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &predictions); err != nil {
		log.Println("cannot decode request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.service.UpdatePredictions(predictions); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *httpHandler) getPrediction(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	matchId := mux.Vars(r)["matchId"]

	predictions, err := h.service.GetPrediction(userId, matchId)

	h.sendResponse(predictions, err, w)
}

func (h *httpHandler) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	if err == prediction.ErrPredictionNotFound {
		log.Println("prediction not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
