package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type AdminUserHandlerImpl struct {
	adminModel model.AdminModel
}

func NewAdminUserHandler(adminModel model.AdminModel) AdminUserHandler {
	return &AdminUserHandlerImpl{adminModel: adminModel}
}

func (h *AdminUserHandlerImpl) HandleGetAdmin(c *gin.Context) {
	req := &messages.GetAdminRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminUserHandler, interr.AdminUserHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	found, err := h.adminModel.FindOne(req.ID)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewGetAdminResponse(found)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *AdminUserHandlerImpl) HandleCreateAdmin(c *gin.Context) {
	req := &messages.CreateAdminRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminUserHandler, interr.AdminUserHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	adding := &entity.AdminUser{
		Name:       req.Name,
		ScreenName: req.ScreenName,
		Password:   req.Password,
	}
	added, err := h.adminModel.Add(adding)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewCreateAdminResponse(added)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *AdminUserHandlerImpl) HandleUpdateAdmin(c *gin.Context) {
	req := &messages.UpdateAdminRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminUserHandler, interr.AdminUserHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	updating := &entity.AdminUser{
		Model:      gorm.Model{ID: req.ID},
		Name:       req.Name,
		ScreenName: req.ScreenName,
		Password:   req.Password,
	}
	updated, err := h.adminModel.Update(updating)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewUpdateAdminResponse(updated)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *AdminUserHandlerImpl) HandleDeleteAdmin(c *gin.Context) {
	req := &messages.DeleteAdminRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.AdminUserHandler, interr.AdminUserHandler_FailedBindJson, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	err = h.adminModel.Delete(req.ID)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewDeleteAdminResponse()
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}
