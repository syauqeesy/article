package handler

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
)

type accountHandler handler

func (h *accountHandler) Oauth(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Account.Oauth(r.PathValue("provider"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    result.State,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   10 * 60,
	})

	common.WriteHttpRedirect(w, r, result.ConsentPageUrl)
}

func (h *accountHandler) OauthCallback(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Account.OauthCallback(r.Context(), r.PathValue("provider"), r.URL.Query().Get("code"), r.URL.Query().Get("state"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((15 * time.Minute).Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
	})

	common.WriteHttpRedirect(w, r, result.RedirectUrl)
}
