package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type ArticleContentRepository interface {
	FindById(ctx context.Context, id string) (*model.ArticleContent, error)
	FindBySlug(ctx context.Context, slug string) (*model.ArticleContent, error)
	CreateTx(ctx context.Context, tx *gorm.DB, articleContent *model.ArticleContent) error
	Update(ctx context.Context, articleContent *model.ArticleContent) error
	DeleteTx(ctx context.Context, tx *gorm.DB, articleContent *model.ArticleContent) error
}

type articleContentRepository repository

func (r *articleContentRepository) FindById(ctx context.Context, id string) (*model.ArticleContent, error) {
	articleContent := &model.ArticleContent{}

	q := r.Database.WithContext(ctx).Where("id = ?", id).Where("deleted_at IS NULL").First(&articleContent)
	if q.Error != nil {
		return nil, q.Error
	}

	return articleContent, nil
}

func (r *articleContentRepository) FindBySlug(ctx context.Context, slug string) (*model.ArticleContent, error) {
	articleContent := &model.ArticleContent{}

	q := r.Database.WithContext(ctx).Where("slug = ?", slug).Where("deleted_at IS NULL").First(&articleContent)
	if q.Error != nil {
		return nil, q.Error
	}

	return articleContent, nil
}

func (r *articleContentRepository) CreateTx(ctx context.Context, tx *gorm.DB, articleContent *model.ArticleContent) error {
	q := tx.WithContext(ctx).Create(articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleContentRepository) Update(ctx context.Context, articleContent *model.ArticleContent) error {
	err := articleContent.SetUpdatedAt()
	if err != nil {
		return err
	}

	q := r.Database.WithContext(ctx).Save(&articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleContentRepository) DeleteTx(ctx context.Context, tx *gorm.DB, articleContent *model.ArticleContent) error {
	var err error

	err = articleContent.SetUpdatedAt()
	if err != nil {
		return err
	}

	err = articleContent.SetDeletedAt()
	if err != nil {
		return err
	}

	q := tx.WithContext(ctx).Save(&articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
