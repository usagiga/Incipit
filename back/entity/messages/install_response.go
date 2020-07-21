package messages

import (
	"net/http"
)

type InstallResponse struct {
	BaseResponse
}

func NewInstallResponse() (resp *InstallResponse) {
	return &InstallResponse{
		BaseResponse: BaseResponse{Type: "install", Details: nil},
	}
}

func (resp *InstallResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
