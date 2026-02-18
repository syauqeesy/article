package service

import (
	"context"

	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/middleware"
	"ahmadsyauqi.dev/article/model"
	"ahmadsyauqi.dev/article/payload"
	"gorm.io/gorm"
)

type DashboardArticleService interface {
	List(ctx context.Context) ([]*payload.ArticleInfo, error)
	Show(ctx context.Context, id string) (*payload.ArticleInfo, error)
	Create(ctx context.Context, request *payload.CreateArticleContent) (*payload.ArticleInfo, error)
	Update(ctx context.Context, id string, request *payload.UpdateArticleContent) (*payload.ArticleInfo, error)
	Delete(ctx context.Context, id string) error
	ChangeStatus(ctx context.Context, id string, request *payload.ChangeArticleStatus) (*payload.ArticleInfo, error)
}

type dashboardArticleService service

func (s *dashboardArticleService) List(ctx context.Context) ([]*payload.ArticleInfo, error) {
	articleInfos := make([]*payload.ArticleInfo, 0)

	articles, err := s.Repository.Article.Find(ctx)
	if err != nil {
		return nil, err
	}

	for _, article := range articles {
		articleInfos = append(articleInfos, article.GetInfo())
	}

	return articleInfos, nil
}

func (s *dashboardArticleService) Show(ctx context.Context, id string) (*payload.ArticleInfo, error) {
	article, err := s.Repository.Article.FindById(ctx, id)
	if err != nil {
		return nil, exception.ArticleNotFound
	}

	return article.GetInfo(), nil
}

func (s *dashboardArticleService) Create(ctx context.Context, request *payload.CreateArticleContent) (*payload.ArticleInfo, error) {
	accountId, ok := middleware.GetSubject(ctx)
	if !ok {
		return nil, exception.Unauthorized
	}

	existingSlug, err := s.Repository.ArticleContent.FindBySlug(ctx, request.Slug)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingSlug != nil {
		return nil, exception.SlugAlreadyExists
	}

	article, err := model.NewArticle(accountId)
	if err != nil {
		return nil, err
	}

	articleContent, err := model.NewArticleContent(article.GetId(), request.Title, request.Slug, request.Summary, request.Content)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Tx.WithTx(ctx, func(tx *gorm.DB) error {
		var err error

		err = s.Repository.Article.CreateTx(ctx, tx, article)
		if err != nil {
			return err
		}

		err = s.Repository.ArticleContent.CreateTx(ctx, tx, articleContent)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	account, err := s.Repository.Account.FindById(ctx, accountId)
	if err != nil {
		return nil, err
	}

	err = article.SetAccount(account)
	if err != nil {
		return nil, err
	}

	err = article.SetArticleContent(articleContent)
	if err != nil {
		return nil, err
	}

	return article.GetInfo(), err
}

func (s *dashboardArticleService) Update(ctx context.Context, id string, request *payload.UpdateArticleContent) (*payload.ArticleInfo, error) {
	article, err := s.Repository.Article.FindById(ctx, id)
	if err != nil {
		return nil, exception.ArticleNotFound
	}

	existingSlug, err := s.Repository.ArticleContent.FindBySlug(ctx, request.Slug)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if article.ArticleContent.GetSlug() != request.Slug && existingSlug != nil {
		return nil, exception.SlugAlreadyExists
	}

	err = article.ArticleContent.SetTitle(request.Title)
	if err != nil {
		return nil, err
	}

	err = article.ArticleContent.SetSlug(request.Slug)
	if err != nil {
		return nil, err
	}

	err = article.ArticleContent.SetSummary(request.Summary)
	if err != nil {
		return nil, err
	}

	err = article.ArticleContent.SetContent(request.Content)
	if err != nil {
		return nil, err
	}

	err = s.Repository.ArticleContent.Update(ctx, article.ArticleContent)
	if err != nil {
		return nil, err
	}

	return article.GetInfo(), nil
}

func (s *dashboardArticleService) Delete(ctx context.Context, id string) error {
	article, err := s.Repository.Article.FindById(ctx, id)
	if err != nil {
		return exception.ArticleNotFound
	}

	err = s.Repository.Tx.WithTx(ctx, func(tx *gorm.DB) error {
		var err error

		err = s.Repository.ArticleContent.DeleteTx(ctx, tx, article.ArticleContent)
		if err != nil {
			return err
		}

		err = s.Repository.Article.DeleteTx(ctx, tx, article)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *dashboardArticleService) ChangeStatus(ctx context.Context, id string, request *payload.ChangeArticleStatus) (*payload.ArticleInfo, error) {
	article, err := s.Repository.Article.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = article.SetStatus(request.Status)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Article.Update(ctx, article)
	if err != nil {
		return nil, err
	}

	return article.GetInfo(), nil
}
