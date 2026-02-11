package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Permission struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	Code      string `gorm:"column:code;varchar(128);not null"`
	Name      string `gorm:"column:name;varchar(128);not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`
}

func (Permission) TableName() string {
	return "permissions"
}

func NewPermission(code string, name string) (*Permission, error) {
	permission := &Permission{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = permission.SetCode(code)
	if err != nil {
		return nil, err
	}

	err = permission.SetName(name)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (m *Permission) SetCode(code string) error {
	validate := validator.New()

	err := validate.Var(code, "required,min=3,max=128")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Code = code

	return nil
}

func (m *Permission) SetName(name string) error {
	validate := validator.New()

	err := validate.Var(name, "required,min=3,max=128")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Name = name

	return nil
}

func (m *Permission) GetId() string {
	return m.Id
}

func (m *Permission) GetCode() string {
	return m.Code
}

func (m *Permission) GetName() string {
	return m.Name
}

func (m *Permission) GetCreatedAt() int64 {
	return m.CreatedAt
}
