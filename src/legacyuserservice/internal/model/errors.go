package model

import (
	"errors"
)

var (
	ErrLegacyUserNotFound = errors.New("legacyUser not found")
	ErrLegacyLoginFailed  = errors.New("legacy login failed")
)
