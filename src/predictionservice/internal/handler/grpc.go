package handler

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"log"
)

type predictionServiceServer struct {
	service interfaces.Service
}

func NewPredictionServiceServer(service interfaces.Service) (*predictionServiceServer, error) {
	log.Print("Registered predictionServiceServer handler")

	return &predictionServiceServer{
		service: service,
	}, nil
}

func (p *predictionServiceServer) GetPrediction(ctx context.Context, req *gen.PredictionRequest) (*gen.Prediction, error) {
	prediction, err := p.service.GetPrediction(req.UserId, req.MatchId)
	if err != nil {
		return nil, err
	}

	return model.PredictionToGrpc(prediction), nil
}

func (p *predictionServiceServer) GetPredictionSummary(ctx context.Context, req *gen.IdRequest) (*gen.MatchPredictionSummary, error) {
	predictionSummary, err := p.service.GetMatchPredictionSummary(req.Id)
	if err != nil {
		return nil, err
	}

	return model.MatchPredictionSummaryToGrpc(predictionSummary), nil
}
