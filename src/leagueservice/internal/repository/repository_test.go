package league

import (
	"context"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

const (
	pin     = int64(12345)
	name1   = "League of champions"
	name2   = "🏆🏆🏆🏆🏆🏆"
	userId1 = "1"
	userId2 = "2"
)

func TestRepository_GetLeagueByPin(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns not found if league cannot be found", func(t *testing.T) {
		cleanupDb(repository)

		league, err := repository.GetLeagueByPin(pin)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
		assert.Nil(t, league)
	})

	t.Run("gets the specified league", func(t *testing.T) {
		defer cleanupDb(repository)

		league := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1, userId2},
		}

		createLeague(fromLeague(league), repository, t)

		result, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestRepository_GetLeaguesByUserId(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("returns a slice of the leagues the user has joined", func(t *testing.T) {
		defer cleanupDb(repository)

		league1 := model.League{
			Pin:   1,
			Name:  name1,
			Users: []string{userId1, userId2},
		}

		createLeague(fromLeague(league1), repository, t)

		league2 := model.League{
			Pin:   2,
			Name:  name2,
			Users: []string{userId1},
		}

		createLeague(fromLeague(league2), repository, t)

		result, err := repository.GetLeaguesByUserId(userId1)
		require.NoError(t, err)

		assert.Equal(t, &league1, result[0])
		assert.Equal(t, &league2, result[1])
	})

	t.Run("returns an empty slice if the user has no leagues", func(t *testing.T) {
		defer cleanupDb(repository)

		league1 := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}

		createLeague(fromLeague(league1), repository, t)

		result, err := repository.GetLeaguesByUserId(userId1)
		require.NoError(t, err)

		assert.Equal(t, 0, len(result))
	})
}

func TestRepository_AddLeague(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("adds league to db", func(t *testing.T) {
		defer cleanupDb(repository)

		league := model.League{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1},
		}

		err := repository.AddLeague(league)
		require.NoError(t, err)

		result, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestRepository_JoinLeague(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("adds user to existing league", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, repository, t)

		err := repository.JoinLeague(pin, userId1)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, userId1, league.Users[1])
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(repository)

		err := repository.JoinLeague(pin, userId1)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not re-add user or return error if user is already in the league", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, repository, t)

		err := repository.JoinLeague(pin, userId1)
		require.NoError(t, err)

		err = repository.JoinLeague(pin, userId1)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, userId1, league.Users[1])
		assert.Equal(t, 2, len(league.Users))
	})
}

func TestRepository_LeaveLeague(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("removes user from league", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId1, userId2},
		}, repository, t)

		err := repository.LeaveLeague(pin, userId1)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, 1, len(league.Users))
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(repository)

		err := repository.LeaveLeague(pin, userId1)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not remove anyone from league and does not return error if user is not in the league", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, repository, t)

		err := repository.LeaveLeague(pin, userId1)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, userId2, league.Users[0])
		assert.Equal(t, 1, len(league.Users))
	})
}

func TestRepository_RenameLeague(t *testing.T) {
	repository := newTestRepository(t)
	defer repository.Close()

	t.Run("renames league", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, repository, t)

		err := repository.RenameLeague(pin, name2)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, name2, league.Name)
	})

	t.Run("returns error if league does not exist", func(t *testing.T) {
		cleanupDb(repository)

		err := repository.RenameLeague(pin, name2)
		require.Error(t, err)

		assert.Equal(t, model.ErrLeagueNotFound, err)
	})

	t.Run("does not return error if the new name is the same as the old name", func(t *testing.T) {
		defer cleanupDb(repository)

		createLeague(&leagueEntity{
			Pin:   pin,
			Name:  name1,
			Users: []string{userId2},
		}, repository, t)

		err := repository.RenameLeague(pin, name1)
		require.NoError(t, err)

		league, err := repository.GetLeagueByPin(pin)
		require.NoError(t, err)

		assert.Equal(t, pin, league.Pin)
		assert.Equal(t, name1, league.Name)
	})
}

func newTestRepository(t *testing.T) *repository {
	err := os.Setenv("MONGO_PORT", "27017")
	require.NoError(t, err)
	err = os.Setenv("MONGO_HOST", "localhost")
	require.NoError(t, err)
	err = os.Setenv("MONGO_SCHEME", "mongodb")
	require.NoError(t, err)

	repository, err := NewRepository()
	require.NoError(t, err)

	return repository
}

func createLeague(l *leagueEntity, repository *repository, t *testing.T) {
	_, err := repository.
		client.
		Database(db).
		Collection(collection).
		InsertOne(
			context.Background(),
			l,
		)

	require.NoError(t, err)
}

func cleanupDb(repository *repository) {
	_, _ = repository.
		client.
		Database(db).
		Collection(collection).
		DeleteMany(
			context.Background(),
			bson.M{},
		)
}
