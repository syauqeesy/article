package repository

import (
	"context"
	"time"

	"ahmadsyauqi.dev/article/model"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, refreshToken *model.RefreshToken) error
	DeleteByAccountId(ctx context.Context, accountId string) error
}

type refreshTokenRepository repository

func (r *refreshTokenRepository) Create(ctx context.Context, refreshToken *model.RefreshToken) error {
	q := r.Database.WithContext(ctx).Create(refreshToken)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *refreshTokenRepository) DeleteByAccountId(ctx context.Context, accountId string) error {
	q := r.Database.WithContext(ctx).Model(&model.RefreshToken{}).Where("account_id = ?", accountId).Update("deleted_at = ?", time.Now().UTC().UnixMilli())
	if q.Error != nil {
		return q.Error
	}

	return nil
}
