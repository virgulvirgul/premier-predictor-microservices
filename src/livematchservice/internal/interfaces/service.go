//go:generate mockgen -destination=../service/mocks/mock_service.go -package=livemocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/interfaces Service

package interfaces

import (
	"context"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"time"
)

type Service interface {
	GetMatchSummary(ctx context.Context, req model.PredictionRequest) (*model.MatchSummary, error)
	GetMatchFacts(id string) (*common.MatchFacts, error)
	GetUpcomingMatches() (map[time.Time][]common.MatchFacts, error)
}
