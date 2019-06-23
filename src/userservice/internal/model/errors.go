package model

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidRequestData = errors.New("invalid request data")
)
