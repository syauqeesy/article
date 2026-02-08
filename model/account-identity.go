package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AccountIdentity struct {
	Id             string `gorm:"column:id;type:char(36);primaryKey"`
	AccountId      string `gorm:"column:account_id;type:char(36);not null"`
	Provider       string `gorm:"column:provider;type:varchar(32);not null"`
	ProviderUserId string `gorm:"column:provider_user_id;type:varchar(128);not null"`
	CreatedAt      int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt      *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt      *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Account Account `gorm:"foreignKey:AccountId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AccountIdentity) TableName() string {
	return "account_identities"
}

func NewAccountIdentity(accountId string, provider string, providerUserId string) (*AccountIdentity, error) {
	accountIdentity := &AccountIdentity{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = accountIdentity.SetAccountId(accountId)
	if err != nil {
		return nil, err
	}

	err = accountIdentity.SetProvider(provider)
	if err != nil {
		return nil, err
	}

	err = accountIdentity.SetProviderUserId(providerUserId)
	if err != nil {
		return nil, err
	}

	return accountIdentity, nil
}

func (m *AccountIdentity) SetAccountId(id string) error {
	validate := validator.New()

	err := validate.Var(id, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.AccountId = id

	return nil
}

func (m *AccountIdentity) SetProvider(provider string) error {
	validate := validator.New()

	err := validate.Var(provider, "required,min=1,max=32")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	if provider != "google" {
		return exception.InvalidOauthProvider
	}

	m.Provider = provider

	return nil
}

func (m *AccountIdentity) SetProviderUserId(providerUserId string) error {
	validate := validator.New()

	err := validate.Var(providerUserId, "required,min=1,max=128")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ProviderUserId = providerUserId

	return nil
}

func (m *AccountIdentity) GetId() string {
	return m.Id
}

func (m *AccountIdentity) GetAccountId() string {
	return m.AccountId
}

func (m *AccountIdentity) GetProvider() string {
	return m.Provider
}

func (m *AccountIdentity) GetProviderUserId() string {
	return m.ProviderUserId
}
