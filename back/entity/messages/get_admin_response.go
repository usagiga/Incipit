package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type GetAdminResponse struct {
	BaseResponse

	AdminUser AdminUser `json:"admin_user"`
}

func NewGetAdminResponse(user *entity.AdminUser) (resp Response) {
	resUser := AdminUser{
		ID:         user.ID,
		Name:       user.Name,
		ScreenName: user.Password,
	}

	return &GetAdminResponse{
		BaseResponse: BaseResponse{
			Type:    "get_admin",
			Details: nil,
		},
		AdminUser: resUser,
	}
}

func (resp *GetAdminResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
