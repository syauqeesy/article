package payload

type ArticleAssetInfo struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
	CreatedAt int64  `json:"created_at"`
}

type SignInfo struct {
	Id        string `json:"id"`
	ObjectKey string `json:"object_key"`
	UploadUrl string `json:"upload_url"`
	PublicUrl string `json:"public_url"`
}

type SignUploadUrlArticleAsset struct {
	ArticleId   string `json:"article_id"`
	ContentType string `json:"content_type"`
}
