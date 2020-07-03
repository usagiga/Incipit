package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type UpdateLinkResponse struct {
	BaseResponse

	Link Link `json:"link"`
}

func NewUpdateLinkResponse(link *entity.Link) (resp *UpdateLinkResponse) {
	resLink := Link{
		ID:  link.ID,
		URL: link.URL,
	}

	return &UpdateLinkResponse{
		BaseResponse: BaseResponse{Type: "update_link", Details: nil},
		Link:         resLink,
	}
}

func (resp *UpdateLinkResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
