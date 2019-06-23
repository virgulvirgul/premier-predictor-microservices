package util

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strings"
)

func CreateRequestMetadata(ctx context.Context) (context.Context, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("can't create request metadata")
	}

	if len(meta["token"]) != 1 {
		return nil, errors.New("can't create request metadata")
	}

	return metadata.AppendToOutgoingContext(context.Background(), "token", meta["token"][0]), nil
}

func GetErrorMessage(err error) string {
	return err.Error()[:strings.IndexByte(err.Error(), ':')]
}