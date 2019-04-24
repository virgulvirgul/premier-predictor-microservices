//go:generate mockgen -destination=./mocks/mock_notification_service_client.go -package=notificationmocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen NotificationServiceClient

package notification

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/notification/mocks"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestNotifier_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	notificationClient := notificationmocks.NewMockNotificationServiceClient(ctrl)

	notifier, err := NewNotifier(notificationClient)
	assert.NoError(t, err)

	const title = "title"
	const message = "message"

	notification := model.Notification{
		Title:   title,
		Message: message,
	}

	userId1 := "1"
	userId2 := "2"
	userId3 := "3"

	tokenMap := map[string][]string{
		"token": {"token"},
	}

	ctx := metadata.NewIncomingContext(context.Background(), tokenMap)

	t.Run("If one user id is specified then notification is sent to single recipient", func(t *testing.T) {
		req := &gen.SingleRequest{
			UserId: userId1,
			Notification: &gen.Notification{
				Title:   notification.Title,
				Message: notification.Message,
			},
		}
		notificationClient.EXPECT().Send(gomock.Any(), req).Return(&empty.Empty{}, nil)

		err := notifier.Send(ctx, notification, userId1)
		assert.NoError(t, err)
	})

	t.Run("If multiple user ids are specified then notification is sent to group of recipients", func(t *testing.T) {
		ids := []string{userId1, userId2, userId3}

		req := &gen.GroupRequest{
			UserIds: ids,
			Notification: &gen.Notification{
				Title:   notification.Title,
				Message: notification.Message,
			},
		}
		notificationClient.EXPECT().SendToGroup(gomock.Any(), req).Return(&empty.Empty{}, nil)

		err := notifier.Send(ctx, notification, ids...)
		assert.NoError(t, err)
	})

	t.Run("If no user ids are specified then notification is sent to all", func(t *testing.T) {
		req := &gen.Notification{
			Title:   notification.Title,
			Message: notification.Message,
		}
		notificationClient.EXPECT().SendToAll(gomock.Any(), req).Return(&empty.Empty{}, nil)

		err := notifier.Send(ctx, notification)
		assert.NoError(t, err)
	})

	t.Run("An error is returned if there is a problem with sending notifications", func(t *testing.T) {
		req := &gen.SingleRequest{
			UserId: userId1,
			Notification: &gen.Notification{
				Title:   notification.Title,
				Message: notification.Message,
			},
		}
		e := errors.New("notification request failed")

		notificationClient.EXPECT().Send(gomock.Any(), req).Return(&empty.Empty{}, e)

		err := notifier.Send(ctx, notification, userId1)
		assert.Error(t, err)
		assert.Equal(t, e, err)
	})
}
