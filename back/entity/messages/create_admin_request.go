package messages

type CreateAdminRequest struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Password   string `json:"password"`
}
