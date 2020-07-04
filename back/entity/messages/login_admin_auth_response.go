package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type LoginAdminAuthResponse struct {
	BaseResponse

	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

func NewLoginAdminAuthResponse(accToken *entity.AccessToken, refToken *entity.RefreshToken) (resp *LoginAdminAuthResponse) {
	resAccToken := AccessToken{Token: accToken.Token}
	resRefToken := RefreshToken{Token: refToken.Token}

	return &LoginAdminAuthResponse{
		BaseResponse: BaseResponse{Type: "login_admin", Details: nil},
		AccessToken:  resAccToken,
		RefreshToken: resRefToken,
	}
}

func (resp *LoginAdminAuthResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
