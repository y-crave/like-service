package service

import (
	"context"
	"github.com/google/uuid"
	"like-service/internal/domain"
)

type LikeService interface {
	SetReaction(ctx context.Context, reaction *domain.Reaction) error
	GetReaction(ctx context.Context, reactionID uuid.UUID) (domain.Reaction, error)
	GetReactionsByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.Reaction, error)
}

type LikeRepo interface {
	Create(ctx context.Context, reaction *domain.Reaction) error
	Read(ctx context.Context, reactionID uuid.UUID) (domain.Reaction, error)
	ReadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.Reaction, error)
}
