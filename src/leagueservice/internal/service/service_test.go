package league

import (
	"github.com/cshep4/premier-predictor-microservices/src/common/timer/mocks"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/repository/mocks"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var e = errors.New("error")

const (
	id1   = "🆔"
	id2   = "⚽️"
	id3   = "🤙️"
	pin1  = int64(12345)
	pin2  = int64(67890)
	name1 = "League of champions"
	name2 = "🏆🏆🏆🏆🏆🏆"

	count       = int64(1234)
	overallRank = int64(123)
	league1Rank = int64(12)
	league2Rank = int64(1)
)

var (
	users1 = []string{id1, id2}
	users2 = []string{id1, id3}
)

func TestService_GetUsersLeagueList(t *testing.T) {
	leagues := []*model.League{
		{
			Pin:   pin1,
			Name:  name1,
			Users: users1,
		},
		{
			Pin:   pin2,
			Name:  name2,
			Users: users2,
		},
	}

	e := errors.New("error")

	t.Run("returns error if there is a problem getting info for a league from db", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := leaguemocks.NewMockRepository(ctrl)
		userService := usermocks.NewMockUserService(ctrl)
		timer := timermocks.NewMockTime(ctrl)

		service, err := NewService(repository, userService, timer)
		require.NoError(t, err)

		repository.EXPECT().GetLeaguesByUserId(id1).MaxTimes(1).Return(nil, e)
		userService.EXPECT().GetUserCount().MaxTimes(1).Return(count, nil)
		userService.EXPECT().GetOverallRank(id1).MaxTimes(1).Return(overallRank, nil)

		result, err := service.GetUsersLeagueList(id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting overall rank", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := leaguemocks.NewMockRepository(ctrl)
		userService := usermocks.NewMockUserService(ctrl)
		timer := timermocks.NewMockTime(ctrl)

		service, err := NewService(repository, userService, timer)
		require.NoError(t, err)

		repository.EXPECT().GetLeaguesByUserId(id1).MaxTimes(1).Return(leagues, nil)
		userService.EXPECT().GetOverallRank(id1).Return(int64(0), e)
		userService.EXPECT().GetUserCount().MaxTimes(1).Return(count, nil)
		userService.EXPECT().GetLeagueRank(id1, users1).MaxTimes(1).Return(league1Rank, nil)
		userService.EXPECT().GetLeagueRank(id1, users2).MaxTimes(1).Return(league2Rank, nil)

		result, err := service.GetUsersLeagueList(id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting overall user count", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := leaguemocks.NewMockRepository(ctrl)
		userService := usermocks.NewMockUserService(ctrl)
		timer := timermocks.NewMockTime(ctrl)

		service, err := NewService(repository, userService, timer)
		require.NoError(t, err)

		repository.EXPECT().GetLeaguesByUserId(id1).MaxTimes(1).Return(leagues, nil)
		userService.EXPECT().GetOverallRank(id1).MaxTimes(1).Return(overallRank, nil)
		userService.EXPECT().GetUserCount().Return(int64(0), e)
		userService.EXPECT().GetLeagueRank(id1, users1).MaxTimes(1).Return(league1Rank, nil)
		userService.EXPECT().GetLeagueRank(id1, users2).MaxTimes(1).Return(league2Rank, nil)

		result, err := service.GetUsersLeagueList(id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns error if there is a problem getting rank for a league", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := leaguemocks.NewMockRepository(ctrl)
		userService := usermocks.NewMockUserService(ctrl)
		timer := timermocks.NewMockTime(ctrl)

		service, err := NewService(repository, userService, timer)
		require.NoError(t, err)

		repository.EXPECT().GetLeaguesByUserId(id1).MaxTimes(1).Return(leagues, nil)
		userService.EXPECT().GetUserCount().MaxTimes(1).Return(count, nil)
		userService.EXPECT().GetOverallRank(id1).MaxTimes(1).Return(overallRank, nil)
		userService.EXPECT().GetLeagueRank(id1, users1).Return(int64(0), e)
		userService.EXPECT().GetLeagueRank(id1, users2).MaxTimes(1).Return(league2Rank, nil)

		result, err := service.GetUsersLeagueList(id1)
		require.Error(t, err)

		assert.Empty(t, result)
	})

	t.Run("returns standings overview from the user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repository := leaguemocks.NewMockRepository(ctrl)
		userService := usermocks.NewMockUserService(ctrl)
		timer := timermocks.NewMockTime(ctrl)

		service, err := NewService(repository, userService, timer)
		require.NoError(t, err)

		repository.EXPECT().GetLeaguesByUserId(id1).Return(leagues, nil)
		userService.EXPECT().GetUserCount().Return(count, nil)
		userService.EXPECT().GetOverallRank(id1).Return(overallRank, nil)
		userService.EXPECT().GetLeagueRank(id1, users1).Return(league1Rank, nil)
		userService.EXPECT().GetLeagueRank(id1, users2).Return(league2Rank, nil)

		expectedResult := &model.StandingsOverview{
			OverallLeagueOverview: model.OverallLeagueOverview{
				Rank:      overallRank,
				UserCount: count,
			},
			UserLeagues: []model.LeagueOverview{
				{
					Rank:       league1Rank,
					LeagueName: name1,
					Pin:        pin1,
				},
				{
					Rank:       league2Rank,
					LeagueName: name2,
					Pin:        pin2,
				},
			},
		}

		result, err := service.GetUsersLeagueList(id1)
		require.NoError(t, err)

		assert.Equal(t, expectedResult.OverallLeagueOverview, result.OverallLeagueOverview)

		if result.UserLeagues[0].Pin == pin1 {
			assert.Equal(t, expectedResult.UserLeagues[0], result.UserLeagues[0])
			assert.Equal(t, expectedResult.UserLeagues[1], result.UserLeagues[1])
		} else {
			assert.Equal(t, expectedResult.UserLeagues[1], result.UserLeagues[0])
			assert.Equal(t, expectedResult.UserLeagues[0], result.UserLeagues[1])
		}
	})
}

func TestService_JoinUserLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	league := &model.League{
		Pin:   pin1,
		Name:  name1,
		Users: []string{id2, id3},
	}

	users := []string{id2, id3, id1}

	t.Run("Returns error if it cannot get league info", func(t *testing.T) {
		repository.EXPECT().GetLeagueByPin(pin1).Return(nil, e)

		result, err := service.JoinUserLeague(id1, pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Returns error if it cannot get league rank", func(t *testing.T) {
		repository.EXPECT().GetLeagueByPin(pin1).Return(league, nil)
		userService.EXPECT().GetLeagueRank(id1, users).Return(int64(0), e)

		result, err := service.JoinUserLeague(id1, pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Returns error if user cannot join league", func(t *testing.T) {
		repository.EXPECT().GetLeagueByPin(pin1).Return(league, nil)
		userService.EXPECT().GetLeagueRank(id1, users).Return(league1Rank, nil)
		repository.EXPECT().JoinLeague(pin1, id1).Return(e)

		result, err := service.JoinUserLeague(id1, pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets league info and joins", func(t *testing.T) {
		leagueOverview := &model.LeagueOverview{
			Rank:       league1Rank,
			LeagueName: name1,
			Pin:        pin1,
		}

		repository.EXPECT().GetLeagueByPin(pin1).Return(league, nil)
		userService.EXPECT().GetLeagueRank(id1, users).Return(league1Rank, nil)
		repository.EXPECT().JoinLeague(pin1, id1).Return(nil)

		result, err := service.JoinUserLeague(id1, pin1)
		require.NoError(t, err)

		assert.Equal(t, leagueOverview, result)
	})
}

func TestService_AddUserLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	const currentTime = 1512345678912
	const pin = int64(12345678912)
	timer.EXPECT().Now().AnyTimes().Return(time.Unix(0, currentTime*int64(time.Millisecond)))

	league := model.League{
		Pin:   pin,
		Name:  name1,
		Users: []string{id1},
	}

	t.Run("returns error if league cannot be added", func(t *testing.T) {
		repository.EXPECT().AddLeague(league).Return(e)

		result, err := service.AddUserLeague(id1, name1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("returns league once added", func(t *testing.T) {
		repository.EXPECT().AddLeague(league).Return(nil)

		result, err := service.AddUserLeague(id1, name1)
		require.NoError(t, err)

		assert.Equal(t, &league, result)
	})
}

func TestService_LeaveUserLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().LeaveLeague(pin1, id1).Return(e)

		err := service.LeaveUserLeague(id1, pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
	})

	t.Run("Adds user to league", func(t *testing.T) {
		repository.EXPECT().LeaveLeague(pin1, id1).Return(nil)

		err := service.LeaveUserLeague(id1, pin1)
		require.NoError(t, err)
	})
}

func TestService_RenameUserLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem", func(t *testing.T) {
		repository.EXPECT().RenameLeague(pin1, name2).Return(e)

		err := service.RenameUserLeague(pin1, name2)
		require.Error(t, err)

		assert.Equal(t, e, err)
	})

	t.Run("Renames league", func(t *testing.T) {
		repository.EXPECT().RenameLeague(pin1, name2).Return(nil)

		err := service.RenameUserLeague(pin1, name2)
		require.NoError(t, err)
	})
}

func TestService_GetLeagueTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	league := &model.League{
		Pin:   pin1,
		Users: users1,
	}

	t.Run("Returns error if there is a problem getting league", func(t *testing.T) {
		repository.EXPECT().GetLeagueByPin(pin1).Return(nil, e)

		result, err := service.GetLeagueTable(pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Returns error if there is a problem getting users", func(t *testing.T) {
		repository.EXPECT().GetLeagueByPin(pin1).Return(league, nil)
		userService.EXPECT().GetLeagueUsers(users1).Return(nil, e)

		result, err := service.GetLeagueTable(pin1)
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets users and returns them sorted by points", func(t *testing.T) {
		u1 := &model.LeagueUser{
			Score: 10,
		}
		u2 := &model.LeagueUser{
			Score: 30,
		}
		u3 := &model.LeagueUser{
			Score: 20,
		}

		repository.EXPECT().GetLeagueByPin(pin1).Return(league, nil)
		userService.EXPECT().GetLeagueUsers(users1).Return([]*model.LeagueUser{u1, u2, u3}, nil)

		result, err := service.GetLeagueTable(pin1)
		require.NoError(t, err)

		expectedResult := []*model.LeagueUser{u2, u3, u1}

		assert.Equal(t, expectedResult, result)
	})
}

func TestService_GetOverallLeagueTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := leaguemocks.NewMockRepository(ctrl)
	userService := usermocks.NewMockUserService(ctrl)
	timer := timermocks.NewMockTime(ctrl)

	service, err := NewService(repository, userService, timer)
	require.NoError(t, err)

	t.Run("Returns error if there is a problem getting users", func(t *testing.T) {
		userService.EXPECT().GetAllUsers().Return(nil, e)

		result, err := service.GetOverallLeagueTable()
		require.Error(t, err)

		assert.Equal(t, e, err)
		assert.Empty(t, result)
	})

	t.Run("Gets users and returns them sorted by points", func(t *testing.T) {
		u1 := &model.LeagueUser{
			Score: 10,
		}
		u2 := &model.LeagueUser{
			Score: 30,
		}
		u3 := &model.LeagueUser{
			Score: 20,
		}

		userService.EXPECT().GetAllUsers().Return([]*model.LeagueUser{u1, u2, u3}, nil)

		result, err := service.GetOverallLeagueTable()
		require.NoError(t, err)

		expectedResult := []*model.LeagueUser{u2, u3, u1}

		assert.Equal(t, expectedResult, result)
	})
}
