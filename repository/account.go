package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindById(ctx context.Context, id string) (*model.Account, error)
	FindByEmail(ctx context.Context, email string) (*model.Account, error)
	CreateTx(ctx context.Context, tx *gorm.DB, account *model.Account) error
}

type accountRepository repository

func (r *accountRepository) FindById(ctx context.Context, id string) (*model.Account, error) {
	account := &model.Account{}

	q := r.Database.WithContext(ctx).Where("id = ?", id).Where("deleted_at IS NULL").First(&account)
	if q.Error != nil {
		return nil, q.Error
	}

	return account, nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*model.Account, error) {
	account := &model.Account{}

	q := r.Database.WithContext(ctx).Where("email = ?", email).Where("deleted_at IS NULL").First(&account)
	if q.Error != nil {
		return nil, q.Error
	}

	return account, nil
}

func (r *accountRepository) CreateTx(ctx context.Context, tx *gorm.DB, account *model.Account) error {
	q := tx.WithContext(ctx).Create(&account)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
