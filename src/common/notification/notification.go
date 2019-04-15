package notification

import (
	"context"
	"errors"
	gen "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/cshep4/premier-predictor-microservices/src/common/interfaces"
	"github.com/cshep4/premier-predictor-microservices/src/common/model"
	"google.golang.org/grpc/metadata"
)

type notifier struct {
	client gen.NotificationServiceClient
}

func NewNotifier(client gen.NotificationServiceClient) (interfaces.Notifier, error) {
	return &notifier{
		client: client,
	}, nil
}

func (n *notifier) Send(ctx context.Context, notification model.Notification, userIds ...string) error {
	not := &gen.Notification{
		Title:   notification.Title,
		Message: notification.Message,
	}

	metadata, err := n.createRequestMetadata(ctx)
	if err != nil {
		return err
	}

	switch len(userIds) {
	case 0:
		_, err = n.client.SendToAll(metadata, not)
	case 1:
		req := &gen.SingleRequest{
			UserId:       userIds[0],
			Notification: not,
		}
		_, err = n.client.Send(metadata, req)
	default:
		req := &gen.GroupRequest{
			UserIds:      userIds,
			Notification: not,
		}
		_, err = n.client.SendToGroup(metadata, req)
	}

	return err
}


func (n *notifier) createRequestMetadata(ctx context.Context) (context.Context, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("can't create request metadata")
	}

	if len(meta["token"]) != 1 {
		return nil, errors.New("can't create request metadata")
	}

	return metadata.AppendToOutgoingContext(context.Background(), "token", meta["token"][0]), nil
}