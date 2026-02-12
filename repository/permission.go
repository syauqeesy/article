package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type PermissionRepository interface {
	FindByCode(ctx context.Context, code string) (*model.Permission, error)
}

type permissionRepository repository

func (r *permissionRepository) FindByCode(ctx context.Context, code string) (*model.Permission, error) {
	permission := &model.Permission{}

	q := r.Database.WithContext(ctx).Where("code = ?", code).Where("deleted_at IS NULL").First(&permission)
	if q.Error != nil {
		return nil, q.Error
	}

	return permission, nil
}
