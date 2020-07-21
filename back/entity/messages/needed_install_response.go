package messages

import (
	"net/http"
)

// NeededInstallResponse represents one of response which is sent to client
type NeededInstallResponse struct {
	BaseResponse
}

// GetHTTPStatusCode returns HTTP status code determine from its internal error code
func (resp *NeededInstallResponse) GetHTTPStatusCode() int {
	return http.StatusServiceUnavailable
}

// NewErrorResponse returns ErrorResponse which is made by error
func NewNeededInstallResponse() (resp Response) {
	return &NeededInstallResponse{
		BaseResponse: BaseResponse{Type: "needed_install", Details: nil},
	}
}
