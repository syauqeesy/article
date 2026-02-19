package handler

import (
	"net/http"
	"strconv"

	"ahmadsyauqi.dev/article/common"
)

type articleHandler handler

func (h *articleHandler) List(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		common.HttpErrorHandler(w, common.CreateException(http.StatusBadRequest, "Invalid Page"), nil)
		return
	}

	result, err := h.Service.Article.List(r.Context(), page)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Get List Article Success", result)
}

func (h *articleHandler) Show(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Article.Show(r.Context(), r.PathValue("slug"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Show Article Success", result)
}

func (h *articleHandler) View(w http.ResponseWriter, r *http.Request) {
	err := h.Service.Article.View(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "View Article Success", nil)
}
