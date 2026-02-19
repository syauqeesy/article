package service

import (
	"context"

	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/payload"
)

type ArticleService interface {
	List(ctx context.Context, page int) (*payload.ArticlePaginationResponse, error)
	Show(ctx context.Context, slug string) (*payload.ArticleInfo, error)
	View(ctx context.Context, id string) error
}

type articleService service

func (s *articleService) List(ctx context.Context, page int) (*payload.ArticlePaginationResponse, error) {
	articles, total, err := s.Repository.Article.FindPaginate(ctx, page)
	if err != nil {
		return nil, err
	}

	articleInfos := make([]*payload.ArticleInfo, 0)

	for _, article := range articles {
		articleInfos = append(articleInfos, article.GetInfo())
	}

	return &payload.ArticlePaginationResponse{
		Articles:  articleInfos,
		Page:      page,
		TotalPage: total,
	}, nil
}

func (s *articleService) Show(ctx context.Context, slug string) (*payload.ArticleInfo, error) {
	article, err := s.Repository.Article.FindBySlug(ctx, slug)
	if err != nil || article.GetArticleContent() == nil {
		return nil, exception.ArticleNotFound
	}

	return article.GetInfo(), nil
}

func (s *articleService) View(ctx context.Context, id string) error {
	article, err := s.Repository.Article.FindById(ctx, id)
	if err != nil {
		return exception.ArticleNotFound
	}

	err = article.SetViews()
	if err != nil {
		return err
	}

	err = s.Repository.Article.Update(ctx, article)
	if err != nil {
		return err
	}

	return nil
}
