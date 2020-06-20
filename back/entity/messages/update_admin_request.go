package messages

type UpdateAdminRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Password   string `json:"password"`
}
