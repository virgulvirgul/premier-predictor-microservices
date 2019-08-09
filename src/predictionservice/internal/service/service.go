package prediction

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"time"
)

type service struct {
	repository     interfaces.Repository
	fixtureService interfaces.FixtureService
}

func NewService(repository interfaces.Repository, fixtureService interfaces.FixtureService) (interfaces.Service, error) {
	return &service{
		repository:     repository,
		fixtureService: fixtureService,
	}, nil
}

func (s *service) GetFixturesWithPredictions(id string) ([]model.FixturePrediction, error) {
	fixturesChan := s.getFixturesAsync()
	predictionsChan := s.getPredictionsAsync(id)

	fResult := <-fixturesChan
	if fResult.err != nil {
		return nil, fResult.err
	}
	fixtures := fResult.result

	pResult := <-predictionsChan
	if pResult.err != nil {
		return nil, pResult.err
	}
	predictions := pResult.result

	fixturePredictions := []model.FixturePrediction{}
	for _, f := range fixtures {
		fp := model.FixturePrediction{
			UserId:     id,
			Id:         f.Id,
			HomeTeam:   f.HomeTeam,
			AwayTeam:   f.AwayTeam,
			HomeResult: f.HomeGoals,
			AwayResult: f.AwayGoals,
			Played:     f.Played,
			DateTime:   f.DateTime,
			Matchday:   f.Matchday,
		}

		for _, p := range predictions {
			if p.MatchId == f.Id {
				fp.HomeGoals = &p.HomeGoals
				fp.AwayGoals = &p.AwayGoals
				break
			}
		}

		fixturePredictions = append(fixturePredictions, fp)
	}

	return fixturePredictions, nil
}

func (s *service) getFixturesAsync() chan fixturesResult {
	fixturesChan := make(chan fixturesResult)
	go func() { fixturesChan <- s.getFixtures() }()
	return fixturesChan
}

func (s *service) getFixtures() fixturesResult {
	r, e := s.fixtureService.GetMatches()
	return fixturesResult{result: r, err: e}
}

func (s *service) getPredictionsAsync(id string) chan predictionsResult {
	predictionsChan := make(chan predictionsResult)
	go func() { predictionsChan <- s.getPredictions(id) }()
	return predictionsChan
}

func (s *service) getPredictions(id string) predictionsResult {
	r, e := s.repository.GetPredictionsByUserId(id)
	return predictionsResult{result: r, err: e}
}

func (s *service) GetPredictorData(id string) (*model.PredictorData, error) {
	fixturePredictionsChan := s.getFixturesWithPredictionsAsync(id)
	formChan := s.getTeamFormAsync()

	fixturePredictionsResult := <-fixturePredictionsChan
	if fixturePredictionsResult.err != nil {
		return nil, fixturePredictionsResult.err
	}
	fixturePredictions := fixturePredictionsResult.result

	formResult := <-formChan
	if formResult.err != nil {
		return nil, formResult.err
	}
	forms := formResult.result

	return &model.PredictorData{
		Predictions: fixturePredictions,
		Forms:       forms,
	}, nil
}

func (s *service) getFixturesWithPredictionsAsync(id string) chan fixturePredictionsResult {
	fixturePredictionsChan := make(chan fixturePredictionsResult)
	go func() { fixturePredictionsChan <- s.getFixturesWithPredictions(id) }()
	return fixturePredictionsChan
}

func (s *service) getFixturesWithPredictions(id string) fixturePredictionsResult {
	r, e := s.GetFixturesWithPredictions(id)
	return fixturePredictionsResult{result: r, err: e}
}

func (s *service) getTeamFormAsync() chan formResult {
	formChan := make(chan formResult)
	go func() { formChan <- s.getTeamForm() }()
	return formChan
}

func (s *service) getTeamForm() formResult {
	r, e := s.fixtureService.GetTeamForm()
	return formResult{result: r, err: e}
}

func (s *service) GetUsersPastPredictions(id string) ([]model.FixturePrediction, error) {
	fixturePredictions, err := s.GetFixturesWithPredictions(id)
	if err != nil {
		return nil, err
	}

	var fp []model.FixturePrediction
	for _, f := range fixturePredictions {
		if time.Now().UTC().After(f.DateTime) {
			fp = append(fp, f)
		}
	}

	return fp, nil
}

func (s *service) UpdatePredictions(predictions []common.Prediction) error {
	futureFixtures, err := s.fixtureService.GetFutureFixtures()
	if err != nil {
		return err
	}

	var validPredictions []common.Prediction
	for _, p := range predictions {
		if _, ok := futureFixtures[p.MatchId]; ok {
			validPredictions = append(validPredictions, p)
		}
	}

	return s.repository.UpdatePredictions(validPredictions)
}

func (s *service) GetPrediction(userId, matchId string) (*common.Prediction, error) {
	return s.repository.GetPrediction(userId, matchId)
}

func (s *service) GetMatchPredictionSummary(id string) (*common.MatchPredictionSummary, error) {
	homeWins, draw, awayWins, err := s.repository.GetMatchPredictionSummary(id)
	if err != nil {
		return nil, err
	}

	return &common.MatchPredictionSummary{
		HomeWin: homeWins,
		Draw:    draw,
		AwayWin: awayWins,
	}, nil
}
