package messages

type LoginAdminAuthRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
