//go:generate mockgen -destination=../user/mocks/mock_user.go -package=usermocks github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/interfaces UserService

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

type UserService interface {
	GetAllUsers() ([]*model.LeagueUser, error)
	GetLeagueUsers(ids []string) ([]*model.LeagueUser, error)
	GetOverallRank(id string) (int64, error)
	GetLeagueRank(id string, ids []string) (int64, error)
	GetUserCount() (int64, error)
}
