package service

import (
	"context"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/model"
	"ahmadsyauqi.dev/article/payload"
	"golang.org/x/oauth2"
)

type AccountService interface {
	Oauth(provider string) (*payload.OauthResponse, error)
	OauthCallback(ctx context.Context, provider string, code string, state string) (*payload.OauthCallbackResponse, error)
	Detail(ctx context.Context, id string) (*payload.AccountInfo, error)
	AuthRefresh(ctx context.Context, token string) (string, error)
	Logout(ctx context.Context, token string) error
}

type accountService service

func (s *accountService) Oauth(provider string) (*payload.OauthResponse, error) {
	if provider != "google" {
		return nil, exception.InvalidOauthProvider
	}

	state, err := common.GenerateRandomState(32)
	if err != nil {
		return nil, err
	}

	oauth := common.SetupOauthConfig(s.Configuration.Oauth.ClientId, s.Configuration.Oauth.ClientSecret, s.Configuration.Oauth.RedirectUrl)

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

	oauth := common.SetupOauthConfig(s.Configuration.Oauth.ClientId, s.Configuration.Oauth.ClientSecret, s.Configuration.Oauth.RedirectUrl)
	idTokenPayload, err := common.ExchangeOauthCode(ctx, oauth, code, s.Configuration.Oauth.ClientId)
	if err != nil {
		return nil, err
	}

	subject, email, name, err := common.GetOauthClaims(idTokenPayload)
	if err != nil {
		return nil, err
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

	accessToken, err := common.IssueAuthenticationToken(accountIdentity.GetAccountId(), s.Configuration.Authentication.Secret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := common.GenerateRandomState(32)
	if err != nil {
		return nil, err
	}

	refreshTokenModel, err := model.NewRefreshToken(accountIdentity.GetAccountId(), s.Configuration.Authentication.RefreshPepper, refreshToken, time.Now().Add(7*24*time.Hour).UTC().UnixMilli())
	if err != nil {
		return nil, err
	}

	err = s.Repository.RefreshToken.Create(ctx, refreshTokenModel)
	if err != nil {
		return nil, err
	}

	return &payload.OauthCallbackResponse{
		RedirectUrl:  s.Configuration.Application.Url,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *accountService) Detail(ctx context.Context, id string) (*payload.AccountInfo, error) {
	account, err := s.Repository.Account.FindById(ctx, id)
	if err != nil {
		return nil, exception.AccountNotFound
	}

	return account.GetInfo(), nil
}

func (s *accountService) AuthRefresh(ctx context.Context, token string) (string, error) {
	hashedRefreshToken, err := model.HashRefreshToken(s.Configuration.Authentication.RefreshPepper, token)
	if err != nil {
		return "", err
	}

	refreshToken, err := s.Repository.RefreshToken.FindByToken(ctx, hashedRefreshToken)
	if err != nil {
		return "", exception.Unauthorized
	}

	if refreshToken.IsExpired() {
		err = s.Repository.RefreshToken.Delete(ctx, refreshToken)
		if err != nil {
			return "", err
		}

		return "", exception.Unauthorized
	}

	accessToken, err := common.IssueAuthenticationToken(refreshToken.GetAccountId(), s.Configuration.Authentication.Secret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *accountService) Logout(ctx context.Context, token string) error {
	hashedRefreshToken, err := model.HashRefreshToken(s.Configuration.Authentication.RefreshPepper, token)
	if err != nil {
		return err
	}

	refreshToken, err := s.Repository.RefreshToken.FindByToken(ctx, hashedRefreshToken)
	if err != nil {
		return exception.Unauthorized
	}

	err = s.Repository.RefreshToken.Delete(ctx, refreshToken)
	if err != nil {
		return err
	}

	return nil
}
