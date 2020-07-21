package messages

import (
	"net/http"
)

// RedundantInstallResponse represents one of response which is sent to client
type RedundantInstallResponse struct {
	BaseResponse
}

// GetHTTPStatusCode returns HTTP status code determine from its internal error code
func (resp *RedundantInstallResponse) GetHTTPStatusCode() int {
	return http.StatusGone
}

// NewErrorResponse returns ErrorResponse which is made by error
func NewRedundantInstallResponse() (resp Response) {
	return &RedundantInstallResponse{
		BaseResponse: BaseResponse{Type: "redundant_install", Details: nil},
	}
}
