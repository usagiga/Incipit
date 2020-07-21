package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/entity/messages"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestInstallerHandlerImpl_HandleInstall(t *testing.T) {
	ih := NewInstallHandler(&installModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", ih.HandleInstall)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "install", ReqBodyStr: `{"name":"valid", "screen_name":"valid", "password":"valid"}`},  // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "install", ReqBodyStr: `{"invalid":"invalid"}`},                                        // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`},                                                        // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"name":"invalid","screen_name":"invalid","password":"invalid"}`}, // Invalid URL
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		w := httptest.NewRecorder()

		// Exec
		reqBodyReader := strings.NewReader(v.ReqBodyStr)
		req := httptest.NewRequest("GET", "/test", reqBodyReader)

		router.ServeHTTP(w, req)

		// Check status code
		expectedCode := v.ExpectedStatusCode
		actualCode := w.Code
		if actualCode != expectedCode {
			t.Errorf("Case %d: Not valid HTTP status code of its response: Expected: %d, Actual: %d", caseNum, expectedCode, actualCode)
		}

		// Check response type
		resBodyBytes := w.Body.Bytes()
		res := &messages.BaseResponse{}
		err := json.Unmarshal(resBodyBytes, &res)
		if err != nil {
			t.Errorf("Case %d: Can't unmarshal response", caseNum)
		}

		expectedRespType := v.ExpectedResponseType
		actualRespType := res.Type
		if actualRespType != expectedRespType {
			t.Errorf("Case %d: Not valid response body:\nExpected: %s\nActual: %s", caseNum, expectedRespType, actualRespType)
		}
	}
}

type installModelStub struct{}

func (m *installModelStub) CreateNewAdmin(name, screenName, password string) (
	accToken *entity.AccessToken,
	refToken *entity.RefreshToken,
	err error,
) {
	if name == "invalid" {
		return nil, nil, errors.New("invalid cred")
	}

	return &entity.AccessToken{
			Model:       gorm.Model{ID: 1},
			Token:       "valid",
			ExpiredAt:   time.Now(),
			AdminUserID: 1,
			AdminUser: entity.AdminUser{
				Model:      gorm.Model{ID: 1},
				Name:       "valid",
				ScreenName: "valid",
				Password:   "valid",
			},
		},
		&entity.RefreshToken{
			Model:       gorm.Model{ID: 1},
			Token:       "valid",
			ExpiredAt:   time.Now(),
			AdminUserID: 1,
			AdminUser: entity.AdminUser{
				Model:      gorm.Model{ID: 1},
				Name:       "valid",
				ScreenName: "valid",
				Password:   "valid",
			},
		},
		nil
}

func (m *installModelStub) IsNeededInstall() (isNeeded bool, err error) {
	panic("implement me")
}
