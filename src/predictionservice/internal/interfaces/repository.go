//go:generate mockgen -destination=../repository/mocks/mock_repository.go -package=predictionmocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces Repository

package interfaces

import "github.com/cshep4/premier-predictor-microservices/src/common/model"

type Repository interface {
	GetUpcomingMatches() ([]*model.MatchFacts, error)
	GetMatchFacts(id string) (*model.MatchFacts, error)
}
