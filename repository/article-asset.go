package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type ArticleAssetRepository interface {
	FindById(ctx context.Context, id string) (*model.ArticleAsset, error)
	Create(ctx context.Context, articleAsset *model.ArticleAsset) error
	Update(ctx context.Context, articleAsset *model.ArticleAsset) error
	Delete(ctx context.Context, articleAsset *model.ArticleAsset) error
}

type articleAssetRepository repository

func (r *articleAssetRepository) FindById(ctx context.Context, id string) (*model.ArticleAsset, error) {
	articleAsset := &model.ArticleAsset{}

	q := r.Database.WithContext(ctx).Where("id = ?", id).Where("deleted_at IS NULL").First(&articleAsset)
	if q.Error != nil {
		return nil, q.Error
	}

	return articleAsset, nil
}

func (r *articleAssetRepository) Create(ctx context.Context, articleAsset *model.ArticleAsset) error {
	q := r.Database.WithContext(ctx).Create(articleAsset)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleAssetRepository) Update(ctx context.Context, articleAsset *model.ArticleAsset) error {
	err := articleAsset.SetUpdatedAt()
	if err != nil {
		return err
	}

	q := r.Database.WithContext(ctx).Save(&articleAsset)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleAssetRepository) Delete(ctx context.Context, articleAsset *model.ArticleAsset) error {
	var err error

	err = articleAsset.SetUpdatedAt()
	if err != nil {
		return err
	}

	err = articleAsset.SetDeletedAt()
	if err != nil {
		return err
	}

	q := r.Database.WithContext(ctx).Save(&articleAsset)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
