package likecontroller

import (
	"time"
)

type JSONLike struct {
	ID         string    `json:"id"`
	FromUserID string    `json:"from_user_id"`
	ToUserID   string    `json:"to_user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
