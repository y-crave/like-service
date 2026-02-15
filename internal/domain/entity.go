package domain

import (
	"github.com/google/uuid"
	"time"
)

type ReactionType uint8

const (
	ReactionTypeLike ReactionType = iota
	ReactionTypeDislike
)

type Reaction struct {
	ID         uuid.UUID
	FromUserID uuid.UUID
	ToUserID   uuid.UUID

	Type ReactionType

	CreatedAt time.Time
}

type Match struct {
	ID         uuid.UUID
	UserFirst  uuid.UUID
	UserSecond uuid.UUID
	CreatedAt  time.Time
}

