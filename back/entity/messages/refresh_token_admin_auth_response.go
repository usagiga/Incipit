package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type RefreshTokenAdminAuthResponse struct {
	BaseResponse

	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

func NewRefreshTokenAdminAuthResponse(accToken *entity.AccessToken, refToken *entity.RefreshToken) (resp *RefreshTokenAdminAuthResponse) {
	resAccToken := AccessToken{Token: accToken.Token}
	resRefToken := RefreshToken{Token: refToken.Token}

	return &RefreshTokenAdminAuthResponse{
		BaseResponse: BaseResponse{Type: "refresh_token_admin", Details: nil},
		AccessToken:  resAccToken,
		RefreshToken: resRefToken,
	}
}

func (resp *RefreshTokenAdminAuthResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
