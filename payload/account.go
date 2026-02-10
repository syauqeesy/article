package payload

type OauthResponse struct {
	ConsentPageUrl string
	State          string
}

type OauthCallbackResponse struct {
	RedirectUrl  string
	Token        string
	RefreshToken string
}

type AccountInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
}
