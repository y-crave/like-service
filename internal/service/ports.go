package service

import (
	"context"
	"github.com/google/uuid"
	"like-service/internal/domain"
)

type LikeService interface {
	SetLike(ctx context.Context, like domain.Like) error
}

type LikeRepo interface {
	Create(ctx context.Context, like domain.Like) error
	Read(ctx context.Context, likeID uuid.UUID) (domain.Like, error)
}
