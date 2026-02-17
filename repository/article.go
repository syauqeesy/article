package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
)

type ArticleRepository interface {
	Find(ctx context.Context) ([]*model.Article, error)
	FindById(ctx context.Context, id string) (*model.Article, error)
	Create(ctx context.Context, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	Delete(ctx context.Context, article *model.Article) error
}

type articleRepository repository

func (r *articleRepository) Find(ctx context.Context) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)

	q := r.Database.WithContext(ctx).Preload("ArticleContent").Preload("Account").Where("deleted_at IS NULL").Find(&articles)
	if q.Error != nil {
		return nil, q.Error
	}

	return articles, nil
}

func (r *articleRepository) FindById(ctx context.Context, id string) (*model.Article, error) {
	article := &model.Article{}

	q := r.Database.WithContext(ctx).Preload("ArticleContent").Preload("Account").Where("id = ?", id).Where("deleted_at IS NULL").First(&article)
	if q.Error != nil {
		return nil, q.Error
	}

	return article, nil
}

func (r *articleRepository) Create(ctx context.Context, article *model.Article) error {
	q := r.Database.WithContext(ctx).Create(article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleRepository) Update(ctx context.Context, article *model.Article) error {
	article.SetUpdatedAt()

	q := r.Database.WithContext(ctx).Save(&article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleRepository) Delete(ctx context.Context, article *model.Article) error {
	article.SetUpdatedAt()

	article.SetDeletedAt()

	q := r.Database.WithContext(ctx).Save(&article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
