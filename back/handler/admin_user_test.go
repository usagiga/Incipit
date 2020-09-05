package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/entity/messages"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdminUserHandlerImpl_HandleCreateAdmin(t *testing.T) {
	ai := NewAdminUserHandler(&adminModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", ai.HandleCreateAdmin)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "create_admin", ReqBodyStr: `{"name":"valid","screen_name":"valid","password":"valid"}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "create_admin", ReqBodyStr: `{"invalid":"invalid"}`},                                     // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`},                                                          // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"name":"invalid","screen_name":"invalid","password":"invalid"}`},  // Invalid URL
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

func TestAdminUserHandlerImpl_HandleGetAdmin(t *testing.T) {
	ai := NewAdminUserHandler(&adminModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", ai.HandleGetAdmin)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "get_admin"},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		w := httptest.NewRecorder()

		// Exec
		req := httptest.NewRequest("GET", "/test", nil)

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

func TestAdminUserHandlerImpl_HandleUpdateAdmin(t *testing.T) {
	ai := NewAdminUserHandler(&adminModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", ai.HandleUpdateAdmin)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "update_admin", ReqBodyStr: `{"id":1}`},              // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "update_admin", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`},                      // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"id":1000}`},                  // Invalid URL
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

func TestAdminUserHandlerImpl_HandleDeleteAdmin(t *testing.T) {
	ai := NewAdminUserHandler(&adminModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", ai.HandleDeleteAdmin)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "delete_admin", ReqBodyStr: `{"id":1}`},              // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "delete_admin", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`},                      // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"id":1000}`},                  // Invalid URL
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

type adminModelStub struct{}

func (m *adminModelStub) Add(user *entity.AdminUser) (added *entity.AdminUser, err error) {
	if user.Name == "invalid" {
		return nil, errors.New("invalid user")
	}

	return &entity.AdminUser{Name: "valid", ScreenName: "valid", Password: "valid"}, nil
}

func (m *adminModelStub) FindOne(id uint) (user *entity.AdminUser, err error) {
	if id == 1000 {
		return nil, errors.New("invalid user")
	}

	return &entity.AdminUser{Name: "valid", ScreenName: "valid", Password: "valid"}, nil
}

func (m *adminModelStub) FindOneByName(name string) (user *entity.AdminUser, err error) {
	panic("implement me")
}

func (m *adminModelStub) Find() (users []entity.AdminUser, err error) {
	return []entity.AdminUser{{Name: "valid", ScreenName: "valid", Password: "valid"}}, nil
}

func (m *adminModelStub) Update(updating *entity.AdminUser) (updated *entity.AdminUser, err error) {
	if updating.ID == 1000 {
		return nil, errors.New("invalid user")
	}

	return &entity.AdminUser{Name: "valid", ScreenName: "valid", Password: "valid"}, nil
}

func (m *adminModelStub) Delete(id uint) (err error) {
	if id == 1000 {
		return errors.New("invalid user")
	}

	return nil
}
