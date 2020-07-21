package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type InstallResponse struct {
	BaseResponse

	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

func NewInstallResponse(accToken *entity.AccessToken, refToken *entity.RefreshToken) (resp *InstallResponse) {
	resAccToken := AccessToken{Token: accToken.Token}
	resRefToken := RefreshToken{Token: refToken.Token}

	return &InstallResponse{
		BaseResponse: BaseResponse{Type: "install", Details: nil},
		AccessToken: resAccToken,
		RefreshToken: resRefToken,
	}
}

func (resp *InstallResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
