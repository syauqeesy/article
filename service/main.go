package service

import (
	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
	Storage       *common.Storage
}

type Service struct {
	Account                AccountService
	DashboardArticle       DashboardArticleService
	DashbboardArticleAsset DashboardArticleAssetService
	Article                ArticleService
}

func New(configuration *configuration.Configuration, repository *repository.Repository, storage *common.Storage) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
		Storage:       storage,
	}

	return &Service{
		Account:                (*accountService)(svc),
		DashboardArticle:       (*dashboardArticleService)(svc),
		DashbboardArticleAsset: (*dashboardArticleAssetService)(svc),
		Article:                (*articleService)(svc),
	}
}
