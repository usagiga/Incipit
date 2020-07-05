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
	// Authorization
	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminAuthModel &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.AdminAuthModel_ExpiredToken {
		return http.StatusForbidden
	}

	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminAuthModel &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.AdminAuthModel_UnmatchPassword {
		return http.StatusForbidden
	}

	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminAuthModel &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.AdminAuthModel_FailedToFindUser {
		return http.StatusForbidden
	}

	// JSON Binding
	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminUserHandler &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.AdminUserHandler_FailedBindJson {
		return http.StatusBadRequest
	}

	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminAuthHandler &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.AdminAuthHandler_FailedBindJson {
		return http.StatusBadRequest
	}

	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.LinkHandler &&
		interr.SecondaryErrorCode(resp.SecondaryErrorCode) == interr.LinkHandler_FailedBindJson {
		return http.StatusBadRequest
	}

	// Validation
	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.AdminUserValidation  {
		return http.StatusBadRequest
	}

	if interr.PrimaryErrorCode(resp.PrimaryErrorCode) == interr.LinkValidation  {
		return http.StatusBadRequest
	}


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
