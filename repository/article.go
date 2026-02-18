package repository

import (
	"context"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Find(ctx context.Context) ([]*model.Article, error)
	FindById(ctx context.Context, id string) (*model.Article, error)
	CreateTx(ctx context.Context, tx *gorm.DB, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	DeleteTx(ctx context.Context, tx *gorm.DB, article *model.Article) error
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

func (r *articleRepository) CreateTx(ctx context.Context, tx *gorm.DB, article *model.Article) error {
	q := tx.WithContext(ctx).Create(article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleRepository) Update(ctx context.Context, article *model.Article) error {
	err := article.SetUpdatedAt()
	if err != nil {
		return err
	}

	q := r.Database.WithContext(ctx).Save(&article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (r *articleRepository) DeleteTx(ctx context.Context, tx *gorm.DB, article *model.Article) error {
	var err error

	err = article.SetUpdatedAt()
	if err != nil {
		return err
	}

	err = article.SetDeletedAt()
	if err != nil {
		return err
	}

	q := tx.WithContext(ctx).Save(&article)
	if q.Error != nil {
		return q.Error
	}

	return nil
}
