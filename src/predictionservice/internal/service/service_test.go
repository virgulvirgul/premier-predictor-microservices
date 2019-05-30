package prediction

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/fixture/mocks"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	userId  = "1"
	matchId = "1"
	matchId2 = "2"
	teamOne = "Team 1"
	teamTwo = "Team 2"
)

var (
	now = time.Now()
)

func TestService_GetFixturesWithPredictions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fixtureService := fixturemocks.NewMockFixtureService(ctrl)
	repository := predictionmocks.NewMockRepository(ctrl)

	service, err := NewService(repository, fixtureService)
	require.NoError(t, err)

	t.Run("Gets fixtures and predictions and merges them", func(t *testing.T) {
		fixtures := []common.Fixture{
			{
				Id: matchId,
				HomeTeam: teamOne,
				AwayTeam: teamTwo,
				HomeGoals: 1,
				AwayGoals: 1,
				Played: 1,
				DateTime: now,
				Matchday: 1,
			},
			{
				Id: matchId2,
				HomeTeam: teamTwo,
				AwayTeam: teamOne,
				HomeGoals: 1,
				AwayGoals: 2,
				Played: 1,
				DateTime: now,
				Matchday: 2,
			},
		}

		predictions := []common.Prediction{
			{
				MatchId:   matchId,
				UserId:    userId,
				HomeGoals: 1,
				AwayGoals: 1,
			},
		}

		fixtureService.EXPECT().GetMatches().Return(fixtures)
		repository.EXPECT().
	})

	t.Run("Error if there is a problem getting fixtureService from FixtureService", func(t *testing.T) {

	})

	t.Run("Error if there is a problem getting predictions from db", func(t *testing.T) {

	})
}
