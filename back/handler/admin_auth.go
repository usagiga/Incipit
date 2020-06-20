package handler

import (
	"github.com/gin-gonic/gin"
)

type AdminAuthHandlerImpl struct {
}

func NewAdminAuthHandler() AdminAuthHandler {
	return &AdminAuthHandlerImpl{}
}

func (h *AdminAuthHandlerImpl) HandleLogin(c *gin.Context) {
	panic("implement me")
}

func (h *AdminAuthHandlerImpl) HandleRefreshToken(c *gin.Context) {
	panic("implement me")
}
