package handler

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
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
		MaxAge:   int((10 * time.Minute).Seconds()),
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
		MaxAge:   int((1 * time.Hour).Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((7 * 24 * time.Hour).Seconds()),
	})

	common.WriteHttpRedirect(w, r, result.RedirectUrl)
}

func (h *accountHandler) Detail(w http.ResponseWriter, r *http.Request) {
	result, err := h.Service.Account.Detail(r.Context(), r.PathValue("id"))
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	common.WriteHttpResponse(w, http.StatusOK, "Get Detail Account Success", result)
}

func (h *accountHandler) AuthRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("refresh_token")
	if err != nil {
		common.HttpErrorHandler(w, exception.Unauthorized, nil)
	}

	result, err := h.Service.Account.AuthRefresh(r.Context(), token.Value)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int((1 * time.Hour).Seconds()),
	})

	common.WriteHttpResponse(w, http.StatusOK, "Refresh Access Token Success", nil)
}

func (h *accountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("refresh_token")
	if err != nil {
		common.HttpErrorHandler(w, exception.Unauthorized, nil)
	}

	err = h.Service.Account.Logout(r.Context(), token.Value)
	if err != nil {
		common.HttpErrorHandler(w, err, nil)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})

	common.WriteHttpResponse(w, http.StatusOK, "Logout Success", nil)
}
