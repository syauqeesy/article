package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type AccountPermissionRepository interface {
	FindByAccountIdAndPermissionId(ctx context.Context, accountId string, permissionId string) (*model.AccountPermission, error)
}

type accountPermissionRepository repository

func (r *accountPermissionRepository) FindByAccountIdAndPermissionId(ctx context.Context, accountId string, permissionId string) (*model.AccountPermission, error) {
	accountPermission := &model.AccountPermission{}

	q := r.Database.WithContext(ctx).Where("account_id = ?", accountId).Where("permission_id", permissionId).First(&accountPermission)
	if q.Error != nil {
		return nil, q.Error
	}

	return accountPermission, nil
}
