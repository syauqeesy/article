package payload

type ArticleContentInfo struct {
	Id        string `json:"id"`
	Language  string `json:"language"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}
