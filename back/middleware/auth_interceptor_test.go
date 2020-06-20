package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/entity/messages"
	"golang.org/x/xerrors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthInterceptorImpl_Handle(t *testing.T) {
	ai := &AuthInterceptorImpl{adminAuthModel: &adminAuthModelStub{}}

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ExpectedUser         *entity.AdminUser
		TestingToken         string
	}{
		// In testing, AdminAuthModel is Stub. So it returns only 200 or 500 as status code.
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingToken: ""},
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingToken: "1234567890123456789012345678901"}, // 31 chars
		{ExpectedStatusCode: http.StatusOK, ExpectedResponseType: "", ExpectedUser: &entity.AdminUser{
			Name:       "test_admin",
			ScreenName: "test_admin",
			Password:   "test_admin",
		}, TestingToken: "12345678901234567890123456789012"},                                                                                   // 32 chars
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingToken: "123456789012345678901234567890123"}, // 33 chars
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		expectedCode := v.ExpectedStatusCode
		expectedRespType := v.ExpectedResponseType
		w := httptest.NewRecorder()
		_, router := gin.CreateTestContext(w)

		req := httptest.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+v.TestingToken)

		// Exec
		router.GET("/test", ai.Handle)
		router.ServeHTTP(w, req)

		// Validate operation
		actualCode := w.Code
		actualRespBodyBytes := w.Body.Bytes()
		if len(actualRespBodyBytes) == 0 && expectedRespType == "" {
			continue
		}

		actualResp := &messages.BaseResponse{}
		err := json.Unmarshal(actualRespBodyBytes, actualResp)
		if err != nil {
			t.Errorf("Case %d: Can't unmarshal response JSON", caseNum)
			continue
		}
		actualRespType := actualResp.Type

		if actualCode != expectedCode {
			t.Errorf("Case %d: Not valid HTTP status code of its response: Expected: %d, Actual: %d", caseNum, expectedCode, actualCode)
		}
		if actualRespType != expectedRespType {
			t.Errorf("Case %d: Not valid response body:\nExpected: %s\nActual: %s", caseNum, expectedRespType, actualRespType)
		}
	}
}

type adminAuthModelStub struct{}

func (m *adminAuthModelStub) Authorize(accTokenStr string) (user *entity.AdminUser, err error) {
	if len(accTokenStr) != 32 {
		return nil, xerrors.New("something wrong")
	}

	return &entity.AdminUser{
		Name:       "test_admin",
		ScreenName: "test_admin",
		Password:   "test_admin",
	}, nil
}

func (m *adminAuthModelStub) Login(name, password string) (
	accToken *entity.AccessToken,
	refToken *entity.RefreshToken,
	err error,
) {
	panic("implement me")
}

func (m *adminAuthModelStub) RenewAccessToken(refTokenStr string) (
	accToken *entity.AccessToken,
	refToken *entity.RefreshToken,
	err error,
) {
	panic("implement me")
}
