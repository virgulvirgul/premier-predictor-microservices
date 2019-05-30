package model

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type MatchPredictionSummary struct {
	HomeWin int32 `json:"homeWin,omitempty"`
	Draw    int32 `json:"draw,omitempty"`
	AwayWin int32 `json:"awayWin,omitempty"`
}

func MatchPredictionSummaryFromGrpc(predictionSummary *model.MatchPredictionSummary) *MatchPredictionSummary {
	return &MatchPredictionSummary{
		HomeWin: predictionSummary.HomeWin,
		Draw:    predictionSummary.Draw,
		AwayWin: predictionSummary.AwayWin,
	}
}

func MatchPredictionSummaryToGrpc(predictionSummary *MatchPredictionSummary) *model.MatchPredictionSummary {
	return &model.MatchPredictionSummary{
		HomeWin: predictionSummary.HomeWin,
		Draw:    predictionSummary.Draw,
		AwayWin: predictionSummary.AwayWin,
	}
}

type Prediction struct {
	UserId    string `json:"userId,omitempty"`
	MatchId   string `json:"matchId,omitempty"`
	HomeGoals int32  `json:"hGoals,omitempty"`
	AwayGoals int32  `json:"aGoals,omitempty"`
}

func PredictionFromGrpc(prediction *model.Prediction) *Prediction {
	return &Prediction{
		UserId:    prediction.UserId,
		MatchId:   prediction.MatchId,
		HomeGoals: prediction.HGoals,
		AwayGoals: prediction.AGoals,
	}
}

func PredictionToGrpc(prediction *Prediction) *model.Prediction {
	return &model.Prediction{
		UserId:  prediction.UserId,
		MatchId: prediction.MatchId,
		HGoals:  prediction.HomeGoals,
		AGoals:  prediction.AwayGoals,
	}
}
