package messages

type AdminUser struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}
