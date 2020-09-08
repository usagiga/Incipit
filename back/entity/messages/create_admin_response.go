package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type CreateAdminResponse struct {
	BaseResponse

	AdminUser AdminUser `json:"admin_user"`
}

func NewCreateAdminResponse(user *entity.AdminUser) (resp *CreateAdminResponse) {
	resUser := AdminUser{
		ID:         user.ID,
		Name:       user.Name,
		ScreenName: user.ScreenName,
	}

	return &CreateAdminResponse{
		BaseResponse: BaseResponse{Type: "create_admin", Details: nil},
		AdminUser:    resUser,
	}
}

func (resp *CreateAdminResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
