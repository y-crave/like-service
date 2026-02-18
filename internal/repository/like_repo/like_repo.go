package likerepo

import (
	"context"
	"errors"
	"fmt"
	"like-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type likeRepo struct {
	db *gorm.DB
}

func NewLikeRepo(db *gorm.DB) *likeRepo {
	return &likeRepo{
		db: db,
	}
}

func (r *likeRepo) Create(ctx context.Context, reaction *domain.Reaction) error {

	gormReaction := toGormReaction(reaction)
	err := r.db.WithContext(ctx).Create(&gormReaction).Error

	if err != nil {
		return fmt.Errorf("%w, gorm error: %v", domain.InternalError, err)
	}

	return nil
}

func (r *likeRepo) Read(ctx context.Context, reactionID uuid.UUID) (domain.Reaction, error) {

	var gormReaction GormReaction

	err := r.db.WithContext(ctx).Where("id = ?", reactionID).First(&gormReaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Reaction{}, domain.ErrLikeNotFound
		}

		return domain.Reaction{}, fmt.Errorf("%w, gorm error: %v", domain.InternalError, err)

	}

	return toDomainReaction(&gormReaction), nil
}

func (r *likeRepo) ReadByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]domain.Reaction, error) {
	var gormReactions []GormReaction

	err := r.db.WithContext(ctx).
		Model(&GormReaction{}).
		Where("from_user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&gormReactions).Error

	if err != nil {
		return nil, fmt.Errorf("%w, gorm error: %v", domain.InternalError, err)
	}

	domainReactions := make([]domain.Reaction, 0, len(gormReactions))

	for i := 0; i < len(gormReactions); i++ {
		domainReactions[i] = toDomainReaction(&gormReactions[i])
	}

	return domainReactions, nil
}
