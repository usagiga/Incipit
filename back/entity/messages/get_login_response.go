package messages

import "net/http"

type GetLoginResponse struct {
	BaseResponse
}

func NewGetLoginResponse() (resp *GetLoginResponse) {
	return &GetLoginResponse{
		BaseResponse: BaseResponse{Type: "is_login", Details: nil},
	}
}

func (resp *GetLoginResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
