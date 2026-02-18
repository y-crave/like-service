package service

import (
	"context"
	"like-service/internal/domain"
	"github.com/google/uuid"
)

type likeService struct {
	reactionRepo LikeRepo
}

func NewLikeService(reactionRepo LikeRepo) *likeService {
	return &likeService{
		reactionRepo: reactionRepo,
	}
}

func (s *likeService) SetReaction(ctx context.Context, reaction *domain.Reaction) error {
	return s.reactionRepo.Create(ctx, reaction)
}

func (s *likeService) GetReaction(ctx context.Context, reactionID uuid.UUID) (domain.Reaction, error) {
	return domain.Reaction{}, nil
}
