package live

import (
	"context"
	"errors"
	"github.com/ahl5esoft/golang-underscore"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"sort"
	"time"
)

type service struct {
	repository interfaces.Repository
	predictor  interfaces.Predictor
}

func NewService(repository interfaces.Repository, predictor interfaces.Predictor) (interfaces.Service, error) {
	return &service{
		repository: repository,
		predictor:  predictor,
	}, nil
}

type predictionResult struct {
	result *common.Prediction
	err    error
}
type matchPredictionSummaryResult struct {
	result *common.MatchPredictionSummary
	err    error
}

type matchFactsResult struct {
	result *common.MatchFacts
	err    error
}

func (s *service) GetMatchSummary(ctx context.Context, req model.PredictionRequest) (*model.MatchSummary, error) {
	predictionChan := make(chan predictionResult)
	matchPredictionSummaryChan := make(chan matchPredictionSummaryResult)
	matchFactsChan := make(chan matchFactsResult)

	go func() {
		r, e := s.predictor.GetPrediction(ctx, req)
		predictionChan <- predictionResult{result: r, err: e}
	}()

	go func() {
		r, e := s.predictor.GetPredictionSummary(ctx, req.MatchId)
		matchPredictionSummaryChan <- matchPredictionSummaryResult{result: r, err: e}
	}()

	go func() {
		r, e := s.repository.GetMatchFacts(req.MatchId)
		matchFactsChan <- matchFactsResult{result: r, err: e}
	}()

	prediction := <-predictionChan
	if prediction.err != nil {
		return nil, prediction.err
	}

	matchPredictionSummary := <-matchPredictionSummaryChan
	if matchPredictionSummary.err != nil {
		return nil, matchPredictionSummary.err
	}

	matchFacts := <-matchFactsChan
	if matchFacts.err != nil {
		return nil, matchFacts.err
	}

	return &model.MatchSummary{
		Match:             matchFacts.result,
		PredictionSummary: matchPredictionSummary.result,
		Prediction:        prediction.result,
	}, nil
}

func (s *service) GetMatchFacts(id string) (*common.MatchFacts, error) {
	return s.repository.GetMatchFacts(id)
}

func (s *service) GetUpcomingMatches() (map[time.Time][]common.MatchFacts, error) {
	matches, err := s.repository.GetUpcomingMatches()
	if err != nil {
		return nil, err
	}

	sort.Sort(common.MatchFactsSlice(matches))

	upcomingMatches := make(map[time.Time][]common.MatchFacts)

	underscore.Chain(matches).
		Group(s.groupByMatchDate).
		Value(&upcomingMatches)

	if err := recover(); err != nil {
		return nil, errors.New("could not map matches")
	}

	return upcomingMatches, nil
}

func (s *service) groupByMatchDate(m common.MatchFacts, _ int) time.Time {
	return m.MatchDate
}
