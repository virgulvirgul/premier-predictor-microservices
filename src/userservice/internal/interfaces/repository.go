//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=usermocks github.com/cshep4/premier-predictor-microservices/src/userservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

type Repository interface {
	GetUserById(id string) (*model.User, error)
	UpdateUserInfo(userInfo model.UserInfo) error
	UpdatePassword(id, password string) error
	GetAllUsers() ([]*model.User, error)
	GetAllUsersByIds(ids []string) ([]*model.User, error)
	IsEmailTakenByADifferentUser(id, email string) bool
	GetOverallRank(id string) (int64, error)
	GetRankForGroup(id string, ids []string) (int64, error)
	GetUserCount() (int64, error)
}
