//go:generate mockgen -destination=../service/mocks/mock_service.go -package=predictionmocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/interfaces Service

package interfaces

type Service interface {
}
