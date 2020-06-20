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

func TestInstallInterceptorImpl_Handle(t *testing.T) {
	// Declare test cases
	testCases := []struct {
		ExpectedStatusCode    int
		ExpectedResponseType  string
		TestingInstallerModel model.InstallerModel
	}{
		{ExpectedStatusCode: http.StatusOK, ExpectedResponseType: "", TestingInstallerModel: installerModelStub{isNeeded: false, err: nil}},
		{ExpectedStatusCode: http.StatusServiceUnavailable, ExpectedResponseType: "required_install", TestingInstallerModel: installerModelStub{isNeeded: true, err: nil}},
		{ExpectedStatusCode: http.StatusInternalServerError, ExpectedResponseType: "error", TestingInstallerModel: installerModelStub{isNeeded: false, err: xerrors.New("something wrong")}},
	}

	// Do test
	for i, v := range testCases {
		caseNum := i + 1
		expectedCode := v.ExpectedStatusCode
		expectedRespType := v.ExpectedResponseType
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ii := &InstallInterceptorImpl{installerModel: v.TestingInstallerModel}

		// Exec
		ii.Handle(ctx)

		// Validate operation
		actualCode := w.Code
		actualRespBodyBytes := []byte(w.Body.String())
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

		if expectedCode != http.StatusOK && !ctx.IsAborted() {
			t.Errorf("Case %d: Not aborted response", caseNum)
		}
		if actualCode != expectedCode {
			t.Errorf("Case %d: Not valid HTTP status code of its response: Expected: %d, Actual: %d", caseNum, expectedCode, actualCode)
		}
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
