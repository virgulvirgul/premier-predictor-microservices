//go:generate mockgen -destination=../service/mocks/mock_service.go -package=usermocks github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces Service

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

type Service interface {
	GetUser(id string) (*model.User, error)
	UpdateUserInfo(userDetails model.UserInfo) error
	UpdatePassword(updatePassword model.UpdatePassword) error
	GetUserScore(id string) (int, error)
	GetAllUsers() ([]*model.User, error)
	GetAllUsersByIds(ids []string) ([]*model.User, error)
	GetRankForGroup(id string, ids []string) (int64, error)
	GetOverallRank(id string) (int64, error)
	GetUserCount() (int64, error)
}
