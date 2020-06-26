package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type UpdateAdminResponse struct {
	BaseResponse

	AdminUser AdminUser `json:"admin_user"`
}

func NewUpdateAdminResponse(user *entity.AdminUser) (resp *UpdateAdminResponse) {
	resUser := AdminUser{
		ID:         user.ID,
		Name:       user.Name,
		ScreenName: user.Password,
	}

	return &UpdateAdminResponse{
		BaseResponse: BaseResponse{Type: "update_admin", Details: nil},
		AdminUser:    resUser,
	}
}

func (resp *UpdateAdminResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
