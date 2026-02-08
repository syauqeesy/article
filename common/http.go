package common

import (
	"encoding/json"
	"net/http"
)

type HttpJsonResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HttpErrorHandler(w http.ResponseWriter, exception error, payload any) {
	w.Header().Set("Content-Type", "application/json")

	var message string = http.StatusText(http.StatusInternalServerError)

	switch convertedException := (exception).(type) {
	case *ApplicationException:
		w.WriteHeader(convertedException.HttpStatusCode)

		message = exception.Error()
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	err := json.NewEncoder(w).Encode(&HttpJsonResponse{
		Message: message,
		Data:    payload,
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
		HttpErrorHandler(w, CreateException(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)), nil)
	}
}

func WriteHttpRedirect(w http.ResponseWriter, request *http.Request, url string) {
	http.Redirect(w, request, url, http.StatusFound)
}
