package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/model"
	"ahmadsyauqi.dev/article/payload"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

type AccountService interface {
	Oauth(provider string) (*payload.OauthResponse, error)
	OauthCallback(ctx context.Context, provider string, code string, state string) (*payload.OauthCallbackResponse, error)
}

type accountService service

func generateRandomState(n int) (string, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func issueAuthenticationToken(subject string, ttl time.Duration, secret string) (*string, *int64, error) {
	now := time.Now().UTC()

	expiresAt := now.Add(ttl).UTC().UnixMilli()

	claims := jwt.MapClaims{
		"sub":  subject,
		"iat":  now.UTC(),
		"exp":  now.Add(ttl).UTC(),
		"type": "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedAccessToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, nil, err
	}

	return &signedAccessToken, &expiresAt, nil
}

func (s *accountService) Oauth(provider string) (*payload.OauthResponse, error) {
	if provider != "google" {
		return nil, exception.InvalidOauthProvider
	}

	state, err := generateRandomState(32)
	if err != nil {
		return nil, err
	}

	oauth := &oauth2.Config{
		ClientID:     s.Configuration.Oauth.ClientId,
		ClientSecret: s.Configuration.Oauth.ClientSecret,
		RedirectURL:  s.Configuration.Oauth.RedirectUrl,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	url := oauth.AuthCodeURL(state, oauth2.AccessTypeOnline)

	return &payload.OauthResponse{
		ConsentPageUrl: url,
		State:          state,
	}, nil
}

func (s *accountService) OauthCallback(ctx context.Context, provider string, code string, state string) (*payload.OauthCallbackResponse, error) {
	if provider != "google" {
		return nil, exception.InvalidOauthProvider
	}

	oauth := &oauth2.Config{
		ClientID:     s.Configuration.Oauth.ClientId,
		ClientSecret: s.Configuration.Oauth.ClientSecret,
		RedirectURL:  s.Configuration.Oauth.RedirectUrl,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	token, err := oauth.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	idToken := token.Extra("id_token").(string)
	if idToken == "" {
		return nil, exception.Unauthorized
	}

	idPayload, err := idtoken.Validate(ctx, idToken, s.Configuration.Oauth.ClientId)
	if err != nil {
		return nil, err
	}

	subject := idPayload.Subject

	email := idPayload.Claims["email"].(string)
	if email == "" {
		return nil, exception.Unauthorized
	}

	name := idPayload.Claims["name"].(string)
	if name == "" {
		return nil, exception.Unauthorized
	}

	accountIdentity, _ := s.Repository.AccountIdentity.FindByProviderAndProviderUserId(ctx, provider, subject)
	if accountIdentity == nil {
		account, err := model.NewAccount(email, name)
		if err != nil {
			return nil, err
		}

		err = s.Repository.Account.Create(ctx, account)
		if err != nil {
			return nil, err
		}

		accountIdentity, err = model.NewAccountIdentity(account.GetId(), provider, subject)
		if err != nil {
			return nil, err
		}

		err = s.Repository.AccountIdentity.Create(ctx, accountIdentity)
		if err != nil {
			return nil, err
		}
	}

	accessToken, _, err := issueAuthenticationToken(accountIdentity.GetAccountId(), 1*time.Hour, s.Configuration.Authentication.Secret)
	if err != nil {
		return nil, err
	}

	refreshToken, expiresAt, err := issueAuthenticationToken(accountIdentity.GetAccountId(), 7*24*time.Hour, s.Configuration.Authentication.RefreshSecret)
	if err != nil {
		return nil, err
	}

	refreshTokenModel, err := model.NewRefreshToken(accountIdentity.GetAccountId(), *refreshToken, *expiresAt)
	if err != nil {
		return nil, err
	}

	err = s.Repository.RefreshToken.Create(ctx, refreshTokenModel)
	if err != nil {
		return nil, err
	}

	return &payload.OauthCallbackResponse{
		RedirectUrl:  s.Configuration.Application.Url,
		Token:        *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}
