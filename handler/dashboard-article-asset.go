package handler

import (
	"encoding/json"
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/payload"
)

type dashboardArticleAssetHandler handler

func (h *dashboardArticleAssetHandler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request payload.SignUploadUrlArticleAsset

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
	}

	result, err := h.Service.DashbboardArticleAsset.Sign(r.Context(), &request)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Sign Upload Url Success", result)
}

func (h *dashboardArticleAssetHandler) Complete(w http.ResponseWriter, r *http.Request) {
	err := h.Service.DashbboardArticleAsset.Complete(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Complete Upload Success", nil)
}

func (h *dashboardArticleAssetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.Service.DashbboardArticleAsset.Delete(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Delete Asset Success", nil)
}
