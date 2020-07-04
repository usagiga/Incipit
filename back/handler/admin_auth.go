package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type AdminAuthHandlerImpl struct {
	adminAuthModel model.AdminAuthModel
}

func NewAdminAuthHandler(adminAuthModel model.AdminAuthModel) AdminAuthHandler {
	return &AdminAuthHandlerImpl{adminAuthModel: adminAuthModel}
}

func (h *AdminAuthHandlerImpl) HandleLogin(c *gin.Context) {
	req := &messages.LoginAdminAuthRequest{}
	err := c.BindJSON(req)
	if err != nil {
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
	err := c.BindJSON(req)
	if err != nil {
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
