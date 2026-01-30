package likerepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"like-service/internal/domain"
	"like-service/internal/service"
)

type likeRepo struct {
	db *gorm.DB
}

func NewLikeRepo(db *gorm.DB) service.LikeRepo {
	return &likeRepo{
		db: db,
	}
}

func (r *likeRepo) Create(ctx context.Context, like domain.Like) error {

	gormLike := toGormLike(like)

	err := r.db.WithContext(ctx).Create(&gormLike).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrLikeAlreadyExists
		}

		return fmt.Errorf("%w, gorm error: %v", domain.InternalError, err)
	}

	return nil
}

func (r *likeRepo) Read(ctx context.Context, likeID uuid.UUID) (domain.Like, error) {

	var gormLike GormLike

	err := r.db.WithContext(ctx).Where("id = ?", likeID).First(&gormLike).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Like{}, domain.ErrLikeNotFound
		}

		return domain.Like{}, fmt.Errorf("%w, gorm error: %v", domain.InternalError, err)

	}

	return toDomainLike(gormLike), nil
}
