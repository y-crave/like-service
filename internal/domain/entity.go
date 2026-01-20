package domain

import (
	"github.com/google/uuid"
	"time"
)

type Like struct{
	ID uuid.UUID
	FromUserID uuid.UUID
	ToUserID uuid.UUID
	CreatedAt time.Time
}

type Dislike struct{
	ID uuid.UUID
	FromUserID uuid.UUID
	ToUserID uuid.UUID
	CreatedAt time.Time
}

type Match struct{
	ID uuid.UUID
	UserFirst uuid.UUID
	UserSecond uuid.UUID
	CreatedAt time.Time
}