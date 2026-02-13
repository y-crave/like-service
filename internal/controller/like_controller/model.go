package likecontroller

import (
	"time"
)

type JSONReaction struct {
	ID         string    `json:"id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	Type       uint8     `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
}
