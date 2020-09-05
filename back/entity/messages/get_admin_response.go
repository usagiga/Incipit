package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type GetAdminResponse struct {
	BaseResponse

	AdminUsers []AdminUser `json:"admin_users"`
}

func NewGetAdminResponse(users []entity.AdminUser) (resp Response) {
	resUsers := make([]AdminUser, len(users))

	for i, user := range users {
		resUser := AdminUser{
			ID:         user.ID,
			Name:       user.Name,
			ScreenName: user.Password,
		}

		resUsers[i] = resUser
	}

	return &GetAdminResponse{
		BaseResponse: BaseResponse{
			Type:    "get_admin",
			Details: nil,
		},
		AdminUsers: resUsers,
	}
}

func (resp *GetAdminResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
