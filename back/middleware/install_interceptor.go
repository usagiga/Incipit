package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type InstallInterceptorImpl struct {
	installerModel model.InstallerModel
}

func NewInstallInterceptor(installerModel model.InstallerModel) InstallInterceptor {
	return &InstallInterceptorImpl{installerModel: installerModel}
}

func (m *InstallInterceptorImpl) HandleNeededInstall(c *gin.Context) {
	isNeeded, err := m.installerModel.IsNeededInstall()
	if err != nil {
		resp := messages.NewErrorResponse(err)
		sCode := resp.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, resp)
		return
	}

	if isNeeded {
		resp := messages.NewNeededInstallResponse()
		sCode := resp.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, resp)
		return
	}
}

func (m *InstallInterceptorImpl) HandleRedundantInstall(c *gin.Context) {
	isNeeded, err := m.installerModel.IsNeededInstall()
	if err != nil {
		resp := messages.NewErrorResponse(err)
		sCode := resp.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, resp)
		return
	}

	if !isNeeded {
		resp := messages.NewRedundantInstallResponse()
		sCode := resp.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, resp)
		return
	}
}
