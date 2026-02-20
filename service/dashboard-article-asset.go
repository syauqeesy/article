package service

import (
	"context"
	"fmt"

	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/model"
	"ahmadsyauqi.dev/article/payload"
	"github.com/google/uuid"
)

type DashboardArticleAssetService interface {
	Sign(ctx context.Context, request *payload.SignUploadUrlArticleAsset) (*payload.SignInfo, error)
	Complete(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type dashboardArticleAssetService service

func (s *dashboardArticleAssetService) Sign(ctx context.Context, request *payload.SignUploadUrlArticleAsset) (*payload.SignInfo, error) {
	article, err := s.Repository.Article.FindById(ctx, request.ArticleId)
	if err != nil {
		return nil, exception.ArticleNotFound
	}

	if request.ContentType != "image/webp" {
		return nil, exception.InvalidArticleAssetContentType
	}

	objectKey := fmt.Sprintf("articles/%s/images/%s.webp", article.GetId(), uuid.New().String())

	articleAsset, err := model.NewArticleAsset(article.GetId(), objectKey, request.ContentType)
	if err != nil {
		return nil, err
	}

	uploadUrl, publicUrl, err := s.Storage.SignedPutURL(objectKey, request.ContentType)
	if err != nil {
		return nil, err
	}

	err = s.Repository.ArticleAsset.Create(ctx, articleAsset)
	if err != nil {
		return nil, err
	}

	return &payload.SignInfo{
		Id:        articleAsset.GetId(),
		ObjectKey: objectKey,
		UploadUrl: uploadUrl,
		PublicUrl: publicUrl,
	}, nil
}

func (s *dashboardArticleAssetService) Complete(ctx context.Context, id string) error {
	articleAsset, err := s.Repository.ArticleAsset.FindById(ctx, id)
	if err != nil {
		return err
	}

	object, err := s.Storage.ObjectAttrs(ctx, articleAsset.GetObjectKey())
	if err != nil {
		return err
	}

	err = articleAsset.SetContentSize(object.Size)
	if err != nil {
		return err
	}

	err = articleAsset.SetStatus(model.ArticleAssetStatusAttached)
	if err != nil {
		return err
	}

	err = s.Repository.ArticleAsset.Update(ctx, articleAsset)
	if err != nil {
		return err
	}

	return nil
}

func (s *dashboardArticleAssetService) Delete(ctx context.Context, id string) error {
	articleAsset, err := s.Repository.ArticleAsset.FindById(ctx, id)
	if err != nil {
		return err
	}

	err = s.Storage.DeleteObject(ctx, articleAsset.GetObjectKey())
	if err != nil {
		return err
	}

	err = s.Repository.ArticleAsset.Delete(ctx, articleAsset)
	if err != nil {
		return err
	}

	return nil
}
