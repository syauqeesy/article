package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type RefreshTokenRepository interface {
	FindByToken(ctx context.Context, token string) (*model.RefreshToken, error)
	Create(ctx context.Context, refreshToken *model.RefreshToken) error
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

func (r *refreshTokenRepository) Create(ctx context.Context, refreshToken *model.RefreshToken) error {
	q := r.Database.WithContext(ctx).Create(refreshToken)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *refreshTokenRepository) Delete(ctx context.Context, refreshToken *model.RefreshToken) error {
	refreshToken.SetUpdatedAt()

	refreshToken.SetDeletedAt()

	q := r.Database.WithContext(ctx).Save(&refreshToken)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
