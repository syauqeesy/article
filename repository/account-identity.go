package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type AccountIdentityRepository interface {
	FindByProviderAndProviderUserId(ctx context.Context, provider string, providerUserId string) (*model.AccountIdentity, error)
	Create(ctx context.Context, accountIdentity *model.AccountIdentity) error
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

func (r *accountIdentityRepository) Create(ctx context.Context, accountIdentity *model.AccountIdentity) error {
	q := r.Database.WithContext(ctx).Create(accountIdentity)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
