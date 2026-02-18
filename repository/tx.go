package repository

import (
	"context"

	"gorm.io/gorm"
)

type TxFunc func(tx *gorm.DB) error

type TxRepository interface {
	WithTx(ctx context.Context, fn TxFunc) error
}

type txRepository repository

func (r *txRepository) WithTx(ctx context.Context, fn TxFunc) error {
	var err error

	tx := r.Database.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback().Error
			panic(r)
		}
		if err != nil {
			_ = tx.Rollback().Error
		}
	}()

	if err = fn(tx); err != nil {
		return err
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
