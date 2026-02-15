package service

import (
	"context"
	"github.com/google/uuid"
	"like-service/internal/domain"
)

type LikeService interface {
	SetReaction(ctx context.Context, reaction *domain.Reaction) error
}

type LikeRepo interface {
	Create(ctx context.Context, reaction *domain.Reaction) error
	Read(ctx context.Context, reactionID uuid.UUID) (domain.Reaction, error)
}
