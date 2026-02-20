package model

import (
	"fmt"
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/payload"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	ArticleAssetStatusPending  = "pending"
	ArticleAssetStatusAttached = "attached"
)

type ArticleAsset struct {
	Id          string `gorm:"column:id;type:char(36);primaryKey"`
	ArticleId   string `gorm:"column:article_id;char(36);not null"`
	ObjectKey   string `gorm:"column:object_key;text;not null"`
	ContentType string `gorm:"column:content_type;varchar(128);not null"`
	ContentSize int64  `gorm:"column:content_size;bigint;not null"`
	Status      string `gorm:"column:status;varchar(32);not null"`
	CreatedAt   int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt   *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt   *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Article *Article `gorm:"foreignKey:ArticleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ArticleAsset) TableName() string {
	return "article_assets"
}

func (m *ArticleAsset) GetInfo(bucket string) *payload.ArticleAssetInfo {
	articleAssetInfo := &payload.ArticleAssetInfo{
		Id:        m.Id,
		Url:       fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, m.ObjectKey),
		Type:      m.ContentType,
		Size:      m.ContentSize,
		CreatedAt: m.CreatedAt,
	}

	return articleAssetInfo
}

func NewArticleAsset(articleId string, objectKey string, contentType string) (*ArticleAsset, error) {
	articleAsset := &ArticleAsset{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = articleAsset.SetArticleId(articleId)
	if err != nil {
		return nil, err
	}

	err = articleAsset.SetObjectKey(objectKey)
	if err != nil {
		return nil, err
	}

	err = articleAsset.SetContentType(contentType)
	if err != nil {
		return nil, err
	}

	err = articleAsset.SetContentSize(0)
	if err != nil {
		return nil, err
	}

	err = articleAsset.SetStatus(ArticleAssetStatusPending)
	if err != nil {
		return nil, err
	}

	return articleAsset, nil
}

func (m *ArticleAsset) SetArticleId(id string) error {
	validate := validator.New()

	err := validate.Var(id, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ArticleId = id

	return nil
}

func (m *ArticleAsset) SetObjectKey(objectKey string) error {
	validate := validator.New()

	err := validate.Var(objectKey, "required,min=3")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ObjectKey = objectKey

	return nil
}

func (m *ArticleAsset) SetContentType(contentType string) error {
	validate := validator.New()

	err := validate.Var(contentType, "required,min=3,max=128")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ContentType = contentType

	return nil
}

func (m *ArticleAsset) SetContentSize(contentSize int64) error {
	m.ContentSize = contentSize

	return nil
}

func (m *ArticleAsset) SetStatus(status string) error {
	validate := validator.New()

	err := validate.Var(status, "required,min=3,max=32")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	if status != ArticleAssetStatusPending && status != ArticleAssetStatusAttached {
		return exception.InvalidArticleAssetStatus
	}

	m.Status = status

	return nil
}

func (m *ArticleAsset) SetUpdatedAt() error {
	updatedAt := time.Now().UTC().UnixMilli()

	m.UpdatedAt = &updatedAt

	return nil
}

func (m *ArticleAsset) SetDeletedAt() error {
	deletedAt := time.Now().UTC().UnixMilli()

	m.DeletedAt = &deletedAt

	return nil
}

func (m *ArticleAsset) GetId() string {
	return m.Id
}

func (m *ArticleAsset) GetArticleId() string {
	return m.ArticleId
}

func (m *ArticleAsset) GetObjectKey() string {
	return m.ObjectKey
}

func (m *ArticleAsset) GetContentType() string {
	return m.ContentType
}

func (m *ArticleAsset) GetContentSize() int64 {
	return m.ContentSize
}

func (m *ArticleAsset) GetStatus() string {
	return m.Status
}

func (m *ArticleAsset) GetCreatedAt() int64 {
	return m.CreatedAt
}
