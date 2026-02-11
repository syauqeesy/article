package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AccountPermission struct {
	Id           string `gorm:"column:id;type:char(36);primaryKey"`
	AccountId    string `gorm:"column:account_id;char(36);not null"`
	PermissionId string `gorm:"column:permission_id;char(36);not null"`
	CreatedAt    int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt    *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt    *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Account    *Account    `gorm:"foreignKey:AccountId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission *Permission `gorm:"foreignKey:PermissionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AccountPermission) TableName() string {
	return "account_permissions"
}

func NewAccountPermission(accountId string, permissionId string) (*AccountPermission, error) {
	accountPermission := &AccountPermission{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = accountPermission.SetAccountId(accountId)
	if err != nil {
		return nil, err
	}

	err = accountPermission.SetPermissionId(permissionId)
	if err != nil {
		return nil, err
	}

	return accountPermission, nil
}

func (m *AccountPermission) SetAccountId(accountId string) error {
	validate := validator.New()

	err := validate.Var(accountId, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.AccountId = accountId

	return nil
}

func (m *AccountPermission) SetPermissionId(permissionId string) error {
	validate := validator.New()

	err := validate.Var(permissionId, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.PermissionId = permissionId

	return nil
}

func (m *AccountPermission) GetId() string {
	return m.Id
}

func (m *AccountPermission) GetAccountId() string {
	return m.AccountId
}

func (m *AccountPermission) GetPermissionId() string {
	return m.PermissionId
}

func (m *AccountPermission) GetCreatedAt() int64 {
	return m.CreatedAt
}
