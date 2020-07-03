package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type GetLinkByShortIDResponse struct {
	BaseResponse

	Link Link `json:"link"`
}

func NewGetLinkByShortIDResponse(link *entity.Link) (resp *GetLinkByShortIDResponse) {
	resLink := Link{
		ID:  link.ID,
		URL: link.URL,
	}

	return &GetLinkByShortIDResponse{
		BaseResponse: BaseResponse{Type: "get_link_by_short_id", Details: nil},
		Link:         resLink,
	}
}

func (resp *GetLinkByShortIDResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
