package common

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

func GenerateRandomState(n int) (string, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func SetupOauthConfig(clientId string, clientSecret string, redirectUrl string) *oauth2.Config {
	oauth := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return oauth
}

func ExchangeOauthCode(ctx context.Context, oauth *oauth2.Config, code string, clientId string) (*idtoken.Payload, error) {
	token, err := oauth.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	idToken := token.Extra("id_token").(string)
	if idToken == "" {
		return nil, CreateException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	idTokenPayload, err := idtoken.Validate(ctx, idToken, clientId)
	if err != nil {
		return nil, err
	}

	return idTokenPayload, nil
}

func GetOauthClaims(idTokenPayload *idtoken.Payload) (string, string, string, error) {
	subject := idTokenPayload.Subject

	email := idTokenPayload.Claims["email"].(string)
	if email == "" {
		return "", "", "", CreateException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	name := idTokenPayload.Claims["name"].(string)
	if name == "" {
		return "", "", "", CreateException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	return subject, email, name, nil
}
