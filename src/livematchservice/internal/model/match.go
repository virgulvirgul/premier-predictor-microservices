package model

import (
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	_ "github.com/golang/protobuf/proto"
	"time"
)

type MatchSummary struct {
	Match             *model.MatchFacts
	PredictionSummary *model.MatchPredictionSummary
	Prediction        *model.Prediction
}

func MatchSummaryFromGrpc(matchFacts *gen.MatchSummary) *MatchSummary {
	return &MatchSummary{
		Match:             model.MatchFactsFromGrpc(matchFacts.Match),
		PredictionSummary: model.MatchPredictionSummaryFromGrpc(matchFacts.PredictionSummary),
		Prediction:        model.PredictionFromGrpc(matchFacts.Prediction),
	}
}

func MatchSummaryToGrpc(matchFacts *MatchSummary) *gen.MatchSummary {
	return &gen.MatchSummary{
		Match:             model.MatchFactsToGrpc(matchFacts.Match),
		PredictionSummary: model.MatchPredictionSummaryToGrpc(matchFacts.PredictionSummary),
		Prediction:        model.PredictionToGrpc(matchFacts.Prediction),
	}
}

func ToUpcomingMatchesResponse(upcomingMatches map[time.Time][]*model.MatchFacts) *gen.UpcomingMatchesResponse {
	matches := make(map[string]*gen.MatchFactsList)

	for k, v := range upcomingMatches {
		date := k.Format("01-02-2006")

		var matchFacts []*gen.MatchFacts
		for i := range v {
			matchFacts = append(matchFacts, model.MatchFactsToGrpc(v[i]))
		}

		matches[date].Matches = matchFacts
	}

	return &gen.UpcomingMatchesResponse{
		Matches: matches,
	}
}
