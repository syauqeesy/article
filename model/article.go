package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/payload"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	ArticleStatusDraft     = "draft"
	ArticleStatusPublished = "published"
)

type Article struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	AccountId string `gorm:"column:account_id;char(36);not null"`
	Status    string `gorm:"column:status;varchar(32);not null"`
	Views     int32  `gorm:"column:views;int;not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Account        *Account        `gorm:"foreignKey:AccountId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ArticleContent *ArticleContent `gorm:"foreignKey:ArticleId"`
}

func (Article) TableName() string {
	return "articles"
}

func (m *Article) GetInfo() *payload.ArticleInfo {
	article := &payload.ArticleInfo{
		Id:             m.Id,
		Account:        *m.Account.GetInfo(),
		Status:         m.Status,
		Views:          m.Views,
		ArticleContent: *m.ArticleContent.GetInfo(),
	}

	return article
}

func NewArticle(accountId string) (*Article, error) {
	article := &Article{
		Id:        uuid.New().String(),
		Views:     0,
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = article.SetAccountId(accountId)
	if err != nil {
		return nil, err
	}

	err = article.SetStatus(ArticleStatusDraft)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (m *Article) SetAccountId(id string) error {
	validate := validator.New()

	err := validate.Var(id, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.AccountId = id

	return nil
}

func (m *Article) SetStatus(status string) error {
	validate := validator.New()

	err := validate.Var(status, "required,min=3,max=32")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	if status != ArticleStatusDraft && status != ArticleStatusPublished {
		return exception.InvalidArticleStatus
	}

	m.Status = status

	return nil
}

func (m *Article) SetViews() error {
	m.Views = m.Views + 1

	return nil
}

func (m *Article) SetUpdatedAt() error {
	updatedAt := time.Now().UTC().UnixMilli()

	m.UpdatedAt = &updatedAt

	return nil
}

func (m *Article) SetDeletedAt() error {
	deletedAt := time.Now().UTC().UnixMilli()

	m.DeletedAt = &deletedAt

	return nil
}

func (m *Article) SetAccount(account *Account) error {
	m.Account = account

	return nil
}

func (m *Article) SetArticleContent(articleContent *ArticleContent) error {
	m.ArticleContent = articleContent

	return nil
}

func (m *Article) GetId() string {
	return m.Id
}

func (m *Article) GetAccountId() string {
	return m.AccountId
}

func (m *Article) GetStatus() string {
	return m.Status
}

func (m *Article) GetViews() int32 {
	return m.Views
}

func (m *Article) GetCreatedAt() int64 {
	return m.CreatedAt
}

func (m *Article) GetAccount() *Account {
	return m.Account
}

func (m *Article) GetArticleContent() *ArticleContent {
	return m.ArticleContent
}
