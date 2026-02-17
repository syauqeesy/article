package service

import (
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/repository"
)

type service struct {
	Configuration *configuration.Configuration
	Repository    *repository.Repository
}

type Service struct {
	Account          AccountService
	DashboardArticle DashboardArticleService
}

func New(configuration *configuration.Configuration, repository *repository.Repository) *Service {
	svc := &service{
		Configuration: configuration,
		Repository:    repository,
	}

	return &Service{
		Account:          (*accountService)(svc),
		DashboardArticle: (*dashboardArticleService)(svc),
	}
}
