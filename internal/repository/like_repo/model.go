package likerepo

import (
	"github.com/google/uuid"
	"time"
)

type ReactionType uint8

const (
	ReactionTypeLike ReactionType = iota
	ReactionTypeDislike
)

type GormReaction struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	FromUserID uuid.UUID `gorm:"type:uuid;not null"`
	ToUserID   uuid.UUID `gorm:"type:uuid;not null"`

	Type ReactionType `gorm:"not null"`

	CreatedAt time.Time
}

func (GormReaction) TableName() string {
	return "reactions"
}
