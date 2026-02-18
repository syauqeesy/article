package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type AccountIdentityRepository interface {
	FindByProviderAndProviderUserId(ctx context.Context, provider string, providerUserId string) (*model.AccountIdentity, error)
	CreateTx(ctx context.Context, tx *gorm.DB, accountIdentity *model.AccountIdentity) error
}

type accountIdentityRepository repository

func (r *accountIdentityRepository) FindByProviderAndProviderUserId(ctx context.Context, provider string, providerUserId string) (*model.AccountIdentity, error) {
	accountIdentity := &model.AccountIdentity{}

	q := r.Database.WithContext(ctx).Where("provider = ?", provider).Where("provider_user_id = ?", providerUserId).Where("deleted_at IS NULL").First(&accountIdentity)
	if q.Error != nil {
		return nil, q.Error
	}

	return accountIdentity, nil
}

func (r *accountIdentityRepository) CreateTx(ctx context.Context, tx *gorm.DB, accountIdentity *model.AccountIdentity) error {
	q := tx.WithContext(ctx).Create(accountIdentity)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
