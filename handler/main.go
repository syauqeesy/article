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
	Account *accountHandler
}

func New(mux *http.ServeMux, configuration *configuration.Configuration, service *service.Service, repository *repository.Repository) *Handler {
	handler := &handler{
		Configuration: configuration,
		Service:       service,
		Repository:    repository,
	}

	h := &Handler{
		Account: (*accountHandler)(handler),
	}

	auth := http.NewServeMux()
	auth.Handle("POST /refresh", middleware.RefreshTokenValidation(configuration)(http.HandlerFunc(h.Account.AuthRefresh)))

	oauth := http.NewServeMux()
	oauth.HandleFunc("GET /{provider}", h.Account.Oauth)
	oauth.HandleFunc("GET /{provider}/callback", h.Account.OauthCallback)

	auth.Handle("/oauth/", http.StripPrefix("/oauth", oauth))

	account := http.NewServeMux()
	account.Handle("GET /{id}", middleware.HasPermission("article.create", repository.AccountPermission)(http.HandlerFunc(h.Account.Detail)))

	mux.Handle("/auth/", http.StripPrefix("/auth", auth))
	mux.Handle("/account/", middleware.Authentication(configuration)(http.StripPrefix("/account", account)))

	return h
}
