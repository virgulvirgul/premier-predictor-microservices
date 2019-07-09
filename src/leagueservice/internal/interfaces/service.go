//go:generate mockgen -destination=../service/mocks/mock_service.go -package=leaguemocks github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces Service

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

type Service interface {
	GetUsersLeagueList(id string) (*model.StandingsOverview, error)
	JoinUserLeague(id string, pin int64) (*model.LeagueOverview, error)
	AddUserLeague(id, name string) (*model.League, error)
	LeaveUserLeague(id string, pin int64) error
	RenameUserLeague(pin int64, name string) error
	GetLeagueTable(pin int64) ([]*model.LeagueUser, error)
	GetOverallLeagueTable() ([]*model.LeagueUser, error)
}
