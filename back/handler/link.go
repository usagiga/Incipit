package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/usagiga/Incipit/back/entity"
	interr "github.com/usagiga/Incipit/back/entity/errors"
	"github.com/usagiga/Incipit/back/entity/messages"
	"github.com/usagiga/Incipit/back/model"
)

type LinkHandlerImpl struct {
	linkModel model.LinkModel
}

func NewLinkHandler(linkModel model.LinkModel) LinkHandler {
	return &LinkHandlerImpl{
		linkModel: linkModel,
	}
}

func (h *LinkHandlerImpl) HandleGetLink(c *gin.Context) {
	founds, err := h.linkModel.Find()
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewGetLinkResponse(founds)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *LinkHandlerImpl) HandleGetLinkByShortURL(c *gin.Context) {
	req := &messages.GetLinkByShortIDRequest{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind query parameter", interr.LinkHandler, interr.LinkHandler_FailedBindQuery, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	found, err := h.linkModel.FindOneByShortID(req.ShortID)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewGetLinkByShortIDResponse(found)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *LinkHandlerImpl) HandleCreateLink(c *gin.Context) {
	req := &messages.CreateLinkRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.LinkHandler, interr.LinkHandler_FailedBindQuery, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	creating := &entity.Link{
		URL: req.URL,
	}
	created, err := h.linkModel.Add(creating)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewCreateLinkResponse(created)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *LinkHandlerImpl) HandleUpdateLink(c *gin.Context) {
	req := &messages.UpdateLinkRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.LinkHandler, interr.LinkHandler_FailedBindQuery, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	updating := &entity.Link{
		Model: gorm.Model{ID: req.ID},
		URL: req.URL,
	}
	updated, err := h.linkModel.Update(updating)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewUpdateLinkResponse(updated)
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}

func (h *LinkHandlerImpl) HandleDeleteLink(c *gin.Context) {
	req := &messages.DeleteLinkRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		err = interr.NewDistinctError("Can't bind JSON", interr.LinkHandler, interr.LinkHandler_FailedBindQuery, nil)

		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	err = h.linkModel.Delete(req.ID)
	if err != nil {
		res := messages.NewErrorResponse(err)
		sCode := res.GetHTTPStatusCode()

		c.AbortWithStatusJSON(sCode, res)
		return
	}

	res := messages.NewDeleteLinkResponse()
	sCode := res.GetHTTPStatusCode()

	c.JSON(sCode, res)
}
