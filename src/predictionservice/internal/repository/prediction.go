package prediction

import "github.com/cshep4/premier-predictor-microservices/src/common/model"

type predictionEntity struct {
	UserId    string `bson:"userId,omitempty"`
	MatchId   string `bson:"matchId,omitempty"`
	HomeGoals int    `bson:"hGoals,omitempty"`
	AwayGoals int    `bson:"aGoals,omitempty"`
}

func fromPrediction(prediction *model.Prediction) *predictionEntity {
	return &predictionEntity{
		UserId:    prediction.UserId,
		MatchId:   prediction.MatchId,
		HomeGoals: prediction.HomeGoals,
		AwayGoals: prediction.AwayGoals,
	}
}

func toPrediction(prediction *predictionEntity) *model.Prediction {
	return &model.Prediction{
		UserId:    prediction.UserId,
		MatchId:   prediction.MatchId,
		HomeGoals: prediction.HomeGoals,
		AwayGoals: prediction.AwayGoals,
	}
}
