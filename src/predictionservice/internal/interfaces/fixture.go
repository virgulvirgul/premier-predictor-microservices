//go:generate mockgen -destination=../fixture/mocks/mock_fixture.go -package=fixturemocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces FixtureService

package interfaces

import (
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
)

type FixtureService interface {
	GetMatches() ([]common.Fixture, error)
	GetTeamForm() (map[string]model.TeamForm, error)
	GetFutureFixtures() (map[string]string, error)
}
