package likerepo

import (
	"like-service/internal/domain"
)

func toDomainLike(like GormLike) domain.Like {
	return domain.Like{
		ID:         like.ID,
		FromUserID: like.FromUserID,
		ToUserID:   like.ToUserID,
		CreatedAt:  like.CreatedAt,
	}
}

func toGormLike(like domain.Like) GormLike {
	return GormLike{
		ID:         like.ID,
		FromUserID: like.FromUserID,
		ToUserID:   like.ToUserID,
		CreatedAt:  like.CreatedAt,
	}
}
