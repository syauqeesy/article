package repository

import (
	"context"
	"math"

	"ahmadsyauqi.dev/article/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Find(ctx context.Context) ([]*model.Article, error)
	FindPaginate(ctx context.Context, page int) ([]*model.Article, int, error)
	FindById(ctx context.Context, id string) (*model.Article, error)
	FindBySlug(ctx context.Context, slug string) (*model.Article, error)
	CreateTx(ctx context.Context, tx *gorm.DB, article *model.Article) error
	Update(ctx context.Context, article *model.Article) error
	DeleteTx(ctx context.Context, tx *gorm.DB, article *model.Article) error
}

type articleRepository repository

func (r *articleRepository) Find(ctx context.Context) ([]*model.Article, error) {
	articles := make([]*model.Article, 0)

	q := r.Database.WithContext(ctx).Preload("Account").Preload("ArticleContent").Where("deleted_at IS NULL").Order("created_at DESC").Find(&articles)
	if q.Error != nil {
		return nil, q.Error
	}

	return articles, nil
}

func (r *articleRepository) FindPaginate(ctx context.Context, page int) ([]*model.Article, int, error) {
	const limit = 10

	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	var total int64

	q := r.Database.WithContext(ctx).Model(&model.Article{}).Count(&total)
	if q.Error != nil {
		return nil, 0, q.Error
	}

	articles := make([]*model.Article, 0)

	q = r.Database.WithContext(ctx).Preload("Account").Preload("ArticleContent").Where("status = ?", model.ArticleStatusPublished).Where("deleted_at IS NULL").Order("created_at DESC").Limit(limit).Offset(offset).Find(&articles)
	if q.Error != nil {
		return nil, 0, q.Error
	}

	return articles, int(math.Ceil(float64(total) / 10.0)), nil
}

func (r *articleRepository) FindById(ctx context.Context, id string) (*model.Article, error) {
	article := &model.Article{}

	q := r.Database.WithContext(ctx).Preload("ArticleContent").Preload("Account").Where("id = ?", id).Where("deleted_at IS NULL").First(&article)
	if q.Error != nil {
		return nil, q.Error
	}

	return article, nil
}

func (r *articleRepository) FindBySlug(ctx context.Context, slug string) (*model.Article, error) {
	article := &model.Article{}

	q := r.Database.WithContext(ctx).Preload("ArticleContent", func(db *gorm.DB) *gorm.DB {
		return db.Where("slug = ?", slug)
	}).Preload("Account").Where("status = ?", model.ArticleStatusPublished).Where("deleted_at IS NULL").First(&article)
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
