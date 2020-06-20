package handler

import (
	"github.com/gin-gonic/gin"
)

type LinkHandlerImpl struct {

}

func NewLinkHandler() LinkHandler {
	return &LinkHandlerImpl{}
}

func (h *LinkHandlerImpl) HandleGetLink(c *gin.Context) {
	panic("implement me")
}

func (h *LinkHandlerImpl) HandleGetLinkByShortURL(c *gin.Context) {
	panic("implement me")
}

func (h *LinkHandlerImpl) HandleCreateLink(c *gin.Context) {
	panic("implement me")
}

func (h *LinkHandlerImpl) HandleUpdateLink(c *gin.Context) {
	panic("implement me")
}

func (h *LinkHandlerImpl) HandleDeleteLink(c *gin.Context) {
	panic("implement me")
}
