package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type GetLinkResponse struct {
	BaseResponse

	Link Link `json:"link"`
}

func NewGetLinkResponse(link *entity.Link) (resp *GetLinkResponse) {
	resLink := Link{
		ID:  link.ID,
		URL: link.URL,
	}

	return &GetLinkResponse{
		BaseResponse: BaseResponse{Type: "get_link", Details: nil},
		Link:         resLink,
	}
}

func (resp *GetLinkResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
