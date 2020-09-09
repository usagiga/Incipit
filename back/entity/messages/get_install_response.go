package messages

import "net/http"

type GetInstallResponse struct {
	BaseResponse
}

func NewGetInstallResponse() (resp *GetInstallResponse) {
	return &GetInstallResponse{
		BaseResponse: BaseResponse{Type: "is_installed", Details: nil},
	}
}

func (resp *GetInstallResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
