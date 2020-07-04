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
)

func TestLinkHandlerImpl_HandleCreateLink(t *testing.T) {
	li := NewLinkHandler(&linkModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", li.HandleCreateLink)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "create_link", ReqBodyStr: `{"url":"valid"}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "create_link", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`}, // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"url":"invalid"}`}, // Invalid URL
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

func TestLinkHandlerImpl_HandleGetLink(t *testing.T) {
	li := NewLinkHandler(&linkModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", li.HandleGetLink)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "get_link", ReqBodyStr: `{"id":1}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "get_link", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`}, // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"id":1000}`}, // Invalid URL
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

func TestLinkHandlerImpl_HandleGetLinkByShortURL(t *testing.T) {
	li := NewLinkHandler(&linkModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", li.HandleGetLinkByShortURL)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "get_link_by_short_id", ReqBodyStr: `{"short_id":"1"}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "get_link_by_short_id", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`}, // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"short_id":"0"}`}, // Invalid URL
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

func TestLinkHandlerImpl_HandleUpdateLink(t *testing.T) {
	li := NewLinkHandler(&linkModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", li.HandleUpdateLink)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "update_link", ReqBodyStr: `{"id":1}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "update_link", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`}, // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"id":1000}`}, // Invalid URL
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

func TestLinkHandlerImpl_HandleDeleteLink(t *testing.T) {
	li := NewLinkHandler(&linkModelStub{})

	_, router := gin.CreateTestContext(nil)
	router.GET("/test", li.HandleDeleteLink)

	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode   int
		ExpectedResponseType string
		ReqBodyStr           string
	}{
		{ExpectedStatusCode: 200, ExpectedResponseType: "delete_link", ReqBodyStr: `{"id":1}`}, // Valid
		{ExpectedStatusCode: 200, ExpectedResponseType: "delete_link", ReqBodyStr: `{"invalid":"invalid"}`}, // Unexpected JSON. There's no wrong point *in handler*, model will raise error
		{ExpectedStatusCode: 400, ExpectedResponseType: "error", ReqBodyStr: `invalid`}, // Invalid JSON
		{ExpectedStatusCode: 500, ExpectedResponseType: "error", ReqBodyStr: `{"id":1000}`}, // Invalid URL
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

type linkModelStub struct{}

func (m *linkModelStub) Add(link *entity.Link) (added *entity.Link, err error) {
	if link.URL == "invalid" {
		return nil, errors.New("invalid link")
	}
	return &entity.Link{Model: gorm.Model{ID:1}, URL: "valid"}, nil
}

func (m *linkModelStub) FindOne(id uint) (link *entity.Link, err error) {
	if id == 1000 {
		return nil, errors.New("invalid link")
	}
	return &entity.Link{Model: gorm.Model{ID:1}, URL: "valid"}, nil
}

func (m *linkModelStub) FindOneByShortID(shortId string) (link *entity.Link, err error) {
	if shortId == "0" {
		return nil, errors.New("invalid link")
	}
	return &entity.Link{Model: gorm.Model{ID:1}, URL: "valid"}, nil
}

func (m *linkModelStub) Find() (links []entity.Link, err error) {
	panic("implement me")
}

func (m *linkModelStub) Update(updating *entity.Link) (updated *entity.Link, err error) {
	if updating.ID == 1000 {
		return nil, errors.New("invalid link")
	}
	return &entity.Link{Model: gorm.Model{ID:1}, URL: "valid"}, nil
}

func (m *linkModelStub) Delete(id uint) (err error) {
	if id == 1000 {
		return errors.New("invalid link")
	}
	return nil
}
