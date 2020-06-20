package messages

import (
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"golang.org/x/xerrors"
	"net/http"
)

// ErrorResponse represents one of response which is sent to client
type ErrorResponse struct {
	BaseResponse

	PrimaryErrorCode   int `json:"p_code"`
	SecondaryErrorCode int `json:"s_code"`
}

// GetHTTPStatusCode returns HTTP status code determine from its internal error code
func (resp *ErrorResponse) GetHTTPStatusCode() int {
	// Temporary
	return http.StatusInternalServerError
}

// NewErrorResponse returns ErrorResponse which is made by error
func NewErrorResponse(err error) (resp Response) {
	var dError *interr.DistinctError
	ok := xerrors.As(err, &dError)
	if !ok {
		// Unknown error
		return &ErrorResponse{
			BaseResponse:       BaseResponse{Type: "error", Details: nil},
			PrimaryErrorCode:   0,
			SecondaryErrorCode: 0,
		}
	}

	// Known error
	return &ErrorResponse{
		BaseResponse:       BaseResponse{Type: "error", Details: dError.Detail},
		PrimaryErrorCode:   int(dError.PrimaryCode),
		SecondaryErrorCode: int(dError.SecondaryCode),
	}
}
