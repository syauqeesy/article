package handler

import (
	"encoding/json"
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/payload"
)

type dashboardArticleHandler handler

func (h *dashboardArticleHandler) List(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.DashboardArticle.List(r.Context())
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Get List Article Success", result)
}

func (h *dashboardArticleHandler) Show(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.DashboardArticle.Show(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Show Article Success", result)
}

func (h *dashboardArticleHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request payload.CreateArticleContent

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
	}

	result, err := h.Service.DashboardArticle.Create(r.Context(), &request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Create Article Success", result)
}

func (h *dashboardArticleHandler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request payload.UpdateArticleContent

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
	}

	result, err := h.Service.DashboardArticle.Update(r.Context(), r.PathValue("id"), &request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Update Article Success", result)
}

func (h *dashboardArticleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.Service.DashboardArticle.Delete(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Delete Article Success", nil)
}

func (h *dashboardArticleHandler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request payload.ChangeArticleStatus

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
	}

	result, err := h.Service.DashboardArticle.ChangeStatus(r.Context(), r.PathValue("id"), &request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Change Article Status Success", result)
}
