package service

import (
	"context"
	"like-service/internal/domain"
)

type likeService struct {
	likeRepo LikeRepo
}

func NewLikeService(likeRepo LikeRepo) LikeService {
	return &likeService{
		likeRepo: likeRepo,
	}
}

func (s *likeService) SetLike(ctx context.Context, like domain.Like) error {
	return s.likeRepo.Create(ctx, like)
}
