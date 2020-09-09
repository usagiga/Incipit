package handler

import (
	"github.com/gin-gonic/gin"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type InstallerHandlerImpl struct {
	installerModel model.InstallerModel
}

func NewInstallHandler(installerModel model.InstallerModel) InstallerHandler {
	return &InstallerHandlerImpl{installerModel: installerModel}
}

func (h *InstallerHandlerImpl) HandleIsInstalled(c *gin.Context) {
	res := messages.NewGetInstallResponse()
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *InstallerHandlerImpl) HandleInstall(c *gin.Context) {
	req := &messages.InstallRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminAuthHandler, interr.AdminAuthHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	accToken, refToken, err := h.installerModel.CreateNewAdmin(req.Name, req.ScreenName, req.Password)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewInstallResponse(accToken, refToken)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}
