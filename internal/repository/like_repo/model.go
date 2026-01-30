package likerepo

import (
	"github.com/google/uuid"
	"time"
)

type GormLike struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FromUserID uuid.UUID
	ToUserID   uuid.UUID
	CreatedAt  time.Time
}
