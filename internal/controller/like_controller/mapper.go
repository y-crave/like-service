package likecontroller

import (
	"github.com/google/uuid"
	"like-service/internal/domain"
)

func toDomainReaction(reaction JSONReaction) (domain.Reaction, error) {
	id, err := uuid.Parse(reaction.ID)
	if err != nil {
		return domain.Reaction{}, err
	}

	fromUserID, err := uuid.Parse(reaction.FromUserID)
	if err != nil {
		return domain.Reaction{}, err
	}

	toUserID, err := uuid.Parse(reaction.ToUserID)
	if err != nil {
		return domain.Reaction{}, err
	}

	return domain.Reaction{
		ID:         id,
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Type:       domain.ReactionType(reaction.Type),
		CreatedAt:  reaction.CreatedAt,
	}, nil
}
