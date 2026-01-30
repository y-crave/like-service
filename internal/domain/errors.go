package domain

import (
	"errors"
)

var (
	ErrLikeAlreadyExists = errors.New("like already exists")
	ErrLikeNotFound      = errors.New("like not found")
	InternalError        = errors.New("internal error")
)
