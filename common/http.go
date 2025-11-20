package common

import (
	"encoding/json"
	"net/http"
)

type HttpJsonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HttpErrorHandler(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	err := json.NewEncoder(w).Encode(&HttpJsonResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Data:    nil,
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func WriteHttpResponse(w http.ResponseWriter, status int, message string, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&HttpJsonResponse{
		Message: message,
		Data:    payload,
	})
	if err != nil {
		HttpErrorHandler(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
