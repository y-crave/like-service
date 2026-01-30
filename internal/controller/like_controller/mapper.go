package likecontroller

import (
	"github.com/google/uuid"
	"like-service/internal/domain"
)

func toDomainLike(like JSONLike) (domain.Like, error) {
	id, err := uuid.Parse(like.ID)
	if err != nil {

	}

	fromUserID, err := uuid.Parse(like.FromUserID)
	if err != nil {

	}

	toUserID, err := uuid.Parse(like.ToUserID)
	if err != nil {

	}

	return domain.Like{
		ID:         id,
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		CreatedAt:  like.CreatedAt,
	}, nil
}
