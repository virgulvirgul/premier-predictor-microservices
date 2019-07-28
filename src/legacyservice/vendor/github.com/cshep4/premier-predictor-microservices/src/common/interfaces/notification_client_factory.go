//go:generate mockgen -destination=../factory/mocks/mock_notification_client_factory.go -package=factorymocks github.com/cshep4/premier-predictor-microservices/src/common/interfaces NotificationClientFactory

package interfaces

import "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"

type NotificationClientFactory interface {
	NewNotificationClient() (model.NotificationServiceClient, error)
	CloseConnection() error
}
