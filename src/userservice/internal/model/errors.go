package model

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserNotInGroup     = errors.New("user not in group")
)
