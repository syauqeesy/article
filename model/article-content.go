package model

import (
	"net/http"
	"time"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/payload"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	ArticleContentLanguageEn = "EN"
)

type ArticleContent struct {
	Id        string `gorm:"column:id;type:char(36);primaryKey"`
	ArticleId string `gorm:"column:article_id;char(36);not null"`
	Language  string `gorm:"column:language;char(2);not null"`
	Title     string `gorm:"column:title;varchar(256);not null"`
	Slug      string `gorm:"column:slug;varchar(256);not null"`
	Summary   string `gorm:"column:summary;varchar(4096);not null"`
	Content   string `gorm:"column:content;text;not null"`
	CreatedAt int64  `gorm:"column:created_at;type:bigint;not null"`
	UpdatedAt *int64 `gorm:"column:updated_at;type:bigint;default:null"`
	DeletedAt *int64 `gorm:"column:deleted_at;type:bigint;default:null"`

	Article *Article `gorm:"foreignKey:ArticleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ArticleContent) TableName() string {
	return "article_contents"
}

func (m *ArticleContent) GetInfo() *payload.ArticleContentInfo {
	articleContentInfo := &payload.ArticleContentInfo{
		Id:        m.Id,
		Language:  m.Language,
		Title:     m.Title,
		Slug:      m.Slug,
		Summary:   m.Summary,
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}

	return articleContentInfo
}

func NewArticleContent(articleId string, title string, slug string, summary string, content string) (*ArticleContent, error) {
	articleContent := &ArticleContent{
		Id:        uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixMilli(),
	}

	var err error

	err = articleContent.SetArticleId(articleId)
	if err != nil {
		return nil, err
	}

	err = articleContent.SetLanguage()
	if err != nil {
		return nil, err
	}

	err = articleContent.SetTitle(title)
	if err != nil {
		return nil, err
	}

	err = articleContent.SetSlug(slug)
	if err != nil {
		return nil, err
	}

	err = articleContent.SetSummary(summary)
	if err != nil {
		return nil, err
	}

	err = articleContent.SetContent(content)
	if err != nil {
		return nil, err
	}

	return articleContent, nil
}

func (m *ArticleContent) SetArticleId(id string) error {
	validate := validator.New()

	err := validate.Var(id, "required,uuid")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.ArticleId = id

	return nil
}

func (m *ArticleContent) SetLanguage() error {
	m.Language = ArticleContentLanguageEn

	return nil
}

func (m *ArticleContent) SetTitle(title string) error {
	validate := validator.New()

	err := validate.Var(title, "required,min=3,max=256")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Title = title

	return nil
}

func (m *ArticleContent) SetSlug(slug string) error {
	validate := validator.New()

	err := validate.Var(slug, "required,min=3,max=256")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Slug = slug

	return nil
}

func (m *ArticleContent) SetSummary(summary string) error {
	validate := validator.New()

	err := validate.Var(summary, "required,min=3,max=4096")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Summary = summary

	return nil
}

func (m *ArticleContent) SetContent(content string) error {
	validate := validator.New()

	err := validate.Var(content, "required,min=3")
	if err != nil {
		return common.CreateException(http.StatusBadRequest, err.Error())
	}

	m.Content = content

	return nil
}

func (m *ArticleContent) SetUpdatedAt() error {
	updatedAt := time.Now().UTC().UnixMilli()

	m.UpdatedAt = &updatedAt

	return nil
}

func (m *ArticleContent) SetDeletedAt() error {
	deletedAt := time.Now().UTC().UnixMilli()

	m.DeletedAt = &deletedAt

	return nil
}

func (m *ArticleContent) GetId() string {
	return m.Id
}

func (m *ArticleContent) GetArticleId() string {
	return m.ArticleId
}

func (m *ArticleContent) GetLanguage() string {
	return m.Language
}

func (m *ArticleContent) GetTitle() string {
	return m.Title
}

func (m *ArticleContent) GetSlug() string {
	return m.Slug
}

func (m *ArticleContent) GetSummary() string {
	return m.Summary
}

func (m *ArticleContent) GetContent() string {
	return m.Content
}

func (m *ArticleContent) GetCreatedAt() int64 {
	return m.CreatedAt
}
