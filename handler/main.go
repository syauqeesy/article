package handler

import (
	"net/http"

	"ahmadsyauqi.dev/article/configuration"
	"ahmadsyauqi.dev/article/service"
)

type handler struct {
	Service       *service.Service
	Configuration *configuration.Configuration
}

type Handler struct {
	Account *accountHandler
}

func New(mux *http.ServeMux, configuration *configuration.Configuration, service *service.Service) *Handler {
	handler := &handler{
		Configuration: configuration,
		Service:       service,
	}

	h := &Handler{
		Account: (*accountHandler)(handler),
	}

	return h
}

