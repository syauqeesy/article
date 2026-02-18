package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	FindByToken(ctx context.Context, token string) (*model.RefreshToken, error)
	CreateTx(ctx context.Context, tx *gorm.DB, refreshToken *model.RefreshToken) error
	Delete(ctx context.Context, refreshToken *model.RefreshToken) error
}

type refreshTokenRepository repository

func (r *refreshTokenRepository) FindByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	refreshToken := &model.RefreshToken{}

	q := r.Database.WithContext(ctx).Where("token = ?", token).Where("deleted_at IS NULL").First(&refreshToken)
	if q.Error != nil {
		return nil, q.Error
	}

	return refreshToken, nil
}

func (r *refreshTokenRepository) CreateTx(ctx context.Context, tx *gorm.DB, refreshToken *model.RefreshToken) error {
	q := tx.WithContext(ctx).Create(refreshToken)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *refreshTokenRepository) Delete(ctx context.Context, refreshToken *model.RefreshToken) error {
	var err error

	err = refreshToken.SetUpdatedAt()
	if err != nil {
		return err
	}

	err = refreshToken.SetDeletedAt()
	if err != nil {
		return err
	}

	q := r.Database.WithContext(ctx).Save(&refreshToken)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
