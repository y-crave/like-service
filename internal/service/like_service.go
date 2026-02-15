package service

import (
	"context"
	"like-service/internal/domain"
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
