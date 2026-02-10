package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RefreshToken struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	AccountId string `gorm:"column:account_id;char(36);not null"`
	Token     string `gorm:"column:token;text;not null"`
	ExpiresAt int64  `gorm:"column:expires_at;type:bigint;not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Account Account `gorm:"foreignKey:AccountId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RefreshToken) TableName() string {
	return "refresh_tokens"
}

func NewRefreshToken(accountId string, pepper string, token string, expiresAt int64) (*RefreshToken, error) {
	refreshToken := &RefreshToken{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = refreshToken.SetAccountId(accountId)
	if err != nil {
		return nil, err
	}

	err = refreshToken.SetToken(pepper, token)
	if err != nil {
		return nil, err
	}

	err = refreshToken.SetExpiresAt(expiresAt)
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func HashRefreshToken(pepper string, refreshToken string) (string, error) {
	pepperInBytes, err := base64.StdEncoding.DecodeString(pepper)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, pepperInBytes)
	mac.Write([]byte(refreshToken))
	macSum := mac.Sum(nil)

	return hex.EncodeToString(macSum), nil
}

func (m *RefreshToken) IsExpired() bool {
	return time.UnixMilli(m.GetExpiresAt()).Before(time.Now())
}

func (m *RefreshToken) SetAccountId(accountId string) error {
	validate := validator.New()

	err := validate.Var(accountId, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.AccountId = accountId

	return nil
}

func (m *RefreshToken) SetToken(pepper string, token string) error {
	validate := validator.New()

	err := validate.Var(token, "required")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	hashedToken, err := HashRefreshToken(pepper, token)
	if err != nil {
		return err
	}

	fmt.Println(hashedToken)
	m.Token = hashedToken

	return nil
}

func (m *RefreshToken) SetExpiresAt(expiresAt int64) error {
	validate := validator.New()
	err := validate.Var(expiresAt, "required,numeric")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ExpiresAt = expiresAt

	return nil
}

func (m *RefreshToken) SetUpdatedAt() error {
	updatedAt := time.Now().UTC().UnixMilli()

	m.UpdatedAt = &updatedAt

	return nil
}

func (m *RefreshToken) SetDeletedAt() error {
	deletedAt := time.Now().UTC().UnixMilli()

	m.DeletedAt = &deletedAt

	return nil
}

func (m *RefreshToken) GetId() string {
	return m.Id
}

func (m *RefreshToken) GetAccountId() string {
	return m.AccountId
}

func (m *RefreshToken) GetToken() string {
	return m.Token
}

func (m *RefreshToken) GetExpiresAt() int64 {
	return m.ExpiresAt
}
