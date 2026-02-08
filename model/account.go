package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Account struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	Email     string `gorm:"column:email;varchar(128);not null"`
	Name      string `gorm:"column:name;varchar(256);not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`
}

func (Account) TableName() string {
	return "accounts"
}

func NewAccount(email string, name string) (*Account, error) {
	account := &Account{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = account.SetEmail(email)
	if err != nil {
		return nil, err
	}

	err = account.SetName(name)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (m *Account) SetEmail(email string) error {
	validate := validator.New()

	err := validate.Var(email, "required,email,min=3,max=128")
	if err != nil {
		common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Email = email

	return nil
}

func (m *Account) SetName(name string) error {
	validate := validator.New()

	err := validate.Var(name, "required,min=1,max=256")
	if err != nil {
		common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Name = name

	return nil
}

func (m *Account) GetId() string {
	return m.Id
}

func (m *Account) GetEmail() string {
	return m.Email
}

func (m *Account) GetName() string {
	return m.Name
}
