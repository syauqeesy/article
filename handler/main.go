package handler

import (
	"net/http"

	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/middleware"
	"ahmadsyauqi.dev/article/repository"
	"ahmadsyauqi.dev/article/service"
)

type handler struct {
	Service       *service.Service
	Repository    *repository.Repository
	Configuration *configuration.Configuration
}

type Handler struct {
	Account          *accountHandler
	DashboardArticle *dashboardArticleHandler
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
	}

	auth := http.NewServeMux()
	auth.HandleFunc("POST /refresh", h.Account.AuthRefresh)
	auth.Handle("POST /logout", middleware.Authentication(configuration)(http.HandlerFunc(h.Account.Logout)))

	oauth := http.NewServeMux()
	oauth.HandleFunc("GET /{provider}", h.Account.Oauth)
	oauth.HandleFunc("GET /{provider}/callback", h.Account.OauthCallback)

	auth.Handle("/oauth/", http.StripPrefix("/oauth", oauth))

	account := http.NewServeMux()
	account.HandleFunc("GET /{id}", h.Account.Detail)

	articleDashboard := http.NewServeMux()
	articleDashboard.Handle("GET /", middleware.HasPermission("article.list", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.List)))
	articleDashboard.Handle("GET /{id}", middleware.HasPermission("article.show", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.Show)))
	articleDashboard.Handle("POST /", middleware.HasPermission("article.create", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.Create)))
	articleDashboard.Handle("PUT /{id}", middleware.HasPermission("article.update", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.Update)))
	articleDashboard.Handle("DELETE /{id}", middleware.HasPermission("article.delete", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.Delete)))
	articleDashboard.Handle("POST /status/{id}", middleware.HasPermission("article.status", repository.Permission, repository.AccountPermission)(http.HandlerFunc(h.DashboardArticle.ChangeStatus)))

	mux.Handle("/auth/", http.StripPrefix("/auth", auth))
	mux.Handle("/account/", middleware.Authentication(configuration)(http.StripPrefix("/account", account)))
	mux.Handle("/dashboard/article/", middleware.Authentication(configuration)(http.StripPrefix("/dashboard/article", articleDashboard)))

	return h
}
