package handler

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/middleware"
	"ahmadsyauqi.dev/article/repository"
	"ahmadsyauqi.dev/article/service"
	"golang.org/x/time/rate"
)

type handler struct {
	Service       *service.Service
	Repository    *repository.Repository
	Configuration *configuration.Configuration
}

type Handler struct {
	Account          *accountHandler
	DashboardArticle *dashboardArticleHandler
	Article          *articleHandler
}

func New(mux *http.ServeMux, configuration *configuration.Configuration, service *service.Service, repository *repository.Repository) *Handler {
	handler := &handler{
		Configuration: configuration,
		Service:       service,
		Repository:    repository,
	}

	h := &Handler{
		Account:          (*accountHandler)(handler),
		DashboardArticle: (*dashboardArticleHandler)(handler),
		Article:          (*articleHandler)(handler),
	}

	authenticationMiddleware := middleware.Authentication(configuration)
	hasPermissionMiddleware := func(permission string, handlerFunction http.HandlerFunc) http.Handler {
		hasPermissionMiddleware := middleware.HasPermission(permission, repository.Permission, repository.AccountPermission)

		return authenticationMiddleware(hasPermissionMiddleware(handlerFunction))
	}
	viewArticleRateLimit := common.NewIPRateLimiter(rate.Every(2*time.Hour), 1, 2*time.Hour, true)

	mux.HandleFunc("POST /auth/refresh", h.Account.AuthRefresh)
	mux.Handle("POST /auth/logout", authenticationMiddleware(http.HandlerFunc(h.Account.Logout)))

	mux.HandleFunc("GET /auth/oauth/{provider}", h.Account.Oauth)
	mux.HandleFunc("GET /auth/oauth/{provider}/callback", h.Account.OauthCallback)

	mux.Handle("GET /account/{id}", authenticationMiddleware(http.HandlerFunc(h.Account.Detail)))

	mux.Handle("GET /dashboard/article", hasPermissionMiddleware("article.list", http.HandlerFunc(h.DashboardArticle.List)))
	mux.Handle("GET /dashboard/article/{id}", hasPermissionMiddleware("article.show", http.HandlerFunc(h.DashboardArticle.Show)))
	mux.Handle("POST /dashboard/article", hasPermissionMiddleware("article.create", http.HandlerFunc(h.DashboardArticle.Create)))
	mux.Handle("PUT /dashboard/article/{id}", hasPermissionMiddleware("article.update", http.HandlerFunc(h.DashboardArticle.Update)))
	mux.Handle("DELETE /dashboard/article/{id}", hasPermissionMiddleware("article.delete", http.HandlerFunc(h.DashboardArticle.Delete)))
	mux.Handle("POST /dashboard/article/status/{id}", hasPermissionMiddleware("article.status", http.HandlerFunc(h.DashboardArticle.ChangeStatus)))

	mux.HandleFunc("GET /article", h.Article.List)
	mux.HandleFunc("GET /article/{slug}", h.Article.Show)
	mux.Handle("POST /article/view/{id}", viewArticleRateLimit.Middleware(http.HandlerFunc(h.Article.View)))

	return h
}
