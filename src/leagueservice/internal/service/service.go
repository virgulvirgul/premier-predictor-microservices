package league

import (
	"github.com/cshep4/premier-predictor-microservices/src/common/timer"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"
	"sort"
	"time"
)

const (
	timeSubtractor = 1500000000000
)

type service struct {
	repository  interfaces.Repository
	userService interfaces.UserService
	time        timer.Time
}

func NewService(repository interfaces.Repository, userService interfaces.UserService, time timer.Time) (interfaces.Service, error) {
	return &service{
		repository:  repository,
		userService: userService,
		time:        time,
	}, nil
}

type userLeaguesResult struct {
	result []*model.League
	err    error
}
type countResult struct {
	result int64
	err    error
}
type leagueOverviewResult struct {
	result model.LeagueOverview
	err    error
}

func (s *service) GetUsersLeagueList(id string) (*model.StandingsOverview, error) {
	userLeaguesChan := make(chan userLeaguesResult)
	userCountChan := make(chan countResult)
	overallRankChan := make(chan countResult)

	go func() {
		r, err := s.repository.GetLeaguesByUserId(id)
		userLeaguesChan <- userLeaguesResult{result: r, err: err}
	}()
	go func() {
		r, err := s.userService.GetOverallRank(id)
		overallRankChan <- countResult{result: r, err: err}
	}()
	go func() {
		r, err := s.userService.GetUserCount()
		userCountChan <- countResult{result: r, err: err}
	}()

	userLeagues := <-userLeaguesChan
	if userLeagues.err != nil {
		return nil, userLeagues.err
	}

	leagueOverviewChans := make(chan leagueOverviewResult, len(userLeagues.result))

	leagueOverviews := []model.LeagueOverview{}
	for _, l := range userLeagues.result {
		go func(league *model.League) {
			leagueOverview := model.LeagueOverview{
				Pin:        league.Pin,
				LeagueName: league.Name,
			}

			rank, err := s.userService.GetLeagueRank(id, league.Users)

			leagueOverview.Rank = rank

			leagueOverviewChans <- leagueOverviewResult{result: leagueOverview, err: err}
		}(l)
	}

	for range userLeagues.result {
		leagueResult := <-leagueOverviewChans
		if leagueResult.err != nil {
			return nil, leagueResult.err
		}

		leagueOverviews = append(leagueOverviews, leagueResult.result)
	}

	overallRank := <-overallRankChan
	if overallRank.err != nil {
		return nil, overallRank.err
	}

	userCount := <-userCountChan
	if userCount.err != nil {
		return nil, userCount.err
	}

	return &model.StandingsOverview{
		UserLeagues: leagueOverviews,
		OverallLeagueOverview: model.OverallLeagueOverview{
			Rank:      overallRank.result,
			UserCount: userCount.result,
		},
	}, nil
}

func (s *service) JoinUserLeague(id string, pin int64) (*model.LeagueOverview, error) {
	league, err := s.repository.GetLeagueByPin(pin)
	if err != nil {
		return nil, err
	}

	ids := append(league.Users, id)

	rank, err := s.userService.GetLeagueRank(id, ids)
	if err != nil {
		return nil, err
	}

	err = s.repository.JoinLeague(pin, id)
	if err != nil {
		return nil, err
	}

	return &model.LeagueOverview{
		LeagueName: league.Name,
		Pin:        league.Pin,
		Rank:       rank,
	}, nil
}

func (s *service) AddUserLeague(id, name string) (*model.League, error) {
	league := model.League{
		Pin:   s.time.Now().UnixNano()/int64(time.Millisecond) - timeSubtractor,
		Name:  name,
		Users: []string{id},
	}

	err := s.repository.AddLeague(league)
	if err != nil {
		return nil, err
	}

	return &league, nil
}

func (s *service) LeaveUserLeague(id string, pin int64) error {
	return s.repository.LeaveLeague(pin, id)
}

func (s *service) RenameUserLeague(pin int64, name string) error {
	return s.repository.RenameLeague(pin, name)
}

func (s *service) GetLeagueTable(pin int64) ([]*model.LeagueUser, error) {
	league, err := s.repository.GetLeagueByPin(pin)
	if err != nil {
		return nil, err
	}

	users, err := s.userService.GetLeagueUsers(league.Users)
	if err != nil {
		return nil, err
	}

	return s.sortUsers(users), nil
}

func (s *service) GetOverallLeagueTable() ([]*model.LeagueUser, error) {
	users, err := s.userService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return s.sortUsers(users), nil
}

func (s *service) sortUsers(users []*model.LeagueUser) model.LeagueUserSlice {
	leagueUsers := model.LeagueUserSlice(users)
	sort.Sort(leagueUsers)

	return leagueUsers
}
