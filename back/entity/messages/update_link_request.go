package messages

type UpdateLinkRequest struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}
