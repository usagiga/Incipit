package messages

import (
	"github.com/usagiga/Incipit/back/entity"
	"net/http"
)

type GetLinkResponse struct {
	BaseResponse

	Links []Link `json:"links"`
}

func NewGetLinkResponse(links []entity.Link) (resp *GetLinkResponse) {
	resLinks := make([]Link, len(links))

	for i, link := range links {
		resLink := Link{
			ID:      link.ID,
			ShortID: link.GetShortID(),
			URL:     link.URL,
		}

		resLinks[i] = resLink
	}

	return &GetLinkResponse{
		BaseResponse: BaseResponse{Type: "get_link", Details: nil},
		Links:        resLinks,
	}
}

func (resp *GetLinkResponse) GetHTTPStatusCode() int {
	return http.StatusOK
}
