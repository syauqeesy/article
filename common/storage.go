package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	cloud_storage "cloud.google.com/go/storage"
)

type Storage struct {
	Bucket  string
	BaseUrl string
	Client  *cloud_storage.Client
}

func NewStorage(ctx context.Context, bucket string) (*Storage, error) {
	client, err := cloud_storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	baseUrl := fmt.Sprintf("https://storage.googleapis.com/%s", bucket)

	fmt.Println("Cloud storage client initialized")

	return &Storage{
		Bucket:  bucket,
		BaseUrl: baseUrl,
		Client:  client,
	}, nil
}

func (s *Storage) SignedPutURL(objectKey string, contentType string) (string, string, error) {
	var (
		publicUrl string
		err       error
	)

	objectKey = strings.TrimPrefix(objectKey, "/")

	opts := &cloud_storage.SignedURLOptions{
		Scheme:      cloud_storage.SigningSchemeV4,
		Method:      "PUT",
		Expires:     time.Now().Add(10 * time.Minute),
		ContentType: contentType,
		Headers:     []string{"Cache-Control: public, max-age=31536000, immutable"},
	}

	uploadUrl, err := s.Client.Bucket(s.Bucket).SignedURL(objectKey, opts)
	if err != nil {
		return "", "", fmt.Errorf("bucket.SignedUrl: %w", err)
	}

	publicUrl = fmt.Sprintf("%s/%s", s.BaseUrl, objectKey)

	return uploadUrl, publicUrl, nil
}

func (s *Storage) Close() error {
	return s.Client.Close()
}

func (s *Storage) ObjectAttrs(ctx context.Context, objectKey string) (*cloud_storage.ObjectAttrs, error) {
	return s.Client.Bucket(s.Bucket).Object(objectKey).Attrs(ctx)
}

func (s *Storage) DeleteObject(ctx context.Context, objectKey string) error {
	return s.Client.Bucket(s.Bucket).Object(objectKey).Delete(ctx)
}
