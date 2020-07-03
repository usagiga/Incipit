package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type CreateLinkResponse struct {
	BaseResponse

	Link Link `json:"link"`
}

func NewCreateLinkResponse(link *entity.Link) (resp *CreateLinkResponse) {
	resLink := Link{
		ID:  link.ID,
		URL: link.URL,
	}

	return &CreateLinkResponse{
		BaseResponse: BaseResponse{Type: "create_link", Details: nil},
		Link:         resLink,
	}
}

func (resp *CreateLinkResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
