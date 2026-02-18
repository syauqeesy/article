package payload

type ArticleInfo struct {
	Id             string             `json:"id"`
	Account        AccountInfo        `json:"account"`
	Status         string             `json:"status"`
	Views          int32              `json:"views"`
	ArticleContent ArticleContentInfo `json:"article_content"`
	CreatedAt      int64              `json:"created_at"`
}

type CreateArticleContent struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type UpdateArticleContent struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type ChangeArticleStatus struct {
	Status string `json:"status"`
}

type ArticlePaginationResponse struct {
	Articles  []*ArticleInfo `json:"articles"`
	Page      int            `json:"page"`
	TotalPage int            `json:"total_page"`
}
