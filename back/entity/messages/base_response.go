package messages

// BaseResponse represents base of responses
type BaseResponse struct {
	Type    string      `json:"type"`
	Details interface{} `json:"details"`
}
