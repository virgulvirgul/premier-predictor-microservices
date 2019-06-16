package prediction

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

const (
	userId   = "userId"
	userId2  = "userId2"
	matchId  = "matchId"
	matchId2 = "matchId2"
)

func Test_Repository(t *testing.T) {
	err := os.Setenv("MONGO_PORT", "27017")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	assert.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	assert.NoError(t, err)

	repository, err := NewRepository()
	assert.NoError(t, err)

	createPrediction := func(p *predictionEntity) {
		_, err = repository.
			client.
			Database(db).
			Collection(collection).
			InsertOne(
				context.Background(),
				p,
			)

		require.NoError(t, err)
	}

	cleanupDb := func() {
		_, _ = repository.
			client.
			Database(db).
			Collection(collection).
			DeleteMany(
				context.Background(),
				bson.M{},
			)
	}

	t.Run("GetPrediction", func(t *testing.T) {
		t.Run("gets prediction by userId and matchId", func(t *testing.T) {
			p := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}

			defer cleanupDb()
			createPrediction(p)

			prediction, err := repository.GetPrediction(userId, matchId)
			require.NoError(t, err)

			expectedResult := toPrediction(p)

			assert.Equal(t, expectedResult, prediction)
		})

		t.Run("returns error if not found", func(t *testing.T) {
			prediction, err := repository.GetPrediction(userId, matchId)
			require.Error(t, err)

			assert.Nil(t, prediction)
			assert.Equal(t, ErrPredictionNotFound, err)
		})
	})

	t.Run("GetPredictionsByUserId", func(t *testing.T) {
		t.Run("gets all predictions by userId", func(t *testing.T) {
			p1 := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}
			p2 := &predictionEntity{
				UserId:    userId,
				MatchId:   matchId2,
				HomeGoals: 1,
				AwayGoals: 1,
			}
			p3 := &predictionEntity{
				UserId:    userId2,
				MatchId:   matchId,
				HomeGoals: 1,
				AwayGoals: 1,
			}

			defer cleanupDb()
			createPrediction(p1)
			createPrediction(p2)
			createPrediction(p3)

			expectedResult := []model.Prediction{
				*toPrediction(p1),
				*toPrediction(p2),
			}

			predictions, err := repository.GetPredictionsByUserId(userId)
			require.NoError(t, err)

			assert.Equal(t, expectedResult, predictions)
		})
	})

	t.Run("UpdatePredictions", func(t *testing.T) {
		t.Run("inserts new predictions", func(t *testing.T) {
			predictions := []model.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 2,
					AwayGoals: 1,
				},
				{
					UserId:    userId,
					MatchId:   matchId2,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			}

			defer cleanupDb()

			err := repository.UpdatePredictions(predictions)
			require.NoError(t, err)

			result, err := repository.GetPredictionsByUserId(userId)
			require.NoError(t, err)

			assert.Equal(t, predictions, result)
		})

		t.Run("updates prediction if already exists", func(t *testing.T) {
			defer cleanupDb()

			err := repository.UpdatePredictions([]model.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			})
			require.NoError(t, err)

			predictions := []model.Prediction{
				{
					UserId:    userId,
					MatchId:   matchId,
					HomeGoals: 3,
					AwayGoals: 3,
				},
				{
					UserId:    userId,
					MatchId:   matchId2,
					HomeGoals: 2,
					AwayGoals: 1,
				},
			}

			err = repository.UpdatePredictions(predictions)
			require.NoError(t, err)

			result, err := repository.GetPredictionsByUserId(userId)
			require.NoError(t, err)

			assert.Equal(t, predictions, result)
		})
	})

	t.Run("GetMatchPredictionSummary", func(t *testing.T) {
		t.Run("gets a count of each prediction's result for a specified match	", func(t *testing.T) {
			predictions := []model.Prediction{
				{
					UserId:    "1",
					MatchId:   matchId,
					HomeGoals: 1,
					AwayGoals: 0,
				},
				{
					UserId:    "2",
					MatchId:   matchId2,
					HomeGoals: 1,
					AwayGoals: 1,
				},
				{
					UserId:    "3",
					MatchId:   matchId,
					HomeGoals: 3,
					AwayGoals: 1,
				},
				{
					UserId:    "4",
					MatchId:   matchId,
					HomeGoals: 0,
					AwayGoals: 1,
				},
			}

			err = repository.UpdatePredictions(predictions)
			require.NoError(t, err)

			homeWins, draw, awayWins, err := repository.GetMatchPredictionSummary(matchId)
			require.NoError(t, err)

			assert.Equal(t, 2, homeWins)
			assert.Equal(t, 0, draw)
			assert.Equal(t, 1, awayWins)
		})
	})
}
