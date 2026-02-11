package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type AccountPermissionRepository interface {
	FindByPermissionCodeAndAccountId(ctx context.Context, code string, accountId string) (*model.AccountPermission, error)
}

type accountPermissionRepository repository

func (r *accountPermissionRepository) FindByPermissionCodeAndAccountId(ctx context.Context, code string, accountId string) (*model.AccountPermission, error) {
	accountPermission := &model.AccountPermission{}

	q := r.Database.WithContext(ctx).Preload("Permission", "code = ?", code).Where("account_id = ?", accountId).First(&accountPermission)
	if q.Error != nil {
		return nil, q.Error
	}

	return accountPermission, q.Error
}
