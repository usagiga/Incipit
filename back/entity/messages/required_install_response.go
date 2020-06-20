package messages

import (
	"net/http"
)

// RequiredInstallResponse represents one of response which is sent to client
type RequiredInstallResponse struct {
	BaseResponse
}

// GetHTTPStatusCode returns HTTP status code determine from its internal error code
func (resp *RequiredInstallResponse) GetHTTPStatusCode() int {
	return http.StatusServiceUnavailable
}

// NewErrorResponse returns ErrorResponse which is made by error
func NewRequiredInstallResponse() (resp Response) {
	return &RequiredInstallResponse{
		BaseResponse: BaseResponse{Type: "required_install", Details: nil},
	}
}
