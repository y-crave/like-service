package likerepo

import (
	"like-service/internal/domain"
)

func toDomainReaction(
	reaction *GormReaction,
) domain.Reaction {
	return domain.Reaction{
		ID:         reaction.ID,
		FromUserID: reaction.FromUserID,
		ToUserID:   reaction.ToUserID,

		Type: domain.ReactionType(reaction.Type),

		CreatedAt: reaction.CreatedAt,
	}
}

func toGormReaction(
	reaction *domain.Reaction,
) GormReaction {
	return GormReaction{
		ID:         reaction.ID,
		FromUserID: reaction.FromUserID,
		ToUserID:   reaction.ToUserID,

		Type: ReactionType(reaction.Type),

		CreatedAt: reaction.CreatedAt,
	}
}
