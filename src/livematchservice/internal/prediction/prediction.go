package prediction

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces"
	. "github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
)

type predictor struct {
	client gen.PredictionServiceClient
}

func NewPredictor(client gen.PredictionServiceClient) (interfaces.Predictor, error) {
	return &predictor{
		client: client,
	}, nil
}

func (p *predictor) GetPrediction(ctx context.Context, req PredictionRequest) (*model.Prediction, error) {
	r := &gen.PredictionRequest{
		UserId:  req.UserId,
		MatchId: req.MatchId,
	}

	metadata, err := util.CreateRequestMetadata(ctx)
	if err != nil {
		return nil, err
	}

	prediction, err := p.client.GetPrediction(metadata, r)
	if err != nil {
		return nil, err
	}

	return model.PredictionFromGrpc(prediction), nil
}

func (p *predictor) GetPredictionSummary(ctx context.Context, matchId string) (*model.MatchPredictionSummary, error) {
	r := &gen.IdRequest{
		Id: matchId,
	}

	metadata, err := util.CreateRequestMetadata(ctx)
	if err != nil {
		return nil, err
	}

	predictionSummary, err := p.client.GetPredictionSummary(metadata, r)
	if err != nil {
		return nil, err
	}

	return model.MatchPredictionSummaryFromGrpc(predictionSummary), nil
}
