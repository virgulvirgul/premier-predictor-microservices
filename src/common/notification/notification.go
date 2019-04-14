package notification

import (
	"context"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
)

type notifier struct {
	notificationClientFactory interfaces.NotificationClientFactory
}

func NewNotifier(notificationClientFactory interfaces.NotificationClientFactory) (interfaces.Notifier, error) {
	return &notifier{
		notificationClientFactory: notificationClientFactory,
	}, nil
}

func (n *notifier) Send(notification model.Notification, userIds ...string) error {
	not := &gen.Notification{
		Title:   notification.Title,
		Message: notification.Message,
	}

	ctx := context.Background()

	client, err := n.notificationClientFactory.NewNotificationClient()
	if err != nil {
		return err
	}
	defer n.notificationClientFactory.CloseConnection()

	switch len(userIds) {
	case 0:
		_, err = client.SendToAll(ctx, not)
	case 1:
		req := &gen.SingleRequest{
			UserId:       userIds[0],
			Notification: not,
		}
		_, err = client.Send(ctx, req)
	default:
		req := &gen.GroupRequest{
			UserIds:      userIds,
			Notification: not,
		}
		_, err = client.SendToGroup(ctx, req)
	}

	return err
}
