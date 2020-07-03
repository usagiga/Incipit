package messages

import (
	"net/http"
)

type DeleteLinkResponse struct {
	BaseResponse
}

func NewDeleteLinkResponse() (resp *DeleteLinkResponse) {
	return &DeleteLinkResponse{
		BaseResponse: BaseResponse{Type: "delete_link", Details: nil},
	}
}

func (resp *DeleteLinkResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
