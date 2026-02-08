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
