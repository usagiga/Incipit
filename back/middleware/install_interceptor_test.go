package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
	"golang.org/x/xerrors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInstallInterceptorImpl_HandleNeededInstall(t *testing.T) {
	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode    int
		ExpectedResponseType  string
		TestingInstallerModel model.InstallerModel
	}{
		{ExpectedStatusCode: http.StatusOK, ExpectedResponseType: "", TestingInstallerModel: installerModelStub{isNeeded: false, err: nil}},
		{ExpectedStatusCode: http.StatusServiceUnavailable, ExpectedResponseType: "needed_install", TestingInstallerModel: installerModelStub{isNeeded: true, err: nil}},
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingInstallerModel: installerModelStub{isNeeded: false, err: xerrors.New("something wrong")}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ii := &InstallInterceptorImpl{installerModel: v.TestingInstallerModel}

		// Exec
		ii.HandleNeededInstall(ctx)

		// Check status code
		expectedCode := v.ExpectedStatusCode
		actualCode := w.Code
		if expectedCode != http.StatusOK && !ctx.IsAborted() {
			t.Errorf("Case %d: Not aborted response", caseNum)
		}
		if actualCode != expectedCode {
			t.Errorf("Case %d: Not valid HTTP status code of its response: Expected: %d, Actual: %d", caseNum, expectedCode, actualCode)
		}

		// Check response type
		expectedRespType := v.ExpectedResponseType

		actualRespBodyBytes := []byte(w.Body.String())
		if len(actualRespBodyBytes) == 0 && expectedRespType == "" {
			continue
		}

		res := &messages.BaseResponse{}
		err := json.Unmarshal(actualRespBodyBytes, res)
		if err != nil {
			t.Errorf("Case %d: Can't unmarshal response JSON", caseNum)
			continue
		}

		actualRespType := res.Type

		if actualRespType != expectedRespType {
			t.Errorf("Case %d: Not valid response body:\nExpected: %s\nActual: %s", caseNum, expectedRespType, actualRespType)
		}
	}
}

func TestInstallInterceptorImpl_HandleRedundantInstall(t *testing.T) {
	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode    int
		ExpectedResponseType  string
		TestingInstallerModel model.InstallerModel
	}{
		{ExpectedStatusCode: http.StatusOK, ExpectedResponseType: "", TestingInstallerModel: installerModelStub{isNeeded: true, err: nil}},
		{ExpectedStatusCode: http.StatusGone, ExpectedResponseType: "redundant_install", TestingInstallerModel: installerModelStub{isNeeded: false, err: nil}},
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingInstallerModel: installerModelStub{isNeeded: false, err: xerrors.New("something wrong")}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ii := &InstallInterceptorImpl{installerModel: v.TestingInstallerModel}

		// Exec
		ii.HandleRedundantInstall(ctx)

		// Check status code
		expectedCode := v.ExpectedStatusCode
		actualCode := w.Code
		if expectedCode != http.StatusOK && !ctx.IsAborted() {
			t.Errorf("Case %d: Not aborted response", caseNum)
		}
		if actualCode != expectedCode {
			t.Errorf("Case %d: Not valid HTTP status code of its response: Expected: %d, Actual: %d", caseNum, expectedCode, actualCode)
		}

		// Check response type
		expectedRespType := v.ExpectedResponseType

		actualRespBodyBytes := []byte(w.Body.String())
		if len(actualRespBodyBytes) == 0 && expectedRespType == "" {
			continue
		}

		res := &messages.BaseResponse{}
		err := json.Unmarshal(actualRespBodyBytes, res)
		if err != nil {
			t.Errorf("Case %d: Can't unmarshal response JSON", caseNum)
			continue
		}

		actualRespType := res.Type

		if actualRespType != expectedRespType {
			t.Errorf("Case %d: Not valid response body:\nExpected: %s\nActual: %s", caseNum, expectedRespType, actualRespType)
		}
	}
}

type installerModelStub struct {
	isNeeded bool
	err      error
}

func (m installerModelStub) CreateNewAdmin(name, screenName, password string) (
	accToken *entity.AccessToken,
	refToken *entity.RefreshToken,
	err error,
) {
	panic("implement me")
}

func (m installerModelStub) IsNeededInstall() (isNeeded bool, err error) {
	return m.isNeeded, m.err
}
