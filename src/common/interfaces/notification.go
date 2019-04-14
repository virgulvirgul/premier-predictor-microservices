//go:generate mockgen -destination=../notification/mocks/mock_notification_client.go -package=notificationmocks github.com/cshep4/premier-predictor-microservices/src/common/interfaces Notifier

package interfaces

import (
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
)

type Notifier interface {
	Send(notification model.Notification, userIds ...string) error
}