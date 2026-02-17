package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type ArticleContentRepository interface {
	FindById(ctx context.Context, id string) (*model.ArticleContent, error)
	FindBySlug(ctx context.Context, slug string) (*model.ArticleContent, error)
	Create(ctx context.Context, articleContent *model.ArticleContent) error
	Update(ctx context.Context, articleContent *model.ArticleContent) error
	Delete(ctx context.Context, articleContent *model.ArticleContent) error
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

func (r *articleContentRepository) Create(ctx context.Context, articleContent *model.ArticleContent) error {
	q := r.Database.WithContext(ctx).Create(articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleContentRepository) Update(ctx context.Context, articleContent *model.ArticleContent) error {
	articleContent.SetUpdatedAt()

	q := r.Database.WithContext(ctx).Save(&articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleContentRepository) Delete(ctx context.Context, articleContent *model.ArticleContent) error {
	articleContent.SetUpdatedAt()

	articleContent.SetDeletedAt()

	q := r.Database.WithContext(ctx).Save(&articleContent)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
