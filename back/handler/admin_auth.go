package handler

import (
	"github.com/gin-gonic/gin"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type AdminAuthHandlerImpl struct {
	adminAuthModel model.AdminAuthModel
}

func NewAdminAuthHandler(adminAuthModel model.AdminAuthModel) AdminAuthHandler {
	return &AdminAuthHandlerImpl{adminAuthModel: adminAuthModel}
}

func (h *AdminAuthHandlerImpl) HandleIsLogin(c *gin.Context) {
	res := messages.NewGetLoginResponse()
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *AdminAuthHandlerImpl) HandleLogin(c *gin.Context) {
	req := &messages.LoginAdminAuthRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminAuthHandler, interr.AdminAuthHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	accToken, refToken, err := h.adminAuthModel.Login(req.Name, req.Password)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewLoginAdminAuthResponse(accToken, refToken)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *AdminAuthHandlerImpl) HandleRefreshToken(c *gin.Context) {
	req := &messages.RefreshTokenAdminAuthRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminAuthHandler, interr.AdminAuthHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	accToken, refToken, err := h.adminAuthModel.RenewAccessToken(req.RefreshToken)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewRefreshTokenAdminAuthResponse(accToken, refToken)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}
