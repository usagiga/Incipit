package messages

import (
	"net/http"
)

type DeleteAdminResponse struct {
	BaseResponse
}

func NewDeleteAdminResponse() (resp *DeleteAdminResponse) {
	return &DeleteAdminResponse{
		BaseResponse: BaseResponse{Type: "delete_admin", Details: nil},
	}
}

func (resp *DeleteAdminResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
