package likerepo

import (
	"context"
	"errors"
	"fmt"
	"like-service/internal/domain"
	"like-service/internal/service"

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
