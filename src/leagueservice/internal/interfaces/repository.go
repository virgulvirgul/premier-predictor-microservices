//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=leaguemocks github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

type Repository interface {
	GetLeagueByPin(pin int64) (*model.League, error)
	GetLeaguesByUserId(id string) ([]*model.League, error)
	AddLeague(league model.League) error
	JoinLeague(pin int64, id string) error
	LeaveLeague(pin int64, id string) error
	RenameLeague(pin int64, name string) error
}
