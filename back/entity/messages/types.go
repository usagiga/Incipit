package messages

type AdminUser struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Link struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}

type AccessToken struct {
	Token string
}

type RefreshToken struct {
	Token string
}
